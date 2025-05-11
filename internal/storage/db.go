package storage

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/database/repositories"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
	"math"
	"math/big"
)

type key_db int

const KeyDBStorage key_db = iota

type DBStorage struct {
	flagRepository              *repositories.FlagRepository
	boolVariationRepository     *repositories.FlagVariationRepository[bool]
	stringVariationRepository   *repositories.FlagVariationRepository[string]
	floatVariationRepository    *repositories.FlagVariationRepository[float64]
	intVariationRepository      *repositories.FlagVariationRepository[float64]
	jsonVariationRepository     *repositories.FlagVariationRepository[datatypes.JSON]
	flagRulesRepository         *repositories.RuleRepository
	evaluationContextRepository *repositories.EvaluationContextRepository
	rolloutRepository           *repositories.RolloutRepository
}

var _ Storageable = &DBStorage{}

func NewDBStorage() *DBStorage {
	db := database.GetDB()
	return &DBStorage{
		flagRepository:              &repositories.FlagRepository{DB: db},
		boolVariationRepository:     &repositories.FlagVariationRepository[bool]{DB: db},
		stringVariationRepository:   &repositories.FlagVariationRepository[string]{DB: db},
		floatVariationRepository:    &repositories.FlagVariationRepository[float64]{DB: db},
		intVariationRepository:      &repositories.FlagVariationRepository[float64]{DB: db},
		flagRulesRepository:         &repositories.RuleRepository{DB: db},
		evaluationContextRepository: &repositories.EvaluationContextRepository{DB: db},
		rolloutRepository:           &repositories.RolloutRepository{DB: db},
		jsonVariationRepository:     &repositories.FlagVariationRepository[datatypes.JSON]{DB: db},
	}
}

func getBucketPercentage(keyString string) float64 {
	h := sha256.New()
	h.Write([]byte(keyString))
	hashSum := h.Sum(nil)
	hashSumInt := new(big.Int).SetBytes(hashSum)

	bucketPercentage := float64(hashSumInt.Int64()) / float64(math.MaxInt64)
	retVal := math.Ceil(math.Abs(bucketPercentage) * 100)
	log.Debug().Msgf("Bucket percentage: %f", retVal)
	return retVal
}

func getRolloutVariant(percentage float64, rollout database.Rollout) (string, error) {
	var rolloutConfig payloads.RolloutConfig
	if err := json.Unmarshal([]byte(rollout.Config), &rolloutConfig); err != nil {
		log.Error().Msg("Error unmarshalling rollout config")
		return "", err
	}

	lenVariants := len(rolloutConfig.Variants)
	for n := range lenVariants - 1 {
		swapped := false
		for i := range lenVariants - n - 1 {
			if rolloutConfig.Variants[i].Percentage > rolloutConfig.Variants[i+1].Percentage {
				rolloutConfig.Variants[i], rolloutConfig.Variants[i+1] = rolloutConfig.Variants[i+1], rolloutConfig.Variants[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	type bucket struct {
		start         int
		end           int
		variationUUID string
	}
	var startingPoint int
	var buckets []bucket
	for idx, variant := range rolloutConfig.Variants {
		end := variant.Percentage
		if idx == len(rolloutConfig.Variants)-1 {
			end = 100
		}
		buckets = append(buckets, bucket{
			start:         startingPoint,
			end:           end,
			variationUUID: variant.VariationUUID,
		})
		startingPoint = variant.Percentage
	}

	for _, bucket := range buckets {
		percentageAsInt := int(percentage)
		if bucket.start < percentageAsInt && percentageAsInt <= bucket.end {
			return bucket.variationUUID, nil
		}
	}

	log.Debug().Msgf("rollout config variants: %v", rolloutConfig.Variants)

	return "", nil
}

func (s *DBStorage) EvaluateFlag(key string, payload payloads.EvaluateFlag) (any, error) {
	flagKey, result := s.flagRepository.GetFlagKey(key)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	var value any
	var _err error
	var contextualVariation datatypes.UUID

	if payload.Context.Attributes != nil && flagKey.Enabled {
		go func() {
			_, err := s.evaluationContextRepository.CreateEvaluationContext(payload.Context)
			if err != nil {
				log.Error().Msgf("Error creating evaluation context: %s", err)
			}
		}()

		flagRules, result := s.flagRulesRepository.GetTargetingRules(flagKey.UUID)
		if result.RowsAffected > 0 {
			log.Debug().Msgf("Flag rules: %v", flagRules)
			for _, fr := range flagRules {
				if fr.Attributes != nil {
					var flagRuleContexts []payloads.RuleContext
					if err := json.Unmarshal([]byte(fr.Attributes), &flagRuleContexts); err != nil {
						return nil, err
					}
					if hasMatchingRules(flagRuleContexts, payload.Context.Attributes) {
						contextualVariation = fr.VariationUUID
						break
					}
				}
			}
		}
	}

	if contextualVariation.IsNil() {
		contextualVariation = flagKey.DefaultVariation

		rollout, err := s.rolloutRepository.GetRollout(flagKey.UUID)
		if err != nil {
			return nil, err
		}
		if rollout.Config != nil && flagKey.Enabled && payload.Context.Key != "" {
			keyString := fmt.Sprintf("%x-%s", flagKey.UUID, payload.Context.Key)
			bucketPercentage := getBucketPercentage(keyString)
			rolloutVariation, err := getRolloutVariant(bucketPercentage, rollout)
			if err != nil {
				return nil, err
			}
			contextualVariation = datatypes.UUID(uuid.MustParse(rolloutVariation))
		} else if flagKey.Enabled {
			contextualVariation = flagKey.DefaultEnabledVariation
		}
	}

	switch flagKey.FlagType {
	case "string":
		value, _err = s.stringVariationRepository.GetFlagVariationValue(contextualVariation)
	case "int":
		value, _err = s.intVariationRepository.GetFlagVariationValue(contextualVariation)
	case "float":
		value, _err = s.floatVariationRepository.GetFlagVariationValue(contextualVariation)
	case "bool":
		value, _err = s.boolVariationRepository.GetFlagVariationValue(contextualVariation)
	case "json", "object":
		value, _err = s.jsonVariationRepository.GetFlagVariationValue(contextualVariation)
	}
	return value, _err
}

func (s *DBStorage) UpdateFlag(payload payloads.UpdateFlag) error {
	flagKey, result := s.flagRepository.GetFlagKey(payload.Key)

	if result.RowsAffected == 0 {
		return result.Error
	}
	gormTx := s.flagRepository.DB.Begin()

	gormTx.Begin()
	flagKey.Enabled = payload.Enabled
	flagKey.Name = &payload.Name
	flagKey.Description = &payload.Description
	if payload.DefaultVariation != "" {
		exists, _ := s.verifyFlagVariationExists(flagKey.UUID.String(), payload.DefaultVariation)
		if exists {
			flagKey.DefaultVariation = datatypes.UUID(uuid.MustParse(payload.DefaultVariation))
		} else {
			return fmt.Errorf("flag variation %s does not belong to flag key %s", payload.DefaultVariation, flagKey.UUID.String())
		}
	}
	if payload.DefaultEnabledVariation != "" {
		exists, _ := s.verifyFlagVariationExists(flagKey.UUID.String(), payload.DefaultEnabledVariation)
		if exists {
			flagKey.DefaultEnabledVariation = datatypes.UUID(uuid.MustParse(payload.DefaultEnabledVariation))
		} else {
			return fmt.Errorf("flag variation %s does not belong to flag key %s", payload.DefaultVariation, flagKey.UUID.String())
		}
	}
	_, result = s.flagRepository.UpdateFlagKey(&flagKey, gormTx)
	if result.Error != nil {
		return result.Error
	}
	gormTx.Commit()

	return nil
}

func (s *DBStorage) verifyFlagVariationExists(flagKeyUUID string, flagVariationUUID string) (bool, error) {
	parsedFlagVariationUUID := datatypes.UUID(uuid.MustParse(flagVariationUUID))
	parsedFlagKeyUUID := datatypes.UUID(uuid.MustParse(flagKeyUUID))
	flagKey, result := s.flagRepository.GetFlagKeyByUUID(parsedFlagKeyUUID.String())
	if result.RowsAffected == 0 {
		return false, result.Error
	}

	switch flagKey.FlagType {
	case "bool":
		flagVariation, err := s.boolVariationRepository.GetFlagKeyVariationByUUID(parsedFlagVariationUUID)
		if err != nil {
			return false, err
		}
		return flagVariation.FlagKeyUUID == flagKey.UUID, nil
	case "string":
		flagVariation, err := s.stringVariationRepository.GetFlagKeyVariationByUUID(parsedFlagVariationUUID)
		if err != nil {
			return false, err
		}
		return flagVariation.FlagKeyUUID == flagKey.UUID, nil
	case "float":
		flagVariation, err := s.floatVariationRepository.GetFlagKeyVariationByUUID(parsedFlagVariationUUID)
		if err != nil {
			return false, err
		}
		return flagVariation.FlagKeyUUID == flagKey.UUID, nil
	case "int":
		flagVariation, err := s.intVariationRepository.GetFlagKeyVariationByUUID(parsedFlagVariationUUID)
		if err != nil {
			return false, err
		}
		return flagVariation.FlagKeyUUID == flagKey.UUID, nil
	default:
		return false, nil
	}
}

func (s *DBStorage) PutRule(payload payloads.PutRule) error {
	log.Printf("setting rule: %s\n", payload.Name)
	variationExists, err := s.verifyFlagVariationExists(payload.FlagUUID, payload.VariationUUID)
	if err != nil {
		return err
	}
	if !variationExists {
		return fmt.Errorf("flag variation %s is not a variation of flag %s", payload.VariationUUID, payload.FlagUUID)
	}
	gormTx := s.flagRulesRepository.DB.Begin()

	err = s.flagRulesRepository.SaveTargetingRule(payload, gormTx)
	if err != nil {
		return err
	}

	gormTx.Commit()

	return nil
}

func (s *DBStorage) CreateFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	_, result := s.flagRepository.GetFlagKey(key)
	if result.RowsAffected > 0 {
		return fmt.Errorf("unable to create flag with key %s", key)
	}
	gormTx := s.flagRepository.DB.Begin()

	var uuids []datatypes.UUID
	flagKey, result := s.flagRepository.CreateFlagKey(flagType, key, gormTx)
	if result.Error != nil {
		gormTx.Rollback()
		return result.Error
	}

	if flagType == "bool" && len(variations) > 2 {
		return fmt.Errorf("can only have two variations for boolean flags")
	}

	for _, value := range variations {
		switch flagKey.FlagType {
		case "bool":
			flagVariation, result := s.boolVariationRepository.CreateFlagKeyVariation(flagKey, value, gormTx)
			if result.Error != nil {
				gormTx.Rollback()
				return result.Error
			}
			uuids = append(uuids, flagVariation.UUID)
		case "string":
			flagVariation, result := s.stringVariationRepository.CreateFlagKeyVariation(flagKey, value, gormTx)
			if result.Error != nil {
				gormTx.Rollback()
				return result.Error
			}
			uuids = append(uuids, flagVariation.UUID)
		case "float":
			flagVariation, result := s.floatVariationRepository.CreateFlagKeyVariation(flagKey, value, gormTx)
			if result.Error != nil {
				gormTx.Rollback()
				return result.Error
			}
			uuids = append(uuids, flagVariation.UUID)
		case "int":
			flagVariation, result := s.intVariationRepository.CreateFlagKeyVariation(flagKey, value, gormTx)
			if result.Error != nil {
				gormTx.Rollback()
				return result.Error
			}
			uuids = append(uuids, flagVariation.UUID)
		case "json", "object":
			v, err := json.Marshal(value.Value)
			value.Value = datatypes.JSON(v)
			if err != nil {
				return err
			}
			flagVariation, result := s.jsonVariationRepository.CreateFlagKeyVariation(flagKey, value, gormTx)
			if result.Error != nil {
				gormTx.Rollback()
				return result.Error
			}
			uuids = append(uuids, flagVariation.UUID)
		}
	}

	lenVariations := len(uuids)
	flagKey.DefaultEnabledVariation = uuids[0]
	flagKey.DefaultVariation = uuids[lenVariations-1]

	_, result = s.flagRepository.UpdateFlagKey(&flagKey, gormTx)
	if result.Error != nil {
		return result.Error
	}
	gormTx.Commit()

	return nil
}

func (s *DBStorage) GetFlagWithVariations(key string) (payloads.FlagRepresentation, error) {
	flagKey, result := s.flagRepository.GetFlagKey(key)
	var response payloads.FlagRepresentation
	if result.RowsAffected == 0 {
		return response, result.Error
	}
	if flagKey.Name == nil {
		flagKey.Name = new(string)
	}
	if flagKey.Description == nil {
		flagKey.Description = new(string)
	}
	response.FlagUUID = flagKey.UUID.String()
	response.Key = flagKey.Key
	response.Name = *flagKey.Name
	response.Description = *flagKey.Description
	response.FlagType = flagKey.FlagType
	response.Enabled = flagKey.Enabled
	response.CreatedAt = flagKey.CreatedAt
	response.UpdatedAt = flagKey.UpdatedAt
	rules, result := s.flagRulesRepository.GetTargetingRules(flagKey.UUID)
	if result.Error != nil {
		return payloads.FlagRepresentation{}, result.Error
	}
	var responseRules []payloads.FlagRuleResponse
	for _, rule := range rules {
		responseRules = append(responseRules, payloads.FlagRuleResponse{
			UUID:          rule.UUID.String(),
			Name:          rule.Name,
			VariationUUID: rule.VariationUUID.String(),
			Priority:      rule.Priority,
			RuleContexts:  rule.Attributes,
		})
	}
	response.Rules = responseRules

	var responseRollout payloads.FlagRolloutResponse
	rollout, err := s.rolloutRepository.GetRollout(flagKey.UUID)
	if err != nil {
		return payloads.FlagRepresentation{}, nil
	}
	if rollout.RolloutType != "" {
		responseRollout.RolloutType = rollout.RolloutType
		err = json.Unmarshal(rollout.Config, &responseRollout.Config)
		if err != nil {
			return payloads.FlagRepresentation{}, err
		}
	}

	response.Rollout = responseRollout

	var enabledValue, disabledValue any
	var enabledErr, disabledErr error

	switch flagKey.FlagType {
	case "string":
		enabledValue, enabledErr = s.stringVariationRepository.GetFlagVariationValue(flagKey.DefaultEnabledVariation)
		disabledValue, disabledErr = s.stringVariationRepository.GetFlagVariationValue(flagKey.DefaultVariation)
	case "int":
		enabledValue, enabledErr = s.intVariationRepository.GetFlagVariationValue(flagKey.DefaultEnabledVariation)
		disabledValue, disabledErr = s.intVariationRepository.GetFlagVariationValue(flagKey.DefaultVariation)
	case "float":
		enabledValue, enabledErr = s.floatVariationRepository.GetFlagVariationValue(flagKey.DefaultEnabledVariation)
		disabledValue, disabledErr = s.floatVariationRepository.GetFlagVariationValue(flagKey.DefaultVariation)
	case "bool":
		enabledValue, enabledErr = s.boolVariationRepository.GetFlagVariationValue(flagKey.DefaultEnabledVariation)
		disabledValue, disabledErr = s.boolVariationRepository.GetFlagVariationValue(flagKey.DefaultVariation)
	case "json", "object":
		enabledValue, enabledErr = s.jsonVariationRepository.GetFlagVariationValue(flagKey.DefaultEnabledVariation)
		disabledValue, disabledErr = s.jsonVariationRepository.GetFlagVariationValue(flagKey.DefaultVariation)
	default:
		enabledValue, enabledErr = nil, fmt.Errorf("%s is not a valid flag type", flagKey.FlagType)
		disabledValue, disabledErr = nil, fmt.Errorf("%s is not a valid flag type", flagKey.FlagType)
	}

	if enabledErr != nil {
		return payloads.FlagRepresentation{}, enabledErr
	}
	response.DefaultEnabledVariation = enabledValue
	if disabledErr != nil {
		return payloads.FlagRepresentation{}, disabledErr
	}
	response.DefaultDisabledVariation = disabledValue

	return response, nil
}

func (s *DBStorage) GetFlagVariations(key string) ([]payloads.FlagVariationResponse, error) {
	flagKey, result := s.flagRepository.GetFlagKey(key)
	var response []payloads.FlagVariationResponse
	if result.RowsAffected == 0 {
		return response, result.Error
	}
	switch flagKey.FlagType {
	case "string":
		variations, err := s.stringVariationRepository.GetFlagVariations(flagKey.UUID)
		for _, variation := range variations {
			response = append(response, payloads.FlagVariationResponse{
				UUID:     variation.UUID.String(),
				FlagUUID: variation.FlagKeyUUID.String(),
				Name:     variation.Name,
				Value:    variation.Value,
			})
		}
		if err != nil {
			return []payloads.FlagVariationResponse{}, err
		}
	case "float", "int":
		variations, err := s.intVariationRepository.GetFlagVariations(flagKey.UUID)
		for _, variation := range variations {
			response = append(response, payloads.FlagVariationResponse{
				UUID:     variation.UUID.String(),
				FlagUUID: variation.FlagKeyUUID.String(),
				Name:     variation.Name,
				Value:    variation.Value,
			})
		}
		if err != nil {
			return []payloads.FlagVariationResponse{}, err
		}
	case "bool":
		variations, err := s.boolVariationRepository.GetFlagVariations(flagKey.UUID)
		for _, variation := range variations {
			response = append(response, payloads.FlagVariationResponse{
				UUID:     variation.UUID.String(),
				FlagUUID: variation.FlagKeyUUID.String(),
				Name:     variation.Name,
				Value:    variation.Value,
			})
		}
		if err != nil {
			return []payloads.FlagVariationResponse{}, err
		}
	case "json", "object":
		variations, err := s.jsonVariationRepository.GetFlagVariations(flagKey.UUID)
		for _, variation := range variations {
			response = append(response, payloads.FlagVariationResponse{
				UUID:     variation.UUID.String(),
				FlagUUID: variation.FlagKeyUUID.String(),
				Name:     variation.Name,
				Value:    variation.Value,
			})
		}
		if err != nil {
			return []payloads.FlagVariationResponse{}, err
		}
	default:
		return []payloads.FlagVariationResponse{}, fmt.Errorf("%s is not a valid flag key", flagKey.FlagType)
	}
	return response, nil
}

func (s *DBStorage) UpdateFlagVariation(payload payloads.UpdateFlagVariation) error {
	flagKey, result := s.flagRepository.GetFlagKeyByUUID(payload.FlagKeyUUID)
	gormTx := s.flagRepository.DB.Begin()
	if result.RowsAffected == 0 {
		return fmt.Errorf("flag key does not exist with uuid %s", flagKey.UUID.String())
	}
	variationUUID := datatypes.UUID(uuid.MustParse(payload.UUID))
	flagUUID := datatypes.UUID(uuid.MustParse(payload.FlagKeyUUID))

	switch flagKey.FlagType {
	case "string":
		variation := database.FlagVariation[string]{
			UUID:        variationUUID,
			FlagKeyUUID: flagUUID,
			Name:        payload.Name,
			Value:       payload.Value.(string),
		}
		_, result = s.stringVariationRepository.UpdateFlagVariation(variation, gormTx)
		if result.Error != nil {
			return result.Error
		}
	case "float", "int":
		variation := database.FlagVariation[float64]{
			UUID:        variationUUID,
			FlagKeyUUID: flagUUID,
			Name:        payload.Name,
			Value:       payload.Value.(float64),
		}
		_, result = s.floatVariationRepository.UpdateFlagVariation(variation, gormTx)
		if result.Error != nil {
			return result.Error
		}
	case "bool":
		variation := database.FlagVariation[bool]{
			UUID:        variationUUID,
			FlagKeyUUID: flagUUID,
			Name:        payload.Name,
			Value:       payload.Value.(bool),
		}
		_, result = s.boolVariationRepository.UpdateFlagVariation(variation, gormTx)
	default:
		return fmt.Errorf("flag variation %s is misconfigured", payload.UUID)
	}

	if result.Error != nil {
		gormTx.Rollback()
		return result.Error
	}
	gormTx.Commit()
	return nil
}

func (s *DBStorage) SetupRolloutRule(payload payloads.PutRolloutRule) error {
	if payload.UUID == "" {
		payload.UUID = uuid.New().String()
	}
	var percs int
	for _, variant := range payload.ContextConfig.Variants {
		variationExists, err := s.verifyFlagVariationExists(payload.FlagUUID, variant.VariationUUID)
		if err != nil {
			return err
		}
		if !variationExists {
			return fmt.Errorf("flag variation %s is not a variation of flag %s", variant.VariationUUID, payload.FlagUUID)
		}
		percs += variant.Percentage
	}

	if percs != 100 {
		return fmt.Errorf("sum of variant percentages does not equal 100: %v", payload.ContextConfig.Variants)
	}
	_, err := s.rolloutRepository.SetupRollout(payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBStorage) DeleteRolloutRule(rolloutUUID string) error {
	err := s.rolloutRepository.DeleteRollout(rolloutUUID)
	if err != nil {
		return err
	}
	return nil
}

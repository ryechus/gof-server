package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"

	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/database/repositories"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
)

type key_db int

const KeyDBStorage key_db = iota

type DBStorage struct {
	flagRepository            *repositories.FlagRepository
	boolVariationRepository   *repositories.FlagVariationRepository[bool]
	stringVariationRepository *repositories.FlagVariationRepository[string]
	floatVariationRepository  *repositories.FlagVariationRepository[float64]
	intVariationRepository    *repositories.FlagVariationRepository[float64]
	flagRulesRepository       *repositories.RuleRepository
}

var _ Storageable = &DBStorage{}

func NewDBStorage() *DBStorage {
	db := database.GetDB()
	return &DBStorage{
		flagRepository:            &repositories.FlagRepository{DB: db},
		boolVariationRepository:   &repositories.FlagVariationRepository[bool]{DB: db},
		stringVariationRepository: &repositories.FlagVariationRepository[string]{DB: db},
		floatVariationRepository:  &repositories.FlagVariationRepository[float64]{DB: db},
		intVariationRepository:    &repositories.FlagVariationRepository[float64]{DB: db},
		flagRulesRepository:       &repositories.RuleRepository{DB: db},
	}
}

func hasMatchingRules(flagRuleContexts []payloads.RuleContext, attributes map[string]any) bool {
	for key, value := range attributes {
		for _, ctx := range flagRuleContexts {
			if ctx.Attribute == key && slices.Contains(ctx.Values, value.(string)) {
				return true
			}
		}
	}
	return false
}

func (s *DBStorage) EvaluateFlag(key string, payload payloads.GetFlag) (any, error) {
	flagKey, result := s.flagRepository.GetFlagKey(key)
	log.Printf("evaluating %v", flagKey)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	var value any
	var _err error
	var contextualVariation datatypes.UUID

	if payload.Context.Attributes != nil && flagKey.Enabled {
		log.Println("evaluating context")
		flagRules, result := s.flagRulesRepository.GetTargetingRules(flagKey.UUID)
		if result.RowsAffected > 0 {
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
		if flagKey.Enabled {
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
	}
	return value, _err
}

func (s *DBStorage) UpdateFlag(payload payloads.UpdateFlag) error {
	log.Printf("Updating flag state for key %s to %t\n", payload.Key, payload.Enabled)

	flagKey, result := s.flagRepository.GetFlagKey(payload.Key)

	if result.RowsAffected == 0 {
		return result.Error
	}
	gormTx := s.flagRepository.DB.Begin()

	gormTx.Begin()
	flagKey.Enabled = payload.Enabled
	s.flagRepository.UpdateFlagKey(&flagKey, gormTx)
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
	log.Printf("Setting flag value for key %s\n", key)
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
		}
	}

	lenVariations := len(uuids)
	flagKey.DefaultEnabledVariation = uuids[0]
	flagKey.DefaultVariation = uuids[lenVariations-1]

	s.flagRepository.UpdateFlagKey(&flagKey, gormTx)
	gormTx.Commit()

	return nil
}

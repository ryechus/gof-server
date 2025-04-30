package storage

import (
	"fmt"
	"time"

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

func (s *DBStorage) GetBool(key string) (bool, error) { return GetFlag[bool](key) }
func (s *DBStorage) CreateBoolFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return CreateFlag[bool](key, flagType, variations)
}
func (s *DBStorage) SetBool(key string, value bool) error {
	variations := []payloads.FlagVariation{
		{
			Name:  "",
			Value: value,
		},
		{
			Name:  "",
			Value: value,
		},
	}

	CreateFlag[bool](key, "bool", variations)
	return nil
}

func (s *DBStorage) GetInt(key string) (int64, error) { return GetFlag[int64](key) }
func (s *DBStorage) CreateIntFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return CreateFlag[int64](key, flagType, variations)
}
func (s *DBStorage) SetInt(key string, value int64) error {
	variations := []payloads.FlagVariation{
		{
			Name:  "",
			Value: value,
		},
		{
			Name:  "",
			Value: value,
		},
	}

	CreateFlag[int64](key, "int", variations)
	return nil
}

func (s *DBStorage) GetFloat(key string) (float64, error) { return GetFlag[float64](key) }
func (s *DBStorage) CreateFloatFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return CreateFlag[float64](key, flagType, variations)
}
func (s *DBStorage) SetFloat(key string, value float64) error {
	variations := []payloads.FlagVariation{
		{
			Name:  "",
			Value: value,
		},
		{
			Name:  "",
			Value: value,
		},
	}

	CreateFlag[float64](key, "float", variations)
	return nil
}

func (s *DBStorage) GetString(key string) (string, error) { return GetFlag[string](key) }
func (s *DBStorage) CreateStringFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return CreateFlag[string](key, flagType, variations)
}
func (s *DBStorage) SetString(key, value string) error {
	variations := []payloads.FlagVariation{
		{
			Name:  "",
			Value: value,
		},
		{
			Name:  "",
			Value: value,
		},
	}

	CreateFlag[string](key, "string", variations)
	return nil
}

func (s *DBStorage) EvaluateFlag(key string, payload payloads.GetFlag) (any, error) {
	flagKey, result := s.flagRepository.GetFlagKey(key)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	var value any
	var _err error
	var contextualVariation datatypes.UUID

	if payload.Context.Attributes != nil && flagKey.Enabled {
		flagRules, result := s.flagRulesRepository.GetTargetingRules(flagKey.UUID)
		if result.RowsAffected > 0 {
		outerLoop:
			for _, fr := range flagRules {
				allFlagContexts := []database.TargetingRuleContext{}
				flagRuleContexts, _ := s.flagRulesRepository.GetTargetingRuleContexts(fr.UUID)
				allFlagContexts = append(allFlagContexts, flagRuleContexts...)
				for key, value := range payload.Context.Attributes {
					for _, ctx := range allFlagContexts {
						if ctx.Attribute == key && ctx.Value == value {
							contextualVariation = fr.VariationUUID
							break outerLoop
						}
					}
				}
			}
		}
	}

	if !contextualVariation.IsNil() {
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
	} else { // TODO MAKE THIS BRANCH GO AWAY
		switch flagKey.FlagType {
		case "string":
			value, _err = s.GetString(flagKey.Key)
		case "int":
			value, _err = s.GetFloat(flagKey.Key)
		case "float":
			value, _err = s.GetFloat(flagKey.Key)
		case "bool":
			value, _err = s.GetBool(flagKey.Key)
		}
	}
	return value, _err
}

func (s *DBStorage) UpdateFlag(payload payloads.UpdateFlag) error {
	fmt.Printf("Updating flag state for key %s to %t\n", payload.Key, payload.Enabled)

	flagKey, result := s.flagRepository.GetFlagKey(payload.Key)

	if result.RowsAffected == 0 {
		return result.Error
	}
	flagKey.Enabled = payload.Enabled
	s.flagRepository.UpdateFlagKey(&flagKey)

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
	if payload.UUID == "" {
		fmt.Printf("creating a new rule %s\n", payload.Name)
		variationExists, err := s.verifyFlagVariationExists(payload.FlagUUID, payload.VariationUUID)
		if err != nil {
			return err
		}
		if !variationExists {
			return fmt.Errorf("flag variation %s is not a variation of flag %s", payload.VariationUUID, payload.FlagUUID)
		}

		s.flagRulesRepository.SaveTargetingRule(payload)
	} else {
		fmt.Printf("updating rule %s\n", payload.Name)
	}
	return nil
}

func (s *DBStorage) CreateFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	fmt.Printf("Setting flag value for key %s\n", key)
	_, result := s.flagRepository.GetFlagKey(key)
	if result.RowsAffected > 0 {
		return fmt.Errorf("unable to create flag with key %s", key)
	}

	var uuids []datatypes.UUID
	flagKey, _ := s.flagRepository.CreateFlagKey(flagType, key)

	if flagType == "bool" && len(variations) > 2 {
		return fmt.Errorf("can only have two variations for boolean flags")
	}

	for _, value := range variations {
		switch flagKey.FlagType {
		case "bool":
			flagVariation, _ := s.boolVariationRepository.CreateFlagKeyVariation(flagKey, value)
			uuids = append(uuids, flagVariation.UUID)
		case "string":
			flagVariation, _ := s.stringVariationRepository.CreateFlagKeyVariation(flagKey, value)
			uuids = append(uuids, flagVariation.UUID)
		case "float":
			flagVariation, _ := s.floatVariationRepository.CreateFlagKeyVariation(flagKey, value)
			uuids = append(uuids, flagVariation.UUID)
		case "int":
			flagVariation, _ := s.intVariationRepository.CreateFlagKeyVariation(flagKey, value)
			uuids = append(uuids, flagVariation.UUID)
		}
	}

	lenVariations := len(uuids)
	flagKey.DefaultEnabledVariation = uuids[0]
	flagKey.DefaultVariation = uuids[lenVariations-1]

	s.flagRepository.UpdateFlagKey(&flagKey)

	return nil
}

func CreateFlag[T comparable](key string, flagType string, variations []payloads.FlagVariation) error {
	fmt.Printf("Setting flag value for key %s\n", key)
	db := database.GetDB()
	var flagKey database.FlagKey

	result := db.First(&flagKey, "key = ?", key)
	var uuids []datatypes.UUID

	now := time.Now()
	if result.RowsAffected == 0 {
		flag_key_uuid := datatypes.NewUUIDv4()
		newFlag := database.FlagKey{
			UUID:        flag_key_uuid,
			FlagType:    flagType,
			Key:         key,
			Enabled:     false,
			LastUpdated: &now,
		}
		result = db.Create(newFlag)
		if result.Error != nil {
			return result.Error
		}
		for _, value := range variations {
			variationUUID := datatypes.NewUUIDv4()
			variation := database.FlagVariation[T]{
				UUID:        variationUUID,
				FlagKeyUUID: newFlag.UUID,
				Value:       value.Value.(T),
				Name:        value.Name,
				LastUpdated: &now,
			}
			result = db.Scopes(database.GetTableName(variation)).Create(variation)
			if result.Error != nil {
				return result.Error
			}
			uuids = append(uuids, variationUUID)
		}
		lenVariations := len(uuids)
		newFlag.DefaultEnabledVariation = uuids[0]
		newFlag.DefaultVariation = uuids[lenVariations-1]
		db.Save(newFlag)
	}
	return nil
}

func GetFlag[T comparable](key string) (T, error) {
	db := database.GetDB()
	var flagKey database.FlagKey
	var returnVal T

	result := db.First(&flagKey, "key = ?", key)
	if result.RowsAffected != 0 {
		currentVariation := flagKey.DefaultVariation
		if flagKey.Enabled {
			currentVariation = flagKey.DefaultEnabledVariation
		}
		return GetFlagVariation[T](currentVariation)
	}

	return returnVal, result.Error
}

func GetFlagVariation[T comparable](variationUUID datatypes.UUID) (T, error) {
	db := database.GetDB()
	var flagVariation database.FlagVariation[T]
	scope := db.Scopes(database.GetTableName(flagVariation))

	var returnVal T
	result := scope.First(&flagVariation, "uuid = ?", variationUUID)
	if result.RowsAffected != 0 {
		returnVal = flagVariation.Value
		return returnVal, nil
	}
	return returnVal, result.Error
}

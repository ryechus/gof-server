package repositories

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RuleRepository struct {
	DB *gorm.DB
}

var _ Repository = &RuleRepository{}

func (rr *RuleRepository) GetTargetingRules(flagKeyUUID datatypes.UUID) ([]database.TargetingRule, *gorm.DB) {
	db := rr.DB
	var flagRules []database.TargetingRule
	result := db.Find(&flagRules, "flag_key_uuid = ?", flagKeyUUID.String())

	return flagRules, result
}

func (rr *RuleRepository) GetTargetingRuleContexts(targetingRuleUUID datatypes.UUID) ([]database.TargetingRuleContext, *gorm.DB) {
	db := rr.DB
	var flagRuleContexts []database.TargetingRuleContext
	result := db.Find(&flagRuleContexts, "targeting_rule_uuid = ?", targetingRuleUUID.String())

	return flagRuleContexts, result
}

func (rr *RuleRepository) SaveTargetingRule(payload payloads.PutRule) error {
	db := rr.DB
	variationUUID := datatypes.UUID(uuid.MustParse(payload.VariationUUID))
	flagKeyUUID := datatypes.UUID(uuid.MustParse(payload.FlagUUID))
	jsonRuleContexts, err := json.Marshal(payload.RuleContexts)
	if err != nil {
		return err
	}

	ruleUUID := datatypes.NewUUIDv4()
	if payload.UUID != "" {
		ruleUUID = datatypes.UUID(uuid.MustParse(payload.UUID))
	}

	rule := database.TargetingRule{
		UUID:          ruleUUID,
		Name:          payload.Name,
		FlagKeyUUID:   flagKeyUUID,
		VariationUUID: variationUUID,
		Attributes:    datatypes.JSON([]byte(jsonRuleContexts)),
	}

	db.Save(rule)
	return nil
}

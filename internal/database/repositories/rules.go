package repositories

import (
	"encoding/json"
	"log"

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
	result := db.Raw("SELECT uuid, name, flag_key_uuid, variation_uuid, attributes FROM targeting_rules WHERE flag_key_uuid = ? ORDER BY priority ASC", flagKeyUUID.String()).Scan(&flagRules)

	return flagRules, result
}

// func (rr *RuleRepository) GetTargetingRuleContexts(targetingRuleUUID datatypes.UUID) ([]database.TargetingRuleContext, *gorm.DB) {
// 	db := rr.DB
// 	var flagRuleContexts []database.TargetingRuleContext
// 	result := db.Find(&flagRuleContexts, "targeting_rule_uuid = ?", targetingRuleUUID.String())

// 	return flagRuleContexts, result
// }

func (rr *RuleRepository) SaveTargetingRule(payload payloads.PutRule, tx *gorm.DB) error {
	db := tx
	variationUUID := datatypes.UUID(uuid.MustParse(payload.VariationUUID))
	flagKeyUUID := datatypes.UUID(uuid.MustParse(payload.FlagUUID))
	jsonRuleContexts, err := json.Marshal(payload.RuleContexts)
	if err != nil {
		return err
	}

	ruleUUID := datatypes.NewUUIDv4()
	operation := "insert"
	if payload.UUID != "" {
		ruleUUID = datatypes.UUID(uuid.MustParse(payload.UUID))
		operation = "update"
	}

	rule := database.TargetingRule{
		UUID:          ruleUUID,
		Name:          payload.Name,
		FlagKeyUUID:   flagKeyUUID,
		VariationUUID: variationUUID,
		Priority:      payload.Priority,
		Attributes:    datatypes.JSON([]byte(jsonRuleContexts)),
	}

	var result *gorm.DB
	if operation == "update" {
		query := `
		UPDATE targeting_rules 
		SET name = ?, variation_uuid = ?, attributes = ?, priority = ? 
		WHERE uuid = ?
		`
		log.Println("updating rule")
		result = db.Raw(query, rule.Name, rule.VariationUUID, rule.Attributes, rule.Priority, payload.UUID).Scan(&rule)
	} else {
		query := "INSERT INTO targeting_rules (uuid, name, flag_key_uuid, variation_uuid, attributes, priority) VALUES (?, ?, ?, ?, ?, ?)"
		log.Println("inserting rule")
		result = db.Raw(query, rule.UUID, rule.Name, rule.FlagKeyUUID, rule.VariationUUID, rule.Attributes, rule.Priority).Scan(&rule)
	}

	if result.Error != nil {
		db.Rollback()
		return result.Error
	}
	return nil
}

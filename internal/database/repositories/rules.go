package repositories

import (
	"fmt"

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

func (rr *RuleRepository) SaveTargetingRule(payload payloads.PutRule) {
	db := rr.DB
	variationUUID := datatypes.UUID(uuid.MustParse(payload.VariationUUID))
	flagKeyUUID := datatypes.UUID(uuid.MustParse(payload.FlagUUID))

	rule := database.TargetingRule{
		UUID:          datatypes.NewUUIDv4(),
		Name:          payload.Name,
		FlagKeyUUID:   flagKeyUUID,
		VariationUUID: variationUUID,
	}

	db.Save(rule)
	for _, ctx := range payload.RuleContexts {
		fmt.Println("saving contexts")
		for _, value := range ctx.Values {
			fmt.Println("saving values")
			targetingContext := database.TargetingRuleContext{
				UUID:              datatypes.NewUUIDv4(),
				TargetingRuleUUID: rule.UUID,
				ContextKind:       ctx.ContextKind,
				Attribute:         ctx.Attribute,
				Value:             value,
			}
			db.Save(targetingContext)
		}
	}
}

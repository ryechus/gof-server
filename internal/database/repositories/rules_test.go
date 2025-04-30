package repositories_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/placer14/gof-server/internal/database/repositories"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
)

func TestGetTargetingRules(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagKeyUUID := datatypes.NewUUIDv4()
	columns := []string{}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "targeting_rules" WHERE flag_key_uuid = $1`)).
		WithArgs(flagKeyUUID.String()).
		WillReturnRows(sqlmock.NewRows(columns))

	flagRepository := repositories.RuleRepository{DB: gormDB}

	flagRepository.GetTargetingRules(flagKeyUUID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTargetingRuleContexts(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	targetingRuleUUID := datatypes.NewUUIDv4()
	columns := []string{}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "targeting_rule_contexts" WHERE targeting_rule_uuid = $1`)).
		WithArgs(targetingRuleUUID.String()).
		WillReturnRows(sqlmock.NewRows(columns))

	flagRepository := repositories.RuleRepository{DB: gormDB}

	flagRepository.GetTargetingRuleContexts(targetingRuleUUID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveTargetingRule(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagKeyUUID := datatypes.NewUUIDv4()
	ruleContexts := []payloads.RuleContext{
		{
			ContextKind: "user",
			Attribute:   "email",
			Values:      []string{"example@gmail.com", "example2@example.com"},
		},
	}
	payload := payloads.PutRule{
		FlagUUID:      flagKeyUUID.String(),
		Name:          "test rule",
		VariationUUID: flagKeyUUID.String(),
		RuleContexts:  ruleContexts,
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "targeting_rules" SET "uuid"=$1,"name"=$2,"flag_key_uuid"=$3,"variation_uuid"=$4 WHERE "uuid" = $5`)).
		WithArgs(sqlmock.AnyArg(), payload.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "targeting_rule_contexts" SET "uuid"=$1,"targeting_rule_uuid"=$2,"context_kind"=$3,"attribute"=$4,"operator"=$5,"value"=$6 WHERE "uuid" = $7`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "user", "email", "", "example@gmail.com", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "targeting_rule_contexts" SET "uuid"=$1,"targeting_rule_uuid"=$2,"context_kind"=$3,"attribute"=$4,"operator"=$5,"value"=$6 WHERE "uuid" = $7`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "user", "email", "", "example2@example.com", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	flagRepository := repositories.RuleRepository{DB: gormDB}

	flagRepository.SaveTargetingRule(payload)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

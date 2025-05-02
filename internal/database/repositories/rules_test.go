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
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT uuid, name, flag_key_uuid, variation_uuid, attributes FROM targeting_rules WHERE flag_key_uuid = $1`)).
		WithArgs(flagKeyUUID.String()).
		WillReturnRows(sqlmock.NewRows(columns))

	flagRepository := repositories.RuleRepository{DB: gormDB}

	flagRepository.GetTargetingRules(flagKeyUUID)

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
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO targeting_rules (uuid, name, flag_key_uuid, variation_uuid, attributes) VALUES ($1, $2, $3, $4, $5)")).
		WithArgs(sqlmock.AnyArg(), payload.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{}))
	flagRepository := repositories.RuleRepository{DB: gormDB}

	gormTx := gormDB.Begin()
	flagRepository.SaveTargetingRule(payload, gormTx)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

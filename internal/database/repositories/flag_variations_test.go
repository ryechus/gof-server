package repositories_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/database/repositories"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
)

func TestGetFlagKeyVariationByUUID(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagUUID := datatypes.NewUUIDv4()
	columns := []string{}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "flag_key_bool_variations" WHERE uuid = $1 ORDER BY "flag_key_bool_variations"."uuid" LIMIT $2`)).
		WithArgs(flagUUID.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow())
	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}

	boolFlagRepository.GetFlagKeyVariationByUUID(flagUUID)
}

func TestCreateFlagKeyVariation(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	now := time.Now()
	variationName := "on"
	flagKey := database.FlagKey{
		UUID:                    datatypes.NewUUIDv4(),
		Key:                     "test",
		FlagType:                "bool",
		DefaultVariation:        datatypes.NewUUIDv4(),
		DefaultEnabledVariation: datatypes.NewUUIDv4(),
		Enabled:                 false,
		LastUpdated:             &now,
	}
	flagVariation := payloads.FlagVariation{
		Name:  variationName,
		Value: true,
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "flag_key_bool_variations" ("uuid","flag_key_uuid","name","value","last_updated") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), variationName, true, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}

	boolFlagRepository.CreateFlagKeyVariation(flagKey, flagVariation)
}

func TestGetFlagVariationValue(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagUUID := datatypes.NewUUIDv4()
	columns := []string{}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "flag_key_bool_variations" WHERE uuid = $1 ORDER BY "flag_key_bool_variations"."uuid" LIMIT $2`)).
		WithArgs(flagUUID.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow())
	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}

	boolFlagRepository.GetFlagVariationValue(flagUUID)
}

package repositories_test

import (
	"regexp"
	"testing"

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
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT uuid, flag_key_uuid, name, value, last_updated FROM flag_key_bool_variations WHERE uuid = $1`)).
		WithArgs(flagUUID.String()).
		WillReturnRows(sqlmock.NewRows(columns))
	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}

	boolFlagRepository.GetFlagKeyVariationByUUID(flagUUID)
}

func TestCreateFlagKeyVariation(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	variationName := "on"
	flagKey := database.FlagKey{
		UUID:                    datatypes.NewUUIDv4(),
		Key:                     "test",
		FlagType:                "bool",
		DefaultVariation:        datatypes.NewUUIDv4(),
		DefaultEnabledVariation: datatypes.NewUUIDv4(),
		Enabled:                 false,
	}
	flagVariation := payloads.FlagVariation{
		Name:  variationName,
		Value: true,
	}
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO flag_key_bool_variations (uuid, flag_key_uuid, value, name) VALUES ($1, $2, $3, $4)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), true, variationName).
		WillReturnRows(sqlmock.NewRows([]string{}))

	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}
	gormTx := gormDB.Begin()
	boolFlagRepository.CreateFlagKeyVariation(flagKey, flagVariation, gormTx)
}

func TestGetFlagVariationValue(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagUUID := datatypes.NewUUIDv4()
	columns := []string{}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT uuid, flag_key_uuid, name, value, last_updated FROM flag_key_bool_variations WHERE uuid = $1`)).
		WithArgs(flagUUID.String()).
		WillReturnRows(sqlmock.NewRows(columns).AddRow())
	boolFlagRepository := repositories.FlagVariationRepository[bool]{DB: gormDB}

	boolFlagRepository.GetFlagVariationValue(flagUUID)
}

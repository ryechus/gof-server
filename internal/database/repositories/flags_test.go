package repositories_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/database/repositories"
	"gorm.io/datatypes"
)

func TestGetFlagKey(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagKeyString := "test"
	columns := []string{"uuid", "key", "flag_type", "default_variation", "default_enabled_variation", "enabled", "last_updated"}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT uuid, key, name, description, flag_type, default_variation, default_enabled_variation, enabled, created_at, updated_at FROM flag_keys WHERE key = $1`)).
		WithArgs(flagKeyString).
		WillReturnRows(sqlmock.NewRows(columns))

	flagRepository := repositories.FlagRepository{DB: gormDB}

	flagRepository.GetFlagKey(flagKeyString)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFlagKeyByUUID(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagUUID := datatypes.NewUUIDv4().String()
	columns := []string{"uuid", "key", "flag_type", "default_variation", "default_enabled_variation", "enabled", "last_updated"}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "flag_keys" WHERE uuid = $1 ORDER BY "flag_keys"."uuid" LIMIT $2`)).
		WithArgs(flagUUID, 1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(
			flagUUID,
			"test",
			"string",
			datatypes.NewUUIDv4().String(),
			datatypes.NewUUIDv4().String(),
			false,
			time.Now()))

	flagRepository := repositories.FlagRepository{DB: gormDB}

	flagRepository.GetFlagKeyByUUID(flagUUID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateFlagKey(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO flag_keys (uuid, name, flag_type, key, enabled) VALUES ($1, $2, $3, $4, $5)")).
		WithArgs(sqlmock.AnyArg(), nil, "bool", "test", false).
		WillReturnRows(sqlmock.NewRows([]string{}))

	flagRepository := repositories.FlagRepository{DB: gormDB}

	gormTx := gormDB.Begin()
	flagRepository.CreateFlagKey("bool", "test", gormTx)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateFlagKey(t *testing.T) {
	gormDB, mock, db := getMockedDB(t)
	defer db.Close()

	flagKey := database.FlagKey{
		UUID:                    datatypes.NewUUIDv4(),
		Key:                     "test",
		FlagType:                "string",
		DefaultVariation:        datatypes.NewUUIDv4(),
		DefaultEnabledVariation: datatypes.NewUUIDv4(),
		Enabled:                 false,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("UPDATE flag_keys SET name=$1, description=$2 , default_variation=$3, default_enabled_variation=$4, enabled=$5, updated_at=$6 WHERE uuid = $7")).
		WithArgs(nil, nil, flagKey.DefaultVariation, flagKey.DefaultEnabledVariation, flagKey.Enabled, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{}))

	flagRepository := repositories.FlagRepository{DB: gormDB}

	gormTx := gormDB.Begin()
	flagRepository.UpdateFlagKey(&flagKey, gormTx)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

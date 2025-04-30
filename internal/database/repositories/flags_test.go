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
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "flag_keys" WHERE key = $1 ORDER BY "flag_keys"."uuid" LIMIT $2`)).
		WithArgs(flagKeyString, 1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(
			datatypes.NewUUIDv4().String(),
			"bert",
			"string",
			datatypes.NewUUIDv4().String(),
			datatypes.NewUUIDv4().String(),
			false,
			time.Now()))

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
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "flag_keys" ("uuid","key","flag_type","default_variation","default_enabled_variation","enabled","last_updated") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), "test", "bool", sqlmock.AnyArg(), sqlmock.AnyArg(), false, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	flagRepository := repositories.FlagRepository{DB: gormDB}

	flagRepository.CreateFlagKey("bool", "test")

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
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "flag_keys" SET "key"=$1,"flag_type"=$2,"default_variation"=$3,"default_enabled_variation"=$4,"enabled"=$5,"last_updated"=$6 WHERE "uuid" = $7`)).
		WithArgs(flagKey.Key, flagKey.FlagType, flagKey.DefaultVariation, flagKey.DefaultEnabledVariation, flagKey.Enabled, flagKey.LastUpdated, flagKey.UUID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	flagRepository := repositories.FlagRepository{DB: gormDB}

	flagRepository.UpdateFlagKey(&flagKey)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

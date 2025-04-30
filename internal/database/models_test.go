package database_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/placer14/gof-server/internal/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getMockedDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, _ := gorm.Open(dialector)

	return gormDB, mock, db
}

func TestGetTableName(t *testing.T) {
	gormDB, _, db := getMockedDB(t)
	defer db.Close()

	t.Run("float_table", func(t *testing.T) {
		table := database.GetTableName(database.FlagVariation[float64]{})(gormDB)

		assert.Equal(t, "flag_key_float_variations", table.Statement.Table)
	})
	t.Run("int_table", func(t *testing.T) {
		table := database.GetTableName(database.FlagVariation[float64]{})(gormDB)

		assert.Equal(t, "flag_key_float_variations", table.Statement.Table)
	})
	t.Run("bool_table", func(t *testing.T) {
		table := database.GetTableName(database.FlagVariation[bool]{})(gormDB)

		assert.Equal(t, "flag_key_bool_variations", table.Statement.Table)
	})
	t.Run("string_table", func(t *testing.T) {
		table := database.GetTableName(database.FlagVariation[string]{})(gormDB)

		assert.Equal(t, "flag_key_string_variations", table.Statement.Table)
	})
}

func TestMigrateDB(t *testing.T) {}

package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/sqls"
)

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(sqls.Pragma + sqls.Schema + MockSchema)
	assert.NoError(t, err)
}

func TestDB(t *testing.T) {
	// success
	db := DB()
	assert.NotNil(t, db)
}

func TestString(t *testing.T) {
	// setup
	db := DB()

	// success
	data := String(db, "select name from Notes where n_id=?", 1)
	assert.Equal(t, "alpha", data)
}

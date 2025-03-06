// Package test implements unit-testing helper functions and mock data.
package test

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/lethe/lethe/tools/sqls"
)

// MockSchema is the default additional unit-testing database schema.
const MockSchema = `
	insert into Notes (name) values ('alpha'), ('bravo');
	insert into Pages (note, body) values
		(1, 'Alpha note (old).' || char(10)),
		(1, 'Alpha note.'       || char(10)),
		(2, 'Bravo note.'       || char(10));
`

// DB returns an in-memory database with real and mock schema.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.Pragma + sqls.Schema + MockSchema)
	return db
}

// String returns a string from a database query, panicking on error.
func String(db *sqlx.DB, code string, elems ...any) string {
	var data string
	if err := db.Get(&data, code, elems...); err != nil {
		panic(err)
	}

	return data
}

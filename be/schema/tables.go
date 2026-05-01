package schema

import (
	"github.com/jmoiron/sqlx"
)

func CreateTables(db *sqlx.DB) error {

	db.MustExec("PRAGMA journal_mode=WAL;")
	db.MustExec("PRAGMA busy_timeout = 5000;")
	db.MustExec("PRAGMA foreign_keys = ON;")

	schema := `
	CREATE TABLE IF NOT EXISTS room (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    connection_key VARCHAR(255) DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
	`
	db.MustExec(schema)
	return nil
}

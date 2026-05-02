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
			name TEXT NOT NULL,
			connection_key TEXT DEFAULT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS subjects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			room_id INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (room_id) REFERENCES room(id)
		);`
	db.MustExec(schema)
	return nil
}

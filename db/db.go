package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the database connection and creates the books table if it does not exist.
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite", "./books.db")
	if err != nil {
		return err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		release_date TEXT NOT NULL,
		score REAL NOT NULL,
		summary TEXT
	);`
	_, err = DB.Exec(createTableSQL)
	return err
}

func InitTestDB(path string) error {
	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		release_date TEXT NOT NULL,
		score REAL NOT NULL,
		summary TEXT
	);`
	_, err = DB.Exec(createTableSQL)
	return err
}

package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

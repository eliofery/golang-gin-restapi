package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Не удалось подключиться к БД: " + err.Error())
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Не удалось создать таблицу событий: " + err.Error())
	}
}

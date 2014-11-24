package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func OpenDatabase() {
	connection := os.Getenv("INV_API_DB_CONNECTION_STRING")
	database, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err.Error()) // TODO actually handle
	} else {
		db = database
	}
}

func CloseDatabase() {
	db.Close()
}

func Get() *sql.DB {
	return db
}

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const mysqlUser string = ""
const mysqlPass string = ""
const mysqlDatabase string = ""

func openDatabase() {
	database, err := sql.Open("mysql", mysqlUser+":"+mysqlPass+"@/"+mysqlDatabase)
	if err != nil {
		panic(err.Error()) // TODO actually handle
	} else {
		db = database
	}
}

func closeDatabase() {
	db.Close()
}

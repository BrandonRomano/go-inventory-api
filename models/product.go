package models

import (
	"encoding/json"
	"github.com/brandonromano/inventory_api/database"
	_ "github.com/go-sql-driver/mysql"
)

const TABLE_NAME = "products"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float32
}

func (p *Product) Load(id string) {
	// Getting the DB
	db := database.Get()
	err := db.QueryRow("SELECT * FROM "+TABLE_NAME+" WHERE id = ?", id).
		Scan(&p.Id, &p.Name, &p.Description, &p.Price)

	// There was nothing with that ID
	if err != nil {
		panic(err) // TODO handle
	}
}

func (p *Product) Store() {
	db := database.Get()

	// Creating prepared statement
	stmt, err := db.Prepare("INSERT INTO " + TABLE_NAME +
		"(name,description,price) VALUES(?,?,?)")
	if err != nil {
		panic(err) // Failed to create prepared statement
	}

	// Executing
	_, response_err := stmt.Exec(&p.Name, &p.Description, &p.Price)
	if response_err != nil {
		panic(response_err) // Failed to insert
	}
}

func (p *Product) ToJSON() string {
	json, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(json)
}

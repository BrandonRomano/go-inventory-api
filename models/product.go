package models

import (
	"github.com/brandonromano/inventory_api/database"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

const TABLE_NAME = "products"

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (p *Product) Load(id string) error {
	// Getting the DB
	db := database.Get()
	err := db.QueryRow("SELECT * FROM "+TABLE_NAME+" WHERE id = ?", id).
		Scan(&p.Id, &p.Name, &p.Description, &p.Price)

	// There was nothing with that ID
	if err != nil {
		return err
	}

	// No error
	return nil
}

func (p *Product) Store() (string, error) {
	db := database.Get()

	// Creating prepared statement
	stmt, err := db.Prepare("INSERT INTO " + TABLE_NAME +
		"(name,description,price) VALUES(?,?,?)")
	if err != nil {
		return "", err
	}

	// Executing
	result, response_err := stmt.Exec(&p.Name, &p.Description, &p.Price)
	if response_err != nil {
		return "", response_err
	}

	id, insert_error := result.LastInsertId()
	if insert_error != nil {
		return "", insert_error
	}

	// No error
	idString := strconv.FormatInt(id, 10)
	return idString, nil
}

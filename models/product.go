package models

import (
	"encoding/json"
	"github.com/brandonromano/inventory_api/database"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float32
}

func (p *Product) Load(id string) {
	// Getting the DB
	db := database.Get()
	err := db.QueryRow("SELECT * FROM product WHERE id = ?", id).
		Scan(&p.Id, &p.Name, &p.Description, &p.Price)

	// There was nothing with that ID
	if err != nil {
		panic(err) // TODO handle
	}
}

func (p *Product) Store() {
	// TODO Store p
}

func (p *Product) ToJSON() string {
	json, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(json)
}

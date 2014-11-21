package models

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name        string
	Description string
	Price       float32
}

func (p *Product) Load(id string) {
	p.Name = "MacBook Air"
	p.Description = "This is a MacBook."
	p.Price = 1199.99
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

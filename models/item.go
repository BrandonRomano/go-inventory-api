package models

const ITEM_TABLE_NAME = "items"

type Item struct {
	Id     int64   `json:"id"`
	UserId int64   `json:"user_id"`
	Name   string  `json:"name"`
	Photo  string  `json:"photo"`
	Price  float64 `json:"price"`
	Notes  string  `json:"notes"`
}

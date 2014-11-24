package models

import (
	"github.com/brandonromano/inventory_api/database"
	_ "github.com/go-sql-driver/mysql"
)

const USER_TABLE_NAME = "users"

type User struct {
	Id       int64
	Username string
	Salt     string
	Hash     string
}

func (user *User) AuthenticateUser(auth Auth) bool {
	// Real Hash
	realHash := ""

	// Getting the real hash from the DB
	db := database.Get()
	err := db.QueryRow("SELECT hash FROM"+USER_TABLE_NAME+" WHERE username = ", auth.Username).
		Scan(realHash)

	// No user with that name
	if err != nil {
		return false
	}

	// Checking if hashes exist
	if realHash != auth.Payload {
		return false
	}

	// Successful
	return true
}

// ===== Auth

type Auth struct {
	Username string `json:"username"`
	Payload  string `json:"payload"` // Outgoing this is the salt, incoming this is the hash
}

func (auth *Auth) RequestAuthentication(username string) error {
	// Getting the Auth from the DB
	db := database.Get()
	err := db.QueryRow("SELECT username, salt FROM "+USER_TABLE_NAME+" WHERE username = ?", username).
		Scan(&auth.Username, &auth.Payload)

	// No user with that username
	if err != nil {
		return err
	}

	// No error
	return nil
}

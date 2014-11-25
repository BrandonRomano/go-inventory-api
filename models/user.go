package models

import (
	"errors"
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

/**
 * This function loads the user from the database from
 * either the id or the username field
 */
func (user *User) Load() error {
	// Getting the DB
	db := database.Get()

	// Trying against the id
	if user.Id != 0 {
		return db.QueryRow("SELECT * FROM "+USER_TABLE_NAME+" WHERE id = ?", user.Id).
			Scan(&user.Id, &user.Username, &user.Salt, &user.Hash)
	}

	// Trying against the username
	if user.Username != "" {
		return db.QueryRow("SELECT * FROM "+USER_TABLE_NAME+" WHERE username = ?", user.Username).
			Scan(&user.Id, &user.Username, &user.Salt, &user.Hash)
	}

	// Custom error for invalid user state
	err := errors.New("A unique key has to be set before running user.Load()")
	return err
}

func (user *User) AuthenticateUser(auth Auth) bool {
	// Real Hash
	var realHash string

	// Getting the real hash from the DB
	db := database.Get()
	err := db.QueryRow("SELECT hash FROM "+USER_TABLE_NAME+" WHERE username = ?", auth.Username).
		Scan(&realHash)

	// No user with that name
	if err != nil {
		return false
	}

	// Checking if hash is correct
	if auth.Payload != realHash {
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

func (auth *Auth) PrepareAuth(user User) {
	auth.Username = user.Username
	auth.Payload = user.Salt
}

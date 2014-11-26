package models

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/brandonromano/inventory_api/database"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"time"
)

const SESSION_TABLE_NAME = "sessions"
const SESSION_LENGTH_HOURS = 2
const SESSION_TOKEN_LENGTH = 32

type Session struct {
	Id         int64 `json:"-"`
	UserId     int64 `json:"-"`
	Token      string
	Expiration int64
}

func (s *Session) Store() error {
	db := database.Get()

	// Creating prepared statement
	stmt, err := db.Prepare("INSERT INTO " + SESSION_TABLE_NAME +
		"(user_id, token, expiration) VALUES(?,?,?)")
	if err != nil {
		return err
	}

	// Executing
	result, response_err := stmt.Exec(&s.UserId, &s.Token, &s.Expiration)
	if response_err != nil {
		return response_err
	}

	// Getting ID
	id, insert_error := result.LastInsertId()
	if insert_error != nil {
		return insert_error
	}

	// Setting ID
	s.Id = id

	// No error
	return nil
}

func GetExpiration() int64 {
	currentTime := time.Now().Unix()
	offset := int64(SESSION_LENGTH_HOURS * 60 * 60)
	return currentTime + offset
}

func GenerateToken() string {
	token_bytes := make([]byte, SESSION_TOKEN_LENGTH)
	_, err := io.ReadFull(rand.Reader, token_bytes)
	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(token_bytes)
}

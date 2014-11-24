package models

const SESSION_TABLE_NAME = "sessions"

type Session struct {
	Id         int64
	UserId     string
	Token      string
	Expiration int32
}

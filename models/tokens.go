package models

import "time"

type Token struct {
	id        int
	userId    int
	token     string
	expired   time.Time
	loginAt   time.Time
	ip        string
	userAgent string
}

func (m *Token) TableName() string {
	return "token"
}

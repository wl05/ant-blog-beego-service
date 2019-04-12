package models

type User struct {
	id       int
	username string
	password string
}

func (m *User) TableName() string {
	return TableName("user")
}

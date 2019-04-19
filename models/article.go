package models

import "time"

type Article struct {
	Id        int
	Title     string
	UserId    int
	Content   string
	PublishAt time.Time
	ViewCount int
}

func (m *Article) TableName() string {
	return "article"
}

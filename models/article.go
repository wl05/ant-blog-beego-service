package models

import "time"

type Article struct {
	id        int
	title     string
	userId    int
	content   string
	publishAt time.Time
	viewCount int
}

func (m *Article) TableName() string {
	return "article"
}

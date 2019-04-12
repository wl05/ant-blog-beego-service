package models

type Tag struct {
	id        int
	tagName   string
	articleId int
}

func (m *Tag) TableName() string {
	return TableName("tag")
}

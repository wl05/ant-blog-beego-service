package models

type Tag struct {
	Id        int
	TagName   string
	ArticleId int64
}

func (m *Tag) TableName() string {
	return "tag"
}

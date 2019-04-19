package models

type Category struct {
	Id           int
	CategoryName string
	ArticleId    int
}

func (m *Category) TableName() string {
	return "category"
}

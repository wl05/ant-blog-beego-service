package models

type Category struct {
	id           int
	categoryName string
	articleId    int
}

func (m *Category) TableName() string {
	return "category"
}

package models

type Category struct {
	Id           int    `orm:"column(id);auto" description:"数据库id" json:"id"`
	CategoryName string `orm:"column(categoryName)" description:"categoryName" json:"categoryName"`
}

func (m *Category) TableName() string {
	return "category"
}

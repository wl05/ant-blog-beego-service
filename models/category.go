package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id           int    `orm:"column(id);auto" description:"数据库id" json:"id"`
	CategoryName string `orm:"column(categoryName)" description:"categoryName" json:"categoryName"`
}

func (m *Category) TableName() string {
	return "category"
}
func AddCategory(c string) (id int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var category Category
	category.CategoryName = c
	_, err = o.Insert(&category)
	return
}

func GetCategories() (categories []Category, err error) {
	var _categories []Category
	o := orm.NewOrm()
	_, _err := o.QueryTable("category").All(&_categories, "Id", "CategoryName")
	return _categories, _err
}

func GetCategoryById(id int) (category Category, err error) {
	var _category Category
	o := orm.NewOrm()
	act := o.Raw("select * from category where id = ?", id)
	_err := act.QueryRow(&_category)
	return _category, _err
}

func GetCategoryByCategoryName(categoryName string) (category Category, err error) {
	DB := orm.NewOrm()
	DB.Using("default")
	act := DB.Raw("SELECT * FROM category WHERE categoryName = ?", categoryName)
	var _category Category
	_err := act.QueryRow(&_category)
	return _category, _err
}
func UpdateCategoryByCategoryId(id int, categoryName string) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("UPDATE category SET categoryName = ? WHERE id = ?", categoryName, id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}

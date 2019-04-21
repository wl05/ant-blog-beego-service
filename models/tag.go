package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id      int    `orm:"column(id);auto" description:"数据库id" json:"id"`
	TagName string `orm:"column(tagName)" description:"tagName" json:"tagName"`
}

func (m *Tag) TableName() string {
	return "tag"
}

func AddTag(t string) (id int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var tag Tag
	tag.TagName = t
	_, err = o.Insert(&tag)
	return
}

func GetTags() (tags []Tag, err error) {
	var _tags []Tag
	o := orm.NewOrm()
	_, _err := o.QueryTable("tag").All(&_tags, "Id", "TagName")
	return _tags, _err
}

func GetTagByTagName(tagName string) (tag Tag, err error) {
	DB := orm.NewOrm()
	var r0 orm.RawSeter
	DB.Using("default")
	r0 = DB.Raw("SELECT * FROM tag WHERE tagName = ?", tagName)
	var _tag Tag
	_err := r0.QueryRow(&_tag)
	return _tag, _err
}
func UpdateTagByTagId(id int, tagName string) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("UPDATE tag SET tagName = ? where id = ?", tagName, id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}

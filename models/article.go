package models

import "time"

type Article struct {
	Id         int       `orm:"column(id);auto" description:"数据库id" json:"id"`
	Title      string    `orm:"column(title)" description:"title" json:"title"`
	UserId     int       `orm:"column(userId)" description:"用户id" json:"userId"`
	Content    string    `orm:"column(content)" description:"content" json:"content"`
	PublishAt  time.Time `orm:"column(publishAt)" description:"publishAt" json:"publishAt"`
	ViewCount  int       `orm:"column(viewCount)" description:"viewCount" json:"viewCount"`
	CategoryId int       `orm:"column(categoryId)" description:"分类id" json:"categoryId"`
	TagId      int       `orm:"column(tagId)" description:"标签id" json:"tagId"`
}

func (m *Article) TableName() string {
	return "article"
}

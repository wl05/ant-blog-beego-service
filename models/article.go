package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

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

func CreateArticle(title string, userId int, content string, publishAt time.Time, categoryId int, tagId int) (id int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var article Article
	article.Title = title
	article.UserId = userId
	article.Content = content
	article.PublishAt = publishAt
	article.ViewCount = 0
	article.CategoryId = categoryId
	article.TagId = tagId
	_, err = o.Insert(&article)
	return
}

func GetArticles() (article []Article, err error) {
	var articles []Article
	o := orm.NewOrm()
	_, _err := o.QueryTable("article").All(&articles)
	return articles, _err
}

func GetArticleById(id int) (article Article, err error) {
	var _article Article
	o := orm.NewOrm()
	act := o.Raw("select * from article where id = ?", id)
	_err := act.QueryRow(&_article)
	return _article, _err
}

func DeleteArticleById(id int) (num int64, err error) {
	DB := orm.NewOrm()
	DB.Using("default")
	act := DB.Raw("DELETE FROM article WHERE id = ?", id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}

func UpdateArticleByArticleId(id int, title string, userId int, content string, publishAt time.Time, categoryId int, tagId int) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("UPDATE article SET title = ?,userId = ?,content = ?,publishAt = ?,categoryId = ? ,tagId = ? where id = ?", title, userId, content, publishAt, categoryId, tagId,
		id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}

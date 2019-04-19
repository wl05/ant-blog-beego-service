package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"column(id);auto" description:"数据库id" json:"id"`
	Username string `orm:"column(username);size(20)" description:"用户名" json:"username"`
	Password string `orm:"column(password);size(32)" description:"密码" json:"password"`
}

func (m *User) TableName() string {
	return "user"
}

func AddUser(username string, password string) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var user User
	user.Username = username
	user.Password = password
	id, err = o.Insert(&user)
	return
}

func GetUser() (user []User, err error) {
	var users []User
	o := orm.NewOrm()
	_, _err := o.QueryTable("user").All(&users, "Id", "Username")
	return users, _err
}

func GetUserByUsername(username string) (user User, err error) {
	DB := orm.NewOrm()
	var r0 orm.RawSeter
	DB.Using("default")
	r0 = DB.Raw("SELECT * FROM user WHERE username = ?", username)
	var _user User
	_err := r0.QueryRow(&_user)
	return _user, _err
}
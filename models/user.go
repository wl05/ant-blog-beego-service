package models

import (
	"fmt"
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

func AddUser(m *User) (id int64, err error) {
	fmt.Println(m)
	o := orm.NewOrm()
	o.Using("default")
	id, err = o.Insert(m)
	return
}

func GetUser() []User {
	var _user []User
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw("SELECT * FROM user").QueryRow(&_user)
	if err == nil {
		fmt.Println("出错了")
		return _user
	} else {
		return _user
	}
}

func GetUserByUsername(username string) User {
	var _user User
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw("SELECT * FROM user WHERE username = ?", username).QueryRow(&_user)
	if err == nil {
		fmt.Println("出错了")
		return _user
	} else {
		return _user
	}
}

package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"column(id);auto" description:"数据库id" json:"id"`
	Username string `orm:"column(username);size(20)" description:"用户名" json:"username"`
	Password string `orm:"column(password);size(32)" description:"密码" json:"password"`
}

func (m *User) TableName() string {
	return "user"
}

func AddUser(username string, password string) (id int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var user User
	user.Username = username
	user.Password = password
	_, err = o.Insert(&user)
	return
}

func GetUsers() (user []User, err error) {
	var users []User
	o := orm.NewOrm()
	_, _err := o.QueryTable("user").All(&users, "Id", "Username")
	return users, _err
}

func GetUserById(id int) (user User, err error) {
	var _user User
	o := orm.NewOrm()
	act := o.Raw("select * from user where id = ?", id)
	_err := act.QueryRow(&_user)
	return _user, _err
}

func GetUserByUsername(username string) (user User, err error) {
	DB := orm.NewOrm()
	DB.Using("default")
	act := DB.Raw("SELECT * FROM user WHERE username = ?", username)
	var _user User
	_err := act.QueryRow(&_user)
	return _user, _err
}
func UpdateTagByUserId(id int, username string, password string) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("UPDATE user SET username = ?,password = ? where id = ?", username, password, id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}

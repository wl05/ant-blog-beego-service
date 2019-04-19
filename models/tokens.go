package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ant-blog-beego-service/common/utils"
)

type Token struct {
	Id        int       `orm:"column(id);auto" description:"数据库id" json:"id"`
	UserId    int       `orm:"column(userId)" description:"用户id" json:"userId"`
	Token     string    `orm:"column(token)" description:"token" json:"token"`
	Expired   time.Time `orm:"column(expired)" description:"过期时间" json:"expired"`
	LoginAt   time.Time `orm:"column(loginAt)" description:"登录时间" json:"loginAt"`
	Ip        string    `orm:"column(ip)" description:"ip地址" json:"ip"`
	UserAgent string    `orm:"column(userAgent)" description:"用户代理" json:"userAgent"`
}

func (m *Token) TableName() string {
	return "tokens"
}

func CreateToken(UserId int, Ip string, UserAgent string) (t string, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var token Token
	token.UserId = UserId
	token.Token = utils.RandSeq(20)
	d, err := time.ParseDuration("1d")
	token.Expired = time.Now().Add(d)
	token.LoginAt = time.Now()
	token.Ip = Ip
	token.UserAgent = UserAgent
	_, err = o.Insert(&token)
	return token.Token, err
}

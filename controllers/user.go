package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"strings"
)

// Operations about User
type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("CreateUser", c.CreateUser)
}

// @Title 创建用户
// @Description 创建用户
// @Param   username    formData    string  "ant"  true    "用户名"
// @Param   password    formData    string  "123"    true    "密码"
// @Success 200     {int}       models.User.id models.User.username
// @Failure 403     body is empty
// @router / [post]
func (this *UserController) CreateUser() {
	username := this.GetString("username")
	password := this.GetString("password")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE__SOURCE_DATA__ILLEGAL,
			"msg":  consts.ERROR_DES__SOURCE_DATA__ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	user := &models.User{}
	user.Username = username
	user.Password = password
	id, err := models.AddUser(user)
	if err == nil {

	}

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "创建成功",
		"data": id,
	}
	this.ServeJSON()
	return
}

// @获取用户列表
// @Description 获取用户列表的
// @Success 200 {object}  models.User.id models.User.username  models.User.id models.User
// @router / [get]
func (this *UserController) GetUser() {
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "",
		"data": models.GetUser(),
	}
	this.ServeJSON()
	return
}

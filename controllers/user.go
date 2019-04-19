package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ant-blog-beego-service/common/utils"
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
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router / [post]
func (this *UserController) CreateUser() {
	username := this.GetString("username")
	password := this.GetString("password")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.GetUserByUsername(username)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_EXIST,
			"msg":  consts.ERROR_DES_USER_EXIST,
		}
		this.ServeJSON()
		return
	}

	_, err = models.AddUser(username, utils.Crypto(password))
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "创建成功",
	}
	this.ServeJSON()
	return
}

// @获取用户列表
// @Description 获取用户列表的
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router / [get]
func (this *UserController) GetUser() {
	users, err := models.GetUser()
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"data": users,
	}
	this.ServeJSON()
	return
}

package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/common/utils"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

// Operations about User
type UserController struct {
	//BaseController
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("CreateUser", c.CreateUser)
	c.Mapping("GetUsers", c.GetUsers)
	c.Mapping("UpdateTagByUserId", c.UpdateTagByUserId)
}

// @Title 创建用户
// @Description 创建用户
// @Param   username    formData    string  "ant"  true    "用户名"
// @Param   password    formData    string  "123"    true    "密码"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1103   用户名已存在
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
func (this *UserController) GetUsers() {
	users, err := models.GetUsers()
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

// @Title 更新用户信息
// @Description 更新用户信息
// @Param   id    path    string  "1"  true    "用户id"
// @Param   username    formData    string  "ant"  true    "用户名"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1103   用户名已存在
// @router /:id [put]
func (this *UserController) UpdateTagByUserId() {
	id, _ := this.GetInt(":id")
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

	num, err := models.UpdateTagByUserId(id, username, utils.Crypto(password))
	if err != nil || num == 0 {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "更新成功",
	}
	this.ServeJSON()
	return
}

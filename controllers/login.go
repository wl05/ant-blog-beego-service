package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/common/utils"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// Operations about Logout
type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// @Title 用户登录
// @Description 用户登录
// @Param   username    formData    string  "ant"  true    "用户名"
// @Param   password    formData    string  "123"    true    "密码"
// @Success 200 登录成功
// @Success 1101   参数错误
// @Success 1102   请求出错
// @Success 1104 用户名不存在
// @Success 1105 用户名或者密码错误
// @router / [post]
func (this *LoginController) Login() {
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
	user, err := models.GetUserByUsername(username)
	if err == orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_NOT_EXIST,
			"msg":  consts.ERROR_DES_USER_NOT_EXIST,
		}
		this.ServeJSON()
		return
	}
	if user.Password != utils.Crypto(password) {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USERNAME_OR_PASSWORD_ERROR,
			"msg":  consts.ERROR_DES_USERNAME_OR_PASSWORD_ERROR,
		}
		this.ServeJSON()
		return
	}
	ip := this.Ctx.Input.IP() + ":" + strconv.Itoa(this.Ctx.Input.Port())
	token, err := models.CreateToken(user.Id, ip, this.Ctx.Request.UserAgent())
	this.Ctx.SetCookie("token", token, 24*60*60, "/") // 设置cookie
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登录成功",
		"data": token,
	}
	this.ServeJSON()
	return
}

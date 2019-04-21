package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"fmt"
)

// Operations about Logout
type LogoutController struct {
	BaseController
}

func (c *LogoutController) URLMapping() {
	c.Mapping("LoginOut", c.LoginOut)
}

// @Title 用户登出
// @Description 用户登出
// @Success 200 登录成功
// @Success 1101   参数错误
// @Success 1102   请求出错
// @router / [post]
func (this *LogoutController) LoginOut() {
	_, err := models.DeleteTokenByToken(this.Ctx.GetCookie("token"))
	fmt.Println(err)
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
		"msg":  "登出成功",
	}
	this.ServeJSON()
	return
}

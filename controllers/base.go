package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {
	ip := b.Ctx.Input.IP() + ":" + strconv.Itoa(b.Ctx.Input.Port())
	token, err := models.GetTokenByTokenIpUa(b.Ctx.GetCookie("token"), ip, b.Ctx.Request.UserAgent())
	if err == orm.ErrNoRows {
		models.DeleteTokenByToken(b.Ctx.GetCookie("token"))
		b.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_LOGIN_ERROR,
			"msg":  consts.ERROR_DES_LOGIN_ERROR,
		}
		b.ServeJSON()
		return
	} else {
		if !time.Now().Before(token.Expired) {
			models.DeleteTokenByToken(b.Ctx.GetCookie("token"))
			b.Data["json"] = map[string]interface{}{
				"code": consts.ERROR_CODE_LOGIN_EXPIRED_ERROR,
				"msg":  consts.ERROR_DES_LOGIN_EXPIRED_ERROR,
			}
			b.ServeJSON()
			return
		}
	}
}

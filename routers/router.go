// @APIVersion 1.0.0
// @Title Swagger API
// @Description 使用swagger测试爽歪歪.
// @Contact 2929712050@qq.com
package routers

import (
	"ant-blog-beego-service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/tag",
				beego.NSInclude(
					&controllers.TagController{},
				),
			),
			beego.NSNamespace("/category",
				beego.NSInclude(
					&controllers.CategoryController{},
				),
			),
			beego.NSNamespace("/article",
				beego.NSInclude(
					&controllers.ArticleController{},
				),
			),
			beego.NSNamespace("/login",
				beego.NSInclude(
					&controllers.LoginController{},
				),
			),
			beego.NSNamespace("/logout",
				beego.NSInclude(
					&controllers.LogoutController{},
				),
			),
		)
	beego.AddNamespace(ns)
}

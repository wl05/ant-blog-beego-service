package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "PostArticle",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:LoginController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

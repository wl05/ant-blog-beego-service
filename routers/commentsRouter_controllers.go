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

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetCategories",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "UpdateCategoryByCategoryId",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
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

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:LogoutController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:LogoutController"],
        beego.ControllerComments{
            Method: "LoginOut",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"],
        beego.ControllerComments{
            Method: "AddTag",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetTags",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:TagController"],
        beego.ControllerComments{
            Method: "UpdateTagByTagId",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
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
            Method: "GetUsers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateTagByUserId",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

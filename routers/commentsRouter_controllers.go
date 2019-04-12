package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ant-blog-beego-service/controllers:MainController"] = append(beego.GlobalControllerRouter["ant-blog-beego-service/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetTopic",
            Router: `/topic[get]`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

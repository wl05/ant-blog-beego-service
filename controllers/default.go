package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("GetTopic", c.GetTopic)
}

// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param	key		path 	string	true		"The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /topic[get]
func (c *MainController) GetTopic() {
	c.Ctx.WriteString("hello")
}
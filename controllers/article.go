package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego"
	"strings"
)

// Operations about Article
type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) URLMapping() {
	c.Mapping("PostArticle", c.PostArticle)
}

// @Title 发布文章
// @Description 发布文章
// @Param   title    formData    string  "hello world"  true    "文章标题"
// @Param   userId    formData    int  0    true    "用户id"
// @Param   content    formData    string  "hello wordl content"    true    "文章内容"
// @Param   publishAt    formData    date  ""    true    "发布日期"
// @Param   viewCount    formData    int  100    true    "浏览次数"
// @Success 200     {int}
// @Failure 403     body is empty
// @router / [post]
func (this *ArticleController) PostArticle() {
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
	id, err := models.AddUser(username, password)
	if err == nil {

	}

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "创建成功",
		"data": id,
	}
	this.ServeJSON()
	return
}

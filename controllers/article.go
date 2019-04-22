package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

// Operations about Article
type ArticleController struct {
	BaseController
}

func (c *ArticleController) URLMapping() {
	c.Mapping("CreateArticle", c.CreateArticle)
	c.Mapping("GetArticles", c.GetArticles)
	c.Mapping("GetArticle", c.GetArticle)
	c.Mapping("UpdateArticleByArticleId", c.UpdateArticleByArticleId)
	c.Mapping("DeleteArticleById", c.DeleteArticleById)
}

// @Title 发布文章
// @Description 发布文章
// @Param   title    formData    string  "hello world"  true    "文章标题"
// @Param   userId    formData    int  0    true    "用户id"
// @Param   content    formData    string  "hello world content"    true    "文章内容"
// @Param   publishAt    formData    int  ""    true    "发布日期时间戳"
// @Param   categoryId    formData    int  ""    true    "分类id"
// @Param   tagId    formData    int  100    true    "标签id"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router / [post]
func (this *ArticleController) CreateArticle() {
	title := this.GetString("title")
	userId, _ := this.GetInt("userId")
	content := this.GetString("content")
	publishAt, _ := this.GetInt64("publishAt")
	categoryId, _ := this.GetInt("categoryId")
	tagId, _ := this.GetInt("tagId")
	if strings.TrimSpace(title) == "" ||
		userId == 0 ||
		strings.TrimSpace(content) == "" ||
		reflect.ValueOf(publishAt).IsNil() ||
		categoryId == 0 ||
		tagId == 0 {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.CreateArticle(
		title,
		userId,
		content,
		time.Unix(publishAt, publishAt).Local(),
		categoryId,
		tagId,
	)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_EXIST,
			"msg":  consts.ERROR_DES_USER_EXIST,
		}
		this.ServeJSON()
		return
	}
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

// @获取文章列表
// @Description 获取文章列表
// @Success 200 请求成功
// @Success 1102   请求出错
// @router / [get]
func (this *ArticleController) GetArticles() {
	articles, err := models.GetArticles()
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
		"data": articles,
	}
	this.ServeJSON()
	return
}

// @获取单篇文章信息
// @Description 获取单篇文章信息
// @Success 200 请求成功
// @Success 1102   请求出错
// @router /:id [get]
func (this *ArticleController) GetArticle() {
	id, _ := this.GetInt(":id")
	tag, err := models.GetArticleById(id)
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
		"data": tag,
	}
	this.ServeJSON()
	return
}

// @Title 更新文章
// @Description 更新文章
// @Param   title    formData    string  "hello world"  true    "文章标题"
// @Param   userId    formData    int  0    true    "用户id"
// @Param   content    formData    string  "hello world content"    true    "文章内容"
// @Param   publishAt    formData    int  ""    true    "发布日期时间戳"
// @Param   categoryId    formData    int  ""    true    "分类id"
// @Param   tagId    formData    int  100    true    "标签id"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router /:id [put]
func (this *ArticleController) UpdateArticleByArticleId() {
	id, _ := this.GetInt(":id")
	title := this.GetString("title")
	userId, _ := this.GetInt("userId")
	content := this.GetString("content")
	publishAt, _ := this.GetInt64("publishAt")
	categoryId, _ := this.GetInt("categoryId")
	tagId, _ := this.GetInt("tagId")
	if strings.TrimSpace(title) == "" ||
		userId == 0 ||
		strings.TrimSpace(content) == "" ||
		reflect.ValueOf(publishAt).IsNil() ||
		categoryId == 0 ||
		tagId == 0 {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.UpdateArticleByArticleId(
		id,
		title,
		userId,
		content,
		time.Unix(publishAt, publishAt).Local(),
		categoryId,
		tagId,
	)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_EXIST,
			"msg":  consts.ERROR_DES_USER_EXIST,
		}
		this.ServeJSON()
		return
	}
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
		"msg":  "更新成功",
	}
	this.ServeJSON()
	return
}

// @删除文章
// @Description 删除文章
// @Success 200 请求成功
// @Success 1102   请求出错
// @router /:id [get]
func (this *ArticleController) DeleteArticleById() {
	id, _ := this.GetInt(":id")
	_, err := models.DeleteArticleById(id)
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
		"data": "删除成功",
	}
	this.ServeJSON()
	return
}

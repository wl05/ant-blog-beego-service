package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego/orm"
	"strings"
)

// Operations about Tag
type TagController struct {
	BaseController
}

func (c *TagController) URLMapping() {
	c.Mapping("AddTag", c.AddTag)
	c.Mapping("GetTags", c.GetTags)
	c.Mapping("UpdateTagByTagId", c.UpdateTagByTagId)
	c.Mapping("GetTag", c.GetTag)
	c.Mapping("DeleteTagById", c.DeleteTagById)
}

// @Title 创建标签
// @Description 创建标签
// @Param   tagName    formData    string  "javascript"  true    "标签名字"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1108   标签名已存在
// @router / [post]
func (this *TagController) AddTag() {
	tagName := this.GetString("tagName")
	if strings.TrimSpace(tagName) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.GetTagByTagName(tagName)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_TAG_EXIST,
			"msg":  consts.ERROR_DES_TAG_EXIST,
		}
		this.ServeJSON()
		return
	}

	_, err = models.AddTag(tagName)
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

// @获取标签列表
// @Description 获取标签列表
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router / [get]
func (this *TagController) GetTags() {
	tags, err := models.GetTags()
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
		"data": tags,
	}
	this.ServeJSON()
	return
}

// @获取单个标签
// @Description 获取单个标签
// @Success 200 请求成功
// @Success 1102   请求出错
// @router /:id [get]
func (this *TagController) GetTag() {
	id, _ := this.GetInt(":id")
	tag, err := models.GetTagById(id)
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

// @Title 更新标签
// @Description 更新标签
// @Param   id    path    string  "标签id"  true    "标签id"
// @Param   tagName    formData    string  "javascript"  true    "标签名字"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1108   标签名已存在
// @router /:id [put]
func (this *TagController) UpdateTagByTagId() {
	id, _ := this.GetInt(":id")
	tagName := this.GetString("tagName")

	if strings.TrimSpace(tagName) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.GetTagByTagName(tagName)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_TAG_EXIST,
			"msg":  consts.ERROR_DES_TAG_EXIST,
		}
		this.ServeJSON()
		return
	}

	num, err := models.UpdateTagByTagId(id, tagName)
	if err != nil || num == 0 {
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

// @删除标签
// @Description 删除标签
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router /:id [delete]
func (this *TagController) DeleteTagById() {
	id, _ := this.GetInt(":id")
	_, err := models.DeleteTagById(id)
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

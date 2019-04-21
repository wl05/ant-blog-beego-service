package controllers

import (
	"ant-blog-beego-service/common/consts"
	"ant-blog-beego-service/models"
	"github.com/astaxie/beego/orm"
	"strings"
)

// Operations about Category
type CategoryController struct {
	BaseController
}

func (c *CategoryController) URLMapping() {
	c.Mapping("AddCategory", c.AddCategory)
	c.Mapping("GetCategories", c.GetCategories)
	c.Mapping("UpdateCategoryByCategoryId", c.UpdateCategoryByCategoryId)

}

// @Title 创建分类
// @Description 创建分类
// @Param   categoryName    formData    string  "前端"  true    "分类名字"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1109   分类已存在
// @router / [post]
func (this *CategoryController) AddCategory() {
	categoryName := this.GetString("categoryName")
	if strings.TrimSpace(categoryName) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.GetCategoryByCategoryName(categoryName)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_CATEGORY_EXIST,
			"msg":  consts.ERROR_DES_CATEGORY_EXIST,
		}
		this.ServeJSON()
		return
	}

	_, err = models.AddCategory(categoryName)
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

// @获取分类列表
// @Description 获取分类列表
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @router / [get]
func (this *CategoryController) GetCategories() {
	tags, err := models.GetCategories()
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

// @Title 更新分类
// @Description 更新分类
// @Param   id    path    string  "分类id"  true    "分类id"
// @Param   tagName    formData    string  "javascript"  true    "标签名字"
// @Success 200 请求成功
// @Success 1101   外部传入参数错误
// @Success 1102   请求出错
// @Success 1108   标签名已存在
// @router /:id [put]
func (this *CategoryController) UpdateCategoryByCategoryId() {
	id, _ := this.GetInt(":id")
	categoryName := this.GetString("categoryName")

	if strings.TrimSpace(categoryName) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	_, err := models.GetCategoryByCategoryName(categoryName)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_TAG_EXIST,
			"msg":  consts.ERROR_DES_TAG_EXIST,
		}
		this.ServeJSON()
		return
	}

	num, err := models.UpdateCategoryByCategoryId(id, categoryName)
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

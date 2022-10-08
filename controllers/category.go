package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/service"
)

// CategoryController User Api
type CategoryController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *CategoryController) URLMapping() {
	c.Mapping("GetAllCategory", c.GetAllCategory)
}

// GetAllCategory
// @Title GetAllCategory
// @Description 获取所有分类
// @router /all [get]
func (c *CategoryController) GetAllCategory() {

	c.Data["json"] = service.GetAllCategory()
	c.ServeJSON()
}

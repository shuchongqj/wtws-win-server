package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// CargotoController User Api
type CargotoController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *CargotoController) URLMapping() {
	c.Mapping("GetCargotoList", c.GetCargotoList)
	c.Mapping("AddCargoto", c.AddCargoto)
	c.Mapping("DeleteCargoto", c.DeleteCargoto)
	c.Mapping("UpdateCargoto", c.UpdateCargoto)
	c.Mapping("GetAllCargoto", c.GetAllCargoto)
}

// GetAllCargoto
// @Title GetAllCargoto
// @Description 获取所有的装卸货地点
// @router /all [get]
func (c *CargotoController) GetAllCargoto() {

	c.Data["json"] = service.GetAllCargoto()
	c.ServeJSON()
}

// GetCargotoList
// @Title GetCargotoList
// @Description 获取装卸货地点列表
// @router /list [post]
func (c *CargotoController) GetCargotoList() {
	var reqBody request_entity.CargotoList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.CargotoList](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if reqBody.PageSize == 0 {
		reqBody.PageSize = conf.DEFAULT_PAGE_SIZE
	}

	if reqBody.PageNum == 0 {
		reqBody.PageNum = conf.START_PAGE_NUM
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetCargotoList(&reqBody)
	c.ServeJSON()
}

// AddCargoto
// @Title AddCargoto
// @Description 新增收货单位
// @router / [post]
func (c *CargotoController) AddCargoto() {
	var reqBody request_entity.AddCargoto
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddCargoto](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddCargoto(&reqBody)
	c.ServeJSON()
}

// DeleteCargoto
// @Title DeleteCargoto
// @Description 删除收货单位
// @router / [delete]
func (c *CargotoController) DeleteCargoto() {
	var reqBody request_entity.DeleteCargoto
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteCargoto](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteCargoto(&reqBody)
	c.ServeJSON()
}

// UpdateCargoto
// @Title UpdateCargoto
// @Description 修改收货单位
// @router / [put]
func (c *CargotoController) UpdateCargoto() {
	var reqBody request_entity.UpdateCargoto
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateCargoto](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateCargoto(&reqBody)
	c.ServeJSON()
}

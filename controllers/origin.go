package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// OriginController User Api
type OriginController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *OriginController) URLMapping() {
	c.Mapping("GetAllOriginList", c.GetAllOriginList)
	c.Mapping("GetOriginList", c.GetOriginList)
	c.Mapping("AddOrigin", c.AddOrigin)
	c.Mapping("DeleteOrigin", c.DeleteOrigin)
	c.Mapping("UpdateOrigin", c.UpdateOrigin)
}

// GetAllOriginList
// @Title GetAllOriginList
// @Description 获取所有发货地址列表
// @router /all-list [get]
func (c *OriginController) GetAllOriginList() {

	c.Data["json"] = service.GetAllOriginList()
	c.ServeJSON()
}

// GetOriginList
// @Title GetOriginList
// @Description 获取发货地列表
// @router /list [post]
func (c *OriginController) GetOriginList() {
	var reqBody request_entity.OriginList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.OriginList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetOriginList(&reqBody)
	c.ServeJSON()
}

// AddOrigin
// @Title AddOrigin
// @Description 新增收货单位
// @router / [post]
func (c *OriginController) AddOrigin() {
	var reqBody request_entity.AddOrigin
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddOrigin](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddOrigin(&reqBody)
	c.ServeJSON()
}

// DeleteOrigin
// @Title DeleteOrigin
// @Description 删除收货单位
// @router / [delete]
func (c *OriginController) DeleteOrigin() {
	var reqBody request_entity.DeleteOrigin
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteOrigin](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteOrigin(&reqBody)
	c.ServeJSON()
}

// UpdateOrigin
// @Title UpdateOrigin
// @Description 修改收货单位
// @router / [put]
func (c *OriginController) UpdateOrigin() {
	var reqBody request_entity.UpdateOrigin
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateOrigin](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateOrigin(&reqBody)
	c.ServeJSON()
}

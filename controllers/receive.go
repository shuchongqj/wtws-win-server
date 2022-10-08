package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// ReceiveController User Api
type ReceiveController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *ReceiveController) URLMapping() {
	c.Mapping("GetAllReceiveList", c.GetAllReceiveList)
	c.Mapping("GetReceiveList", c.GetReceiveList)
	c.Mapping("AddReceive", c.AddReceive)
	c.Mapping("DeleteReceive", c.DeleteReceive)
	c.Mapping("UpdateReceive", c.UpdateReceive)
}

// GetAllReceiveList
// @Title GetAllReceiveList
// @Description 获取所有收货地址列表
// @router /all-list [get]
func (c *ReceiveController) GetAllReceiveList() {

	c.Data["json"] = service.GetAllReceiveList()
	c.ServeJSON()
}

// GetReceiveList
// @Title GetReceiveList
// @Description 获取收货地址列表
// @router /list [post]
func (c *ReceiveController) GetReceiveList() {
	var reqBody request_entity.ReceiveList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.ReceiveList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetReceiveList(&reqBody)
	c.ServeJSON()
}

// AddReceive
// @Title AddReceive
// @Description 新增收货单位
// @router / [post]
func (c *ReceiveController) AddReceive() {
	var reqBody request_entity.AddReceive
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddReceive](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddReceive(&reqBody)
	c.ServeJSON()
}

// DeleteReceive
// @Title DeleteReceive
// @Description 删除收货单位
// @router / [delete]
func (c *ReceiveController) DeleteReceive() {
	var reqBody request_entity.DeleteReceive
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteReceive](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteReceive(&reqBody)
	c.ServeJSON()
}

// UpdateReceive
// @Title UpdateReceive
// @Description 修改收货单位
// @router / [put]
func (c *ReceiveController) UpdateReceive() {
	var reqBody request_entity.UpdateReceive
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateReceive](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateReceive(&reqBody)
	c.ServeJSON()
}

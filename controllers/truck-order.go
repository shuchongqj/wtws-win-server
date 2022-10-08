package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// TruckOrderController User Api
type TruckOrderController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *TruckOrderController) URLMapping() {
	c.Mapping("GetTruckOrderList", c.GetTruckOrderList)
	c.Mapping("AddTruckOrder", c.AddTruckOrder)
	c.Mapping("AddSentDirectTruckOrder", c.AddSentDirectTruckOrder)
	c.Mapping("DeleteTruckOrder", c.DeleteTruckOrder)
	c.Mapping("UpdateTruckOrder", c.UpdateTruckOrder)
	c.Mapping("CheckTruckOrder", c.CheckTruckOrder)
	c.Mapping("InvalidTruckOrder", c.InvalidTruckOrder)
	c.Mapping("DownAllTruckOrder", c.DownAllTruckOrder)
	c.Mapping("GetTruckOrderByVehicle", c.GetTruckOrderByVehicle)

}

// GetTruckOrderList
// @Title GetTruckOrderList
// @Description 获取派车单列表
// @router /list [post]
func (c *TruckOrderController) GetTruckOrderList() {
	var reqBody request_entity.TruckOrderList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.TruckOrderList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetTruckOrderList(&reqBody)
	c.ServeJSON()
}

// GetTruckOrderByVehicle
// @Title GetTruckOrderByVehicle
// @Description 获取车辆关联的派车单
// @router /vehicle [get]
func (c *TruckOrderController) GetTruckOrderByVehicle() {
	vehicleNumber := ""
	c.Ctx.Input.Bind(&vehicleNumber, "vehicleNumber")
	if len(vehicleNumber) == 0 {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetTruckOrderByVehicle(vehicleNumber)
	c.ServeJSON()
}

// AddTruckOrder
// @Title AddTruckOrder
// @Description 新增派车单
// @router / [post]
func (c *TruckOrderController) AddTruckOrder() {
	var reqBody request_entity.AddTruckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddTruckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddTruckOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// AddSentDirectTruckOrder
// @Title AddSentDirectTruckOrder
// @Description 新增直发/倒短派车单
// @router /sent-direct [post]
func (c *TruckOrderController) AddSentDirectTruckOrder() {
	var reqBody request_entity.AddSentDirectOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddSentDirectOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddSentDirectOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DeleteTruckOrder
// @Title DeleteTruckOrder
// @Description 删除派车单
// @router / [delete]
func (c *TruckOrderController) DeleteTruckOrder() {
	var reqBody request_entity.DeleteTruckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteTruckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteTruckOrder(&reqBody)
	c.ServeJSON()
}

// UpdateTruckOrder
// @Title UpdateTruckOrder
// @Description 修改派车单
// @router / [put]
func (c *TruckOrderController) UpdateTruckOrder() {
	var reqBody request_entity.UpdateTruckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateTruckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateTruckOrder(&reqBody)
	c.ServeJSON()
}

// CheckTruckOrder
// @Title CheckTruckOrder
// @Description 审核派车单
// @router /check [post]
func (c *TruckOrderController) CheckTruckOrder() {
	var reqBody request_entity.CheckTruckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.CheckTruckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.CheckTruckOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// InvalidTruckOrder
// @Title InvalidTruckOrder
// @Description 作废派车单
// @router /invalid [post]
func (c *TruckOrderController) InvalidTruckOrder() {
	var reqBody request_entity.InvalidTruckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.InvalidTruckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.InvalidTruckOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DownAllTruckOrder
// @Title DownAllTruckOrder
// @Description 下载所有派车单
// @router /down-all [get]
func (c *TruckOrderController) DownAllTruckOrder() {

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	bytes := service.DownAllTruckOrder()
	c.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
	c.Ctx.Output.Body(bytes)
}

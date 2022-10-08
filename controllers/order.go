package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// OrderController User Api
type OrderController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *OrderController) URLMapping() {
	c.Mapping("GetOrderList", c.GetOrderList)
	c.Mapping("AddOrder", c.AddOrder)
	c.Mapping("DeleteOrder", c.DeleteOrder)
	c.Mapping("UpdateOrder", c.UpdateOrder)
	c.Mapping("CheckOrder", c.CheckOrder)
	c.Mapping("InvalidOrder", c.InvalidOrder)
	c.Mapping("DownAllOrder", c.DownAllOrder)
	c.Mapping("GetAllCheckedOrder", c.GetAllCheckedOrder)
}

// GetAllCheckedOrder
// @Title GetAllCheckedOrder
// @Description 获取所有的审核通过的订单
// @router /checked [get]
func (c *OrderController) GetAllCheckedOrder() {

	c.Data["json"] = service.GetAllCheckedOrder()
	c.ServeJSON()
}

// GetOrderInfo
// @Title GetOrderInfo
// @Description 查询订单详情
// @router /info [get]
func (c *OrderController) GetOrderInfo() {
	orderID := 0
	c.Ctx.Input.Bind(&orderID, "orderID")
	c.Data["json"] = service.GetOrderInfo(orderID)
	c.ServeJSON()
}

// GetOrderList
// @Title GetOrderList
// @Description 获取订单列表
// @router /list [post]
func (c *OrderController) GetOrderList() {
	var reqBody request_entity.OrderList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.OrderList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetOrderList(&reqBody)
	c.ServeJSON()
}

// AddOrder
// @Title AddOrder
// @Description 新增订单
// @router / [post]
func (c *OrderController) AddOrder() {
	var reqBody request_entity.AddOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DeleteOrder
// @Title DeleteOrder
// @Description 删除订单
// @router / [delete]
func (c *OrderController) DeleteOrder() {
	var reqBody request_entity.DeleteOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteOrder(&reqBody)
	c.ServeJSON()
}

// UpdateOrder
// @Title UpdateOrder
// @Description 修改订单
// @router / [put]
func (c *OrderController) UpdateOrder() {
	var reqBody request_entity.UpdateOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateOrder(&reqBody)
	c.ServeJSON()
}

// CheckOrder
// @Title CheckOrder
// @Description 审核订单
// @router /check [post]
func (c *OrderController) CheckOrder() {
	var reqBody request_entity.CheckOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.CheckOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.CheckOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// InvalidOrder
// @Title InvalidOrder
// @Description 作废订单
// @router /invalid [post]
func (c *OrderController) InvalidOrder() {
	var reqBody request_entity.InvalidOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.InvalidOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.InvalidOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DownAllOrder
// @Title DownAllOrder
// @Description 下载所有订单
// @router /down-all [get]
func (c *OrderController) DownAllOrder() {

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	bytes := service.DownAllOrder()
	c.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
	c.Ctx.Output.Body(bytes)
}

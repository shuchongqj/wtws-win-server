package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// DriverController User Api
type DriverController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *DriverController) URLMapping() {
	c.Mapping("GetDriverList", c.GetDriverList)
	c.Mapping("AddDriver", c.AddDriver)
	c.Mapping("DeleteDriver", c.DeleteDriver)
	c.Mapping("UpdateDriver", c.UpdateDriver)
	c.Mapping("GetAllDriver", c.GetAllDriver)
	c.Mapping("TruckOrderList", c.TruckOrderList)
}

// GetAllDriver
// @Title GetAllDriver
// @Description 获取所有的司机
// @router /all [get]
func (c *DriverController) GetAllDriver() {
	c.Data["json"] = service.GetAllDriver()
	c.ServeJSON()
}

// GetDriverList
// @Title GetDriverList
// @Description 获取司机列表
// @router /list [post]
func (c *DriverController) GetDriverList() {
	var reqBody request_entity.DriverList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DriverList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetDriverList(&reqBody)
	c.ServeJSON()
}

// AddDriver
// @Title AddDriver
// @Description 新增司机
// @router / [post]
func (c *DriverController) AddDriver() {
	var reqBody request_entity.AddDriver
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddDriver](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if len(reqBody.DriverAccount) == 0 || len(reqBody.VehicleNumber) == 0 || len(reqBody.DriverName) == 0 || len(reqBody.Tel) == 0 {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddDriver(&reqBody)
	c.ServeJSON()
}

// DeleteDriver
// @Title DeleteDriver
// @Description 删除司机
// @router / [delete]
func (c *DriverController) DeleteDriver() {
	var reqBody request_entity.DeleteDriver
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteDriver](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteDriver(&reqBody)
	c.ServeJSON()
}

// UpdateDriver
// @Title UpdateDriver
// @Description 修改司机
// @router / [put]
func (c *DriverController) UpdateDriver() {
	var reqBody request_entity.UpdateDriver
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateDriver](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateDriver(&reqBody)
	c.ServeJSON()
}

// TruckOrderList
// @Title TruckOrderList
// @Description 获取司机的派车单
// @router /truck-order/list [get]
func (c *DriverController) TruckOrderList() {

	pageNum := 1
	pageSize := 10
	c.Ctx.Input.Bind(&pageNum, "pageNum")
	if pageNum == 0 {
		pageNum = 1
	}
	c.Ctx.Input.Bind(&pageSize, "pageSize")
	if pageSize == 0 {
		pageNum = 10
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))
	c.Data["json"] = service.DriverTruckOrderList(pageNum, pageSize, userInfo)
	c.ServeJSON()
}

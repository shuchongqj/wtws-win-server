package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// StationController User Api
type StationController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *StationController) URLMapping() {
	c.Mapping("GetUserStation", c.GetUserStation)
	c.Mapping("GetStationList", c.GetStationList)
	c.Mapping("UpdateStation", c.UpdateStation)
	c.Mapping("AddStation", c.AddStation)
	c.Mapping("DeleteStation", c.DeleteStation)
}

// GetUserStation
// @Title GetUserStation
// @Description 获取用户所属服务站
// @router /user/list [get]
func (c *StationController) GetUserStation() {
	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetUserStation(userInfo.Id)
	c.ServeJSON()
}

// GetStationList
// @Title GetStationList
// @Description 获取服务站列表
// @router /list [post]
func (c *StationController) GetStationList() {

	var reqBody request_entity.StationList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.StationList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetStationList(&reqBody, userInfo.Id)
	c.ServeJSON()
}

// UpdateStation
// @Title UpdateStation
// @Description 更改服务站信息
// @router / [put]
func (c *StationController) UpdateStation() {

	var reqBody request_entity.UpdateStation
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateStation](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateStation(&reqBody, userInfo.Id)
	c.ServeJSON()
}

// AddStation
// @Title AddStation
// @Description 新增服务站
// @router / [post]
func (c *StationController) AddStation() {

	var reqBody request_entity.AddStation
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddStation](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.AddStation(&reqBody)
	c.ServeJSON()
}

// DeleteStation
// @Title DeleteStation
// @Description 删除服务站
// @router / [delete]
func (c *StationController) DeleteStation() {

	var reqBody request_entity.DeleteStation
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteStation](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.DeleteStation(&reqBody)
	c.ServeJSON()
}

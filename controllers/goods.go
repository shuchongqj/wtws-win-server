package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// GoodsController User Api
type GoodsController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *GoodsController) URLMapping() {
	c.Mapping("GetAllGoodsList", c.GetAllGoodsList)
	c.Mapping("GetGoodsList", c.GetGoodsList)
	c.Mapping("AddGoods", c.AddGoods)
	c.Mapping("DeleteGoods", c.DeleteGoods)
	c.Mapping("UpdateGoods", c.UpdateGoods)
}

// GetAllGoodsList
// @Title GetAllGoodsList
// @Description 获取所有货品列表
// @router /all-list [get]
func (c *GoodsController) GetAllGoodsList() {
	c.Data["json"] = service.GetAllGoodsList()
	c.ServeJSON()
}

// GetGoodsList
// @Title GetGoodsList
// @Description 获取货品列表
// @router /list [post]
func (c *GoodsController) GetGoodsList() {
	var reqBody request_entity.GoodsList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.GoodsList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetGoodsList(&reqBody)
	c.ServeJSON()
}

// AddGoods
// @Title AddGoods
// @Description 新增收货单位
// @router / [post]
func (c *GoodsController) AddGoods() {
	var reqBody request_entity.AddGoods
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddGoods](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddGoods(&reqBody)
	c.ServeJSON()
}

// DeleteGoods
// @Title DeleteGoods
// @Description 删除收货单位
// @router / [delete]
func (c *GoodsController) DeleteGoods() {
	var reqBody request_entity.DeleteGoods
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteGoods](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteGoods(&reqBody)
	c.ServeJSON()
}

// UpdateGoods
// @Title UpdateGoods
// @Description 修改收货单位
// @router / [put]
func (c *GoodsController) UpdateGoods() {
	var reqBody request_entity.UpdateGoods
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateGoods](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateGoods(&reqBody)
	c.ServeJSON()
}

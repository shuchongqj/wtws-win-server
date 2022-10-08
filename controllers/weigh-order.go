package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// WeighOrderController User Api
type WeighOrderController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *WeighOrderController) URLMapping() {
	c.Mapping("GetWeighOrderList", c.GetWeighOrderList)
	c.Mapping("AddWeighOrder", c.AddWeighOrder)
	c.Mapping("DeleteWeighOrder", c.DeleteWeighOrder)
	c.Mapping("UpdateWeighOrder", c.UpdateWeighOrder)
	c.Mapping("CheckWeighOrder", c.WarehouseCheckWeighOrder)
	c.Mapping("InvalidWeighOrder", c.InvalidWeighOrder)
	c.Mapping("DownAllWeighOrder", c.DownAllWeighOrder)
	c.Mapping("ScanVehicle", c.ScanVehicle)
	c.Mapping("TareWight", c.TareWight)
	c.Mapping("CheckWareHouseGoods", c.CheckWareHouseGoods)
	c.Mapping("WaitFinishOrder", c.WaitFinishOrder)
	c.Mapping("FinishWeighOrder", c.FinishWeighOrder)
}

// GetWeighOrderList
// @Title GetWeighOrderList
// @Description 获取过磅单列表
// @router /list [post]
func (c *WeighOrderController) GetWeighOrderList() {
	var reqBody request_entity.WeighOrderList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.WeighOrderList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetWeighOrderList(&reqBody)
	c.ServeJSON()
}

// AddWeighOrder
// @Title AddWeighOrder
// @Description 新增过磅单
// @router / [post]
func (c *WeighOrderController) AddWeighOrder() {
	var reqBody request_entity.AddWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddWeighOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DeleteWeighOrder
// @Title DeleteWeighOrder
// @Description 删除过磅单
// @router / [delete]
func (c *WeighOrderController) DeleteWeighOrder() {
	var reqBody request_entity.DeleteWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteWeighOrder(&reqBody)
	c.ServeJSON()
}

// UpdateWeighOrder
// @Title UpdateWeighOrder
// @Description 修改过磅单
// @router / [put]
func (c *WeighOrderController) UpdateWeighOrder() {
	var reqBody request_entity.UpdateWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.UpdateWeighOrder(&reqBody)
	c.ServeJSON()
}

// WarehouseCheckWeighOrder
// @Title WarehouseCheckWeighOrder
// @Description 仓库确认过磅单
// @router /warehouse-check [post]
func (c *WeighOrderController) WarehouseCheckWeighOrder() {
	var reqBody request_entity.WarehouseCheckWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.WarehouseCheckWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.WarehouseCheckWeighOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// InvalidWeighOrder
// @Title InvalidWeighOrder
// @Description 作废过磅单
// @router /invalid [post]
func (c *WeighOrderController) InvalidWeighOrder() {
	var reqBody request_entity.InvalidWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.InvalidWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.InvalidWeighOrder(&reqBody, userInfo)
	c.ServeJSON()
}

// DownAllWeighOrder
// @Title DownAllWeighOrder
// @Description 下载所有过磅单
// @router /down-all [get]
func (c *WeighOrderController) DownAllWeighOrder() {

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	bytes := service.DownAllWeighOrder()
	c.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
	c.Ctx.Output.Body(bytes)
}

// ScanVehicle
// @Title ScanVehicle
// @Description 过磅单识别车牌上报数据
// @router /scan-vehicle [post]
func (c *WeighOrderController) ScanVehicle() {

	var reqBody request_entity.ScanVehicle
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.ScanVehicle](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
		c.ServeJSON()
		return
	}

	if len(reqBody.TransportTime) == 0 {
		reqBody.TransportTime = time.Now().Format("2006-01-02 15:04:05")
	}

	c.Data["json"] = service.ScanVehicle(&reqBody)
	c.ServeJSON()
}

// TareWight
// @Title TareWight
// @Description 过磅单上磅测量皮重
// @router /tare-wight [post]
func (c *WeighOrderController) TareWight() {

	var reqBody request_entity.TareWight
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.TareWight](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.TareWight(&reqBody)
	c.ServeJSON()
}

// CheckWareHouseGoods
// @Title CheckWareHouseGoods
// @Description 检查仓库是否已经选择货品
// @router /ware-house-check-goods [post]
func (c *WeighOrderController) CheckWareHouseGoods() {

	var reqBody request_entity.CheckWareHouseGoods
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.CheckWareHouseGoods](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.CheckWareHouseGoods(&reqBody)
	c.ServeJSON()
}

// WaitFinishOrder
// @Title WaitFinishOrder
// @Description 等待完成
// @router /wait-finish [post]
func (c *WeighOrderController) WaitFinishOrder() {
	var reqBody request_entity.WaitFinish
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.WaitFinish](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = service.WaitFinish(&reqBody)
	c.ServeJSON()
}

// FinishWeighOrder
// @Title FinishWeighOrder
// @Description 完成过磅
// @router /finish [post]
func (c *WeighOrderController) FinishWeighOrder() {
	var reqBody request_entity.FinishWeighOrder
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.FinishWeighOrder](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = service.FinishWeighOrder(&reqBody)
	c.ServeJSON()
}

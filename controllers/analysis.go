package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
)

// AnalysisController User Api
type AnalysisController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *AnalysisController) URLMapping() {
	c.Mapping("GetAnalysisDetail", c.GetAnalysisDetail)
	c.Mapping("GetAnalysisOrderTypeDetail", c.GetAnalysisOrderTypeDetail)
}

// GetAnalysisDetail
// @Title GetAnalysisDetail
// @Description 获取首页数据分析详情
// @router /detail [get]
func (c *AnalysisController) GetAnalysisDetail() {
	var dateType string
	if getDataErr := c.Ctx.Input.Bind(&dateType, "dateType"); getDataErr != nil {
		logs.Error("参数异常 err:", getDataErr.Error())
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if dateType != conf.ANALYSIS_DATE_TYPE_DAY &&
		dateType != conf.ANALYSIS_DATE_TYPE_MONTH &&
		dateType != conf.ANALYSIS_DATE_TYPE_YEAR {
		logs.Error("[service]  请求的dateType参数错误")
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.GetAnalysisDetail(dateType)
	c.ServeJSON()
}

// GetAnalysisOrderTypeDetail
// @Title GetAnalysisOrderTypeDetail
// @Description 获取首页数据订单分析详情
// @router /order-type-detail [get]
func (c *AnalysisController) GetAnalysisOrderTypeDetail() {
	var dateType string
	var orderType string
	if getDataErr := c.Ctx.Input.Bind(&dateType, "dateType"); getDataErr != nil {
		logs.Error("参数异常 err:", getDataErr.Error())
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if getOrderTypeErr := c.Ctx.Input.Bind(&orderType, "orderType"); getOrderTypeErr != nil {
		logs.Error("参数异常 err:", getOrderTypeErr.Error())
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if dateType != conf.ANALYSIS_DATE_TYPE_DAY &&
		dateType != conf.ANALYSIS_DATE_TYPE_MONTH &&
		dateType != conf.ANALYSIS_DATE_TYPE_YEAR {
		logs.Error("[service]  请求的dateType参数错误")
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	if orderType != conf.ANALYSIS_ORDER_TYPE_PURCHASE &&
		orderType != conf.ANALYSIS_ORDER_TYPE_SALE &&
		orderType != conf.ANALYSIS_ORDER_TYPE_SENT_DIRECT &&
		orderType != conf.ANALYSIS_ORDER_TYPE_TRUCK &&
		orderType != conf.ANALYSIS_ORDER_TYPE_WEIGH {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.GetAnalysisOrderTypeDetail(dateType, orderType)
	c.ServeJSON()
}

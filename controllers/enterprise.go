package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
)

// EnterpriseController User Api
type EnterpriseController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *EnterpriseController) URLMapping() {
	c.Mapping("GetUserEnterprises", c.GetUserEnterprises)
}

// GetUserEnterprises
// @Title GetUserEnterprises
// @Description 获取用户所属企业
// @router /user/list [get]
func (c *EnterpriseController) GetUserEnterprises() {
	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetUserEnterprise(userInfo.Id)
	c.ServeJSON()
}

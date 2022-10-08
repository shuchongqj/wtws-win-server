package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// RoleController User Api
type RoleController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *RoleController) URLMapping() {
	c.Mapping("GetRoleFunctions", c.GetRoleFunctions)
	c.Mapping("GetAllRoleList", c.GetAllRoleList)
	c.Mapping("AddRole", c.AddRole)
	c.Mapping("AddRoleFunctions", c.AddRoleFunctions)
}

// GetRoleFunctions
// @Title GetRoleFunctions
// @Description 获取角色的功能
// @router /functions [get]
func (c *RoleController) GetRoleFunctions() {
	var roleId int
	if getDataErr := c.Ctx.Input.Bind(&roleId, "roleId"); getDataErr != nil {
		logs.Error("参数异常 err:", getDataErr.Error())
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}
	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetRoleFunctions(roleId)
	c.ServeJSON()
}

// GetAllRoleList
// @Title GetAllRoleList
// @Description 获取所有角色的权限列表
// @router /all-list [get]
func (c *RoleController) GetAllRoleList() {
	c.Data["json"] = service.GetAllRoleList()
	c.ServeJSON()
}

// DeleteRole
// @Title DeleteRole
// @Description 删除权限
// @router / [delete]
func (c *RoleController) DeleteRole() {
	var reqBody request_entity.DeleteRole
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteRole](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	//userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.DeleteRole(&reqBody)
	c.ServeJSON()
}

// AddRole
// @Title AddRole
// @Description 新增权限
// @router / [post]
func (c *RoleController) AddRole() {
	var reqBody request_entity.AddRole
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddRole](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddRole(reqBody, userInfo)
	c.ServeJSON()
}

// AddRoleFunctions
// @Title AddRoleFunctions
// @Description 新增权限
// @router /functions [post]
func (c *RoleController) AddRoleFunctions() {
	var reqBody request_entity.AddRoleFunctions
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddRoleFunctions](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.AddRoleFunctions(&reqBody, userInfo)
	c.ServeJSON()
}

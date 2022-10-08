package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
	request_entity "wtws-server/service-struct/request-entity"
	"wtws-server/validata"
)

// UserController User Api
type UserController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *UserController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("UpdatePwd", c.UpdatePwd)
	c.Mapping("GetUserInfo", c.GetUserInfo)
	c.Mapping("GetUserLIst", c.GetUserLIst)
	c.Mapping("AddUserInfo", c.AddUserInfo)
	c.Mapping("DeleteUser", c.DeleteUser)
	c.Mapping("UpdateUser", c.UpdateUser)
	c.Mapping("UpdateUserLoginName", c.UpdateUserLoginName)
	c.Mapping("RestUserPwd", c.RestUserPwd)
	c.Mapping("LogOut", c.LogOut)
	c.Mapping("UpdateUserRole", c.UpdateUserRole)
}

// Login
// @Title Login
// @Description 登录
// @router /login [post]
func (c *UserController) Login() {

	var reqBody request_entity.UserLogin
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UserLogin](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.Login(reqBody)
	c.ServeJSON()

}

// UpdatePwd
// @Title UpdatePwd
// @Description 更改密码
// @router /update-pwd [post]
func (c *UserController) UpdatePwd() {
	var reqBody request_entity.UserUpdatePwd
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UserUpdatePwd](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	if reqBody.OldPassWord != userInfo.PassWord {
		c.Data["json"] = common.ResponseStatus(-4, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.UpdatePwd(reqBody, userInfo)
	c.ServeJSON()

}

// GetUserInfo
// @Title GetUserInfo
// @Description 根据用户token获取用户详情
// @router /user-info [get]
func (c *UserController) GetUserInfo() {
	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	if userInfo.Id <= 0 {
		c.Data["json"] = common.ResponseStatus(-1, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.GetUserInfo(userInfo)
	c.ServeJSON()
}

// GetUserLIst
// @Title GetUserLIst
// @Description 获取用户列表
// @router /list [post]
func (c *UserController) GetUserLIst() {
	var reqBody request_entity.UserList
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UserList](c.Ctx.Input.RequestBody); reqErr != nil {
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

	c.Data["json"] = service.GetUserList(&reqBody)
	c.ServeJSON()
}

// AddUserInfo
// @Title AddUserInfo
// @Description 创建用户
// @router /add [post]
func (c *UserController) AddUserInfo() {
	var reqBody request_entity.AddUserInfo
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.AddUserInfo](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.AddUser(&reqBody)
	c.ServeJSON()
}

// DeleteUser
// @Title DeleteUser
// @Description 删除用户
// @router / [delete]
func (c *UserController) DeleteUser() {
	var reqBody request_entity.DeleteUser
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.DeleteUser](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.DeleteUser(&reqBody)
	c.ServeJSON()
}

// UpdateUser
// @Title UpdateUser
// @Description 修改用户
// @router / [put]
func (c *UserController) UpdateUser() {
	var reqBody request_entity.UpdateUser
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateUser](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.UpdateUser(&reqBody)
	c.ServeJSON()
}

// UpdateUserLoginName
// @Title UpdateUserLoginName
// @Description 修改用户登录账号
// @router /login-name [put]
func (c *UserController) UpdateUserLoginName() {
	var reqBody request_entity.UpdateUserLoginName
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateUserLoginName](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.UpdateUserLoginName(&reqBody)
	c.ServeJSON()
}

// UpdateUserRole
// @Title UpdateUserRole
// @Description 修改用户权限
// @router /role [put]
func (c *UserController) UpdateUserRole() {
	var reqBody request_entity.UpdateUserRole
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.UpdateUserRole](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.UpdateUserRole(&reqBody)
	c.ServeJSON()
}

// RestUserPwd
// @Title RestUserPwd
// @Description 重置用户密码
// @router /reset-pwd [post]
func (c *UserController) RestUserPwd() {
	var reqBody request_entity.RestUserPwd
	var reqErr error
	if reqBody, reqErr = validata.CommonValidate[request_entity.RestUserPwd](c.Ctx.Input.RequestBody); reqErr != nil {
		c.Data["json"] = common.ResponseStatus(-99, "", nil)
		c.ServeJSON()
		return
	}

	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	adminRoleIDStr := conf.ADMIN_ROLE_ID
	adminRoleID, _ := strconv.Atoi(adminRoleIDStr)
	if userInfo.Id != adminRoleID {
		c.Data["json"] = common.ResponseStatus(-9, "", nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = service.RestUserPwd(&reqBody, userInfo)
	c.ServeJSON()
}

// LogOut
// @Title LogOut
// @Description 用户登出
// @router /logout [post]
func (c *UserController) LogOut() {
	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.LogOut(userInfo.Id)
	c.ServeJSON()
}

package service

import (
	"strconv"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	"wtws-server/models"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetRoleFunctions(roleId int) common_struct.ResponseStruct {
	roleFunctions := []wtws_mysql.FunctionInfo{}
	//if userRole, err := wtws_mysql.GetSUserRoleByUserId(userId); err == nil {
	//	roleFunctions = wtws_mysql.GetAllFunctionsByRoleID(userRole.RoleId)
	//}
	if roleId == 0 {
		adminRoleIdStr := conf.ADMIN_ROLE_ID
		roleId, _ = strconv.Atoi(adminRoleIdStr)
	}
	roleFunctions = wtws_mysql.GetAllFunctionsByRoleID(roleId)
	return common.ResponseStatus(0, "", roleFunctions)
}

func GetAllRoleList() common_struct.ResponseStruct {
	roleList := wtws_mysql.GetAllRoleList()
	roleListDto := []dto.RoleListItemData{}
	for _, role := range roleList {
		roleListDto = append(roleListDto, dto.RoleListItemData{
			RoleId:      role.Id,
			RoleName:    role.RoleName,
			Description: role.Description,
		})
	}
	return common.ResponseStatus(0, "", roleListDto)
}

func AddRole(data request_entity.AddRole, user *wtws_mysql.SUser) common_struct.ResponseStruct {

	adminRoleIdStr := conf.ADMIN_ROLE_ID
	adminRoleId, _ := strconv.Atoi(adminRoleIdStr)

	var userRole *wtws_mysql.SUserRole
	var err error
	if userRole, err = wtws_mysql.GetSUserRoleByUserId(user.Id); err != nil || userRole.Id != adminRoleId {
		return common.ResponseStatus(-8, "", nil)
	}

	if dbRole, err := wtws_mysql.GetSRoleByRoleName(data.RoleName); dbRole != nil || err == nil {
		return common.ResponseStatus(-6, "", nil)
	}

	if roleID, err := wtws_mysql.AddSRole(&wtws_mysql.SRole{
		RoleName:   data.RoleName,
		RoleType:   1,
		IsDeleted:  int8(data.RoleType),
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	}); err != nil || roleID <= 0 {
		return common.ResponseStatus(-13, "", nil)
	} else {
		return common.ResponseStatus(12, "", dto.AddRoleResult{RoleId: int(roleID)})
	}
}

func AddRoleFunctions(data *request_entity.AddRoleFunctions, user *wtws_mysql.SUser) common_struct.ResponseStruct {

	adminRoleIdStr := conf.ADMIN_ROLE_ID
	adminRoleId, _ := strconv.Atoi(adminRoleIdStr)

	var userRole *wtws_mysql.SUserRole
	var err error
	if userRole, err = wtws_mysql.GetSUserRoleByUserId(user.Id); err != nil || userRole.Id != adminRoleId {
		return common.ResponseStatus(-8, "", nil)
	}

	if _, err := wtws_mysql.DelRoleFunctionByRoleID(data.RoleFunctionList[0].RoleID); err != nil {
		return common.ResponseStatus(500, "", nil)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(data.RoleFunctionList))

	for _, roleFunctionItem := range data.RoleFunctionList {
		go func(roleFunctionItem request_entity.AddRoleFunctionItem) {
			wtws_mysql.AddSRoleFunction(&wtws_mysql.SRoleFunction{
				RoleId:     roleFunctionItem.RoleID,
				FunctionId: roleFunctionItem.FunctionID,
				InsertTime: time.Now(),
				UpdateTime: time.Now(),
			})
			wg.Done()
		}(roleFunctionItem)
	}

	wg.Wait()

	//删除之前权限的用户登录记录
	userRoleList, _ := wtws_mysql.GetSUserRoleByRoleId(data.RoleFunctionList[0].RoleID)
	wg2 := &sync.WaitGroup{}
	wg2.Add(len(userRoleList))
	for _, userRole := range userRoleList {
		go func(userId int) {
			models.RedisDel(strconv.Itoa(userId))
			wg2.Done()
		}(userRole.UserId)
	}

	wg2.Wait()

	return common.ResponseStatus(0, "", nil)
}

// DeleteRole 删除角色
func DeleteRole(data *request_entity.DeleteRole) common_struct.ResponseStruct {
	list, _ := wtws_mysql.GetSUserRoleByRoleId(data.RoleID)
	if len(list) > 0 {
		return common.ResponseStatus(-11, "", nil)
	}
	if delRoleErr, num := wtws_mysql.UpdateByRoleID(&wtws_mysql.SRole{
		Id:        data.RoleID,
		IsDeleted: conf.IS_DELETE,
	}, []string{"IsDeleted"}); delRoleErr != nil || num == 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

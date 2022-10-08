package request_entity

type AddRoleFunctionItem struct {
	RoleID     int `json:"roleID" validate:"required"`     //用户角色
	FunctionID int `json:"functionID" validate:"required"` // 权限ID
}
type AddRoleFunctions struct {
	RoleFunctionList []AddRoleFunctionItem `json:"roleFunctionList" validate:"required"` //权限list
}

type AddRole struct {
	RoleName string `json:"roleName" validate:"required"`
	RoleType int    `json:"roleType" validate:"required"`
}

type DeleteRole struct {
	RoleID int `json:"roleID" validate:"required"`
}

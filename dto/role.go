package dto

type RoleListItemData struct {
	RoleId      int    `json:"roleID"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}

type AddRoleResult struct {
	RoleId int `json:"roleID"`
}

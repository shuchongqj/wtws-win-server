package dto

import (
	service_struct "wtws-server/service-struct"
)

type Login service_struct.Authorization

type UserInfo struct {
	UserId      int    `json:"userId"`
	RoleID      int    `json:"roleID"`
	LoginName   string `json:"loginName"`
	WorkNo      string `json:"workNo"`
	DisplayName string `json:"displayName"`
	JobTitle    string `json:"jobTitle"`
	HeadPicUrl  string `json:"headPicUrl"`
}

type UserListItem struct {
	UserID      int    `json:"userID"`
	WorkNo      string `json:"workNo"`
	UserType    int    `json:"userType"`
	DisplayName string `json:"displayName"`
	LoginName   string `json:"loginName" description:"账号"`
	RoleID      int    `json:"roleID"`
	RoleName    string `json:"roleName"`
	PhoneTel    string `json:"phoneTel"`
	Tel         string `json:"tel"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	HeadPicURL  string `json:"headPicUrl"`
	JobTitle    string `json:"jobTitle"`
	BirthDate   string `json:"birthDate"`
	InsertTime  string `json:"insertTime"`
}

type UserList struct {
	List  []UserListItem `json:"list"`
	Count int            `json:"count"`
}

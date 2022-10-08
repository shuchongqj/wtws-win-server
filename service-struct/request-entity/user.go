package request_entity

type UserLogin struct {
	LoginName string `json:"loginName" validate:"required"` //账号
	Password  string `json:"password" validate:"required"`  //密码
}

type UserUpdatePwd struct {
	OldPassWord string `json:"oldPassWord" validate:"required"` //账号
	Password    string `json:"password" validate:"required"`    //密码
}

type UserRoleInfoRequest struct {
	UrID   string `json:"urID" `                      // 用户权限标识
	UserID string `json:"userID" `                    // 用户ID
	RoleID string `json:"roleID" validate:"required"` // 角色ID
}

type InsertUser struct {
	UserID           string                `json:"userID"   `                              // 系统用户标识
	WorkNo           string                `json:"workNo"      validate:"required"`        // 工号
	UserLoginAuth    string                `json:"userLoginAuth"      validate:"required"` // 登录权限
	LoginName        string                `json:"loginName"   validate:"required" `       // 系统用户编号
	DisplayName      string                `json:"displayName" validate:"required" `       // 显示名称
	PhoneTel         string                `json:"phoneTel"    validate:"required" `       // 手机号码
	UserRoleInfoList []UserRoleInfoRequest `json:"userRoleInfoList"  validate:"required"`
}

type UserList struct {
	RoleID      int    `json:"roleID"`
	DisplayName string `json:"displayName"`
	LoginName   string `json:"loginName"`
	PhoneTel    string `json:"phoneTel"`
	WorkNo      string `json:"workNo"`
	IsExport    int8   `json:"isExport"`
	PageNum     int    `json:"pageNum" validate:"required"`
	PageSize    int    `json:"pageSize" validate:"required"`
}

//type Stations struct {
//	StationID   int    `json:"stationID"`
//	StationName string `json:"stationName"`
//}

type AddUserInfo struct {
	WorkNo      string `json:"workNo" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	LoginName   string `json:"loginName" validate:"required"`
	RoleID      int    `json:"roleID" validate:"required"`
	//Stations    []Stations `json:"stations" validate:"required"`
	UserType  int    `json:"user_type" validate:"required"`
	PhoneTel  string `json:"phoneTel"`
	Tel       string `json:"tel"`
	BirthDate string `json:"birthDate"`
	Gender    int    `json:"gender"`
	JobTitle  string `json:"jobTitle"`
	Email     string `json:"email"`
}

type DeleteUser struct {
	UserIDs []int `json:"userIDs"`
}

type UpdateUser struct {
	UserID      int    `json:"userID" validate:"required"`
	WorkNo      string `json:"workNo" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	LoginName   string `json:"loginName" validate:"required"`
	PhoneTel    string `json:"phoneTel" validate:"required"`
	Tel         string `json:"tel" validate:"required"`
	BirthDate   string `json:"birthDate" validate:"required"`
	Gender      int    `json:"gender" validate:"required"`
	JobTitle    string `json:"jobTitle" validate:"required"`
	Email       string `json:"email" validate:"required"`
	RoleID      int    `json:"roleID" validate:"required"`
	//Stations    []Stations `json:"stations" validate:"required"`
	UserType int `json:"user_type" validate:"required"`
}

type UpdateUserLoginName struct {
	UserID    int    `json:"userID"`
	LoginName string `json:"loginName"`
}

type RestUserPwd struct {
	UserID int `json:"userID"`
}

type UpdateUserRole struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
}

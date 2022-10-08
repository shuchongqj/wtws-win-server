package service

import (
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func GetUserEnterprise(userId int) common_struct.ResponseStruct {

	enterprises := wtws_mysql.GetUserEnterprise(userId)
	return common.ResponseStatus(0, "", enterprises)
}

package common

import (
	"encoding/json"
	common_struct "wtws-server/common/common-struct"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func ResponseStatus(_code int16, _message string, _result interface{}) common_struct.ResponseStruct {
	result := common_struct.ResponseStructMap[_code]

	if len(_message) > 0 {
		result.Message = _message
	}
	if _result != nil {
		result.Result = _result
	}

	return result
}

// GetContextUserInfo 获取上下文中的userInfo
func GetContextUserInfo(userInfoInterface interface{}) (user *wtws_mysql.SUser, err error) {
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfoInterface); err != nil {
		return nil, err
	}

	user = &wtws_mysql.SUser{}

	if err = json.Unmarshal(userInfoBytes, &user); err != nil {
		return nil, err
	}

	return user, err

}

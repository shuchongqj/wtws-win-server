package service

import (
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func GetAllCategory() common_struct.ResponseStruct {

	if categoryList, err := wtws_mysql.GetAllCategory(); err != nil {
		return common.ResponseStatus(0, "", dto.CategoryList{
			List:  []wtws_mysql.GCategory{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.CategoryList{
			List:  categoryList,
			Count: len(categoryList),
		})
	}

}

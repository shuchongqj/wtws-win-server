package service

import (
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetAllCargoto() common_struct.ResponseStruct {

	return common.ResponseStatus(0, "", wtws_mysql.GetAllCargoto())

}

func GetCargotoList(data *request_entity.CargotoList) common_struct.ResponseStruct {

	if cargotoList, count, err := wtws_mysql.GetCargotoList(data.PageNum, data.PageSize, data.CargotoName, data.Code); err != nil {
		return common.ResponseStatus(0, "", dto.CargotoList{
			List:  []wtws_mysql.OCargoto{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.CargotoList{
			List:  cargotoList,
			Count: count,
		})
	}

}

func AddCargoto(data *request_entity.AddCargoto) common_struct.ResponseStruct {

	if id, err := wtws_mysql.AddOCargoto(&wtws_mysql.OCargoto{
		Name:       data.CargotoName,
		Code:       data.Code,
		IsDelete:   conf.UN_DELETE,
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	}); err != nil || id <= 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", nil)
}

func DeleteCargoto(data *request_entity.DeleteCargoto) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.CargotoIDs))

	deleteErrs := []error{}

	for _, id := range data.CargotoIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByCargotoId(&wtws_mysql.OCargoto{
				Id:         id,
				IsDelete:   conf.IS_DELETE,
				UpdateTime: time.Now(),
			}, []string{"IsDelete", "UpdateTime"}); err != nil {
				deleteErrs = append(deleteErrs, err)
			}
			wg.Done()
		}(id)
	}

	wg.Wait()

	if len(deleteErrs) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

func UpdateCargoto(data *request_entity.UpdateCargoto) common_struct.ResponseStruct {

	if err := wtws_mysql.UpdateByCargotoId(&wtws_mysql.OCargoto{
		Id:         data.CargotoID,
		Name:       data.CargotoName,
		Code:       data.Code,
		UpdateTime: time.Now(),
	}, []string{"Name", "Code", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

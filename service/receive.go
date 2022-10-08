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

func GetAllReceiveList() common_struct.ResponseStruct {
	list := wtws_mysql.GetAllReceiveList()
	return common.ResponseStatus(0, "", dto.ReceiveList{list, len(list)})
}

func GetReceiveList(data *request_entity.ReceiveList) common_struct.ResponseStruct {

	if receiveList, count, err := wtws_mysql.GetReceiveList(data.PageNum, data.PageSize, data.Type, data.ReceiveName, data.ContactName, data.Tel, data.Address); err != nil {
		return common.ResponseStatus(0, "", dto.ReceiveList{
			List:  []wtws_mysql.OReceive{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.ReceiveList{
			List:  receiveList,
			Count: count,
		})
	}

}

func AddReceive(data *request_entity.AddReceive) common_struct.ResponseStruct {

	if id, err := wtws_mysql.AddOReceive(&wtws_mysql.OReceive{
		StationId:   conf.DEFAULT_STATION_ID,
		Name:        data.ReceiveName,
		ContactName: data.ContactName,
		Tel:         data.Tel,
		Address:     data.Address,
		Type:        int8(data.Type),
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		IsDelete:    conf.UN_DELETE,
		InsertTime:  time.Now(),
		UpdateTime:  time.Now(),
	}); err != nil || id <= 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", nil)
}

func DeleteReceive(data *request_entity.DeleteReceive) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.ReceiveIDs))

	deleteErrs := []error{}

	for _, id := range data.ReceiveIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByReceiveId(&wtws_mysql.OReceive{
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

func UpdateReceive(data *request_entity.UpdateReceive) common_struct.ResponseStruct {

	if err := wtws_mysql.UpdateByReceiveId(&wtws_mysql.OReceive{
		Id:          data.ReceiveID,
		Name:        data.ReceiveName,
		ContactName: data.ContactName,
		Tel:         data.Tel,
		Address:     data.Address,
		Type:        int8(data.Type),
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		IsDelete:    conf.UN_DELETE,
		UpdateTime:  time.Now(),
	}, []string{"Name", "ContactName", "Tel", "Address", "Type", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

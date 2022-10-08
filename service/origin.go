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

func GetAllOriginList() common_struct.ResponseStruct {
	list := wtws_mysql.GetAllOriginList()
	return common.ResponseStatus(0, "", dto.OriginList{list, len(list)})
}

func GetOriginList(data *request_entity.OriginList) common_struct.ResponseStruct {

	if originList, count, err := wtws_mysql.GetOriginList(data.PageNum, data.PageSize, data.Type, data.OriginName, data.ContactName, data.Tel, data.Address); err != nil {
		return common.ResponseStatus(0, "", dto.OriginList{
			List:  []wtws_mysql.OOrigin{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.OriginList{
			List:  originList,
			Count: count,
		})
	}

}

func AddOrigin(data *request_entity.AddOrigin) common_struct.ResponseStruct {

	if id, err := wtws_mysql.AddOOrigin(&wtws_mysql.OOrigin{
		StationId:   conf.DEFAULT_STATION_ID,
		Name:        data.OriginName,
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

func DeleteOrigin(data *request_entity.DeleteOrigin) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.OriginIDs))

	deleteErrs := []error{}

	for _, id := range data.OriginIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByOriginId(&wtws_mysql.OOrigin{
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

func UpdateOrigin(data *request_entity.UpdateOrigin) common_struct.ResponseStruct {

	if err := wtws_mysql.UpdateByOriginId(&wtws_mysql.OOrigin{
		Id:          data.OriginID,
		Name:        data.OriginName,
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

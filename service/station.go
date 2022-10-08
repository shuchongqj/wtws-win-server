package service

import (
	"strconv"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetUserStation(userId int) common_struct.ResponseStruct {
	station := wtws_mysql.GetUserStation(userId)
	return common.ResponseStatus(0, "", station)
}

func GetAllStation(userId int) common_struct.ResponseStruct {
	station := wtws_mysql.GetUserStation(userId)
	return common.ResponseStatus(0, "", station)
}

func GetStationList(data *request_entity.StationList, userId int) common_struct.ResponseStruct {

	stationList, count := wtws_mysql.GetStationList(
		userId, data.Name,
		data.Address,
		data.ContactPerson,
		data.ContactTel,
		data.EnterpriseName,
		data.EnterpryTel,
		data.PageNum,
		data.PageSize)

	return common.ResponseStatus(0, "", dto.StationList{
		List:  stationList,
		Count: count,
	})

}

func UpdateStation(data *request_entity.UpdateStation, userId int) common_struct.ResponseStruct {

	if checkUserStation := wtws_mysql.CheckUserStation(userId, data.StationID); !checkUserStation {
		return common.ResponseStatus(-9, "", nil)
	}

	if updateErr := wtws_mysql.UpdateStationByID(&wtws_mysql.OStation{
		Id:            data.StationID,
		Name:          data.StationName,
		Address:       data.StationAddress,
		ContactPerson: data.ContactPerson,
		ContactTel:    data.ContactTel,
		Longitude:     data.Longitude,
		Latitude:      data.Latitude,
		IsDelete:      conf.UN_DELETE,
		UpdateTime:    time.Now(),
		Province:      data.Province,
		City:          data.City,
		Area:          data.Area,
	}, []string{"Name", "Address", "ContactPerson", "ContactTel", "Longitude", "Latitude", "UpdateTime", "Province", "City", "Area"}); updateErr != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	if updateStationEnterpriseErr := wtws_mysql.UpdateStationEnterprise(data.StationID, data.EnterpriseIDs); updateStationEnterpriseErr != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)

}

func AddStation(data *request_entity.AddStation) common_struct.ResponseStruct {
	var stationID int64
	var addStationErr error

	if stationID, addStationErr = wtws_mysql.AddOStation(&wtws_mysql.OStation{
		Name:          data.Name,
		Address:       data.Address,
		ContactPerson: data.ContactPerson,
		ContactTel:    data.ContactTel,
		Longitude:     data.Longitude,
		Latitude:      data.Latitude,
		IsDelete:      conf.UN_DELETE,
		InsertTime:    time.Now(),
		UpdateTime:    time.Now(),
		Province:      data.Province,
		City:          data.City,
		Area:          data.Area,
	}); addStationErr != nil || stationID <= 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(data.EnterpriseIDs))

	addOStationEnterpriseErrs := []error{}

	for _, enterpriseID := range data.EnterpriseIDs {
		go func() {
			if _, addOStationEnterpriseErr := wtws_mysql.AddOStationEnterprise(&wtws_mysql.OStationEnterprise{
				StationId:    int(stationID),
				EnterpriseId: enterpriseID,
				InsertTime:   time.Now(),
				UpdateTime:   time.Now(),
			}); addOStationEnterpriseErr != nil {
				addOStationEnterpriseErrs = append(addOStationEnterpriseErrs, addOStationEnterpriseErr)
			}

			wg.Done()
		}()

	}

	wg.Wait()

	if len(addOStationEnterpriseErrs) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	adminIdStr := conf.ADMIN_ROLE_ID
	adminID, _ := strconv.Atoi(adminIdStr)
	wtws_mysql.AddSUserStation(&wtws_mysql.SUserStation{
		StationId:  int(stationID),
		UserId:     adminID,
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	})

	return common.ResponseStatus(12, "", nil)
}

func DeleteStation(data *request_entity.DeleteStation) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.StationIDs))

	pseudoDeleteErr := []error{}

	for _, stationID := range data.StationIDs {
		go func(stationID int) {
			if updateErr := wtws_mysql.UpdateStationByID(&wtws_mysql.OStation{
				Id:         stationID,
				IsDelete:   conf.IS_DELETE,
				UpdateTime: time.Now(),
			}, []string{"IsDelete", "UpdateTime"}); updateErr != nil {
				pseudoDeleteErr = append(pseudoDeleteErr, updateErr)
			}

			wg.Done()

		}(stationID)
	}

	wg.Wait()

	if len(pseudoDeleteErr) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

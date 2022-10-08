package service

import (
	"github.com/beego/beego/v2/core/logs"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetAllDriver() common_struct.ResponseStruct {
	return common.ResponseStatus(0, "", wtws_mysql.GetAllDriver())
}

func GetDriverList(data *request_entity.DriverList) common_struct.ResponseStruct {

	if driverList, count, err := wtws_mysql.GetDriverList(data.PageNum, data.PageSize, data.VehicleNumber, data.DriverName, data.IDCardNo, data.Tel, data.BankUserName, data.BankNo); err != nil {
		return common.ResponseStatus(0, "", dto.DriverList{
			List:  []wtws_mysql.SDriverList{},
			Count: 0,
		})
	} else {

		//for i := range driverList {
		//	driverList[i].DriverPass = "******"
		//}

		return common.ResponseStatus(0, "", dto.DriverList{
			List:  driverList,
			Count: count,
		})
	}

}

func checkAddDriverData(data *request_entity.AddDriver) (bool, bool, bool) {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	var checkDriverLoginNameOk, checkDriverVehicleNumOk, checkDriverTelOk bool

	go func() {
		checkDriverLoginNameOk = wtws_mysql.CheckRegistUserInfo(data.WorkNo, data.DriverAccount)
		wg.Done()
	}()

	go func() {
		checkDriverVehicleNumOk = wtws_mysql.CheckAddDriverVehicleNum(data.VehicleNumber)
		wg.Done()
	}()

	go func() {
		checkDriverTelOk = wtws_mysql.CheckAddDriverTel(data.Tel)
		wg.Done()
	}()

	wg.Wait()

	return checkDriverLoginNameOk, checkDriverVehicleNumOk, checkDriverTelOk
}

func AddDriver(data *request_entity.AddDriver) common_struct.ResponseStruct {

	checkDriverLoginNameOk, checkDriverVehicleNumOk, checkDriverTelOk := checkAddDriverData(data)
	if !checkDriverLoginNameOk {
		return common.ResponseStatus(-6, "该账号/工号已经被使用，请重新输入", nil)
	}

	if !checkDriverVehicleNumOk {
		return common.ResponseStatus(-6, "该 车牌 已经被使用，请重新输入", nil)
	}

	if !checkDriverTelOk {
		return common.ResponseStatus(-6, "该 司机联系电话 已经被使用，请重新输入", nil)
	}

	driverUser := wtws_mysql.SUser{
		LoginName:   data.DriverAccount,
		PassWord:    conf.DEFAULT_PASS_WORD,
		WorkNo:      data.WorkNo,
		BirthDate:   "1000-01-01",
		UserType:    conf.USER_DRIVER_DEFAULT_TYPE,
		Status:      conf.DEFAULT_USER_STATUS,
		IsDelete:    conf.UN_DELETE,
		DisplayName: data.DriverName,
		Gender:      int8(data.Gender),
		PhoneTel:    data.Tel,
		Tel:         "",
		JobTitle:    conf.USER_DRIVER_TITLE,
		Email:       data.Email,
		HeadPicUrl:  "",
		InsertTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	if len(data.BirthDate) > 0 {
		var birthDate time.Time
		var birthDateErr error
		if birthDate, birthDateErr = time.ParseInLocation("2006-01-02T15:04:05Z", data.BirthDate, time.Local); birthDateErr != nil {
			logs.Error("[service]  解析新增司机用户的出生日期失败，失败信息:", birthDateErr.Error())
		} else {
			driverUser.BirthDate = birthDate.Format("2006-01-02")
		}
	}

	var userId int64
	var addUserErr error
	if userId, addUserErr = wtws_mysql.AddSUser(&driverUser); addUserErr != nil || userId <= 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	wg2 := &sync.WaitGroup{}
	wg2.Add(3)
	var addDriverErr, addSUserRoleErr, addSUserStationErr error

	go func() {
		_, addDriverErr = wtws_mysql.AddSDriver(&wtws_mysql.SDriver{
			UserID:         int(userId),
			DriverName:     data.DriverName,
			IdCardNo:       data.IDCardNo,
			VehicleNumber:  data.VehicleNumber,
			LimitTotalLoad: float32(data.LimitTotalLoad),
			Length:         data.Length,
			BankUserName:   data.BankUserName,
			BankName:       data.BankName,
			BankNo:         data.BankNo,

			Tel:        data.Tel,
			IsValid:    conf.IS_VALID,
			IsDelete:   conf.UN_DELETE,
			InsertTime: time.Now(),
			UpdateTime: time.Now(),
		})
		wg2.Done()
	}()

	go func() {

		_, addSUserRoleErr = wtws_mysql.AddSUserRole(&wtws_mysql.SUserRole{
			UserId:     int(userId),
			RoleId:     conf.USER_DRIVER_ROLE_ID,
			InsertTime: time.Now(),
			UpdateTime: time.Now(),
		})
		wg2.Done()
	}()

	go func() {
		_, addSUserStationErr = wtws_mysql.AddSUserStation(&wtws_mysql.SUserStation{
			StationId:  conf.DEFAULT_STATION_ID,
			UserId:     int(userId),
			InsertTime: time.Time{},
			UpdateTime: time.Time{},
		})
		wg2.Done()
	}()

	wg2.Wait()

	if addDriverErr != nil || addSUserStationErr != nil || addSUserRoleErr != nil {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", nil)
}

func DeleteDriver(data *request_entity.DeleteDriver) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.DriverIDs))

	deleteErrs := []error{}

	for _, id := range data.DriverIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateSDriverById(&wtws_mysql.SDriver{
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

func UpdateDriver(data *request_entity.UpdateDriver) common_struct.ResponseStruct {

	var checkDriverVehicleNumOk, checkDriverTelOk bool
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		checkDriverVehicleNumOk = wtws_mysql.CheckOtherVehicleNum(data.DriverID, data.VehicleNumber)
		wg.Done()
	}()
	go func() {
		checkDriverTelOk = wtws_mysql.CheckOtherDriverTel(data.DriverID, data.Tel)
		wg.Done()
	}()
	wg.Wait()
	if !checkDriverVehicleNumOk {
		return common.ResponseStatus(-15, "该 车牌 已经被使用，请重新输入", nil)
	}
	if !checkDriverTelOk {
		return common.ResponseStatus(-15, "该 司机联系电话 已经被使用，请重新输入", nil)
	}

	var userID int
	//判断司机是否有关联的用户
	if data.UpdateUserInfo {
		if driver, _ := wtws_mysql.GetSDriverById(data.DriverID); driver.UserID == 0 && len(data.DriverAccount) > 0 {
			//当司机没有关联用户，并且request请求中有账号信息时，才新增用户并关联

			if !wtws_mysql.CheckRegistUserInfo(data.WorkNo, data.DriverAccount) {
				return common.ResponseStatus(-15, "该 账号/工号 已经被使用，请重新输入", nil)
			}

			driverUser := wtws_mysql.SUser{
				LoginName:   data.DriverAccount,
				PassWord:    conf.DEFAULT_PASS_WORD,
				WorkNo:      data.WorkNo,
				UserType:    conf.USER_DRIVER_DEFAULT_TYPE,
				Status:      conf.DEFAULT_USER_STATUS,
				IsDelete:    conf.UN_DELETE,
				DisplayName: data.DriverName,
				Gender:      int8(data.Gender),
				BirthDate:   "1000-01-01",
				PhoneTel:    data.Tel,
				Tel:         "",
				JobTitle:    conf.USER_DRIVER_TITLE,
				Email:       data.Email,
				HeadPicUrl:  "",
				InsertTime:  time.Now(),
				UpdateTime:  time.Now(),
			}

			if len(data.BirthDate) > 0 {
				var birthDate time.Time
				var birthDateErr error
				if birthDate, birthDateErr = time.ParseInLocation("2006-01-02T15:04:05Z", data.BirthDate, time.Local); birthDateErr != nil {
					logs.Error("[service]  解析新增司机用户的出生日期失败，失败信息:", birthDateErr.Error())
				} else {
					driverUser.BirthDate = birthDate.Format("2006-01-02")
				}
			}

			var userId int64
			var addUserErr error

			if userId, addUserErr = wtws_mysql.AddSUser(&driverUser); addUserErr != nil || userId <= 0 {
				return common.ResponseStatus(500, "", nil)
			}

			userID = int(userId)

			wg1 := &sync.WaitGroup{}
			wg1.Add(2)

			var addSUserRoleErr, addSUserStationErr error

			go func() {

				_, addSUserRoleErr = wtws_mysql.AddSUserRole(&wtws_mysql.SUserRole{
					UserId:     int(userId),
					RoleId:     conf.USER_DRIVER_ROLE_ID,
					InsertTime: time.Now(),
					UpdateTime: time.Now(),
				})
				wg1.Done()
			}()

			go func() {
				_, addSUserStationErr = wtws_mysql.AddSUserStation(&wtws_mysql.SUserStation{
					StationId:  conf.DEFAULT_STATION_ID,
					UserId:     int(userId),
					InsertTime: time.Time{},
					UpdateTime: time.Time{},
				})
				wg1.Done()
			}()

			wg1.Wait()

			if addSUserStationErr != nil || addSUserRoleErr != nil {
				return common.ResponseStatus(500, "", nil)
			}

		} else if driver.UserID != 0 {

			//如果司机已经关联了用户信息，则分情况修改用户信息

			//若账号为空，删除用户信息
			if len(data.DriverAccount) == 0 {
				userID = 0
				if updateDriverUserErr := wtws_mysql.UpdateSUserById(&wtws_mysql.SUser{
					Id:         driver.UserID,
					IsDelete:   conf.IS_DELETE,
					UpdateTime: time.Now(),
				}, []string{"IsDelete", "UpdateTime"}); updateDriverUserErr != nil {
					return common.ResponseStatus(-15, "", nil)
				}
			} else if !wtws_mysql.CheckOtherUserByWorkNoAndLoginName(driver.UserID, data.WorkNo, data.DriverAccount) {
				//若账号不为空，判断账号和工号是否已经存在
				return common.ResponseStatus(-15, "该账号/工号已经被使用，请重新输入", nil)
			} else {
				userID = driver.UserID
				//若司机存在关联账号并且账号不为空，工号和账号合规，则修改用户信息
				driverUserData := &wtws_mysql.SUser{
					Id:          driver.UserID,
					LoginName:   data.DriverAccount,
					WorkNo:      data.WorkNo,
					DisplayName: data.DriverName,
					Gender:      int8(data.Gender),
					BirthDate:   "1000-01-01",
					PhoneTel:    data.Tel,
					Email:       data.Email,
					UpdateTime:  time.Now(),
				}
				if len(data.BirthDate) > 0 {
					var birthDate time.Time
					var birthDateErr error
					if birthDate, birthDateErr = time.ParseInLocation("2006-01-02T15:04:05Z", data.BirthDate, time.Local); birthDateErr == nil {
						driverUserData.BirthDate = birthDate.Format("2006-01-02")
					} else {
						logs.Error("[service]  解析新增司机用户的出生日期失败，失败信息:", birthDateErr.Error())
					}
				}

				if updateDriverUserErr := wtws_mysql.UpdateSUserById(driverUserData, []string{
					"LoginName", "WorkNo", "DisplayName", "Gender", "BirthDate", "PhoneTel", "Email"}); updateDriverUserErr != nil {
					return common.ResponseStatus(-15, "", nil)
				}
			}

		}
	}

	if data.UpdateDriverInfo {
		updateDriverData := &wtws_mysql.SDriver{
			UserID:         userID,
			DriverName:     data.DriverName,
			IdCardNo:       data.IDCardNo,
			VehicleNumber:  data.VehicleNumber,
			LimitTotalLoad: float32(data.LimitTotalLoad),
			Length:         data.Length,
			BankUserName:   data.BankUserName,
			BankName:       data.BankName,
			BankNo:         data.BankNo,
			Id:             data.DriverID,
			Tel:            data.Tel,
			UpdateTime:     time.Now(),
		}
		if err := wtws_mysql.UpdateSDriverById(updateDriverData, []string{
			"DriverName", "UserID", "IdCardNo", "VehicleNumber", "LimitTotalLoad", "Length",
			"BankUserName", "BankName", "BankNo", "Tel", "UpdateTime"}); err != nil {
			return common.ResponseStatus(-15, "", nil)
		}
	}

	return common.ResponseStatus(14, "", nil)
}

func DriverTruckOrderList(pageNum int, pageSize int, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	if userInfo.UserType == conf.USER_MANAGER_DEFAULT_TYPE {
		return GetTruckOrderList(&request_entity.TruckOrderList{
			PageNum:  pageNum,
			PageSize: pageSize,
			Status:   conf.TRUCK_ORDER_STATUS_PASS,
		})
	} else if dbDriver, getDriverErr := wtws_mysql.GetSDriverByUserID(userInfo.Id); getDriverErr != nil || dbDriver == nil {
		return common.ResponseStatus(0, "", dto.TruckOrderList{})
	} else {
		truckOrderList, _ := wtws_mysql.GetCheckedTruckOrderByVehicleNumber(dbDriver.VehicleNumber, conf.TRUCK_ORDER_STATUS_PASS)
		return common.ResponseStatus(0, "", dto.TruckOrderList{
			List:  truckOrderList,
			Count: len(truckOrderList),
		})
	}
}

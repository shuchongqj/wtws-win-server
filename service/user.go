package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/wangcong0918/sunrise/utils/jwt"
	"regexp"
	"strconv"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	"wtws-server/models"
	wtws_mongodb "wtws-server/models/wtws-mongodb"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func Login(data request_entity.UserLogin) common_struct.ResponseStruct {
	loginName := data.LoginName
	password := data.Password

	var tokenString = ""
	var userInfo *wtws_mysql.SUser
	var err error

	if userInfo, err = wtws_mysql.GetLoginUserInfo(loginName, password); err != nil {
		return common.ResponseStatus(-2, "", nil)

	}
	//判断用户密码是否正确
	jwtInfo := new(jwt.User)
	jwtInfo.UserID = fmt.Sprintf("%d", userInfo.Id)
	token, jwtErr := jwt.JwtGenerateToken(jwtInfo, 720*time.Hour)
	if jwtErr != nil {
		return common.ResponseStatus(500, "", nil)
	}

	tokenString = token

	wg := sync.WaitGroup{}
	wg.Add(2)

	var setUserRedisErr, insertUserLoginLogErr error

	go func() {
		setUserRedisErr = models.RedisDelAndSet(fmt.Sprintf("%d", userInfo.Id), tokenString) //添加新的token 有效期一个月
		wg.Done()
	}()

	go func() {
		insertUserLoginLogErr = wtws_mongodb.InsertUserLoginLog(userInfo.Id, loginName, userInfo.WorkNo)
		wg.Done()
	}()

	wg.Wait()

	if setUserRedisErr != nil || insertUserLoginLogErr != nil {
		return common.ResponseStatus(500, "", nil)
	} else {
		return common.ResponseStatus(0, "", dto.Login{
			Authorization: tokenString,
		})
	}
}

// UpdatePwd 更新用户密码
func UpdatePwd(data request_entity.UserUpdatePwd, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	oldPassWord := data.OldPassWord
	password := data.Password

	if oldPassWord != userInfo.PassWord {
		return common.ResponseStatus(-4, "", nil)
	}

	userInfo.PassWord = password

	if err := wtws_mysql.UpdateById(userInfo, []string{"PassWord"}); err != nil {
		return common.ResponseStatus(-4, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func GetUserInfo(userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	userInfoDTO := dto.UserInfo{
		UserId:      userInfo.Id,
		LoginName:   userInfo.LoginName,
		WorkNo:      userInfo.WorkNo,
		DisplayName: userInfo.DisplayName,
		JobTitle:    userInfo.JobTitle,
		HeadPicUrl:  userInfo.HeadPicUrl,
	}

	var userRole *wtws_mysql.SUserRole
	var err error

	if userRole, err = wtws_mysql.GetSUserRoleByUserId(userInfo.Id); err != nil && userRole == nil {
		return common.ResponseStatus(-8, "", nil)
	}

	userInfoDTO.RoleID = userRole.RoleId

	return common.ResponseStatus(0, "", userInfoDTO)

}

func GetUserList(data *request_entity.UserList) common_struct.ResponseStruct {

	if userListParams, count, getUserListErr := wtws_mysql.GetUserList(data); getUserListErr != nil {
		return common.ResponseStatus(500, "", nil)
	} else {

		userList := []dto.UserListItem{}

		for _, itemParam := range userListParams {
			userListItem := dto.UserListItem{
				WorkNo:      itemParam["work_no"].(string),
				DisplayName: itemParam["display_name"].(string),
				LoginName:   itemParam["login_name"].(string),
				RoleName:    itemParam["role_name"].(string),
				PhoneTel:    itemParam["phone_tel"].(string),
				Tel:         itemParam["tel"].(string),
				Email:       itemParam["email"].(string),
				HeadPicURL:  itemParam["head_pic_url"].(string),
				JobTitle:    itemParam["job_title"].(string),
				BirthDate:   itemParam["birth_date"].(string),
				InsertTime:  itemParam["insert_time"].(string),
			}

			var atoiErr error
			userListItem.UserID, _ = strconv.Atoi(itemParam["user_id"].(string))
			if userListItem.UserType, atoiErr = strconv.Atoi(itemParam["user_type"].(string)); atoiErr != nil {
				userListItem.UserType = conf.USER_MANAGER_DEFAULT_TYPE
			}

			var genderType int
			if genderType, atoiErr = strconv.Atoi(itemParam["gender"].(string)); atoiErr != nil {
				genderType = 3
			}

			userListItem.Gender = conf.GENDER_ENUM[genderType]

			userListItem.RoleID, _ = strconv.Atoi(itemParam["role_id"].(string))

			userList = append(userList, userListItem)
		}

		return common.ResponseStatus(0, "", dto.UserList{
			List:  userList,
			Count: count,
		})
	}
}

func AddUser(data *request_entity.AddUserInfo) common_struct.ResponseStruct {

	if !wtws_mysql.CheckRegistUserInfo(data.WorkNo, data.LoginName) {
		return common.ResponseStatus(-6, "当前工号和账号对应的用户已存在", nil)
	}

	birthDateStr := "0001-01-01"
	if len(data.BirthDate) > 0 {
		var birthDate time.Time
		var birthDateErr error
		if birthDate, birthDateErr = time.ParseInLocation("2006-01-02T15:04:05Z", data.BirthDate, time.Local); birthDateErr != nil {
			logs.Error("[service]  解析新增用户的出生日期失败，失败信息:", birthDateErr.Error())
			return common.ResponseStatus(-99, "", nil)
		}
		birthDateStr = birthDate.Format("2006-01-02")
	}

	var addUserID int64
	var err error

	addUserInfo := wtws_mysql.SUser{
		LoginName:   data.LoginName,
		PassWord:    conf.DEFAULT_PASS_WORD,
		WorkNo:      data.WorkNo,
		UserType:    conf.USER_MANAGER_DEFAULT_TYPE,
		Status:      conf.DEFAULT_USER_STATUS,
		IsDelete:    conf.UN_DELETE,
		DisplayName: data.DisplayName,
		Gender:      int8(data.Gender),
		BirthDate:   birthDateStr,
		PhoneTel:    data.PhoneTel,
		Tel:         data.Tel,
		JobTitle:    data.JobTitle,
		Email:       data.Email,
		InsertTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if addUserID, err = wtws_mysql.AddSUser(&addUserInfo); err != nil || addUserID <= 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	userRole := wtws_mysql.SUserRole{
		UserId:     int(addUserID),
		RoleId:     data.RoleID,
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if userRoleID, addUserRoleErr := wtws_mysql.AddSUserRole(&userRole); addUserRoleErr != nil || userRoleID <= 0 {
		logs.Error("[service]  给用户关联权限失败，失败信息:", addUserRoleErr.Error())
		return common.ResponseStatus(500, "", nil)
	}

	return common.ResponseStatus(12, "", addUserID)

}

func DeleteUser(data *request_entity.DeleteUser) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	pseudoDeleteUserErrs := []error{}

	go func() {
		subWg := &sync.WaitGroup{}
		subWg.Add(len(data.UserIDs))
		for _, userId := range data.UserIDs {
			go func(userId int) {
				if updateErr := wtws_mysql.UpdateById(&wtws_mysql.SUser{
					Id:       userId,
					IsDelete: conf.IS_DELETE,
				}, []string{"IsDelete"}); updateErr != nil {
					pseudoDeleteUserErrs = append(pseudoDeleteUserErrs, updateErr)
				}

				wtws_mysql.FindAndUpdateUserDriver(userId)

				subWg.Done()
			}(userId)

		}
		subWg.Wait()

		wg.Done()
	}()

	go func() {

		subWg2 := &sync.WaitGroup{}
		subWg2.Add(len(data.UserIDs))

		for _, userId := range data.UserIDs {
			go func(userId int) {
				models.RedisDel(fmt.Sprintf("%s", userId))
				subWg2.Done()
			}(userId)
		}

		subWg2.Wait()

		wg.Done()
	}()

	wg.Wait()

	if len(pseudoDeleteUserErrs) > 0 {
		return common.ResponseStatus(500, "", nil)
	}

	return common.ResponseStatus(16, "", nil)

}

func UpdateUser(data *request_entity.UpdateUser) common_struct.ResponseStruct {

	if existUser, getUserErr := wtws_mysql.GetSUserById(data.UserID); getUserErr != nil || existUser == nil {
		return common.ResponseStatus(-5, "", nil)
	}

	var birthDateStr string

	exp1 := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}) \d{2}:\d{2}:\d{2}$`)
	requestUrlArr := exp1.FindStringSubmatch(data.BirthDate)
	if len(requestUrlArr) > 1 {
		birthDateStr = requestUrlArr[1]
	} else {
		if birthDate, birthDateErr := time.ParseInLocation("2006-01-02T15:04:05Z", data.BirthDate, time.Local); birthDateErr != nil {
			logs.Error("[service]  解析新增用户的出生日期失败，失败信息:", birthDateErr.Error())
			return common.ResponseStatus(-99, "", nil)
		} else {

			birthDateStr = birthDate.Format("2006-01-02")
		}
	}

	if err := wtws_mysql.UpdateById(&wtws_mysql.SUser{
		Id:          data.UserID,
		LoginName:   data.LoginName,
		WorkNo:      data.WorkNo,
		UserType:    int16(data.UserType),
		DisplayName: data.DisplayName,
		Gender:      int8(data.Gender),
		BirthDate:   birthDateStr,
		PhoneTel:    data.PhoneTel,
		Tel:         data.Tel,
		JobTitle:    data.JobTitle,
		Email:       data.Email,
		UpdateTime:  time.Now(),
	}, []string{"LoginName", "WorkNo", "UserType", "DisplayName", "Gender", "BirthDate", "PhoneTel", "Tel", "JobTitle", "Email", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	} else {
		return common.ResponseStatus(14, "", nil)
	}

}

func UpdateUserLoginName(data *request_entity.UpdateUserLoginName) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var updateUserLoginNameErr, delRedisUserAuthErr error

	go func() {
		updateUserLoginNameErr = wtws_mysql.UpdateById(&wtws_mysql.SUser{
			Id:         data.UserID,
			LoginName:  data.LoginName,
			UpdateTime: time.Now(),
		}, []string{"LoginName", "UpdateTime"})
		wg.Done()
	}()

	go func() {

		delRedisUserAuthErr = models.RedisDel(fmt.Sprintf("%d", data.UserID))

		wg.Done()
	}()

	wg.Wait()

	if updateUserLoginNameErr != nil {
		logs.Error("[service] 修改用户的登录账号失败，失败信息:", updateUserLoginNameErr.Error())
		return common.ResponseStatus(-15, "", nil)
	}

	if delRedisUserAuthErr != nil {
		logs.Error("[service]  删除用户的redis token信息失败，失败信息：", delRedisUserAuthErr.Error())
		return common.ResponseStatus(-16, "", nil)
	}

	return common.ResponseStatus(14, "", nil)

}

func RestUserPwd(data *request_entity.RestUserPwd, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var restUserPwdErr, delRedisUserAuthErr error

	go func() {
		restUserPwdErr = wtws_mysql.UpdateById(&wtws_mysql.SUser{
			Id:         data.UserID,
			PassWord:   conf.DEFAULT_PASS_WORD,
			UpdateTime: time.Now(),
		}, []string{"PassWord", "UpdateTime"})
		wg.Done()
	}()

	go func() {

		delRedisUserAuthErr = models.RedisDel(fmt.Sprintf("%d", data.UserID))

		wg.Done()
	}()

	wg.Wait()

	if restUserPwdErr != nil {
		logs.Error("[service] 重置用户密码失败，用户ID:", data.UserID, "\t失败信息:", restUserPwdErr.Error())
		return common.ResponseStatus(-15, "", nil)
	}

	if delRedisUserAuthErr != nil {
		logs.Error("[service]  删除用户的redis token信息失败，用户ID：", data.UserID, "\t失败信息：", delRedisUserAuthErr.Error())
		return common.ResponseStatus(-16, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func LogOut(userId int) common_struct.ResponseStruct {

	if delRedisErr := models.RedisDel(fmt.Sprintf("%d", userId)); delRedisErr != nil {
		return common.ResponseStatus(-16, "", nil)
	}
	return common.ResponseStatus(0, "", nil)
}

func UpdateUserRole(data *request_entity.UpdateUserRole) common_struct.ResponseStruct {
	if err := wtws_mysql.DeleteUserRoles(data.UserID); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	if userRoleID, err := wtws_mysql.AddSUserRole(&wtws_mysql.SUserRole{
		UserId:     data.UserID,
		RoleId:     data.RoleID,
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	}); err != nil || userRoleID <= 0 {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(0, "", nil)
}

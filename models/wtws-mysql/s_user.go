package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"
	request_entity "wtws-server/service-struct/request-entity"

	"github.com/astaxie/beego/orm"
)

type SUser struct {
	Id          int       `orm:"column(user_id);auto" description:"系统用户标识" json:"userID"`
	LoginName   string    `orm:"column(login_name);size(16)" description:"登录名称" json:"loginName"`
	PassWord    string    `orm:"column(pass_word);size(32)" description:"登录密码通过加密的二进制保存" json:"passWord"`
	WorkNo      string    `orm:"column(work_no);size(8)" description:"工号" json:"workNo"`
	UserType    int16     `orm:"column(user_type)" description:"登录权限 1-后台" json:"userType"`
	Status      int16     `orm:"column(status)" description:"账号状态 1-正常  2-禁用" json:"status"`
	IsDelete    int16     `orm:"column(is_delete)" description:"是否删除  1-未删除 2-删除" json:"isDelete"`
	DisplayName string    `orm:"column(display_name);size(32);null" description:"显示名称" json:"displayName"`
	Gender      int8      `orm:"column(gender);null" description:"性别 1-男 2-女 3-未知" json:"gender"`
	BirthDate   string    `orm:"column(birth_date);size(32);null" description:"生日" json:"birthDate"`
	PhoneTel    string    `orm:"column(phone_tel);size(16);null" description:"手机号码" json:"phoneTel"`
	Tel         string    `orm:"column(tel);size(16);null" description:"座机号码" json:"tel"`
	JobTitle    string    `orm:"column(job_title);size(32);null" description:"职位名字" json:"jobTitle"`
	Email       string    `orm:"column(email);size(50);null" description:"电子邮箱" json:"email"`
	HeadPicUrl  string    `orm:"column(head_pic_url);size(128);null" description:"头像" json:"headPicUrl"`
	InsertTime  time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间" json:"insertTime"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间" json:"updateTime"`
}

func (t *SUser) TableName() string {
	return "s_user"
}

func init() {
	orm.RegisterModel(new(SUser))
}

// AddSUser insert a new SUser into database and returns
// last inserted Id on success.
func AddSUser(m *SUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSUserById retrieves SUser by Id. Returns error if
// Id doesn't exist
func GetSUserById(id int) (v *SUser, err error) {
	o := orm.NewOrm()
	v = &SUser{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSUser retrieves all SUser matches certain condition. Returns empty list if
// no records exist
func GetAllSUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSUser updates SUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSUserById(m *SUser, cols []string) (err error) {
	o := orm.NewOrm()
	v := SUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, cols...); err == nil {
			fmt.Println("[mysql] Number of records s_user updated in database:", num)
		}
	}
	return

}

// DeleteSUser deletes SUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSUser(id int) (err error) {
	o := orm.NewOrm()
	v := SUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetLoginUserInfo 获取登录的用户信息
func GetLoginUserInfo(loginName string, password string) (v *SUser, err error) {
	o := orm.NewOrm()
	v = &SUser{LoginName: loginName, PassWord: password, Status: 1, IsDelete: 1}
	if err = o.Read(v, "LoginName", "PassWord", "Status", "IsDelete"); err != nil {
		return v, err
	}
	return v, nil
}

func UpdateById(u *SUser, cols []string) (err error) {
	o := orm.NewOrm()
	v := SUser{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of s_user records updated in database:", num)
		}
	}
	return err
}

// GetUserListInfo 获取用户列表
func GetUserList(userInfo *request_entity.UserList) (userListParams []orm.Params, countInfo int, err error) {

	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("su.*,sr.role_name,sr.role_id").
		From("s_user as su").
		InnerJoin("s_user_role as sur").
		On("sur.user_id = su.user_id").
		InnerJoin("s_role as sr").
		On("sr.role_id = sur.role_id")

	qd.Where("su.is_delete = 1")

	if userInfo.RoleID > 0 {
		qd.And(fmt.Sprintf("sur.role_id = %d", userInfo.RoleID))
	}

	if len(userInfo.DisplayName) > 0 {
		qd.And("su.display_name like '%" + userInfo.DisplayName + "%'")
	}

	if len(userInfo.LoginName) > 0 {
		qd.And("su.login_name like '%" + userInfo.LoginName + "%'")
	}

	if len(userInfo.PhoneTel) > 0 {
		qd.And("su.phone_tel like '%" + userInfo.PhoneTel + "%'")
	}

	if len(userInfo.WorkNo) > 0 {
		qd.And("su.work_no like '%" + userInfo.WorkNo + "%'")
	}

	qd.GroupBy("su.user_id").OrderBy("su.update_time desc")

	countSql := qd.String()

	if userInfo.IsExport == 1 {

	} else {
		qd.Limit(userInfo.PageSize).Offset((userInfo.PageNum - 1) * userInfo.PageSize)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	userList := []orm.Params{}
	userCount := 0
	var getUserListErr, getUserListCountErr error

	go func() {
		o := orm.NewOrm()
		sql := qd.String()
		if _, getUserListErr = o.Raw(sql).Values(&userList); getUserListErr != nil {
			logs.Error("[mysql]  查询用户列表数据失败，失败信息:", getUserListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		o := orm.NewOrm()
		countUserList := []orm.Params{}
		if _, getUserListCountErr = o.Raw(countSql).Values(&countUserList); getUserListCountErr != nil {
			logs.Error("[mysql]  查询用户列表总数失败，失败信息:", getUserListCountErr.Error())
		} else {
			userCount = len(countUserList)
		}
		wg.Done()
	}()

	wg.Wait()

	if getUserListErr != nil {
		return []orm.Params{}, 0, getUserListErr
	}
	if getUserListCountErr != nil {
		return []orm.Params{}, 0, getUserListCountErr
	}

	return userList, userCount, nil
}

func CheckRegistUserInfo(workNo string, loginName string) bool {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("s_user").Where("is_delete = 1").And("(work_no = ? OR login_name = ?)")
	sql := qd.String()
	if num, err := o.Raw(sql, workNo, loginName).QueryRows(&[]SUser{}); err != nil {
		logs.Error("[mysql]  查询workNo和loginName对应的用户信息失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func CheckRegistLoginName(loginName string) bool {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_user").
		Where("is_delete = 1").
		And("(login_name = ?)")
	sql := qd.String()
	if num, err := o.Raw(sql, loginName).QueryRows(&[]SUser{}); err != nil {
		logs.Error("[mysql]  查询loginName对应的用户信息失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false

}

func CheckOtherUserByWorkNoAndLoginName(userID int, workNo string, loginName string) bool {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_user").
		Where("is_delete = 1").
		And("user_id != ?").
		And("(work_no = ? OR login_name = ?)")

	sql := qd.String()
	if num, err := o.Raw(sql, userID, workNo, loginName).QueryRows(&[]SUser{}); err != nil {
		logs.Error("[mysql]  查询workNo和loginName对应的用户信息失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func FindAndUpdateUserDriver(userId int) error {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Update("s_driver").Set("user_id = null").Where("user_id = ?")
	sql := qd.String()

	if _, updateErr := o.Raw(sql, userId).Exec(); updateErr != nil {
		logs.Error("[mysql]  删除用户关联的司机user_id失败，失败信息：:", updateErr.Error())
		return updateErr
	} else {
		return nil
	}

}

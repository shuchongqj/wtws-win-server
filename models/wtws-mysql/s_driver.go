package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"
)

type SDriver struct {
	Id int `orm:"column(driver_id);auto" description:"id" json:"driverID"`
	//DriverAccount  string    `orm:"column(driver_account);size(32)" description:"司机账号" json:"driverAccount"`
	//DriverPass     string    `orm:"column(driver_pass);size(32)" description:"密码" json:"driverPass"`
	UserID         int       `orm:"column(user_id)" description:"关联的用户ID" json:"userID"`
	DriverName     string    `orm:"column(driver_name);size(32)" description:"司机名字" json:"driverName"`
	IdCardNo       string    `orm:"column(id_card_no);size(18)" description:"身份证号" json:"idCardNo"`
	Tel            string    `orm:"column(tel);size(11)" description:"电话" json:"tel"`
	VehicleNumber  string    `orm:"column(vehicle_number);size(16)" description:"车牌号" json:"vehicleNumber"`
	LimitTotalLoad float32   `orm:"column(limit_total_load);null" description:"车辆载重" json:"limitTotalLoad"`
	Length         string    `orm:"column(length);size(32);null" description:"车总长度" json:"length"`
	BankUserName   string    `orm:"column(bank_user_name);size(32);null" description:"银行卡开户人姓名" json:"bankUserName"`
	BankName       string    `orm:"column(bank_name);size(32);null" description:"银行" json:"bankName"`
	BankNo         string    `orm:"column(bank_no);size(200);null" description:"银行卡" json:"bankNo"`
	IsValid        int8      `orm:"column(is_valid)" description:"状态  1-启用  2-禁用" json:"isValid"`
	IsDelete       int8      `orm:"column(is_delete)" description:"是否删除  1-未删除  2-已删除" json:"isDelete"`
	UpdateTime     time.Time `orm:"column(update_time);type(datetime);null" description:"记录更新时间" json:"updateTime"`
	InsertTime     time.Time `orm:"column(insert_time);type(datetime);null" description:"记录插入时间" json:"insertTime"`
}

type SDriverList struct {
	SDriver
	LoginName string `json:"loginName"`
	WorkNo    string `json:"workNo"`
	BirthDate string `json:"birthDate"`
	Email     string `json:"email"`
}

func (t *SDriver) TableName() string {
	return "s_driver"
}

func init() {
	orm.RegisterModel(new(SDriver))
}

// AddSDriver insert a new SDriver into database and returns
// last inserted Id on success.
func AddSDriver(m *SDriver) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSDriverById retrieves SDriver by Id. Returns error if
// Id doesn't exist
func GetSDriverById(id int) (v *SDriver, err error) {
	o := orm.NewOrm()
	v = &SDriver{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSDriver retrieves all SDriver matches certain condition. Returns empty list if
// no records exist
func GetAllSDriver(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SDriver))
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

	var l []SDriver
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

// DeleteSDriver deletes SDriver by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSDriver(id int) (err error) {
	o := orm.NewOrm()
	v := SDriver{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SDriver{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetDriverList(
	pageNum, pageSize int,
	vehicleNumber, driverName, iDCardNo, tel, bankUserName, bankNo string) (list []SDriverList, count int, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("sd.*,su.login_name,su.work_no,su.birth_date,su.email").
		From("s_driver as sd").
		LeftJoin("s_user as su").
		On("sd.user_id = su.user_id").
		And("su.is_delete = 1").
		Where("sd.is_delete = 1")

	//if len(driverAccount) > 0 {
	//	qd.And("su.login_name LIKE '%" + driverAccount + "%'")
	//}

	if len(vehicleNumber) > 0 {
		qd.And("sd.vehicle_number LIKE '%" + vehicleNumber + "%'")
	}

	if len(tel) > 0 {
		qd.And("sd.tel LIKE '%" + tel + "%'")
	}

	if len(driverName) > 0 {
		qd.And("sd.driver_name LIKE '%" + driverName + "%'")
	}

	if len(iDCardNo) > 0 {
		qd.And("sd.id_card_no LIKE '%" + iDCardNo + "%'")
	}

	if len(bankNo) > 0 {
		qd.And("sd.bank_no LIKE '%" + bankNo + "%'")
	}
	if len(bankUserName) > 0 {
		qd.And("sd.bank_user_name LIKE '%" + bankUserName + "%'")
	}

	qd.OrderBy(" driver_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []SDriverList{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询司机失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]SDriverList{}); getCountErr != nil {
			logs.Error("[mysql]  查询司机失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []SDriverList{}, 0, getListErr
	}

	if getCountErr != nil {
		return []SDriverList{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateSDriverById(u *SDriver, cols []string) (err error) {
	o := orm.NewOrm()
	v := SDriver{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of s_driver records updated in database:", num)
		}
	}
	return err
}

func CheckAddDriverVehicleNum(vehicleNum string) bool {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And("vehicle_number = ? ")
	sql := qd.String()
	if num, err := o.Raw(sql, vehicleNum).QueryRows(&[]SDriver{}); err != nil {
		logs.Error("[mysql]  查询loginName对应的s_driver失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func CheckOtherVehicleNum(driverID int, vehicleNum string) bool {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And("driver_id != ?").
		And("vehicle_number = ? ")

	sql := qd.String()
	if num, err := o.Raw(sql, driverID, vehicleNum).QueryRows(&[]SDriver{}); err != nil {
		logs.Error("[mysql]  查询loginName对应的s_driver失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func CheckAddDriverTel(tel string) bool {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And(" tel = ? ")
	sql := qd.String()
	if num, err := o.Raw(sql, tel).QueryRows(&[]SDriver{}); err != nil {
		logs.Error("[mysql]  查询loginName对应的s_driver失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func CheckOtherDriverTel(driverID int, tel string) bool {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And("driver_id != ?").
		And(" tel = ? ")
	sql := qd.String()
	if num, err := o.Raw(sql, driverID, tel).QueryRows(&[]SDriver{}); err != nil {
		logs.Error("[mysql]  查询loginName对应的s_driver失败，失败信息：", err.Error())
		return false
	} else if num == 0 {
		return true
	}

	return false
}

func GetAllDriver() (list []SDriver) {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").OrderBy("driver_id asc")

	sql := qd.String()
	list = []SDriver{}
	o.Raw(sql).QueryRows(&list)
	return list
}

func GetSDriverByVehicleNum(vehicleNum string) (driver *SDriver, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And("vehicle_number = ? ")
	driver = &SDriver{}
	if queryErr := o.Raw(qd.String(), vehicleNum).QueryRow(driver); queryErr != nil {
		logs.Error("[mysql]  查询车牌对应的司机信息失败，失败信息：", queryErr.Error())
		return nil, queryErr
	} else if driver == nil {
		logs.Error("[mysql]  未查询到车牌对应的司机信息,车牌号码:", vehicleNum)
		return nil, errors.New("未查询到车牌对应的司机信息")
	}

	return driver, nil
}

func GetSDriverByUserID(userID int) (driver *SDriver, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("s_driver").
		Where("is_delete = 1").
		And("user_id = ? ")
	driver = &SDriver{}
	if queryErr := o.Raw(qd.String(), userID).QueryRow(driver); queryErr != nil {
		logs.Error("[mysql]  查询userID对应的司机信息失败，失败信息：", queryErr.Error())
		return nil, queryErr
	} else if driver == nil {
		logs.Error("[mysql]  未查询到userID对应的司机信息,userID:", userID)
		return nil, errors.New("未查询到userID对应的司机信息")
	}

	return driver, nil
}

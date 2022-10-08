package wtws_mysql

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type SPublicFunction struct {
	Id             int       `orm:"column(function_id);auto" description:"guid"`
	Path           string    `orm:"column(path);size(64);null" description:"权限路径"`
	FunctionName   string    `orm:"column(function_name);size(32)" description:"功能名称"`
	FunctionMethod string    `orm:"column(function_method);size(16);null" description:"请求方法:get、post、delete、put"`
	Description    string    `orm:"column(description);size(64);null" description:"功能描述"`
	IsValid        int8      `orm:"column(is_valid);null" description:"是否有效 1-有效 2-无效"`
	InsertTime     time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间"`
	UpdateTime     time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间"`
}

func (t *SPublicFunction) TableName() string {
	return "s_public_function"
}

func init() {
	orm.RegisterModel(new(SPublicFunction))
}

// AddSPublicFunction insert a new SPublicFunction into database and returns
// last inserted Id on success.
func AddSPublicFunction(m *SPublicFunction) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSPublicFunctionById retrieves SPublicFunction by Id. Returns error if
// Id doesn't exist
func GetSPublicFunctionById(id int) (v *SPublicFunction, err error) {
	o := orm.NewOrm()
	v = &SPublicFunction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateSPublicFunction updates SPublicFunction by Id and returns error if
// the record to be updated doesn't exist
func UpdateSPublicFunctionById(m *SPublicFunction) (err error) {
	o := orm.NewOrm()
	v := SPublicFunction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSPublicFunction deletes SPublicFunction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSPublicFunction(id int) (err error) {
	o := orm.NewOrm()
	v := SPublicFunction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SPublicFunction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllSPublicFunction retrieves all SPublicFunction matches certain condition. Returns empty list if
// no records exist
func GetAllSPublicFunction() (dataArr []SPublicFunction) {
	o := orm.NewOrm()
	qs := o.QueryTable("s_public_function")
	if _, err := qs.All(&dataArr); err != nil {
		dataArr = []SPublicFunction{}
	}
	return dataArr
}

package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EEnterprise struct {
	Id             int       `orm:"column(enterprise_id);auto" description:"主键ID" json:"enterpriseID"`
	Name           string    `orm:"column(name);size(64);null" description:"类别名称" json:"name"`
	ContactPerson  string    `orm:"column(contact_person);size(16);null" description:"联系人" json:"contactPerson"`
	PhoneTel       string    `orm:"column(phone_tel);size(16);null" description:"联系电话" json:"phoneTel"`
	LegalName      string    `orm:"column(legal_name);size(16);null" description:"法人" json:"legalName"`
	EnterpriseType string    `orm:"column(enterprise_type);size(20);null" description:"企业类别" json:"enterpriseType"`
	Longitude      float64   `orm:"column(longitude);null" description:"经度" json:"longitude"`
	Latitude       float64   `orm:"column(latitude);null" description:"纬度" json:"latitude"`
	LicensePic     string    `orm:"column(license_pic);size(128);null" description:"营业执照" json:"licensePic"`
	Remark         string    `orm:"column(remark);size(128);null" description:"备注" json:"remark"`
	IndustryType   string    `orm:"column(industry_type);size(20);null" description:"行业类别" json:"industryType"`
	Area           string    `orm:"column(area);size(16);null" description:"行政区" json:"area"`
	City           string    `orm:"column(city);size(16);null" description:"市" json:"city"`
	Province       string    `orm:"column(province);size(16);null" description:"省" json:"province"`
	Address        string    `orm:"column(address);size(64);null" description:"地址" json:"address"`
	IsDelete       int8      `orm:"column(is_delete);null" description:"是否删除1-未删除 2-删除" json:"isDelete"`
	UpdateTime     time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间" json:"updateTime"`
	InsertTime     time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间" json:"insertTime"`
}

func (t *EEnterprise) TableName() string {
	return "e_enterprise"
}

func init() {
	orm.RegisterModel(new(EEnterprise))
}

// AddEEnterprise insert a new EEnterprise into database and returns
// last inserted Id on success.
func AddEEnterprise(m *EEnterprise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEEnterpriseById retrieves EEnterprise by Id. Returns error if
// Id doesn't exist
func GetEEnterpriseById(id int) (v *EEnterprise, err error) {
	o := orm.NewOrm()
	v = &EEnterprise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEEnterprise retrieves all EEnterprise matches certain condition. Returns empty list if
// no records exist
func GetAllEEnterprise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EEnterprise))
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

	var l []EEnterprise
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

// UpdateEEnterprise updates EEnterprise by Id and returns error if
// the record to be updated doesn't exist
func UpdateEEnterpriseById(m *EEnterprise) (err error) {
	o := orm.NewOrm()
	v := EEnterprise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEEnterprise deletes EEnterprise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEEnterprise(id int) (err error) {
	o := orm.NewOrm()
	v := EEnterprise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EEnterprise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetUserEnterprise 获取用户对应的服务站
func GetUserEnterprise(userId int) (enterprises []EEnterprise) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	//qb.Select("ee.*").
	//	From("e_enterprise as ee").
	//	InnerJoin("o_station_enterprise as ose").
	//	On("ee.enterprise_id = ose.enterprise_id").
	//	InnerJoin("o_station as os").
	//	On("os.station_id = ose.station_id").
	//	InnerJoin("s_user_station as sus").
	//	On("sus.station_id = os.station_id").
	//	Where("sus.user_id = ?").
	//	And("os.is_delete = 1").
	//	And("ee.is_delete = 1").
	//	OrderBy("os.station_id asc")

	qb.Select("ee.*").
		From("e_enterprise as ee").
		InnerJoin("e_user_enterprise as eue").
		On("ee.enterprise_id = eue.enterprise_id").
		Where("eue.user_id = ?").
		And("ee.is_delete = 1").
		OrderBy("ee.update_time desc")

	sql := qb.String()
	enterprises = []EEnterprise{}
	if _, err := o.Raw(sql, userId).QueryRows(&enterprises); err != nil {
		logs.Error("[mysql]  根据userId查询关联的服务站点企业信息失败，失败信息:", err.Error())
		return []EEnterprise{}
	}
	return enterprises
}

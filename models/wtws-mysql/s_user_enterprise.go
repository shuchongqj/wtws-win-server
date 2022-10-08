package wtws_mysql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SUserEnterprise struct {
	Id           int       `orm:"column(us_id);auto" description:"用户客户站点关系标识"`
	UserId       int       `orm:"column(user_id)" description:"用户标识"`
	EnterpriseId int       `orm:"column(enterprise_id)" description:"客户站点标识"`
	InsertTime   time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间"`
}

func (t *SUserEnterprise) TableName() string {
	return "s_user_enterprise"
}

func init() {
	orm.RegisterModel(new(SUserEnterprise))
}

// AddSUserEnterprise insert a new SUserEnterprise into database and returns
// last inserted Id on success.
func AddSUserEnterprise(m *SUserEnterprise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSUserEnterpriseById retrieves SUserEnterprise by Id. Returns error if
// Id doesn't exist
func GetSUserEnterpriseById(id int) (v *SUserEnterprise, err error) {
	o := orm.NewOrm()
	v = &SUserEnterprise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSUserEnterprise retrieves all SUserEnterprise matches certain condition. Returns empty list if
// no records exist
func GetAllSUserEnterprise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SUserEnterprise))
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

	var l []SUserEnterprise
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

// UpdateSUserEnterprise updates SUserEnterprise by Id and returns error if
// the record to be updated doesn't exist
func UpdateSUserEnterpriseById(m *SUserEnterprise) (err error) {
	o := orm.NewOrm()
	v := SUserEnterprise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSUserEnterprise deletes SUserEnterprise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSUserEnterprise(id int) (err error) {
	o := orm.NewOrm()
	v := SUserEnterprise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SUserEnterprise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

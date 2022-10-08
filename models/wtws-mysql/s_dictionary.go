package wtws_mysql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SDictionary struct {
	Id         int       `orm:"column(dictionary_id);pk;auto" description:"主键ID"`
	TypeName   int8      `orm:"column(type_name)" description:"类别名称"`
	Key        string    `orm:"column(key);size(64)" description:"key"`
	Value      string    `orm:"column(value);size(64)" description:"value"`
	IsDisable  int16     `orm:"column(is_disable);null" description:"是否禁用1-可用2-不可用"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime);auto_now_add" description:"记录插入时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);auto_now" description:"记录更新时间"`
}

func (t *SDictionary) TableName() string {
	return "s_dictionary"
}

func init() {
	orm.RegisterModel(new(SDictionary))
}

// AddSDictionary insert a new SDictionary into database and returns
// last inserted Id on success.
func AddSDictionary(m *SDictionary) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSDictionaryById retrieves SDictionary by Id. Returns error if
// Id doesn't exist
func GetSDictionaryById(id int) (v *SDictionary, err error) {
	o := orm.NewOrm()
	v = &SDictionary{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSDictionary retrieves all SDictionary matches certain condition. Returns empty list if
// no records exist
func GetAllSDictionary(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SDictionary))
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

	var l []SDictionary
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

// UpdateSDictionary updates SDictionary by Id and returns error if
// the record to be updated doesn't exist
func UpdateSDictionaryById(m *SDictionary) (err error) {
	o := orm.NewOrm()
	v := SDictionary{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSDictionary deletes SDictionary by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSDictionary(id int) (err error) {
	o := orm.NewOrm()
	v := SDictionary{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SDictionary{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

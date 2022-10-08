package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"time"
	"wtws-server/conf"

	"github.com/astaxie/beego/orm"
)

type GCategory struct {
	Id         int       `orm:"column(category_id);auto" description:"id" json:"categoryID"`
	Name       string    `orm:"column(name);size(64)" description:"分类名称" json:"name"`
	IsDelete   int8      `orm:"column(is_delete)" description:"是否删除 1-未删除 2-已删除" json:"isDelete"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime)" description:"记录创建时间" json:"insertTime"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)" description:"记录更新时间" json:"updateTime"`
}

func (t *GCategory) TableName() string {
	return "g_category"
}

func init() {
	orm.RegisterModel(new(GCategory))
}

// AddGCategory insert a new GCategory into database and returns
// last inserted Id on success.
func AddGCategory(m *GCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGCategoryById retrieves GCategory by Id. Returns error if
// Id doesn't exist
func GetGCategoryById(id int) (v *GCategory, err error) {
	o := orm.NewOrm()
	v = &GCategory{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGCategory retrieves all GCategory matches certain condition. Returns empty list if
// no records exist
func GetAllGCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GCategory))
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

	var l []GCategory
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

// UpdateGCategory updates GCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateGCategoryById(m *GCategory) (err error) {
	o := orm.NewOrm()
	v := GCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGCategory deletes GCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGCategory(id int) (err error) {
	o := orm.NewOrm()
	v := GCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllCategory() (list []GCategory, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("g_category").Where("is_delete = 1")
	sql := qd.String()

	list = []GCategory{}

	if _, queryErr := o.Raw(sql).QueryRows(&list); queryErr != nil {
		logs.Error("[mysql]  查询所有的g_category失败，失败信息:", queryErr.Error())
		return list, queryErr
	}

	return list, nil

}

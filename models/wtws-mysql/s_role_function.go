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

type SRoleFunction struct {
	Id         int       `orm:"column(rf_id);pk;auto" description:"权限关系标识"`
	RoleId     int       `orm:"column(role_id)" description:"角色标识"`
	FunctionId int       `orm:"column(function_id)" description:"功能标示"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"插入时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"更新时间"`
}

func (t *SRoleFunction) TableName() string {
	return "s_role_function"
}

func init() {
	orm.RegisterModel(new(SRoleFunction))
}

// AddSRoleFunction insert a new SRoleFunction into database and returns
// last inserted Id on success.
func AddSRoleFunction(m *SRoleFunction) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(m); err != nil {
		logs.Error("[mysql] 新增s_role_function失败，roleID:", m.RoleId, "\tfunctionID:", m.FunctionId, "\t失败信息：", err.Error())
	}

	return

}

// GetSRoleFunctionById retrieves SRoleFunction by Id. Returns error if
// Id doesn't exist
func GetSRoleFunctionById(id int) (v *SRoleFunction, err error) {
	o := orm.NewOrm()
	v = &SRoleFunction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSRoleFunction retrieves all SRoleFunction matches certain condition. Returns empty list if
// no records exist
func GetAllSRoleFunction(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SRoleFunction))
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

	var l []SRoleFunction
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

// UpdateSRoleFunction updates SRoleFunction by Id and returns error if
// the record to be updated doesn't exist
func UpdateSRoleFunctionById(m *SRoleFunction) (err error) {
	o := orm.NewOrm()
	v := SRoleFunction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSRoleFunction deletes SRoleFunction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSRoleFunction(id int) (err error) {
	o := orm.NewOrm()
	v := SRoleFunction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SRoleFunction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetRoleFunctionsByRoleID 根据roleID查询角色的功能
func GetRoleFunctionsByRoleID(roleID int) (roleFunctions []SRoleFunction, err error) {
	o := orm.NewOrm()
	roleFunctions = []SRoleFunction{}
	if _, err = o.QueryTable("s_role_function").Filter("RoleId", roleID).All(&roleFunctions); err != nil {
		logs.Error("[mysql]  查询role对应的s_role_functions失败，失败信息:", err.Error())
	}
	return roleFunctions, err

}

// DelRoleFunctionByRoleID 根据RoleID删除用户的功能
func DelRoleFunctionByRoleID(roleID int) (int, error) {
	o := orm.NewOrm()
	sql := "DELETE FROM s_role_function WHERE role_id = ?"
	if result, err := o.Raw(sql, roleID).Exec(); err == nil {
		num, _ := result.RowsAffected()
		return int(num), err
	} else {
		logs.Error("[mysql]  根据roleID删除角色功能失败，失败信息：", err.Error())
		return 0, err
	}
}

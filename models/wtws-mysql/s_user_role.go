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

type SUserRole struct {
	Id         int       `orm:"column(ur_id);auto" description:"用户权限标识"`
	UserId     int       `orm:"column(user_id)" description:"用户ID"`
	RoleId     int       `orm:"column(role_id)" description:"角色ID"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null;auto_now"`
}

func (t *SUserRole) TableName() string {
	return "s_user_role"
}

func init() {
	orm.RegisterModel(new(SUserRole))
}

// AddSUserRole insert a new SUserRole into database and returns
// last inserted Id on success.
func AddSUserRole(m *SUserRole) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSUserRoleById retrieves SUserRole by Id. Returns error if
// Id doesn't exist
func GetSUserRoleById(id int) (v *SUserRole, err error) {
	o := orm.NewOrm()
	v = &SUserRole{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSUserRole retrieves all SUserRole matches certain condition. Returns empty list if
// no records exist
func GetAllSUserRole(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SUserRole))
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

	var l []SUserRole
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

// UpdateSUserRole updates SUserRole by Id and returns error if
// the record to be updated doesn't exist
func UpdateSUserRoleById(m *SUserRole) (err error) {
	o := orm.NewOrm()
	v := SUserRole{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSUserRole deletes SUserRole by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSUserRole(id int) (err error) {
	o := orm.NewOrm()
	v := SUserRole{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SUserRole{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetSUserRoleByUserId 查询用户的 用户权限
func GetSUserRoleByUserId(userID int) (v *SUserRole, err error) {
	o := orm.NewOrm()
	v = &SUserRole{UserId: userID}

	if err = o.Read(v, "UserId"); err != nil {
		logs.Error("[mysql]  根据UserID查询s_user_role失败，失败信息:", err.Error())
	}

	return v, err

}

// GetSUserRoleByRoleId 查询角色关联的s_urser_role
func GetSUserRoleByRoleId(roleId int) (v []SUserRole, err error) {
	o := orm.NewOrm()
	v = []SUserRole{}
	if _, err = o.QueryTable("s_user_role").Filter("RoleId", roleId).All(&v); err != nil {
		logs.Error("[mysql]  查询roleID对应是s_user_role失败，失败信息:", err.Error())
		return []SUserRole{}, err
	}

	return v, nil

}

// 删除用户的role
func DeleteUserRoles(userID int) (err error) {
	o := orm.NewOrm()
	sql := `DELETE FROM s_user_role WHERE user_id = ?`
	if res, delErr := o.Raw(sql, userID).Exec(); err != nil {
		logs.Error(fmt.Printf("[mysql]  删除userID: %d 的s_user_role数据失败，失败信息：%s", userID, delErr.Error()))
		return delErr
	} else {
		num, _ := res.RowsAffected()
		logs.Info(fmt.Printf("[mysql]  删除userID: %d 的s_user_role数据成功，共删除 %d 条", userID, num))
		return nil
	}

}

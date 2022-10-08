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

type SRole struct {
	Id          int       `orm:"column(role_id);auto" description:"角色标识"`
	RoleName    string    `orm:"column(role_name);size(20)" description:"角色名称"`
	RoleType    int16     `orm:"column(role_type)" description:"角色类型 1-后台"`
	IsDeleted   int8      `orm:"column(is_deleted)" description:"是否删除  1-未删除 2-删除"`
	Description string    `orm:"column(description);size(128);null" description:"角色描述"`
	InsertTime  time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"插入时间"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"更新时间"`
}

func (t *SRole) TableName() string {
	return "s_role"
}

func init() {
	orm.RegisterModel(new(SRole))
}

// AddSRole insert a new SRole into database and returns
// last inserted Id on success.
func AddSRole(m *SRole) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSRoleById retrieves SRole by Id. Returns error if
// Id doesn't exist
func GetSRoleById(id int) (v *SRole, err error) {
	o := orm.NewOrm()
	v = &SRole{Id: id, IsDeleted: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDeleted"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSRole retrieves all SRole matches certain condition. Returns empty list if
// no records exist
func GetAllSRole(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SRole))
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

	var l []SRole
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

// UpdateSRole updates SRole by Id and returns error if
// the record to be updated doesn't exist
func UpdateSRoleById(m *SRole) (err error) {
	o := orm.NewOrm()
	v := SRole{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSRole deletes SRole by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSRole(id int) (err error) {
	o := orm.NewOrm()
	v := SRole{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SRole{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllRoleList 获取所有的角色列表
func GetAllRoleList() (roles []SRole) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("s_role").Where("is_deleted = 1").OrderBy("update_time desc")
	sql := qd.String()

	roles = []SRole{}
	if _, err := o.Raw(sql).QueryRows(&roles); err != nil {
		roles = []SRole{}
	}

	return roles
}

// GetSRoleByRoleName 根据roleName查询可用的role信息
func GetSRoleByRoleName(roleName string) (role *SRole, err error) {
	o := orm.NewOrm()
	roles := []SRole{}
	if _, err = o.QueryTable("s_role").Filter("IsDeleted", 1).Filter("RoleName", roleName).All(&roles); err != nil || len(roles) == 0 {
		logs.Error("[mysql]  查询roleName对应的角色失败")
		return nil, errors.New("[mysql]  查询roleName对应的角色失败")
	}
	return &roles[0], nil
}

// UpdateByRoleID 根据RoleID修改自定义数据
func UpdateByRoleID(data *SRole, clos []string) (err error, updateNum int64) {
	o := orm.NewOrm()
	v := SRole{Id: data.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(data, clos...); err == nil {
			return nil, num
		}
	}
	return err, 0
}

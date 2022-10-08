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
)

type SFunction struct {
	Id                 int       `orm:"column(function_id);auto" description:"guid" json:"functionID"`
	SystemCode         int       `orm:"column(system_code)" description:"系统编码" json:"systemCode"`
	FunctionCode       int       `orm:"column(function_code)" description:"节点code" json:"functionCode"`
	ParentFunctionCode int       `orm:"column(parent_function_code);null" description:"功能编码父节点" json:"parentFunctionCode"`
	FunctionName       string    `orm:"column(function_name);size(50)" description:"功能名称" json:"functionName"`
	FunctionType       int8      `orm:"column(function_type)" description:"功能类型 1-菜单 2-按钮" json:"functionType"`
	IsValid            int16     `orm:"column(is_valid)" description:"是否有效 1-有效 2-无效" json:"isValid"`
	Index              int16     `orm:"column(index);null" description:"排序号" json:"index"`
	FunctionMethod     string    `orm:"column(function_method);size(32);null" description:"请求方法:get、post、delete、put" json:"functionMethod"`
	FunctionIcon       string    `orm:"column(function_icon);size(64);null" description:"菜单图标" json:"functionIcon"`
	Description        string    `orm:"column(description);size(128);null" description:"功能描述" json:"description"`
	Path               string    `orm:"column(path);size(128);null" description:"权限路径" json:"path"`
	InsertTime         time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间" json:"insertTime"`
	UpdateTime         time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间" json:"updateTime"`
}

type FunctionInfo struct {
	SFunction
	Children []FunctionInfo `json:"children"                `
}

func (t *SFunction) TableName() string {
	return "s_function"
}

func init() {
	orm.RegisterModel(new(SFunction))
}

// AddSFunction insert a new SFunction into database and returns
// last inserted Id on success.
func AddSFunction(m *SFunction) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSFunctionById retrieves SFunction by Id. Returns error if
// Id doesn't exist
func GetSFunctionById(id int) (v *SFunction, err error) {
	o := orm.NewOrm()
	v = &SFunction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSFunction retrieves all SFunction matches certain condition. Returns empty list if
// no records exist
func GetAllSFunction(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SFunction))
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

	var l []SFunction
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

// UpdateSFunction updates SFunction by Id and returns error if
// the record to be updated doesn't exist
func UpdateSFunctionById(m *SFunction) (err error) {
	o := orm.NewOrm()
	v := SFunction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSFunction deletes SFunction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSFunction(id int) (err error) {
	o := orm.NewOrm()
	v := SFunction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SFunction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllFunctionsByRoleID 查询用户所有的权限
func GetAllFunctionsByRoleID(roleID int) []FunctionInfo {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("sf.*").
		From("s_role_function as srf").
		LeftJoin("s_function as sf").
		On("sf.function_id = srf.function_id").
		Where("srf.role_id = ?").
		And("sf.is_valid = 1").
		And("sf.parent_function_code = 0").
		OrderBy("sf.index asc, sf.insert_time desc")

	sql := qd.String()
	sFunctionArr := []SFunction{}
	if parentRoleCount, getParentFunctionErr := o.Raw(sql, roleID).QueryRows(&sFunctionArr); getParentFunctionErr != nil || parentRoleCount == 0 {
		logs.Error("[mysql]  根据roleID查询parent的functions失败，失败信息：", getParentFunctionErr.Error())
		sFunctionArr = []SFunction{}
	}

	functionInfos := []FunctionInfo{}

	wg := &sync.WaitGroup{}
	wg.Add(len(sFunctionArr))

	for _, parentFunctionItem := range sFunctionArr {
		go func(parentFunctionItem SFunction, roleID int) {
			parentFunction := RecursiveFunction(parentFunctionItem, roleID)
			functionInfos = append(functionInfos, parentFunction)
			wg.Done()
		}(parentFunctionItem, roleID)

	}

	wg.Wait()
	return functionInfos

}

// RecursiveFunction 递归遍历父子结构的权限树
func RecursiveFunction(functionItem SFunction, roleID int) FunctionInfo {
	functionInfo := FunctionInfo{
		SFunction: functionItem,
		Children:  []FunctionInfo{},
	}

	subO := orm.NewOrm()
	subQd, _ := orm.NewQueryBuilder("mysql")
	subQd.Select("sf.*").
		From("s_role_function as srf").
		LeftJoin("s_function as sf").
		On("sf.function_id = srf.function_id").
		Where("srf.role_id = ?").
		And("sf.is_valid = 1").
		And("sf.parent_function_code = ?").
		OrderBy("sf.index asc, sf.insert_time desc")
	subSql := subQd.String()

	subFunctionArr := []SFunction{}
	if num, err := subO.Raw(subSql, roleID, functionItem.FunctionCode).QueryRows(&subFunctionArr); err != nil || num == 0 {
		functionInfo.Children = []FunctionInfo{}
		return functionInfo
	} else {
		subWg := &sync.WaitGroup{}
		subWg.Add(len(subFunctionArr))

		subFunctionInfos := []FunctionInfo{}

		for _, subFunctionItem := range subFunctionArr {
			go func(subFunctionItem SFunction, roleID int) {
				subFunctionInfo := RecursiveFunction(subFunctionItem, roleID)
				subFunctionInfos = append(subFunctionInfos, subFunctionInfo)
				subWg.Done()
			}(subFunctionItem, roleID)
		}

		subWg.Wait()

		functionInfo.Children = subFunctionInfos
		return functionInfo
	}

}

// CheckFunctionByUserID 根据userID账号查看是否有权限
func CheckFunctionByUserID(userID int, requestURI string, requestMethod string) bool {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("sf.*").
		From("s_function as sf").
		LeftJoin("s_role_function as srf").
		On("sf.function_id = srf.function_id").
		LeftJoin("s_user_role as sur").
		On("sur.role_id = srf.role_id").
		Where("sur.user_id = ?").
		And("sf.path=?").
		And("sf.function_method = ?").
		And("sf.is_valid = 1")

	sql := qb.String()
	if num, err := o.Raw(sql, userID, requestURI, requestMethod).QueryRows(&[]SFunction{}); err != nil || num == 0 {
		return false
	}

	return true
}

// GetFunctionByPath 根据path查询对应的functions
func GetFunctionByPath(path string) (functions []SFunction) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("s_function").
		Where("path=?").
		And("is_valid = 1")

	sql := qb.String()
	functions = []SFunction{}
	if num, err := o.Raw(sql, path).QueryRows(&functions); err != nil || num == 0 {
		return []SFunction{}
	}

	return functions
}

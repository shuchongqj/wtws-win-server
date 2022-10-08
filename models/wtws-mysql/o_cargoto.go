package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"

	"github.com/astaxie/beego/orm"
)

type OCargoto struct {
	Id         int       `orm:"column(cargoto_id);auto" json:"cargotoID"`
	Name       string    `orm:"column(name);size(32)" description:"装卸货地点中文" json:"cargotoName"`
	Code       string    `orm:"column(code);size(32)" description:"装卸货地点简拼" json:"code"`
	IsDelete   int8      `orm:"column(is_delete);null" description:"是否已删除  1-已删除  2未删除" json:"isDelete"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime);null" description:"记录创建时间" json:"insertTime"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"记录更新时间" json:"updateTime"`
}

func (t *OCargoto) TableName() string {
	return "o_cargoto"
}

func init() {
	orm.RegisterModel(new(OCargoto))
}

// AddOCargoto insert a new OCargoto into database and returns
// last inserted Id on success.
func AddOCargoto(m *OCargoto) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOCargotoById retrieves OCargoto by Id. Returns error if
// Id doesn't exist
func GetOCargotoById(id int) (v *OCargoto, err error) {
	o := orm.NewOrm()
	v = &OCargoto{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOCargoto retrieves all OCargoto matches certain condition. Returns empty list if
// no records exist
func GetAllOCargoto(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OCargoto))
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

	var l []OCargoto
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

// UpdateOCargoto updates OCargoto by Id and returns error if
// the record to be updated doesn't exist
func UpdateOCargotoById(m *OCargoto) (err error) {
	o := orm.NewOrm()
	v := OCargoto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOCargoto deletes OCargoto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOCargoto(id int) (err error) {
	o := orm.NewOrm()
	v := OCargoto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OCargoto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetCargotoList(pageNum, pageSize int, cargotoName, code string) (list []OCargoto, count int, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("o_cargoto").Where("is_delete = 1")

	if len(cargotoName) > 0 {
		qd.And("name LIKE '%" + cargotoName + "%'")
	}

	if len(code) > 0 {
		qd.And("code LIKE '%" + code + "%'")
	}

	qd.OrderBy("cargoto_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OCargoto{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询收货地失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OCargoto{}); getCountErr != nil {
			logs.Error("[mysql]  查询收货地失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OCargoto{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OCargoto{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByCargotoId(u *OCargoto, cols []string) (err error) {
	o := orm.NewOrm()
	v := OCargoto{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of o_cargoto records updated in database:", num)
		}
	}
	return err
}

func GetAllCargoto() (list []OCargoto) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("o_cargoto").
		Where("is_delete = 1").OrderBy("cargoto_id asc")

	list = []OCargoto{}
	o.Raw(qd.String()).QueryRows(&list)
	return list
}

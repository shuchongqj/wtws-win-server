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

type OOrigin struct {
	Id          int       `orm:"column(origin_id);auto" description:"id" json:"originID"`
	StationId   int       `orm:"column(station_id);null" description:"关联服务站ID(为平台化备用)" json:"stationId"`
	ContactName string    `orm:"column(contact_name);size(32)" description:"联系人姓名"  json:"contactName"`
	Name        string    `orm:"column(name);size(32)" description:"单位名称"  json:"name"`
	Tel         string    `orm:"column(tel);size(11)" description:"公司电话" json:"tel"`
	Address     string    `orm:"column(address);size(64)" description:"单位地址" json:"address"`
	Type        int8      `orm:"column(type)" description:"内外部 1-内部  2-外部" json:"type"`
	IsDelete    int8      `orm:"column(is_delete);null" description:"是否删除 1-正常 2-已删除" json:"isDelete"`
	InsertTime  time.Time `orm:"column(insert_time);type(datetime)" description:"记录创建时间" json:"insertTime"`
	Latitude    float64   `orm:"column(latitude);null" description:"纬度" json:"latitude"`
	Longitude   float64   `orm:"column(longitude);null" description:"经度" json:"longitude"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);auto_now" description:"记录更新时间" json:"updateTime"`
}

func (t *OOrigin) TableName() string {
	return "o_origin"
}

func init() {
	orm.RegisterModel(new(OOrigin))
}

// AddOOrigin insert a new OOrigin into database and returns
// last inserted Id on success.
func AddOOrigin(m *OOrigin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOOriginById retrieves OOrigin by Id. Returns error if
// Id doesn't exist
func GetOOriginById(id int) (v *OOrigin, err error) {
	o := orm.NewOrm()
	v = &OOrigin{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOOrigin retrieves all OOrigin matches certain condition. Returns empty list if
// no records exist
func GetAllOOrigin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OOrigin))
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

	var l []OOrigin
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

// UpdateOOrigin updates OOrigin by Id and returns error if
// the record to be updated doesn't exist
func UpdateOOriginById(m *OOrigin) (err error) {
	o := orm.NewOrm()
	v := OOrigin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOOrigin deletes OOrigin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOOrigin(id int) (err error) {
	o := orm.NewOrm()
	v := OOrigin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OOrigin{Id: id}); err == nil {
			logs.Info("Number of records deleted in database:", num)
		}
	}
	return
}

func GetOriginList(pageNum, pageSize, originType int, name, contactName, tel, address string) (list []OOrigin, count int, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("o_origin").Where("is_delete = 1")

	if len(name) > 0 {
		qd.And("name LIKE '%" + name + "%'")
	}

	if len(contactName) > 0 {
		qd.And("contact_name LIKE '%" + contactName + "%'")
	}

	if len(tel) > 0 {
		qd.And("tel LIKE '%" + tel + "%'")
	}

	if originType > 0 {
		qd.And(fmt.Sprintf("type = %d", originType))
	}

	if len(address) > 0 {
		qd.And("address LIKE '%" + address + "%'")
	}

	qd.OrderBy("origin_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OOrigin{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询发货地失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OOrigin{}); getCountErr != nil {
			logs.Error("[mysql]  查询发货地失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OOrigin{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OOrigin{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByOriginId(u *OOrigin, cols []string) (err error) {
	o := orm.NewOrm()
	v := OOrigin{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of o_origin records updated in database:", num)
		}
	}
	return err
}

func GetAllOriginList() []OOrigin {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("o_origin").Where("is_delete = 1")
	sql := qd.String()
	list := []OOrigin{}
	if _, err := o.Raw(sql).QueryRows(&list); err != nil {
		logs.Error("[mysql]  查询所有的发货地址失败，失败信息：", err.Error())
		return []OOrigin{}
	}

	return list
}

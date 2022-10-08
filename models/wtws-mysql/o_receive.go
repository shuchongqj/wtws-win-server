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

type OReceive struct {
	Id          int       `orm:"column(receive_id);auto" description:"id" json:"receiveID"`
	StationId   int       `orm:"column(station_id);null" description:"关联服务站ID(为平台化备用)" json:"stationId"`
	Name        string    `orm:"column(name);size(32)" description:"单位名称" json:"name"`
	ContactName string    `orm:"column(contact_name);size(32)" description:"联系人姓名"  json:"contactName"`
	Tel         string    `orm:"column(tel);size(16)" description:"公司电话" json:"tel"`
	Address     string    `orm:"column(address);size(64)" description:"单位地址" json:"address"`
	Type        int8      `orm:"column(type)" description:"内外部  1-内部  2-外部" json:"type"`
	Latitude    float64   `orm:"column(latitude);null" description:"纬度" json:"latitude"`
	Longitude   float64   `orm:"column(longitude);null" description:"经度" json:"longitude"`
	IsDelete    int8      `orm:"column(is_delete)" description:"是否删除 1-未删除  2-已删除" json:"isDelete"`
	InsertTime  time.Time `orm:"column(insert_time);type(datetime);auto_now_add" description:"记录插入时间" json:"insertTime"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);auto_now" description:"记录更新时间" json:"updateTime"`
}

func (t *OReceive) TableName() string {
	return "o_receive"
}

func init() {
	orm.RegisterModel(new(OReceive))
}

// AddOReceive insert a new OReceive into database and returns
// last inserted Id on success.
func AddOReceive(m *OReceive) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOReceiveById retrieves OReceive by Id. Returns error if
// Id doesn't exist
func GetOReceiveById(id int) (v *OReceive, err error) {
	o := orm.NewOrm()
	v = &OReceive{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOReceive retrieves all OReceive matches certain condition. Returns empty list if
// no records exist
func GetAllOReceive(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OReceive))
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

	var l []OReceive
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

// UpdateOReceive updates OReceive by Id and returns error if
// the record to be updated doesn't exist
func UpdateOReceiveById(m *OReceive) (err error) {
	o := orm.NewOrm()
	v := OReceive{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOReceive deletes OReceive by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOReceive(id int) (err error) {
	o := orm.NewOrm()
	v := OReceive{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&OReceive{Id: id}); err == nil {
			logs.Error(fmt.Sprintf("[mysql]  根据id删除o_receive失败，receive_id:%d 失败信息：%s", id, err.Error()))
		}
	}
	return
}

func GetReceiveList(pageNum, pageSize, receiveType int, name, contactName, tel, address string) (list []OReceive, count int, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("o_receive").Where("is_delete = 1")

	if len(name) > 0 {
		qd.And("name LIKE '%" + name + "%'")
	}

	if len(contactName) > 0 {
		qd.And("contact_name LIKE '%" + contactName + "%'")
	}

	if len(tel) > 0 {
		qd.And("tel LIKE '%" + tel + "%'")
	}

	if receiveType > 0 {
		qd.And(fmt.Sprintf("type = %d", receiveType))
	}

	if len(address) > 0 {
		qd.And("address LIKE '%" + address + "%'")
	}

	qd.OrderBy("receive_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OReceive{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询收货地失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OReceive{}); getCountErr != nil {
			logs.Error("[mysql]  查询收货地失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OReceive{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OReceive{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByReceiveId(u *OReceive, cols []string) (err error) {
	o := orm.NewOrm()
	v := OReceive{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of o_receive records updated in database:", num)
		}
	}
	return err
}

func GetAllReceiveList() []OReceive {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("o_receive").Where("is_delete = 1")
	sql := qd.String()
	list := []OReceive{}
	if _, err := o.Raw(sql).QueryRows(&list); err != nil {
		logs.Error("[mysql]  查询所有的收货地址失败，失败信息：", err.Error())
		return []OReceive{}
	}

	return list
}

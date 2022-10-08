package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
)

type OStationEnterprise struct {
	Id           int       `orm:"column(se_id);auto" description:"企业服务站关联标识"`
	StationId    int       `orm:"column(station_id)" description:"服务站ID"`
	EnterpriseId int       `orm:"column(enterprise_id)" description:"企业ID"`
	InsertTime   time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime);null;auto_now"`
}

func (t *OStationEnterprise) TableName() string {
	return "o_station_enterprise"
}

func init() {
	orm.RegisterModel(new(OStationEnterprise))
}

// AddOStationEnterprise insert a new OStationEnterprise into database and returns
// last inserted Id on success.
func AddOStationEnterprise(m *OStationEnterprise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOStationEnterpriseById retrieves OStationEnterprise by Id. Returns error if
// Id doesn't exist
func GetOStationEnterpriseById(id int) (v *OStationEnterprise, err error) {
	o := orm.NewOrm()
	v = &OStationEnterprise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOStationEnterprise retrieves all OStationEnterprise matches certain condition. Returns empty list if
// no records exist
func GetAllOStationEnterprise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OStationEnterprise))
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

	var l []OStationEnterprise
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

// UpdateOStationEnterprise updates OStationEnterprise by Id and returns error if
// the record to be updated doesn't exist
func UpdateOStationEnterpriseById(m *OStationEnterprise) (err error) {
	o := orm.NewOrm()
	v := OStationEnterprise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOStationEnterprise deletes OStationEnterprise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOStationEnterprise(id int) (err error) {
	o := orm.NewOrm()
	v := OStationEnterprise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OStationEnterprise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func UpdateStationEnterprise(stationID int, enterpriseIDs []int) (err error) {
	o := orm.NewOrm()

	sql := `DELETE FROM o_station_enterprise WHERE station_id = ?`

	if _, deleteErr := o.Raw(sql, stationID).Exec(); deleteErr != nil {
		logs.Error("[mysql]  删除s_station_enterprise失败，stationID:", stationID)
		return errors.New("删除原有o_station_enterprise失败")
	}

	//stationEnterprise := []OStationEnterprise{}
	//
	//qd.Select("*").From("o_station_enterprise").Where("station_id = ?")
	//sql := qd.String()
	//
	//if _, getStationEnterpriseErr := o.Raw(sql, stationID).QueryRows(&stationEnterprise); getStationEnterpriseErr != nil {
	//	return errors.New(fmt.Sprintf("查询stationID: %d 对应的o_station_enterprise失败。失败信息：%s", stationID, getStationEnterpriseErr.Error()))
	//}
	//
	//wg := &sync.WaitGroup{}
	//wg.Add(len(stationEnterprise))
	//
	//deleteErrArr := []error{}
	//
	//for _, e := range stationEnterprise {
	//	go func(stationID int, enterpriseID int) {
	//		o2 := orm.NewOrm()
	//		if num, deleteErr := o2.Delete(&OStationEnterprise{StationId: stationID, EnterpriseId: enterpriseID}, []string{"StationId", "EnterpriseId"}...); deleteErr != nil {
	//			logs.Error("[mysql]  删除s_station_enterprise失败，stationID:", stationID, "\tenterpriseID:", enterpriseID)
	//			deleteErrArr = append(deleteErrArr, deleteErr)
	//		} else {
	//			logs.Info("[mysql]  删除s_station_enterprise成功，stationID:", stationID, "\tenterpriseID:", enterpriseID, "\t删除数量:", num)
	//		}
	//
	//		wg.Done()
	//	}(stationID, e.Id)
	//}
	//
	//wg.Wait()

	//if len(deleteErrArr) > 0 {
	//	return errors.New("删除原有o_station_enterprise失败")
	//}

	wg2 := &sync.WaitGroup{}
	wg2.Add(len(enterpriseIDs))

	addStationEnterpriseErr := []error{}

	for _, enterpriseID := range enterpriseIDs {
		go func(stationID int, enterpriseID int) {
			if oeId, addErr := AddOStationEnterprise(&OStationEnterprise{
				StationId:    stationID,
				EnterpriseId: enterpriseID,
				InsertTime:   time.Now(),
				UpdateTime:   time.Now(),
			}); addErr != nil || oeId <= 0 {
				logs.Error(fmt.Sprintf("[mysql]  新增o_station_enterprise失败，stationID:%d  enterpriseID:%d  错误信息:", stationID, enterpriseID, addErr.Error()))
				addStationEnterpriseErr = append(addStationEnterpriseErr, errors.New("新增o_station_enterprise失败"))
			}

			wg2.Done()
		}(stationID, enterpriseID)

	}

	if len(addStationEnterpriseErr) > 0 {
		return errors.New("新增现在的o_station_enterprise失败")
	}

	return nil

}

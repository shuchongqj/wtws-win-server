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

type OStation struct {
	Id            int       `orm:"column(station_id);auto" description:"主键ID " json:"stationID"`
	Name          string    `orm:"column(name);size(64)" description:"客户站点名称" json:"name"`
	Address       string    `orm:"column(address);size(64)" description:"客户站点地址" json:"address"`
	ContactPerson string    `orm:"column(contact_person);size(32);null" description:"联系人" json:"contactPerson"`
	ContactTel    string    `orm:"column(contact_tel);size(16);null" description:"联系电话" json:"contactTel"`
	Longitude     float64   `orm:"column(longitude)" description:"经度" json:"longitude"`
	Latitude      float64   `orm:"column(latitude)" description:"纬度" json:"latitude"`
	IsDelete      int16     `orm:"column(is_delete);null" description:"是否删除1-未删除 2-删除" json:"isDelete"`
	InsertTime    time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间" json:"insertTime"`
	UpdateTime    time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间" json:"updateTime"`
	Province      string    `orm:"column(province);size(32);null" description:"省" json:"province"`
	City          string    `orm:"column(city);size(32);null" description:"市" json:"city"`
	Area          string    `orm:"column(area);size(32);null" description:"行政区" json:"area"`
}

type StationEnterpriseItem struct {
	EnterpriseID   int    `json:"enterpriseID"`
	EnterPriseName string `json:"enterpriseName"`
}

type StationListStruct struct {
	OStation
	Enterprise []StationEnterpriseItem `json:"enterprise"`
}

func (t *OStation) TableName() string {
	return "o_station"
}

func init() {
	orm.RegisterModel(new(OStation))
}

// AddOStation insert a new OStation into database and returns
// last inserted Id on success.
func AddOStation(m *OStation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOStationById retrieves OStation by Id. Returns error if
// Id doesn't exist
func GetOStationById(id int) (v *OStation, err error) {
	o := orm.NewOrm()
	v = &OStation{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOStation retrieves all OStation matches certain condition. Returns empty list if
// no records exist
func GetAllOStation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OStation))
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

	var l []OStation
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

// UpdateOStation updates OStation by Id and returns error if
// the record to be updated doesn't exist
func UpdateOStationById(m *OStation) (err error) {
	o := orm.NewOrm()
	v := OStation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOStation deletes OStation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOStation(id int) (err error) {
	o := orm.NewOrm()
	v := OStation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OStation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetUserStation 获取用户对应的服务站
func GetUserStation(userId int) (stations []OStation) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("os.*").
		From("o_station as os").
		LeftJoin("s_user_station as sus").
		On("os.station_id = sus.station_id").
		Where("sus.user_id = ?").
		And("os.is_delete = 1").OrderBy("os.station_id asc")

	sql := qb.String()
	stations = []OStation{}
	if count, err := o.Raw(sql, userId).QueryRows(&stations); count == 0 || err != nil {
		if err != nil {
			logs.Error("[mysql]  根据userId查询关联的服务站点失败，失败信息:", err.Error())
		}
		return []OStation{}
	}
	return stations
}

func GetStationList(
	userID int,
	name string,
	address string,
	contactPerson string,
	contactTel string,
	enterpriseName string,
	enterpryTel string,
	pageNum int,
	pageSize int) (stationListParams []StationListStruct, stationCount int) {

	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("os.*").
		From("o_station as os").
		LeftJoin("s_user_station as sus").
		On("os.station_id = sus.station_id").
		And("sus.user_id = ?").
		LeftJoin("o_station_enterprise as ose").
		On("ose.station_id = os.station_id").
		LeftJoin("e_enterprise as ee").
		On("ose.enterprise_id = ee.enterprise_id").
		Where("os.is_delete = 1")

	if len(name) > 0 {
		qd.And(" os.name LIKE '%" + name + "%'")
	}

	if len(address) > 0 {
		qd.And("os.address LIKE '%" + address + "%'")
	}

	if len(contactPerson) > 0 {
		qd.And("os.contact_person LIKE '%" + contactPerson + "%'")
	}

	if len(contactTel) > 0 {
		qd.And("os.contact_tel LIKE '%" + contactTel + "%'")
	}

	if len(enterpriseName) > 0 {
		qd.And("ee.name LIKE '%" + enterpriseName + "%'")
	}

	if len(enterpryTel) > 0 {
		qd.And("ee.phone_tel LIKE '%" + enterpryTel + "%'")
	}

	qd.OrderBy("os.station_id desc")

	countSql := qd.String()

	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)

	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	stationList := []OStation{}
	var stationCountInt64 int64
	stationCount = 0

	var getStationListErr, getStationCountErr error

	go func() {

		o := orm.NewOrm()
		if _, getStationListErr = o.Raw(sql, userID).QueryRows(&stationList); getStationListErr != nil {
			logs.Error("[mysql]  根据条件查询服务站列表失败，失败信息：", getStationListErr.Error())
		}

		wg.Done()
	}()

	go func() {
		o := orm.NewOrm()
		if stationCountInt64, getStationCountErr = o.Raw(countSql, userID).QueryRows(&[]OStation{}); getStationListErr != nil {
			logs.Error("[mysql]  根据条件查询服务站列表数量失败，失败信息：", getStationCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getStationListErr != nil || getStationCountErr != nil {
		return []StationListStruct{}, 0
	}

	stationCount = int(stationCountInt64)

	enterpriseGroup := sync.Map{}

	wg2 := &sync.WaitGroup{}
	wg2.Add(len(stationList))
	for _, station := range stationList {
		go func(station OStation) {
			enterpriseArr := []EEnterprise{}
			qd2, _ := orm.NewQueryBuilder("mysql")
			qd2.Select("ee.*").
				From("e_enterprise as ee").
				LeftJoin("o_station_enterprise as ose").
				On("ose.enterprise_id = ee.enterprise_id").
				Where("ee.is_delete = 1").
				And("ose.station_id = ?")
			sql2 := qd2.String()

			if num, getEnterpriseErr := orm.NewOrm().Raw(sql2, station.Id).QueryRows(&enterpriseArr); getEnterpriseErr != nil || num == 0 {
				enterpriseGroup.Store(station.Id, []StationEnterpriseItem{})
			} else {
				enterpriseNames := []StationEnterpriseItem{}
				for _, enterprise := range enterpriseArr {
					enterpriseNames = append(enterpriseNames, StationEnterpriseItem{
						EnterpriseID:   enterprise.Id,
						EnterPriseName: enterprise.Name,
					})
				}
				enterpriseGroup.Store(station.Id, enterpriseNames)
			}

			wg2.Done()
		}(station)
	}
	wg2.Wait()

	resStationList := []StationListStruct{}
	for _, station := range stationList {
		if enterpriseName, ok := enterpriseGroup.Load(station.Id); ok {
			resStationList = append(resStationList, StationListStruct{
				OStation:   station,
				Enterprise: enterpriseName.([]StationEnterpriseItem),
			})
		} else {
			resStationList = append(resStationList, StationListStruct{
				OStation:   station,
				Enterprise: []StationEnterpriseItem{},
			})
		}
	}

	return resStationList, stationCount

}

func UpdateStationByID(station *OStation, cols []string) (err error) {
	o := orm.NewOrm()
	v := OStation{Id: station.Id, IsDelete: 1}

	// ascertain id exists in the database
	if err = o.Read(&v, "Id", "IsDelete"); err == nil {
		var num int64
		if num, err = o.Update(station, cols...); err == nil {
			logs.Info("[mysql]  成功更新数据库服务站数据，更新数量:", num)
		}
	}
	return err
}

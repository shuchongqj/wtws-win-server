package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"

	"github.com/astaxie/beego/orm"
)

type OrTruckOrder struct {
	Id                 int       `orm:"column(truck_order_id);auto" description:"派车单ID" json:"truckOrderID"`
	TruckOrderNo       string    `orm:"column(truck_order_no);size(32)" description:"派车单编号" json:"truckOrderNo"`
	OrderID            int       `orm:"column(order_id)" description:"关联订单ID" json:"orderID"`
	OrderNo            string    `orm:"column(order_no);size(32);null" description:"关联订单编号" json:"orderNo"`
	Status             int8      `orm:"column(status)" description:"派车单状态 1-待审核 2-审核通过 3-审核拒绝 4-失效 5-结束  6-过磅单完成" json:"status"`
	OrderType          int8      `orm:"column(order_type)" description:"派车单类型  1-采购单  2-销售单  3-直发" json:"orderType"`
	DriverId           int       `orm:"column(driver_id)" description:"司机ID" json:"driverID"`
	DriverName         string    `orm:"column(driver_name);size(32)" description:"司机名字" json:"driverName"`
	DriverTel          string    `orm:"column(driver_tel);size(16)" description:"司机电话" json:"driverTel"`
	VehicleNumber      string    `orm:"column(vehicle_number);size(16)" description:"车牌号" json:"vehicleNumber"`
	LimitTotalLoad     string    `orm:"column(limit_total_load);size(50);null" description:"车辆载重" json:"limitTotalLoad"`
	IsWeightLimit      int8      `orm:"column(is_weight_limit)" description:"是否限重  1-不限重  2-限重" json:"isWeightLimit"`
	DriverTime         time.Time `orm:"column(driver_time);type(datetime)" description:"运输时间" json:"driverTime"`
	ContainerNo        string    `orm:"column(container_no);size(32)" description:"集装箱/柜号" json:"containerNo"`
	ReceiveId          int       `orm:"column(receive_id)" description:"收货单位id" json:"receiveID"`
	ReceiveTel         string    `orm:"column(receive_tel);size(16)" description:"收货电话" json:"receiveTel"`
	ReceiveName        string    `orm:"column(receive_name);size(16)" description:"收货人姓名" json:"receiveName"`
	ReceiveAddress     string    `orm:"column(receive_address);size(64)" description:"收货地址" json:"receiveAddress"`
	OriginId           int       `orm:"column(origin_id)" description:"发货单位id" json:"originID"`
	OriginName         string    `orm:"column(origin_name);size(16)" description:"发件人" json:"originName"`
	OriginTel          string    `orm:"column(origin_tel);size(50)" description:"发货电话" json:"originTel"`
	OriginAddress      string    `orm:"column(origin_address);size(64)" description:"发货地址" json:"originAddress"`
	CargotoId          int       `orm:"column(cargoto_id);null" description:"装卸货地点ID" json:"cargotoID"`
	CargotoName        string    `orm:"column(cargoto_name);size(32);null" description:"装卸货地点名称" json:"cargotoName"`
	GoodsId            int       `orm:"column(goods_id)" description:"货品ID" json:"goodsID"`
	GoodsName          string    `orm:"column(goods_name);size(32)" description:"货品名称" json:"goodsName"`
	GoodsNo            string    `orm:"column(goods_no);size(64)" description:"货品编号" json:"goodsNo"`
	GoodsWeight        float32   `orm:"column(goods_weight)" description:"重量" json:"goodsWeight"`
	GoodsNum           int       `orm:"column(goods_num)" description:"货品数量" json:"goodsNum"`
	GoodsUnit          string    `orm:"column(goods_unit);size(16)" description:"货单单位" json:"goodsUnit"`
	GoodsLoadQuantity  float32   `orm:"column(goods_load_quantity)" description:"货品装车量" json:"goodsLoadQuantity"`
	GoodsSpecification float32   `orm:"column(goods_specification)" description:"货品规格" json:"goodsSpecification"`
	GoodsNote          string    `orm:"column(goods_note);size(64);null" description:"货品备注" json:"goodsNote"`
	GoodsArranged      float32   `orm:"column(goods_arranged)" description:"已运输量/已派车" json:"goodsArranged"`
	GoodsExtraWeight   float32   `orm:"column(goods_extra_weight)" description:"运行额外可配发重量" json:"goodsExtraWeight"`
	BankName           string    `orm:"column(bank_name);size(32);null" description:"银行" json:"bankName"`
	BankNo             string    `orm:"column(bank_no);size(32);null" description:"银行卡号" json:"bankNo"`
	BankUserName       string    `orm:"column(bank_user_name);size(32);null" description:"银行卡开户人姓名" json:"bankUserName"`
	PaymentMethod      int8      `orm:"column(payment_method);null" description:"结算方式  1-供方结算 2-客户结算  3-无需结算" json:"paymentMethod"`
	CreatorId          int       `orm:"column(creator_id)" description:"制单人ID user_id" json:"creatorID"`
	CreatorName        string    `orm:"column(creator_name);size(32)" description:"制单人名字 user_name" json:"creatorName"`
	VerifierId         int       `orm:"column(verifier_id);null" description:"审核人ID user_id" json:"verifierID"`
	VerifierName       string    `orm:"column(verifier_name);size(32);null" description:"审核人名字" json:"verifierName"`
	VerifierNote       string    `orm:"column(verifier_note);size(64);null" description:"审核备注" json:"verifierNote"`
	InvalidId          int       `orm:"column(invalid_id)" description:"作废人的 user_id" json:"invalidId"`
	InvalidName        string    `orm:"column(invalid_name)" description:"作废人的 user_name" json:"invalidName"`
	IsDelete           int8      `orm:"column(is_delete)" description:"是否删除 1-正常 2-已删除" json:"isDelete" json:"isDelete"`
	VerifyTime         time.Time `orm:"column(verify_time);type(datetime);null" description:"派车单审核时间" json:"verifyTime"`
	InvalidTime        time.Time `orm:"column(invalid_time);type(datetime);null" description:"作废时间" json:"invalidTime"`
	InsertTime         time.Time `orm:"column(insert_time);type(datetime);auto_now_add" description:"创建时间" json:"insertTime"`
	UpdateTime         time.Time `orm:"column(update_time);type(datetime);auto_now" description:"记录更新时间" json:"updateTime"`
}

func (t *OrTruckOrder) TableName() string {
	return "or_truck_order"
}

func init() {
	orm.RegisterModel(new(OrTruckOrder))
}

// AddOrTruckOrder insert a new OrTruckOrder into database and returns
// last inserted Id on success.
func AddOrTruckOrder(m *OrTruckOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrTruckOrderById retrieves OrTruckOrder by Id. Returns error if
// Id doesn't exist
func GetOrTruckOrderById(id int) (v *OrTruckOrder, err error) {
	o := orm.NewOrm()
	v = &OrTruckOrder{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrTruckOrder retrieves all OrTruckOrder matches certain condition. Returns empty list if
// no records exist
func GetAllOrTruckOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrTruckOrder))
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

	var l []OrTruckOrder
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

// UpdateOrTruckOrder updates OrTruckOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrTruckOrderById(m *OrTruckOrder) (err error) {
	o := orm.NewOrm()
	v := OrTruckOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrTruckOrder deletes OrTruckOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrTruckOrder(id int) (err error) {
	o := orm.NewOrm()
	v := OrTruckOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrTruckOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetTruckOrderList(
	pageNum, pageSize, orderType, orderStatus, receiveID, originID int,
	truckOrderNo, orderNo, goodsName, goodsNo string,
	insertTimeStr, verifyTimeStr []string) (list []OrTruckOrder, count int, err error) {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_truck_order").Where("is_delete = 1")

	if orderType > 0 {
		qd.And(fmt.Sprintf("order_type = %d", orderType))
	}

	if orderStatus > 0 {
		qd.And(fmt.Sprintf("status = %d", orderStatus))
	}

	if receiveID > 0 {
		qd.And(fmt.Sprintf("receive_id = %d", receiveID))
	}

	if originID > 0 {
		qd.And(fmt.Sprintf("origin_id = %d", originID))
	}

	if len(truckOrderNo) > 0 {
		qd.And("truck_order_no LIKE '%" + truckOrderNo + "%'")
	}

	if len(orderNo) > 0 {
		qd.And("order_no LIKE '%" + orderNo + "%'")
	}

	if len(goodsName) > 0 {
		qd.And("goods_name LIKE '%" + goodsName + "%'")
	}

	if len(goodsNo) > 0 {
		qd.And("goods_no LIKE '%" + goodsNo + "%'")
	}

	if len(insertTimeStr) == 2 {
		qd.And(fmt.Sprintf("insert_time >= '%s 00:00:00' AND insert_time <= '%s 23:59:59'", insertTimeStr[0], insertTimeStr[1]))
	}

	if len(verifyTimeStr) == 2 {
		qd.And(fmt.Sprintf("verify_time >= '%s 00:00:00' AND verify_time <= '%s 23:59:59'", verifyTimeStr[0], verifyTimeStr[1]))
	}

	qd.OrderBy("truck_order_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OrTruckOrder{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询派车单失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OrTruckOrder{}); getCountErr != nil {
			logs.Error("[mysql]  查询派车单失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OrTruckOrder{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OrTruckOrder{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByTruckOrderId(u *OrTruckOrder, cols []string) (err error) {
	o := orm.NewOrm()
	v := OrTruckOrder{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of or_truck_order records updated in database:", num)
		}
	}
	return err
}

func GetAllTruckOrderList() (list []OrTruckOrder) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_truck_order").Where("is_delete = 1")
	list = []OrTruckOrder{}

	o.Raw(qd.String()).QueryRows(&list)

	return list
}

func GetCheckedTruckOrderByVehicleNumber(vehicleNumber string, status int) (list []OrTruckOrder, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("or_truck_order").
		Where("is_delete = 1").
		And("status = 2").
		And("vehicle_number = ?")

	if status > 0 {
		qd.And(fmt.Sprintf("status = %d", status))
	}

	list = []OrTruckOrder{}
	if _, err = o.Raw(qd.String(), vehicleNumber).QueryRows(&list); err != nil {
		logs.Error("[mysql]  根据车牌号查询车辆的派车单失败，失败信息:", err.Error())
		return []OrTruckOrder{}, err
	} else if len(list) == 0 {
		logs.Error(fmt.Sprintf("[mysql]  根据车牌号未查询到对应的派车单，车牌号:%s", vehicleNumber))
		return []OrTruckOrder{}, errors.New("未查询到对应的车牌号")
	}
	return list, nil
}

func GetLatestVehicleTruckOrder(vehicleNumber string, driverTime string) (list []OrTruckOrder, err error) {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select(`oto.*`).
		From("or_truck_order as oto").
		Where("oto.is_delete = 1").And("oto.status = 2")

	if len(vehicleNumber) > 0 {
		qd.And("oto.vehicle_number LIKE '%" + vehicleNumber + "%'")
	}

	if len(driverTime) > 0 {
		timeRegx1 := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}) \d{2}:\d{2}:\d{2}$`)
		regexpArrStr := timeRegx1.FindStringSubmatch(driverTime)
		if len(regexpArrStr) == 2 {
			qd.And(fmt.Sprintf("oto.driver_time >= '%s 00:00:00' AND oto.driver_time <= '%s 23:59:59'", regexpArrStr[1], regexpArrStr[1]))
		} else {
			logs.Error("[mysql]  查询近期派车单的driverTime参数错误，driverTime:", driverTime)
		}
	} else {
		startTimeStr := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
		endTimeStr := time.Now().AddDate(0, 0, 2).Format("2006-01-02")
		qd.And(fmt.Sprintf("oto.driver_time >= '%s 00:00:00' AND oto.driver_time <= '%s 23:59:59'", startTimeStr, endTimeStr))
	}

	qd.OrderBy("driver_time desc")

	list = []OrTruckOrder{}
	if _, queryErr := o.Raw(qd.String()).QueryRows(&list); queryErr != nil {
		logs.Error("[mysql]  查询车辆近期派车单失败，失败信息:", queryErr.Error())
		return []OrTruckOrder{}, queryErr
	}

	return list, nil

}

func GetOrTruckOrdersByInsertTime(startTimeStr string, endTimeStr string) (list []OrTruckOrder, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("or_truck_order").
		Where("is_delete = 1").
		And(fmt.Sprintf("insert_time >= '%s 00:00:00'", startTimeStr)).
		And(fmt.Sprintf("insert_time <= '%s 23:59:59'", endTimeStr))

	list = []OrTruckOrder{}

	if num, err := o.Raw(qd.String()).QueryRows(&list); err != nil {
		logs.Error("[mysql]  根据时间查询派车单失败，失败信息:", err.Error())
		return []OrTruckOrder{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询派车单失败,未查询到对应的派车单")
		return []OrTruckOrder{}, errors.New("根据时间查询派车单失败,未查询到对应的派车单")
	}

	return list, nil

}

func GetAnalysisOrTruckOrdersDayTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%m月%d日") AS dates,
			COUNT(*) as count 
		FROM or_truck_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%m月%d日")
 		ORDER BY DATE_FORMAT(insert_time,"%m月%d日") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询派送单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询派送单失败,未查询到对应的派送单")
		return []orm.Params{}, errors.New("根据时间查询派送单失败,未查询到对应的派送单")
	}

	return list, nil

}

func GetAnalysisOrTruckOrdersMonthTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年%m月") AS dates,
			COUNT(*) as count 
		FROM or_truck_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%Y年%m月")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年%m月") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询派送单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询派送单失败,未查询到对应的派送单")
		return []orm.Params{}, errors.New("根据时间查询派送单失败,未查询到对应的派送单")
	}

	return list, nil

}

func GetAnalysisOrTruckOrdersYearTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年") AS dates,
			COUNT(*) as count 
		FROM or_truck_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%Y年")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询派送单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询派送单失败,未查询到对应的派送单")
		return []orm.Params{}, errors.New("根据时间查询派送单失败,未查询到对应的派送单")
	}

	return list, nil

}

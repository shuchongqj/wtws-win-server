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

type OrWeighOrder struct {
	Id                   int       `orm:"column(weigh_order_id);auto" description:"过磅单ID" json:"weighOrderID"`
	WeighOrderNo         string    `orm:"column(weigh_order_no);size(32)" description:"过磅单号" json:"weighOrderNo"`
	Status               int8      `orm:"column(status)" description:"状态:  1-未完成 2-运输中 3-作废 4-已完成"  json:"status"`
	DriverId             int       `orm:"column(driver_id)" description:"过磅司机ID" json:"driverID"`
	DriverName           string    `orm:"column(driver_name);size(16)" description:"司机姓名" json:"driverName"`
	OrderType            int8      `orm:"column(order_type)" description:"订单类型  1-采购单 2-销售单  3-直发" json:"orderType"`
	IsDelete             int8      `orm:"column(is_delete)" description:"是否删除 1-未删除 2-已删除" json:"isDelete" json:"isDelete"`
	TruckOrderId         int       `orm:"column(truck_order_id);null" description:"预约单id" json:"truckOrderID"`
	TruckOrderNo         string    `orm:"column(truck_order_no);size(32);null" description:"关联派车单号" json:"truckOrderNo"`
	OrderId              int       `orm:"column(order_id);null" description:"关联订单的ID" json:"orderID"`
	OrderNo              string    `orm:"column(order_no);size(32);null" description:"关联订单单号" json:"orderNo"`
	VehicleNumber        string    `orm:"column(vehicle_number);size(16)" description:"过磅车牌号" json:"vehicleNumber"`
	IsWeightLimit        int8      `orm:"column(is_weight_limit);null" description:"是否限重  1-不限重 2-限重" json:"isWeightLimit"`
	LimitTotalLoad       string    `orm:"column(limit_total_load);size(50);null" description:"车辆载重" json:"limitTotalLoad"`
	ContainerNo          string    `orm:"column(container_no);size(32);null" description:"集装箱柜号" json:"containerNo"`
	TareWight            float32   `orm:"column(tare_wight);null" description:"皮重" json:"tareWight"`
	RoughWight           float32   `orm:"column(rough_wight);null" description:"毛重" json:"roughWight"`
	NetWight             float32   `orm:"column(net_wight);null" description:"净重" json:"netWight"`
	TareTime             time.Time `orm:"column(tare_time);type(datetime);null" description:"过磅皮重时间" json:"tareTime"`
	OtherWight           float32   `orm:"column(other_wight);null" description:"扣杂扣重" json:"otherWight"`
	WeighOrderNote       string    `orm:"column(weigh_order_note);size(64);null" description:"过磅备注" json:"weighOrderNote"`
	GoodsName            string    `orm:"column(goods_name);size(32);null" description:"货品" json:"goodsName"`
	GoodsNo              string    `orm:"column(goods_no);size(64);null" description:"货品编号" json:"goodsNo"`
	GoodsBatchNo         string    `orm:"column(goods_batch_no);size(64);null" description:"生产批号" json:"goodsBatchNo"`
	GoodsNum             int       `orm:"column(goods_num);null" description:"货品数量" json:"goodsNum"`
	GoodsSpecification   float32   `orm:"column(goods_specification);null" description:"货品规格" json:"goodsSpecification"`
	GoodsUnit            string    `orm:"column(goods_unit);size(50);null" description:"货品单位" json:"goodsUnit"`
	GoodsWeight          float32   `orm:"column(goods_weight);null" description:"货品重量" json:"goodsWeight"`
	GoodsExtraWeight     float32   `orm:"column(goods_extra_weight);null" description:"允许额外可配发重量" json:"goodsExtraWeight"`
	GoodsCategoryName    string    `orm:"column(goods_category_name);size(64);null" description:"产品类别" json:"goodsCategoryName"`
	GoodsDeductWeight    float32   `orm:"column(goods_deduct_weight);null" description:"每吨扣kg数（吨）" json:"goodsDeductWeight"`
	GoodsBagWeight       float32   `orm:"column(goods_bag_weight);null" description:"编织袋重量" json:"goodsBagWeight"`
	OriginId             int       `orm:"column(origin_id);null" description:"发货单位ID" json:"originID"`
	OriginName           string    `orm:"column(origin_name);size(255);null" description:"发货单位" json:"originName"`
	ReceiveId            int       `orm:"column(receive_id);null" description:"收货单位ID" json:"receiveID"`
	ReceiveName          string    `orm:"column(receive_name);size(64);null" description:"收货单位" json:"receiveName"`
	CargotoId            int       `orm:"column(cargoto_id);null" description:"装卸货地点ID" json:"cargotoID"`
	CargotoName          string    `orm:"column(cargoto_name);size(32);null" description:"装卸货地点" json:"cargotoName"`
	BankName             string    `orm:"column(bank_name);size(32);null" description:"银行" json:"bankName"`
	BankNo               string    `orm:"column(bank_no);size(32);null" description:"银行卡号" json:"bankNo"`
	WarehouseId          int       `orm:"column(warehouse_id);null" description:"仓库管理员user_id" json:"warehouseID"`
	WarehouseName        string    `orm:"column(warehouse_name);size(32);null" description:"仓库管理员姓名" json:"warehouseName"`
	FinancialId          int       `orm:"column(financial_id);null" description:"财务user_id" json:"financialID"`
	FinancialNote        string    `orm:"column(financial_note);size(64);null" description:"财务确认" json:"financialNote"`
	InvalidId            int       `orm:"column(invalid_id)" description:"作废人的 user_id" json:"invalidId"`
	InvalidName          string    `orm:"column(invalid_name)" description:"作废人的 user_name" json:"invalidName"`
	InvalidTime          time.Time `orm:"column(invalid_time);type(datetime);null" description:"作废时间" json:"invalidTime"`
	RoughTime            time.Time `orm:"column(rough_time);type(datetime);null" description:"过磅毛重时间" json:"roughTime"`
	WeightOverTime       time.Time `orm:"column(weight_over_time);type(datetime);null" description:"过磅结束时间" json:"weightOverTime"`
	WarehouseConfirmTime time.Time `orm:"column(warehouse_confirm_time);type(datetime);null" description:"仓库确认时间" json:"warehouseConfirmTime"`
	FinancialConfirmTime time.Time `orm:"column(financial_confirm_time);type(datetime);null" description:"财务确认时间" json:"financialConfirmTime"`
	PrintTime            time.Time `orm:"column(print_time);type(datetime)" description:"打印票据时间" json:"printTime"`
	InsertTime           time.Time `orm:"column(insert_time);type(datetime)" description:"记录创建时间" json:"insertTime"`
	UpdateTime           time.Time `orm:"column(update_time);type(datetime);auto_now_add" description:"记录更新时间" json:"updateTime"`
}

type OrWeighOrderItem struct {
	OrWeighOrder
	DriverTel string `json:"driverTel"`
}

func (t *OrWeighOrder) TableName() string {
	return "or_weigh_order"
}

func init() {
	orm.RegisterModel(new(OrWeighOrder))
}

// AddOrWeighOrder insert a new OrWeighOrder into database and returns
// last inserted Id on success.
func AddOrWeighOrder(m *OrWeighOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrWeighOrderById retrieves OrWeighOrder by Id. Returns error if
// Id doesn't exist
func GetOrWeighOrderById(id int) (v *OrWeighOrder, err error) {
	o := orm.NewOrm()
	v = &OrWeighOrder{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrWeighOrder retrieves all OrWeighOrder matches certain condition. Returns empty list if
// no records exist
func GetAllOrWeighOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrWeighOrder))
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

	var l []OrWeighOrder
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

// UpdateOrWeighOrder updates OrWeighOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrWeighOrderById(m *OrWeighOrder) (err error) {
	o := orm.NewOrm()
	v := OrWeighOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrWeighOrder deletes OrWeighOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrWeighOrder(id int) (err error) {
	o := orm.NewOrm()
	v := OrWeighOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrWeighOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetWeighOrderList(
	pageNum, pageSize, orderType, status, receiveID, originID int,
	weighOrderNo, truckOrderNo, orderNo, vehicleNumber, goodsName, goodsNo string,
	insertTimes, warehouseConfirmTime, financialConfirmTime []string) (list []OrWeighOrder, count int, err error) {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_weigh_order").Where("is_delete = 1")

	if orderType > 0 {
		qd.And(fmt.Sprintf("order_type = %d", orderType))
	}

	if status > 0 {
		qd.And(fmt.Sprintf("status = %d", status))
	}

	if receiveID > 0 {
		qd.And(fmt.Sprintf("receive_id = %d", receiveID))
	}

	if originID > 0 {
		qd.And(fmt.Sprintf("origin_id = %d", originID))
	}

	if len(weighOrderNo) > 0 {
		qd.And("weigh_order_no LIKE '%" + weighOrderNo + "%'")
	}

	if len(truckOrderNo) > 0 {
		qd.And("truck_order_no LIKE '%" + truckOrderNo + "%'")
	}

	if len(orderNo) > 0 {
		qd.And("order_no LIKE '%" + orderNo + "%'")
	}
	if len(vehicleNumber) > 0 {
		qd.And("vehicle_number LIKE '%" + vehicleNumber + "%'")
	}

	if len(goodsName) > 0 {
		qd.And("goods_name LIKE '%" + goodsName + "%'")
	}

	if len(goodsNo) > 0 {
		qd.And("goods_no LIKE '%" + goodsNo + "%'")
	}

	if len(insertTimes) == 2 {
		qd.And(fmt.Sprintf("insert_time >= '%s 00:00:00' AND insert_time <= '%s 23:59:59'", insertTimes[0], insertTimes[1]))
	}

	if len(warehouseConfirmTime) == 2 {
		qd.And(fmt.Sprintf("warehouse_confirm_time >= '%s 00:00:00' AND warehouse_confirm_time <= '%s 23:59:59'", warehouseConfirmTime[0], warehouseConfirmTime[1]))
	}

	if len(financialConfirmTime) == 2 {
		qd.And(fmt.Sprintf("financial_confirm_time >= '%s 00:00:00' AND financial_confirm_time <= '%s 23:59:59'", financialConfirmTime[0], financialConfirmTime[1]))
	}

	qd.OrderBy("weigh_order_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OrWeighOrder{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询订单失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OrWeighOrder{}); getCountErr != nil {
			logs.Error("[mysql]  查询订单失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OrWeighOrder{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OrWeighOrder{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByWeighOrderId(u *OrWeighOrder, cols []string) (err error) {
	o := orm.NewOrm()
	v := OrWeighOrder{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of or_weigh_order records updated in database:", num)
		}
	}
	return err
}

func GetAllWeighOrderList() (list []OrWeighOrder) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_weigh_order").Where("is_delete = 1")
	list = []OrWeighOrder{}

	o.Raw(qd.String()).QueryRows(&list)

	return list
}

func GetAllCheckedWeighOrder() (list []OrWeighOrder) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_weigh_order").Where("is_delete = 1").And("status = 2")
	list = []OrWeighOrder{}

	o.Raw(qd.String()).QueryRows(&list)

	return list
}

func GetWeighOrderByWeighOrderNo(weighOrderNo string) (weighOrder *OrWeighOrderItem, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("owo.*,sd.tel as driver_tel").
		From("or_weigh_order as owo").
		LeftJoin("s_driver as sd").
		On("owo.driver_id = sd.driver_id").
		Where("owo.is_delete = 1").
		And("owo.weigh_order_no = ?")
	weighOrders := []OrWeighOrderItem{}

	if num, err := o.Raw(qd.String(), weighOrderNo).QueryRows(&weighOrders); err != nil {
		logs.Error(fmt.Sprintf("[mysql]  根据过磅单号查找过磅单失败，过磅单号:%s \t失败信息:%s", weighOrderNo, err.Error()))
		return nil, err
	} else if num == 0 || len(weighOrders) == 0 {
		logs.Error(fmt.Sprintf("[mysql]  根据过磅单号查找过磅单失败，未查询到过磅单信息，过磅单号:%s", weighOrderNo))
		return nil, errors.New(fmt.Sprintf("[mysql]  根据过磅单号查找过磅单失败，未查询到过磅单信息，过磅单号:%s", weighOrderNo))
	}

	return &weighOrders[0], err
}

func UpdateOrWeighOrderStatusByWeighOrderNo(orderStatus int, oldStatus int, weighOrderNo string) error {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Update("or_weigh_order").
		Set("status = ?").
		Where("is_delete = 1").
		And("status = ?").
		And("weigh_order_no = ?")

	if execResult, updateErr := o.Raw(qd.String(), orderStatus, oldStatus, weighOrderNo).Exec(); updateErr != nil {
		logs.Error(fmt.Sprintf("[mysql]  更改订单状态失败,单号:%s \t状态:%d \t失败信息:%s", weighOrderNo, orderStatus, updateErr.Error()))
		return updateErr
	} else if affectedRow, _ := execResult.RowsAffected(); affectedRow == 0 {
		logs.Error(fmt.Sprintf("[mysql]  更改订单状态失败,未找到过磅单,单号:%s \t状态:%d \t失败信息:%s", weighOrderNo, orderStatus, updateErr.Error()))
		return errors.New("更改订单状态失败,未找到过磅单")
	}
	return nil
}

func GetOrWeighOrdersByInsertTime(startTimeStr string, endTimeStr string) (list []OrWeighOrder, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("or_weigh_order").
		Where("is_delete = 1").
		And(fmt.Sprintf("insert_time >= '%s 00:00:00'", startTimeStr)).
		And(fmt.Sprintf("insert_time <= '%s 23:59:59'", endTimeStr))

	list = []OrWeighOrder{}

	if num, err := o.Raw(qd.String()).QueryRows(&list); err != nil {
		logs.Error("[mysql]  根据时间查询过磅单失败，失败信息:", err.Error())
		return []OrWeighOrder{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询过磅单失败,未查询到对应的过磅单")
		return []OrWeighOrder{}, errors.New("根据时间查询过磅单失败,未查询到对应的过磅单")
	}

	return list, nil

}

func GetAnalysisOrWeighOrdersDayTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%m月%d日") AS dates,
			COUNT(*) as count 
		FROM or_weigh_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%m月%d日")
 		ORDER BY DATE_FORMAT(insert_time,"%m月%d日") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询过磅单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询过磅单失败,未查询到对应的过磅单")
		return []orm.Params{}, errors.New("根据时间查询过磅单失败,未查询到对应的过磅单")
	}

	return list, nil

}

func GetAnalysisOrWeighOrdersMonthTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年%m月") AS dates,
			COUNT(*) as count 
		FROM or_weigh_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%Y年%m月")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年%m月") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询过磅单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询过磅单失败,未查询到对应的过磅单")
		return []orm.Params{}, errors.New("根据时间查询过磅单失败,未查询到对应的过磅单")
	}

	return list, nil

}

func GetAnalysisOrWeighOrdersYearTimeGroupCount(startTimeStr string, endTimeStr string) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年") AS dates,
			COUNT(*) as count 
		FROM or_weigh_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		GROUP BY DATE_FORMAT(insert_time,"%Y年")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询过磅单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询过磅单失败,未查询到对应的过磅单")
		return []orm.Params{}, errors.New("根据时间查询过磅单失败,未查询到对应的过磅单")
	}

	return list, nil

}

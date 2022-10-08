package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"

	"github.com/astaxie/beego/orm"
)

type OrOrder struct {
	Id                 int       `orm:"column(order_id);auto" description:"订单ID" json:"orderID"`
	StationId          int       `orm:"column(station_id)" description:"所属服务站ID" json:"stationID"`
	OrderType          int8      `orm:"column(order_type)" description:"订单类型  1-采购单 2-出库单" json:"orderType"`
	OrderNo            string    `orm:"column(order_no);size(32)" description:"订单编号" json:"orderNo"`
	OriginId           int       `orm:"column(origin_id)" description:"发货单位id" json:"originID"`
	OriginName         string    `orm:"column(origin_name);size(32)" description:"发货单位" json:"originName"`
	ReceiveId          int       `orm:"column(receive_id)" description:"收货单位id" json:"receiveID"`
	ReceiveName        string    `orm:"column(receive_name);size(32)" description:"收货单位" json:"receiveName"`
	CreatorId          int       `orm:"column(creator_id)" description:"制单员ID user_id" json:"creatorID"`
	CreatorName        string    `orm:"column(creator_name);size(32)" description:"制单员名字 user_name" json:"creatorName"`
	Status             int8      `orm:"column(status)" description:"订单状态  1-等待审核 2-审核通过 3-审核驳回 4-失效 5-结束" json:"status"`
	IsDelete           int8      `orm:"column(is_delete)" description:"是否删除 1-未删除 2-已删除" json:"isDelete"`
	VerifierId         int       `orm:"column(verifier_id);null" description:"审核人ID  user_id" json:"verifierID"`
	VerifierName       string    `orm:"column(verifier_name);size(32);null" description:"审核人名字 user_name" json:"verifierName"`
	VerifierNote       string    `orm:"column(verifier_note);size(64);null" description:"审核备注" json:"verifierNote"`
	InvalidId          int       `orm:"column(invalid_id)" description:"作废人的 user_id" json:"invalidId"`
	InvalidName        string    `orm:"column(invalid_name)" description:"作废人的 user_name" json:"invalidName"`
	GoodsId            int       `orm:"column(goods_id)" description:"货品id" json:"goodsID"`
	GoodsName          string    `orm:"column(goods_name);size(32)" description:"货品名称" json:"goodsName"`
	GoodsNum           int       `orm:"column(goods_num)" description:"数量" json:"goodsNum"`
	GoodsUnit          string    `orm:"column(goods_unit);size(16)" description:"单位" json:"goodsUnit"`
	GoodsWeight        float32   `orm:"column(goods_weight)" description:"重量" json:"goodsWeight"`
	GoodsSpecification float32   `orm:"column(goods_specification)" description:"规格" json:"goodsSpecification"`
	GoodsNo            string    `orm:"column(goods_no);size(32)" description:"货品编号" json:"goodsNo"`
	GoodsMargin        float32   `orm:"column(goods_margin)" description:"产品余量" json:"goodsMargin"`
	GoodsArranged      float32   `orm:"column(goods_arranged)" description:"已派车" json:"goodsArranged"`
	GoodsExtraWeight   float32   `orm:"column(goods_extra_weight)" description:"允许额外可配发重量" json:"goodsExtraWeight"`
	GoodsNote          string    `orm:"column(goods_note);size(64);null" description:"物品备注" json:"goodsNote"`
	InvalidTime        time.Time `orm:"column(invalid_time);type(datetime);null" description:"作废时间" json:"invalidTime"`
	VerifyTime         time.Time `orm:"column(verify_time);type(datetime);null" description:"订单审核时间" json:"verifyTime"`
	UpdateTime         time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"记录更新时间" json:"updateTime"`
	InsertTime         time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"记录插入时间" json:"insertTime"`
}

func (t *OrOrder) TableName() string {
	return "or_order"
}

func init() {
	orm.RegisterModel(new(OrOrder))
}

// AddOrOrder insert a new OrOrder into database and returns
// last inserted Id on success.
func AddOrOrder(m *OrOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrOrderById retrieves OrOrder by Id. Returns error if
// Id doesn't exist
func GetOrOrderById(id int) (v *OrOrder, err error) {
	o := orm.NewOrm()
	v = &OrOrder{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrOrder retrieves all OrOrder matches certain condition. Returns empty list if
// no records exist
func GetAllOrOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrOrder))
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

	var l []OrOrder
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

// UpdateOrOrder updates OrOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrOrderById(m *OrOrder, cols []string) (err error) {
	o := orm.NewOrm()
	v := OrOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err != nil {
		logs.Error("[mysql]  根据orderID更新订单数据失败，失败信息:", err.Error())
		return err
	} else {
		if _, err = o.Update(m, cols...); err != nil {
			logs.Error("[mysql]  根据orderID更新订单数据失败，失败信息:", err.Error())
			return err
		}
	}

	return nil
}

// DeleteOrOrder deletes OrOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrOrder(id int) (err error) {
	o := orm.NewOrm()
	v := OrOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetOrderList(
	pageNum, pageSize, orderType, orderStatus, receiveID, originID int,
	orderNo, goodsName, goodsNo string,
	insertTimeStr, verifyTimeStr []string) (list []OrOrder, count int, err error) {

	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_order").Where("is_delete = 1")

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

	qd.OrderBy(" order_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []OrOrder{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询订单失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]OrOrder{}); getCountErr != nil {
			logs.Error("[mysql]  查询订单失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []OrOrder{}, 0, getListErr
	}

	if getCountErr != nil {
		return []OrOrder{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByOrderId(u *OrOrder, cols []string) (err error) {
	o := orm.NewOrm()
	v := OrOrder{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of or_order records updated in database:", num)
		}
	}
	return err
}

func GetAllOrderList() (list []OrOrder) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_order").Where("is_delete = 1")
	list = []OrOrder{}

	o.Raw(qd.String()).QueryRows(&list)

	return list
}

func GetAllCheckedOrder() (list []OrOrder) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("or_order").Where("is_delete = 1").And("status = 2")
	list = []OrOrder{}

	o.Raw(qd.String()).QueryRows(&list)

	return list
}

func GetOrOrdersByInsertTime(startTimeStr string, endTimeStr string) (list []OrOrder, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").
		From("or_order").
		Where("is_delete = 1").
		And(fmt.Sprintf("insert_time >= '%s 00:00:00'", startTimeStr)).
		And(fmt.Sprintf("insert_time <= '%s 23:59:59'", endTimeStr))

	list = []OrOrder{}

	if num, err := o.Raw(qd.String()).QueryRows(&list); err != nil {
		logs.Error("[mysql]  根据时间查询订单失败，失败信息:", err.Error())
		return []OrOrder{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询订单失败,未查询到对应的订单")
		return []OrOrder{}, errors.New("根据时间查询订单失败,未查询到对应的订单")
	}

	return list, nil

}

func GetAnalysisOrOrdersDayTimeGroupCount(startTimeStr string, endTimeStr string, orderType int) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%m月%d日") AS dates,
			COUNT(*) as count 
		FROM or_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		AND order_type = ` + strconv.Itoa(orderType) + `
		GROUP BY DATE_FORMAT(insert_time,"%m月%d日")
 		ORDER BY DATE_FORMAT(insert_time,"%m月%d日") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询订单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询订单失败,未查询到对应的订单")
		return []orm.Params{}, errors.New("根据时间查询订单失败,未查询到对应的订单")
	}

	return list, nil

}

func GetAnalysisOrOrdersMonthTimeGroupCount(startTimeStr string, endTimeStr string, orderType int) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年%m月") AS dates,
			COUNT(*) as count 
		FROM or_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		AND order_type = ` + strconv.Itoa(orderType) + `
		GROUP BY DATE_FORMAT(insert_time,"%Y年%m月")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年%m月") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询订单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询订单失败,未查询到对应的订单")
		return []orm.Params{}, errors.New("根据时间查询订单失败,未查询到对应的订单")
	}

	return list, nil

}

func GetAnalysisOrOrdersYearTimeGroupCount(startTimeStr string, endTimeStr string, orderType int) (list []orm.Params, err error) {
	o := orm.NewOrm()
	sql := `
		SELECT 
			DATE_FORMAT(insert_time ,"%Y年") AS dates,
			COUNT(*) as count 
		FROM or_order 
		WHERE is_delete = 1 
		AND insert_time >= '` + startTimeStr + ` 00:00:00'
		AND insert_time <= '` + endTimeStr + ` 23:59:59'
		AND order_type = ` + strconv.Itoa(orderType) + `
		GROUP BY DATE_FORMAT(insert_time,"%Y年")
 		ORDER BY DATE_FORMAT(insert_time,"%Y年") DESC `

	list = []orm.Params{}

	if num, err := o.Raw(sql).Values(&list); err != nil {
		logs.Error("[mysql]  根据时间查询订单失败，失败信息:", err.Error())
		return []orm.Params{}, err
	} else if num == 0 {
		logs.Error("[mysql]  根据时间查询订单失败,未查询到对应的订单")
		return []orm.Params{}, errors.New("根据时间查询订单失败,未查询到对应的订单")
	}

	return list, nil

}

package service

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/logs"
	"strconv"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetOrderInfo(orderID int) common_struct.ResponseStruct {
	if dbOrder, err := wtws_mysql.GetOrOrderById(orderID); err != nil || dbOrder == nil {
		return common.ResponseStatus(-1, "", nil)
	} else {
		return common.ResponseStatus(0, "", dbOrder)
	}
}

func GetAllCheckedOrder() common_struct.ResponseStruct {
	checkedOrderList := wtws_mysql.GetAllCheckedOrder()
	return common.ResponseStatus(0, "", checkedOrderList)
}

func GetOrderList(data *request_entity.OrderList) common_struct.ResponseStruct {

	if orderList, count, err := wtws_mysql.GetOrderList(data.PageNum, data.PageSize, data.OrderType,
		data.OrderStatus, data.ReceiveID, data.OriginID, data.OrderNo, data.GoodsName, data.GoodsNo,
		data.InsertTimes, data.VerifyTimes); err != nil {
		return common.ResponseStatus(0, "", dto.OrderList{
			List:  []wtws_mysql.OrOrder{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.OrderList{
			List:  orderList,
			Count: count,
		})
	}

}

func AddOrder(data *request_entity.AddOrder, user *wtws_mysql.SUser) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var receiveData *wtws_mysql.OReceive
	var originData *wtws_mysql.OOrigin

	var getReceiveErr, getOriginErr error

	go func() {
		receiveData, getReceiveErr = wtws_mysql.GetOReceiveById(data.ReceiveID)
		wg.Done()
	}()

	go func() {
		originData, getOriginErr = wtws_mysql.GetOOriginById(data.OriginID)
		wg.Done()
	}()

	wg.Wait()

	if getReceiveErr != nil || getOriginErr != nil || receiveData == nil || originData == nil {
		return common.ResponseStatus(-13, "", nil)
	}

	wg2 := &sync.WaitGroup{}
	wg2.Add(len(data.Goods))

	truckOrderArr := []request_entity.TruckOrders{}

	var addOrderErrs []error

	orderPrefix := "XS"
	if data.OrderType == conf.ORDER_PURCHASE_TYPE {
		orderPrefix = "CG"
	} else if data.OrderType == conf.ORDER_SENT_DIRECT_TYPE {
		orderPrefix = "ZF"
	}
	orderNo := common.GenerateOrderNo(orderPrefix)
	for _, good := range data.Goods {
		go func(good request_entity.Goods) {

			//var dbGood *wtws_mysql.GGoods
			//var getGoodErr error
			//dbGood,getGoodErr = wtws_mysql.GetGGoodsById(good.GoodID)
			//if dbGood == nil || getGoodErr != nil{
			//	logs.Error(fmt.Sprintf("[service]  新增订单时，查询goodID对应的货品失败，失败信息:%s",getGoodErr.Error()))
			//	addOrderErrs =  append(addOrderErrs , getGoodErr)
			//	wg2.Done()
			//	return
			//}

			addOrderData := &wtws_mysql.OrOrder{
				StationId:          conf.DEFAULT_STATION_ID,
				OrderType:          int8(data.OrderType),
				OrderNo:            orderNo,
				OriginId:           originData.Id,
				OriginName:         originData.Name,
				ReceiveId:          receiveData.Id,
				ReceiveName:        receiveData.Name,
				CreatorId:          user.Id,
				CreatorName:        user.DisplayName,
				Status:             conf.ORDER_STATUS_WAIT,
				GoodsId:            good.GoodID,
				GoodsName:          good.GoodsName,
				GoodsNum:           int(good.GoodsNum),
				GoodsUnit:          conf.ORDER_GOODS_UNIT,
				GoodsWeight:        float32(good.GoodsWeight),
				GoodsSpecification: float32(good.GoodsSpecification),
				GoodsNo:            good.GoodsNo,
				GoodsMargin:        float32(good.GoodsWeight),
				GoodsArranged:      0,
				GoodsNote:          good.GoodsNote,
				GoodsExtraWeight:   good.GoodsExtraWeight,
				IsDelete:           conf.UN_DELETE,
				InsertTime:         time.Now(),
				UpdateTime:         time.Now(),
			}
			if data.Status > 0 {
				addOrderData.Status = int8(data.Status)
			}
			if data.VerifierId > 0 && len(data.VerifierName) > 0 && len(data.VerifierNote) > 0 {
				addOrderData.VerifierId = data.VerifierId
				addOrderData.VerifierName = data.VerifierName
				addOrderData.VerifierNote = data.VerifierNote
				addOrderData.VerifyTime = time.Now()
			}
			if id, err := wtws_mysql.AddOrOrder(addOrderData); err != nil || id <= 0 {
				addOrderErrs = append(addOrderErrs, err)
			} else {
				truckOrderArr = append(truckOrderArr, request_entity.TruckOrders{
					OrderID:           int(id),
					GoodsLoadQuantity: float32(good.GoodsWeight),
				})
			}
			wg2.Done()
		}(good)
	}

	wg2.Wait()

	if len(addOrderErrs) > 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", truckOrderArr)
}

func DeleteOrder(data *request_entity.DeleteOrder) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.OrderIDs))

	deleteErrs := []error{}

	for _, id := range data.OrderIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByOrderId(&wtws_mysql.OrOrder{
				Id:         id,
				IsDelete:   conf.IS_DELETE,
				UpdateTime: time.Now(),
			}, []string{"IsDelete", "UpdateTime"}); err != nil {
				deleteErrs = append(deleteErrs, err)
			}
			wg.Done()
		}(id)
	}

	wg.Wait()

	if len(deleteErrs) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

func UpdateOrder(data *request_entity.UpdateOrder) common_struct.ResponseStruct {
	var dbOrder *wtws_mysql.OrOrder
	var err error
	dbOrder, err = wtws_mysql.GetOrOrderById(data.OrderID)
	if err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	updateOrderData := &wtws_mysql.OrOrder{
		Id: data.OrderID,

		OrderType:   int8(data.OrderType),
		ReceiveId:   data.ReceiveID,
		ReceiveName: data.ReceiveName,
		OriginId:    data.OriginID,
		OriginName:  data.OriginName,
		GoodsId:     data.GoodsID,
		GoodsNum:    int(data.GoodsNum),
		GoodsWeight: float32(data.GoodsWeight),
		GoodsNote:   data.GoodsNote,
		GoodsName:   data.GoodsName,

		UpdateTime: time.Now(),
	}

	updateOrderData.GoodsMargin = float32(data.GoodsWeight - dbOrder.GoodsArranged)

	if err := wtws_mysql.UpdateByOrderId(updateOrderData, []string{"OrderType", "ReceiveId", "ReceiveName", "OriginId", "OriginName", "GoodsId", "GoodsNum", "GoodsWeight", "GoodsNote", "GoodsName", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func CheckOrder(data *request_entity.CheckOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {

	if dbOrder, getOrderErr := wtws_mysql.GetOrOrderById(data.OrderID); getOrderErr != nil {
		logs.Error(fmt.Sprintf("[service]  查询订单ID:%d \t对应的订单失败，失败信息:%s", data.OrderID, getOrderErr.Error()))
		return common.ResponseStatus(-15, "", nil)
	} else if dbOrder == nil {
		logs.Error(fmt.Sprintf("[service]  查询订单ID:%d \t对应的订单失败", data.OrderID))
		return common.ResponseStatus(-15, "", nil)
	} else if dbOrder.Status != conf.ORDER_STATUS_WAIT {
		logs.Error(fmt.Sprintf("[service]  当前状态的订单不允许被审核，当前订单状态:%s", conf.ORDER_STATUS_MAP[int(dbOrder.Status)]))
		return common.ResponseStatus(-15, "", nil)
	}

	updateOrderData := &wtws_mysql.OrOrder{
		Id:           data.OrderID,
		VerifierNote: data.CheckNote,
		VerifierId:   userInfo.Id,
		VerifierName: userInfo.DisplayName,
		VerifyTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	if data.CheckType == 1 {
		updateOrderData.Status = conf.ORDER_STATUS_PASS
	} else {
		updateOrderData.Status = conf.ORDER_STATUS_REJECT
	}

	if err := wtws_mysql.UpdateByOrderId(updateOrderData, []string{"Status", "VerifierNote", "VerifierId", "VerifierName", "VerifyTime", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func InvalidOrder(data *request_entity.InvalidOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	if dbOrder, getDbOrderErr := wtws_mysql.GetOrOrderById(data.OrderID); getDbOrderErr != nil || dbOrder == nil {
		logs.Error(fmt.Sprintf("[service]  查询orderID：%d \t对应的订单失败", data.OrderID))
		return common.ResponseStatus(-15, "", nil)
	} else if dbOrder.Status != conf.ORDER_STATUS_WAIT && dbOrder.Status != conf.ORDER_STATUS_PASS {
		logs.Error(fmt.Sprintf("[service]  当前状态的订单不允许作废，当前订单状态:%s", conf.ORDER_STATUS_MAP[int(dbOrder.Status)]))
		return common.ResponseStatus(-15, "", nil)
	}
	updateOrderData := &wtws_mysql.OrOrder{
		Id:          data.OrderID,
		Status:      conf.ORDER_STATUS_FAILURE,
		InvalidId:   userInfo.Id,
		InvalidName: userInfo.DisplayName,
		InvalidTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := wtws_mysql.UpdateByOrderId(updateOrderData, []string{"Status", "InvalidId", "InvalidName", "InvalidTime", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}
	return common.ResponseStatus(14, "", nil)
}

func ExportExcel(dates []wtws_mysql.OrOrder) (buffer *bytes.Buffer) {

	//nowDate := time.Now()
	//nowDateTimeStr := nowDate.Format("20060102150405")
	//sheetName := "order_"+nowDateTimeStr

	wf := excelize.NewFile()
	sheetName := "销售单/采购单"
	wf.SetSheetName("Sheet1", sheetName)

	var orderExcelHeader = []string{"订单单号", "订单类型", "订单状态", "发货单位", "收货单位", "货物名称",
		"货物编号", "货品规格", "货品重量(吨)", "货品件数(件)", "货品余量", "已派车(吨)", "货品备注", "制单员名字",
		"审核员名字", "审核备注", "作废管理员名字", "创建时间", "审核时间", "作废时间"}

	wf.SetSheetRow(sheetName, "A1", &orderExcelHeader)
	for i, v := range dates {

		verifyTimeStr := v.VerifyTime.Format("2006-01-02 15:04:05")
		invalidTimeStr := v.InvalidTime.Format("2006-01-02 15:04:05")

		if verifyTimeStr == "0001-01-01 00:00:00" {
			verifyTimeStr = ""
		}

		if invalidTimeStr == "0001-01-01 00:00:00" {
			invalidTimeStr = ""
		}

		wf.SetSheetRow(sheetName, "A"+strconv.Itoa(i+2), &[]interface{}{
			v.OrderNo, conf.ORDER_TYPE_MAP[int(v.OrderType)], conf.ORDER_STATUS_MAP[int(v.Status)],
			v.OriginName, v.ReceiveName, v.GoodsName, v.GoodsNo, v.GoodsSpecification, v.GoodsWeight,
			v.GoodsNum, v.GoodsMargin, v.GoodsArranged, v.GoodsNote, v.CreatorName, v.VerifierName,
			v.VerifierNote, v.InvalidName, v.InsertTime.Format("2006-01-02 15:04:05"),
			verifyTimeStr, invalidTimeStr})
	}

	buffer, _ = wf.WriteToBuffer()

	return buffer

}

func DownAllOrder() []byte {

	allOrderList := wtws_mysql.GetAllOrderList()

	buffer := ExportExcel(allOrderList)

	return buffer.Bytes()

}

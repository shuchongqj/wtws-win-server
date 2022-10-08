package service

import (
	"bytes"
	"encoding/json"
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

func GetTruckOrderList(data *request_entity.TruckOrderList) common_struct.ResponseStruct {

	if orderList, count, err := wtws_mysql.GetTruckOrderList(data.PageNum, data.PageSize, data.OrderType,
		data.Status, data.ReceiveID, data.OriginID, data.TruckOrderNo, data.OrderNo, data.GoodsName, data.GoodsNo,
		data.InsertTimes, data.VerifyTimes); err != nil {
		return common.ResponseStatus(0, "", dto.TruckOrderList{
			List:  []wtws_mysql.OrTruckOrder{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.TruckOrderList{
			List:  orderList,
			Count: count,
		})
	}

}

func GetTruckOrderByVehicle(vehicleNumber string) common_struct.ResponseStruct {
	var truckOrders []wtws_mysql.OrTruckOrder
	var getTruckOrderErr error
	truckOrders, getTruckOrderErr = wtws_mysql.GetCheckedTruckOrderByVehicleNumber(vehicleNumber, conf.TRUCK_ORDER_STATUS_PASS)
	if getTruckOrderErr != nil || len(truckOrders) == 0 {
		return common.ResponseStatus(0, "", []wtws_mysql.OrTruckOrder{})
	}

	return common.ResponseStatus(0, "", truckOrders)
}

func AddTruckOrder(data *request_entity.AddTruckOrder, user *wtws_mysql.SUser) common_struct.ResponseStruct {

	var dbDriver *wtws_mysql.SDriver
	var getDriverErr error
	dbDriver, getDriverErr = wtws_mysql.GetSDriverById(data.DriverID)
	if dbDriver == nil || getDriverErr != nil {
		logs.Error("[service]  新增派车单时，查询对应的司机信息失败，失败信息:", getDriverErr.Error())
		return common.ResponseStatus(-13, "", nil)
	}

	var dirverTime time.Time
	var parseDriverTimeErr error
	if len(data.DriverTime) > 0 {
		dirverTime, parseDriverTimeErr = time.ParseInLocation("2006-01-02", data.DriverTime, time.Local)
	}

	wg2 := &sync.WaitGroup{}
	wg2.Add(len(data.Orders))
	var addTruckOrderErrs []error

	truckOrderNo := common.GenerateOrderNo("WL")

	for _, order := range data.Orders {
		go func(o request_entity.TruckOrders, d *wtws_mysql.SDriver) {

			var dbOrder *wtws_mysql.OrOrder
			var getOrderErr error
			dbOrder, getOrderErr = wtws_mysql.GetOrOrderById(o.OrderID)
			if dbOrder == nil || getOrderErr != nil {
				logs.Error(fmt.Sprintf("[service]  新增派车单时，查询orderID对应的订单失败，失败信息:%s", getOrderErr.Error()))
				addTruckOrderErrs = append(addTruckOrderErrs, getOrderErr)
				wg2.Done()
				return
			}

			subWg := &sync.WaitGroup{}
			subWg.Add(3)

			var dbReceive *wtws_mysql.OReceive
			var dbOrigin *wtws_mysql.OOrigin
			var dbCargoto *wtws_mysql.OCargoto
			var getReceiveErr, getOriginErr, getCarGotoErr error

			go func() {
				dbReceive, getReceiveErr = wtws_mysql.GetOReceiveById(dbOrder.ReceiveId)
				subWg.Done()
			}()

			go func() {
				dbOrigin, getOriginErr = wtws_mysql.GetOOriginById(dbOrder.OriginId)
				subWg.Done()
			}()

			go func() {
				dbCargoto, getCarGotoErr = wtws_mysql.GetOCargotoById(data.CargotoID)
				subWg.Done()
			}()

			subWg.Wait()

			if dbReceive == nil || getReceiveErr != nil ||
				dbOrigin == nil || getOriginErr != nil ||
				dbCargoto == nil || getCarGotoErr != nil {
				logs.Error(fmt.Sprintf("[service]  新增派车单时，查询对应的收发货地址和装卸货地失败，失败信息:%s", getOrderErr.Error()))
				addTruckOrderErrs = append(addTruckOrderErrs, getOrderErr)
				wg2.Done()
				return
			}

			truckOrderData := &wtws_mysql.OrTruckOrder{
				TruckOrderNo:       truckOrderNo,
				OrderID:            dbOrder.Id,
				OrderNo:            dbOrder.OrderNo,
				Status:             int8(conf.TRUCK_ORDER_STATUS_WAIT),
				OrderType:          dbOrder.OrderType,
				DriverId:           d.Id,
				DriverName:         d.DriverName,
				DriverTel:          d.Tel,
				VehicleNumber:      d.VehicleNumber,
				LimitTotalLoad:     fmt.Sprintf("%f", d.LimitTotalLoad),
				IsWeightLimit:      int8(data.IsWeightLimit),
				ContainerNo:        data.ContainerNo,
				ReceiveId:          dbOrder.ReceiveId,
				ReceiveTel:         dbReceive.Tel,
				ReceiveName:        dbReceive.Name,
				ReceiveAddress:     dbReceive.Address,
				OriginId:           dbOrigin.Id,
				OriginName:         dbOrigin.Name,
				OriginTel:          dbOrigin.Tel,
				OriginAddress:      dbOrigin.Address,
				CargotoId:          data.CargotoID,
				CargotoName:        dbCargoto.Name,
				GoodsId:            dbOrder.GoodsId,
				GoodsName:          dbOrder.GoodsName,
				GoodsWeight:        dbOrder.GoodsWeight,
				GoodsNo:            dbOrder.GoodsNo,
				GoodsNum:           dbOrder.GoodsNum,
				GoodsExtraWeight:   dbOrder.GoodsExtraWeight,
				GoodsUnit:          conf.ORDER_GOODS_UNIT,
				GoodsLoadQuantity:  o.GoodsLoadQuantity,
				GoodsSpecification: dbOrder.GoodsSpecification,
				GoodsNote:          dbOrder.GoodsNote,
				GoodsArranged:      0,
				BankName:           d.BankName,
				BankNo:             d.BankNo,
				BankUserName:       d.BankUserName,
				PaymentMethod:      int8(data.PaymentMethod),
				CreatorId:          user.Id,
				CreatorName:        user.DisplayName,
				IsDelete:           conf.UN_DELETE,
				InsertTime:         time.Now(),
				UpdateTime:         time.Now(),
			}

			if data.Status > 0 {
				truckOrderData.Status = int8(data.Status)
			}

			if data.VerifierId > 0 && len(data.VerifierName) > 0 && len(data.VerifierNote) > 0 {
				truckOrderData.VerifierId = data.VerifierId
				truckOrderData.VerifierName = data.VerifierName
				truckOrderData.VerifierNote = data.VerifierNote
				truckOrderData.VerifyTime = time.Now()
			}

			if parseDriverTimeErr == nil {
				truckOrderData.DriverTime = dirverTime
			}
			if id, err := wtws_mysql.AddOrTruckOrder(truckOrderData); err != nil || id <= 0 {
				addTruckOrderErrs = append(addTruckOrderErrs, err)
			} else {
				dbOrder.GoodsMargin -= o.GoodsLoadQuantity
				dbOrder.GoodsArranged += o.GoodsLoadQuantity
				dbOrder.Status = conf.ORDER_STATUS_HAS_TRUCK
				if updateOrderErr := wtws_mysql.UpdateOrOrderById(dbOrder, []string{"Status", "GoodsMargin", "GoodsArranged"}); updateOrderErr != nil {
					addTruckOrderErrs = append(addTruckOrderErrs, updateOrderErr)
				}
			}
			wg2.Done()

		}(order, dbDriver)
	}

	wg2.Wait()

	if len(addTruckOrderErrs) > 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", nil)
}

func AddSentDirectOrder(data *request_entity.AddSentDirectOrder, user *wtws_mysql.SUser) common_struct.ResponseStruct {
	addOrderData := request_entity.AddOrder{
		OrderType:    data.OrderType,
		OriginID:     data.OriginID,
		ReceiveID:    data.ReceiveID,
		Goods:        data.Goods,
		Status:       conf.ORDER_STATUS_HAS_TRUCK,
		VerifierId:   user.Id,
		VerifierName: user.DisplayName,
		VerifierNote: "直发/倒短派车单",
	}

	addOrderRes := AddOrder(&addOrderData, user)
	if addOrderRes.Code != 12 {
		logs.Error("[service] 给直发派车单创建订单数据失败")
		return common.ResponseStatus(-13, "", nil)
	}

	orderDataInterface := addOrderRes.Result
	orderDataBytes, _ := json.Marshal(orderDataInterface)
	orderDataArr := []request_entity.TruckOrders{}
	json.Unmarshal(orderDataBytes, &orderDataArr)

	return AddTruckOrder(&request_entity.AddTruckOrder{
		Orders:        orderDataArr,
		DriverID:      data.DriverID,
		IsWeightLimit: data.IsWeightLimit,
		CargotoID:     data.CargotoID,
		PaymentMethod: data.PaymentMethod,
		DriverTime:    data.DriverTime,
		ContainerNo:   data.ContainerNo,
		Status:        conf.TRUCK_ORDER_STATUS_PASS,
		VerifierId:    user.Id,
		VerifierName:  user.DisplayName,
		VerifierNote:  "直发/倒短派车单",
	}, user)

}

func DeleteTruckOrder(data *request_entity.DeleteTruckOrder) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.TruckOrderIDs))

	deleteErrs := []error{}

	for _, id := range data.TruckOrderIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByTruckOrderId(&wtws_mysql.OrTruckOrder{
				Id:         id,
				IsDelete:   conf.IS_DELETE,
				UpdateTime: time.Now(),
			}, []string{"IsDelete", "UpdateTime"}); err != nil {
				deleteErrs = append(deleteErrs, err)
				wg.Done()
				return
			}

			if dbTruckOrder, getTruckOrderErr := wtws_mysql.GetOrTruckOrderById(id); getTruckOrderErr != nil {
				logs.Error(fmt.Sprintf("[service]  审核订单，查询truckOrder失败，truckOrderID: %d   失败信息:%s", id, getTruckOrderErr.Error()))
				deleteErrs = append(deleteErrs, getTruckOrderErr)
				wg.Done()
				return
			} else if dbTruckOrder.Status == conf.TRUCK_ORDER_STATUS_WAIT || dbTruckOrder.Status == conf.TRUCK_ORDER_STATUS_PASS {
				updateOrderData := &wtws_mysql.OrOrder{
					Id: dbTruckOrder.OrderID,
				}

				updateOrderData.GoodsMargin += dbTruckOrder.GoodsLoadQuantity
				updateOrderData.GoodsArranged -= dbTruckOrder.GoodsLoadQuantity

				if updateOrderErr := wtws_mysql.UpdateOrOrderById(updateOrderData, []string{"GoodsMargin", "GoodsArranged"}); updateOrderErr != nil {
					deleteErrs = append(deleteErrs, updateOrderErr)
					wg.Done()
					return
				}
			}
			wg.Done()
			return

		}(id)
	}

	wg.Wait()

	if len(deleteErrs) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

func UpdateTruckOrder(data *request_entity.UpdateTruckOrder) common_struct.ResponseStruct {
	//var dbTruckOrder *wtws_mysql.OrTruckOrder
	//var err error
	//dbTruckOrder, err = wtws_mysql.GetOrTruckOrderById(data.TruckOrderID)
	//if err != nil {
	//	return common.ResponseStatus(-15, "", nil)
	//}

	updateTruckOrderData := &wtws_mysql.OrTruckOrder{
		Id: data.TruckOrderID,

		ReceiveId:         data.ReceiveID,
		OriginId:          data.OriginID,
		GoodsId:           data.GoodsID,
		GoodsNum:          int(data.GoodsNum),
		GoodsLoadQuantity: float32(data.GoodsWeight),
		GoodsNote:         data.GoodsNote,
		GoodsName:         data.GoodsName,

		UpdateTime: time.Now(),
	}

	if err := wtws_mysql.UpdateByTruckOrderId(updateTruckOrderData, []string{"TruckOrderType", "ReceiveId", "ReceiveName", "OriginId", "OriginName", "GoodsId", "GoodsNum", "GoodsWeight", "GoodsNote", "GoodsName", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func CheckTruckOrder(data *request_entity.CheckTruckOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {

	var dbTruckOrder *wtws_mysql.OrTruckOrder
	var getTruckOrderErr error
	if dbTruckOrder, getTruckOrderErr = wtws_mysql.GetOrTruckOrderById(data.TruckOrderID); getTruckOrderErr != nil {
		logs.Error(fmt.Sprintf("[service]  审核订单，查询truckOrder失败，truckOrderID: %d   失败信息:%s", data.TruckOrderID, getTruckOrderErr.Error()))
		return common.ResponseStatus(-15, "", nil)
	} else if dbTruckOrder.Status != conf.ORDER_STATUS_WAIT {
		logs.Error("[service]  审核订单，审核的订单非待审核状态")
		return common.ResponseStatus(-15, "审核的订单非待审核状态", nil)
	}

	updateTruckOrderData := &wtws_mysql.OrTruckOrder{
		Id:           data.TruckOrderID,
		VerifierNote: data.CheckNote,
		VerifierId:   userInfo.Id,
		VerifierName: userInfo.DisplayName,
		VerifyTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	if data.CheckType == 1 {
		updateTruckOrderData.Status = conf.TRUCK_ORDER_STATUS_PASS
	} else {
		updateTruckOrderData.Status = conf.TRUCK_ORDER_STATUS_REJECT
	}

	if err := wtws_mysql.UpdateByTruckOrderId(updateTruckOrderData, []string{"Status", "VerifierNote", "VerifierId", "VerifierName", "VerifyTime", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	if updateTruckOrderData.Status == conf.TRUCK_ORDER_STATUS_REJECT {
		updateOrderData := &wtws_mysql.OrOrder{
			Id: dbTruckOrder.OrderID,
		}

		updateOrderData.GoodsMargin += dbTruckOrder.GoodsLoadQuantity
		updateOrderData.GoodsArranged -= dbTruckOrder.GoodsLoadQuantity

		if updateOrderErr := wtws_mysql.UpdateOrOrderById(updateOrderData, []string{"GoodsMargin", "GoodsArranged"}); updateOrderErr != nil {
			return common.ResponseStatus(-15, "", nil)
		}
	}

	return common.ResponseStatus(14, "", nil)
}

func InvalidTruckOrder(data *request_entity.InvalidTruckOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	var dbTruckOrder *wtws_mysql.OrTruckOrder
	var getTruckOrderErr error
	if dbTruckOrder, getTruckOrderErr = wtws_mysql.GetOrTruckOrderById(data.TruckOrderID); getTruckOrderErr != nil || dbTruckOrder == nil {
		logs.Error(fmt.Sprintf("[service]  查询truckOrder失败，truckOrderID: %d   失败信息:%s", data.TruckOrderID, getTruckOrderErr.Error()))
		return common.ResponseStatus(-15, "", nil)
	}
	if dbTruckOrder.Status != conf.TRUCK_ORDER_STATUS_WAIT && dbTruckOrder.Status != conf.TRUCK_ORDER_STATUS_PASS {
		logs.Error(fmt.Sprintf("[service]  当前状态的派车单不允许作废，当前状态：%s", conf.TRUCK_ORDER_STATUS_MAP[int(dbTruckOrder.Status)]))
		return common.ResponseStatus(-15, "", nil)
	}
	updateTruckOrderData := &wtws_mysql.OrTruckOrder{
		Id:          data.TruckOrderID,
		Status:      conf.TRUCK_ORDER_STATUS_FAILURE,
		InvalidId:   userInfo.Id,
		InvalidName: userInfo.DisplayName,
		InvalidTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := wtws_mysql.UpdateByTruckOrderId(updateTruckOrderData, []string{"Status", "InvalidId", "InvalidName", "InvalidTime", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	if dbTruckOrder.Status == conf.TRUCK_ORDER_STATUS_PASS || dbTruckOrder.Status == conf.TRUCK_ORDER_STATUS_WAIT {
		updateOrderData := &wtws_mysql.OrOrder{
			Id: dbTruckOrder.OrderID,
		}

		updateOrderData.GoodsMargin += dbTruckOrder.GoodsLoadQuantity
		updateOrderData.GoodsArranged -= dbTruckOrder.GoodsLoadQuantity

		if updateOrderErr := wtws_mysql.UpdateOrOrderById(updateOrderData, []string{"GoodsMargin", "GoodsArranged"}); updateOrderErr != nil {
			return common.ResponseStatus(-15, "", nil)
		}
	}

	return common.ResponseStatus(14, "", nil)
}

func ExportTruckOrderExcel(dates []wtws_mysql.OrTruckOrder) (buffer *bytes.Buffer) {

	wf := excelize.NewFile()
	sheetName := "派车单"
	wf.SetSheetName("Sheet1", sheetName)

	var truckOderExcelHeader = []string{"派车单单号", "关联订单号", "订单类型", "派车单状态", "司机名字",
		"司机电话", "车牌号", "车辆载重", "是否限重", "运输时间", "集装箱/柜号", "收货单位", "收货电话",
		"收货地址", "发件单位", "发货电话", "发货地址", "装卸货地点名称", "货品名称", "货品编号", "货品数量",
		"货品装车量", "货品规格", "货品备注", "已运输量/已派车", "银行名称", "银行卡号", "银行卡开户人姓名",
		"结算方式", "制单人名字", "制单时间", "审核人名字", "审核备注", "订单审核时间", "作废管理员的", "作废时间"}

	wf.SetColWidth(sheetName, "A", "B", 35)
	wf.SetColWidth(sheetName, "B", "I", 12)
	wf.SetColWidth(sheetName, "J", "J", 20)
	wf.SetColWidth(sheetName, "K", "K", 12)
	wf.SetColWidth(sheetName, "L", "L", 40)
	wf.SetColWidth(sheetName, "M", "N", 12)
	wf.SetSheetRow(sheetName, "A1", &truckOderExcelHeader)
	for i, v := range dates {

		driverTimeStr := v.DriverTime.Format("2006-01-02 15:04:05")
		verifyTimeStr := v.VerifyTime.Format("2006-01-02 15:04:05")
		invalidTimeStr := v.InvalidTime.Format("2006-01-02 15:04:05")
		insertTimeStr := v.InsertTime.Format("2006-01-02 15:04:05")

		if verifyTimeStr == "0001-01-01 00:00:00" {
			verifyTimeStr = ""
		}

		if invalidTimeStr == "0001-01-01 00:00:00" {
			invalidTimeStr = ""
		}

		wf.SetSheetRow(sheetName, "A"+strconv.Itoa(i+2), &[]interface{}{
			v.TruckOrderNo, v.OrderNo, conf.TRUCK_ORDER_TYPE_MAP[int(v.OrderType)], conf.TRUCK_ORDER_STATUS_MAP[int(v.Status)], v.DriverName,
			v.DriverTel, v.VehicleNumber, v.LimitTotalLoad, conf.TRUCK_ORDER_IS_LIMIT_LOAD[int(v.IsWeightLimit)], driverTimeStr, v.ContainerNo, v.ReceiveName, v.ReceiveTel,
			v.ReceiveAddress, v.OriginName, v.OriginTel, v.OriginAddress, v.CargotoName, v.GoodsName, v.GoodsNo, v.GoodsNum,
			v.GoodsLoadQuantity, v.GoodsSpecification, v.GoodsNote, v.GoodsArranged, v.BankName, v.BankNo, v.BankUserName,
			conf.PAYMENT_METHOD_MAP[int(v.PaymentMethod)], v.CreatorName, insertTimeStr, v.VerifierName, v.VerifierNote, verifyTimeStr, v.InvalidName, invalidTimeStr,
		})

	}

	buffer, _ = wf.WriteToBuffer()

	return buffer

}

func DownAllTruckOrder() []byte {

	allTruckOrderList := wtws_mysql.GetAllTruckOrderList()

	buffer := ExportTruckOrderExcel(allTruckOrderList)

	return buffer.Bytes()

}

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

func GetWeighOrderInfo(weighOrderID int) common_struct.ResponseStruct {
	if dbWeighOrder, err := wtws_mysql.GetOrWeighOrderById(weighOrderID); err != nil || dbWeighOrder == nil {
		return common.ResponseStatus(-1, "", nil)
	} else {
		return common.ResponseStatus(0, "", dbWeighOrder)
	}
}

func GetWeighOrderList(data *request_entity.WeighOrderList) common_struct.ResponseStruct {

	if weighOrderList, count, err := wtws_mysql.GetWeighOrderList(data.PageNum, data.PageSize, data.OrderType, data.Status,
		data.ReceiveID, data.OriginID, data.WeighOrderNo, data.TruckOrderNo, data.OrderNo, data.VehicleNumber, data.GoodsName,
		data.GoodsNo, data.InsertTimes, data.WarehouseConfirmTime, data.FinancialConfirmTime); err != nil {
		return common.ResponseStatus(0, "", dto.WeighOrderList{
			List:  []wtws_mysql.OrWeighOrder{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.WeighOrderList{
			List:  weighOrderList,
			Count: count,
		})
	}

}

func AddWeighOrder(data *request_entity.AddWeighOrder, user *wtws_mysql.SUser) common_struct.ResponseStruct {

	return common.ResponseStatus(12, "", nil)
}

func DeleteWeighOrder(data *request_entity.DeleteWeighOrder) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.WeighOrderIDs))

	deleteErrs := []error{}

	for _, id := range data.WeighOrderIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByWeighOrderId(&wtws_mysql.OrWeighOrder{
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

func UpdateWeighOrder(data *request_entity.UpdateWeighOrder) common_struct.ResponseStruct {

	updateWeighOrderData := &wtws_mysql.OrWeighOrder{
		Id: data.WeighOrderID,

		ReceiveId:   data.ReceiveID,
		ReceiveName: data.ReceiveName,
		OriginId:    data.OriginID,
		OriginName:  data.OriginName,
		GoodsNum:    int(data.GoodsNum),
		GoodsWeight: float32(data.GoodsWeight),
		GoodsName:   data.GoodsName,

		UpdateTime: time.Now(),
	}

	if err := wtws_mysql.UpdateByWeighOrderId(updateWeighOrderData, []string{"WeighOrderType", "ReceiveId", "ReceiveName", "OriginId", "OriginName", "GoodsId", "GoodsNum", "GoodsWeight", "GoodsNote", "GoodsName", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func WarehouseCheckWeighOrder(data *request_entity.WarehouseCheckWeighOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {

	updateWeighOrderData := &wtws_mysql.OrWeighOrder{
		Id:                   data.WeighOrderID,
		Status:               conf.WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM,
		GoodsBatchNo:         data.GoodsBatchNo,
		OtherWight:           data.OtherWight,
		WarehouseId:          userInfo.Id,
		WarehouseName:        userInfo.DisplayName,
		WarehouseConfirmTime: time.Now(),
		UpdateTime:           time.Now(),
	}

	updateWeighOrderCols := []string{
		"Status", "GoodsBatchNo", "OtherWight", "WarehouseId", "WarehouseName", "WarehouseConfirmTime", "UpdateTime",
	}

	if data.TruckOrderID > 0 {
		var dbTruckOrder *wtws_mysql.OrTruckOrder
		var dbGoods *wtws_mysql.GGoods
		var dbCategory *wtws_mysql.GCategory
		var getTruckOrderErr, getGoodsErr, getCategoryErr error

		dbTruckOrder, getTruckOrderErr = wtws_mysql.GetOrTruckOrderById(data.TruckOrderID)
		if getTruckOrderErr != nil || dbTruckOrder == nil {
			logs.Error("[service]  查询到过磅单对应的派车单失败，失败信息:", getTruckOrderErr.Error())
			return common.ResponseStatus(-1, "", nil)
		}

		dbGoods, getGoodsErr = wtws_mysql.GetGGoodsById(dbTruckOrder.GoodsId)
		if getGoodsErr != nil || dbGoods == nil {
			logs.Error("[service]  查询过磅单更换的货品信息失败，失败信息:", getGoodsErr.Error())
			return common.ResponseStatus(-1, "", nil)
		}

		dbCategory, getCategoryErr = wtws_mysql.GetGCategoryById(dbGoods.CategoryId)
		if getCategoryErr != nil || dbCategory == nil {
			logs.Error("[service]  查询过磅单更换的货品的类别信息失败，失败信息:", getCategoryErr.Error())
			return common.ResponseStatus(-1, "", nil)
		}
		updateWeighOrderData.DriverId = dbTruckOrder.DriverId
		updateWeighOrderData.DriverName = dbTruckOrder.DriverName

		updateWeighOrderData.OrderType = dbTruckOrder.OrderType
		updateWeighOrderData.TruckOrderId = dbTruckOrder.Id
		updateWeighOrderData.TruckOrderNo = dbTruckOrder.TruckOrderNo
		updateWeighOrderData.OrderId = dbTruckOrder.OrderID
		updateWeighOrderData.OrderNo = dbTruckOrder.OrderNo
		updateWeighOrderData.VehicleNumber = dbTruckOrder.VehicleNumber
		updateWeighOrderData.IsWeightLimit = dbTruckOrder.IsWeightLimit
		updateWeighOrderData.ContainerNo = dbTruckOrder.ContainerNo
		updateWeighOrderData.OriginId = dbTruckOrder.OriginId
		updateWeighOrderData.OriginName = dbTruckOrder.OriginName
		updateWeighOrderData.ReceiveId = dbTruckOrder.ReceiveId
		updateWeighOrderData.ReceiveName = dbTruckOrder.ReceiveName
		updateWeighOrderData.CargotoId = dbTruckOrder.CargotoId
		updateWeighOrderData.CargotoName = dbTruckOrder.CargotoName
		updateWeighOrderData.BankNo = dbTruckOrder.BankNo
		updateWeighOrderData.BankName = dbTruckOrder.BankName

		updateWeighOrderData.GoodsName = dbTruckOrder.GoodsName
		updateWeighOrderData.GoodsNo = dbTruckOrder.GoodsNo
		updateWeighOrderData.GoodsNum = dbTruckOrder.GoodsNum
		updateWeighOrderData.GoodsSpecification = dbTruckOrder.GoodsSpecification
		updateWeighOrderData.GoodsUnit = dbTruckOrder.GoodsUnit
		updateWeighOrderData.GoodsWeight = dbTruckOrder.GoodsWeight
		updateWeighOrderData.GoodsExtraWeight = float32(dbGoods.ExtraWeight)
		updateWeighOrderData.GoodsCategoryName = dbCategory.Name
		updateWeighOrderData.GoodsBagWeight = float32(dbGoods.BagWeight)
		updateWeighOrderData.GoodsDeductWeight = float32(dbGoods.DeductWeight)

		updateWeighOrderCols = append(updateWeighOrderCols, []string{
			"DriverId", "DriverName", "Status", "GoodsName", "GoodsNo", "GoodsNum", "GoodsSpecification",
			"GoodsUnit", "GoodsWeight", "GoodsExtraWeight", "GoodsCategoryName", "GoodsBagWeight", "OrderType",
			"TruckOrderId", "TruckOrderNo", "OrderId", "OrderNo", "VehicleNumber", "IsWeightLimit", "ContainerNo",
			"OriginId", "OriginName", "ReceiveId", "ReceiveName", "CargotoId", "CargotoName", "BankNo", "BankName",
			"GoodsDeductWeight",
		}...)
	}

	if err := wtws_mysql.UpdateByWeighOrderId(updateWeighOrderData, updateWeighOrderCols); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}

func InvalidWeighOrder(data *request_entity.InvalidWeighOrder, userInfo *wtws_mysql.SUser) common_struct.ResponseStruct {
	if dbWeighOrder, getWeighOrderErr := wtws_mysql.GetOrWeighOrderById(data.WeighOrderID); getWeighOrderErr != nil || dbWeighOrder == nil {
		logs.Error(fmt.Sprintf("[service]  未查询到过磅单ID:%d \t对应的派车单。", data.WeighOrderID))
		return common.ResponseStatus(-15, "", nil)
	} else if dbWeighOrder.Status != conf.WEIGH_ORDER_STATUS_PROCESS && dbWeighOrder.Status != conf.WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM {
		logs.Error(fmt.Sprintf("[service]  当前状态的过磅单不允许作废，当前过磅单状态:%s", conf.WEIGH_ORDER_STATUS_MAP[int(dbWeighOrder.Status)]))
		return common.ResponseStatus(-15, "", nil)
	}
	updateWeighOrderData := &wtws_mysql.OrWeighOrder{
		Id:          data.WeighOrderID,
		Status:      conf.WEIGH_ORDER_STATUS_FAILURE,
		InvalidId:   userInfo.Id,
		InvalidName: userInfo.DisplayName,
		InvalidTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := wtws_mysql.UpdateByWeighOrderId(updateWeighOrderData, []string{"Status", "InvalidId", "InvalidName", "InvalidTime", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}
	return common.ResponseStatus(14, "", nil)
}

func ExportWeighOrderExcel(dates []wtws_mysql.OrWeighOrder) (buffer *bytes.Buffer) {
	wf := excelize.NewFile()
	sheetName := "过磅单"
	wf.SetSheetName("Sheet1", sheetName)

	var orderExcelHeader = []string{"过磅单号", "派送单号", "关联订单号", "状态", "司机姓名", "订单类型",
		"过磅车牌号", "车辆最大载重", "是否限重", "集装箱柜号", "毛重", "皮重", "净重", "扣杂扣重", "过磅备注",
		"货品", "货品编号", "生产批号", "货品数量", "货品规格", "货品单位", "货品重量", "允许额外可配发重量",
		"产品类别", "每吨扣kg数（吨）", "编织袋重量", "发货单位", "收货单位", "装卸货地点", "银行", "银行卡号",
		"仓库管理员姓名", "过磅皮重时间", "过磅毛重时间", "过磅结束时间", "仓库确认时间", "打印票据时间", "记录创建时间"}

	wf.SetSheetRow(sheetName, "A1", &orderExcelHeader)
	for i, v := range dates {

		tarTimeStr := v.TareTime.Format("2006-01-02 15:04:05")
		roughTimeStr := v.RoughTime.Format("2006-01-02 15:04:05")
		weightOverTimeStr := v.WeightOverTime.Format("2006-01-02 15:04:05")
		warehouseConfirmTimeStr := v.WarehouseConfirmTime.Format("2006-01-02 15:04:05")
		printTimeStr := v.PrintTime.Format("2006-01-02 15:04:05")

		if tarTimeStr == "0001-01-01 00:00:00" {
			tarTimeStr = ""
		}

		if roughTimeStr == "0001-01-01 00:00:00" {
			roughTimeStr = ""
		}

		if weightOverTimeStr == "0001-01-01 00:00:00" {
			weightOverTimeStr = ""
		}

		if warehouseConfirmTimeStr == "0001-01-01 00:00:00" {
			warehouseConfirmTimeStr = ""
		}

		if printTimeStr == "0001-01-01 00:00:00" {
			printTimeStr = ""
		}

		wf.SetSheetRow(sheetName, "A"+strconv.Itoa(i+2), &[]interface{}{v.WeighOrderNo, v.TruckOrderNo, v.OrderNo,
			conf.WEIGH_ORDER_STATUS_MAP[int(v.Status)], v.DriverName, conf.ORDER_TYPE_MAP[int(v.OrderType)],
			v.VehicleNumber, v.LimitTotalLoad, conf.TRUCK_ORDER_IS_LIMIT_LOAD[int(v.IsWeightLimit)], v.ContainerNo,
			v.RoughWight, v.TareWight, v.NetWight, v.OtherWight, v.WeighOrderNote, v.GoodsName, v.GoodsNo,
			v.GoodsBatchNo, v.GoodsNum, v.GoodsSpecification, v.GoodsUnit, v.GoodsWeight, v.GoodsExtraWeight,
			v.GoodsCategoryName, v.GoodsDeductWeight, v.GoodsBagWeight, v.OriginName, v.ReceiveName, v.CargotoName,
			v.BankName, v.BankNo, v.WarehouseName, tarTimeStr, roughTimeStr, weightOverTimeStr, warehouseConfirmTimeStr,
			printTimeStr, v.InsertTime.Format("2006-01-02 15:04:05")})
	}

	buffer, _ = wf.WriteToBuffer()

	return buffer

}

func DownAllWeighOrder() []byte {

	allWeighOrderList := wtws_mysql.GetAllWeighOrderList()

	buffer := ExportWeighOrderExcel(allWeighOrderList)

	return buffer.Bytes()

}

func ScanVehicle(data *request_entity.ScanVehicle) common_struct.DeveiceResponseStruct {
	truckOrderList, _ := wtws_mysql.GetLatestVehicleTruckOrder(data.Vehicle, data.TransportTime)
	responseData := []dto.ScanVehicleData{}
	for _, tr := range truckOrderList {

		var dbGood *wtws_mysql.GGoods
		var getDbGoodErr error
		if dbGood, getDbGoodErr = wtws_mysql.GetGGoodsById(tr.GoodsId); getDbGoodErr != nil || dbGood == nil {
			logs.Error("[service]  查询派车单关联的货品信息失败，失败信息:", getDbGoodErr.Error())
			continue
		}

		responseData = append(responseData, dto.ScanVehicleData{
			//Mid:      tr.Id,
			//Mly:      conf.ORDER_TYPE_MAP[int(tr.OrderType)],
			//Mno:      tr.OrderNo,
			//Mzdrid:   tr.CreatorId,
			//Mzdrname: tr.CreatorName,
			//Msjid:    tr.DriverId,
			//Msjname:  tr.DriverName,
			//Msjcp:    tr.VehicleNumber,
			//Msjtel:   tr.DriverTel,
			//Msjzd:    tr.LimitTotalLoad,
			//Mshdwid:  tr.ReceiveId,
			//Mshdh:    tr.ReceiveTel,
			//Mshdz:    tr.ReceiveAddress,
			//Mfhdwid:  tr.OriginId,
			//Mfhdh:    tr.OriginTel,
			//Mfhdz:    tr.OriginAddress,
			//Myssj:    fmt.Sprintf("%d", tr.DriverTime.Unix()/1e6),
			//Mshrid:   tr.VerifierId,
			//Mshrname: tr.VerifierName,
			//Mshzt:    conf.DEVICE_TRUCK_ORDER_STATUS_MAP[int(tr.Status)],
			//Mshbz:    tr.VerifierNote,
			//Mcpid:    tr.GoodsId,
			//Mcpname:  tr.GoodsName,
			//Mcpbh:    tr.GoodsNo,
			//Mcpsl:    tr.GoodsNum,
			//Mcpdw:    tr.GoodsUnit,
			//Mcpzl:    tr.GoodsWeight,
			//Mcpgg:    fmt.Sprintf("%f", tr.GoodsSpecification),
			//Mcpbz:    tr.GoodsNote,
			//Myh:      tr.BankName,
			//Myhkh:    tr.BankNo,
			//Mysl:     tr.GoodsArranged,
			//Voidtime: int(tr.InvalidTime.Unix() / 1e6),
			////Over:          0,
			//Xianzhong: int(tr.IsWeightLimit),
			//ID:        tr.Id,
			////Cateid:        ,
			//Prname: tr.GoodsName,
			//Pnum:   fmt.Sprintf("%d", tr.GoodsNum),
			//Spec:   float64(tr.GoodsSpecification),
			//Goodsbianhao:  dbGood.GoodNo,

			Allow:         tr.GoodsExtraWeight,
			Bank:          tr.BankName,
			Banknum:       tr.BankNo,
			Createtime:    tr.InsertTime.Format("2006-01-02 15:04:05"),
			Deduction:     dbGood.DeductWeight,
			Drivername:    tr.DriverName,
			Drivertel:     tr.DriverTel,
			Goodsname:     dbGood.Name,
			Goodsnum:      tr.GoodsNum,
			Goodsspec:     fmt.Sprintf("%f", dbGood.Specification),
			Goodsunit:     tr.GoodsUnit,
			Goodsweight:   tr.GoodsWeight,
			Jid:           tr.CreatorId,
			Makeid:        tr.TruckOrderNo,
			Pmid:          tr.Id,
			Source:        conf.ORDER_TYPE_MAP[int(tr.OrderType)],
			Sourceno:      tr.TruckOrderNo,
			Transporttime: tr.DriverTime.Format("2006-01-02 15:04:05"),
			Vehicle:       tr.VehicleNumber,
			Vehicleload:   tr.LimitTotalLoad,
			Weight:        dbGood.BagWeight,
			Mgh:           tr.ContainerNo,
			Mdd:           tr.CargotoName,
			Mjsfs:         conf.PAYMENT_METHOD_MAP[int(tr.PaymentMethod)],
			Mshdw:         tr.ReceiveName,
			Mfhdw:         tr.OriginName,
			Wckname:       tr.CreatorName,
		})

	}

	return common_struct.DeveiceResponseStruct{
		Status:  "1",
		Message: "获取成功",
		Data:    responseData,
	}

}

func TareWight(data *request_entity.TareWight) common_struct.DeveiceResponseStruct {

	var dbDriver *wtws_mysql.SDriver
	var truckOrderList []wtws_mysql.OrTruckOrder
	var truckOrder wtws_mysql.OrTruckOrder

	var getDriverErr, getTruckOrderListErr error

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		truckOrderList, getTruckOrderListErr = wtws_mysql.GetLatestVehicleTruckOrder(data.Wvehicle, "")
		wg.Done()
	}()

	go func() {
		dbDriver, getDriverErr = wtws_mysql.GetSDriverByVehicleNum(data.Wvehicle)
		wg.Done()
	}()

	wg.Wait()

	if getTruckOrderListErr != nil ||
		truckOrderList == nil ||
		len(truckOrderList) == 0 ||
		getDriverErr != nil ||
		dbDriver == nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "处理失败",
		}
	}

	truckOrder = truckOrderList[0]

	var dbGoods *wtws_mysql.GoodListItem
	var getGoodsErr error
	dbGoods, getGoodsErr = wtws_mysql.GetGoodsAndCategoryByGoodsID(truckOrder.GoodsId)
	if getGoodsErr != nil || dbGoods == nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "处理失败",
		}
	}

	tarTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.GrossTime, time.Local)
	tarWeigh, _ := strconv.ParseFloat(data.TareWight, 32)

	if weighOrderID, addErr := wtws_mysql.AddOrWeighOrder(&wtws_mysql.OrWeighOrder{
		WeighOrderNo:       data.Wno,
		Status:             conf.WEIGH_ORDER_STATUS_PROCESS,
		DriverId:           dbDriver.Id,
		DriverName:         dbDriver.DriverName,
		OrderType:          truckOrder.OrderType,
		IsDelete:           conf.UN_DELETE,
		TruckOrderId:       truckOrder.Id,
		TruckOrderNo:       truckOrder.TruckOrderNo,
		OrderId:            truckOrder.OrderID,
		OrderNo:            truckOrder.OrderNo,
		VehicleNumber:      truckOrder.VehicleNumber,
		IsWeightLimit:      truckOrder.IsWeightLimit,
		LimitTotalLoad:     truckOrder.LimitTotalLoad,
		ContainerNo:        truckOrder.ContainerNo,
		TareWight:          float32(tarWeigh),
		RoughWight:         0,
		NetWight:           0,
		TareTime:           tarTime,
		WeighOrderNote:     data.Wno,
		GoodsName:          truckOrder.GoodsName,
		GoodsNo:            truckOrder.OrderNo,
		GoodsBatchNo:       "",
		GoodsNum:           truckOrder.GoodsNum,
		GoodsSpecification: truckOrder.GoodsSpecification,
		GoodsUnit:          conf.ORDER_GOODS_UNIT,
		GoodsWeight:        truckOrder.GoodsWeight,
		GoodsExtraWeight:   truckOrder.GoodsExtraWeight,
		GoodsCategoryName:  dbGoods.CategoryName,
		GoodsDeductWeight:  float32(dbGoods.DeductWeight),
		GoodsBagWeight:     float32(dbGoods.BagWeight),
		OriginId:           truckOrder.OriginId,
		OriginName:         truckOrder.OriginName,
		ReceiveId:          truckOrder.ReceiveId,
		ReceiveName:        truckOrder.ReceiveName,
		CargotoId:          truckOrder.CargotoId,
		CargotoName:        truckOrder.CargotoName,
		BankName:           truckOrder.BankName,
		BankNo:             truckOrder.BankNo,
		//WarehouseId:        0,
		//WarehouseName:      "",
		//FinancialId:          0,
		//FinancialNote:        "",
		//InvalidId:            0,
		//InvalidName:          "",
		//InvalidTime:          time.Time{},
		//RoughTime:            time.Time{},
		//WeightOverTime:       time.Time{},
		//WarehouseConfirmTime: time.Time{},
		//FinancialConfirmTime: time.Time{},
		InsertTime: time.Now(),
		UpdateTime: time.Now(),
	}); addErr != nil || weighOrderID == 0 {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "处理失败",
			Data:    nil,
		}
	}

	return common_struct.DeveiceResponseStruct{
		Status:  "1",
		Message: "处理成功",
		Data:    nil,
	}
}

func CheckWareHouseGoods(data *request_entity.CheckWareHouseGoods) common_struct.DeveiceResponseStruct {

	var weighOrder *wtws_mysql.OrWeighOrderItem
	var getWeighOrderErr error
	weighOrder, getWeighOrderErr = wtws_mysql.GetWeighOrderByWeighOrderNo(data.Wno)
	if getWeighOrderErr != nil || weighOrder == nil {
		logs.Error(fmt.Sprintf("[service]  根据过磅单号：%s \t查询过磅单失败，失败信息:%s", data.Wno, getWeighOrderErr.Error()))
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "未选择货品",
			Data:    nil,
		}
	} else if weighOrder.Status != conf.WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM && weighOrder.Status != conf.WEIGH_ORDER_STATUS_WAITE_FINISH {
		logs.Error(fmt.Sprintf("[service]  根据过磅单号：%s \t查询到的过磅单未选择货品", data.Wno))
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "未选择货品",
			Data:    nil,
		}
	}

	responseData := dto.WareHouseCheckGoods{
		Wid:        weighOrder.Id,
		Wno:        weighOrder.WeighOrderNo,
		Wdriver:    weighOrder.DriverName,
		DriverTel:  weighOrder.DriverTel,
		Wvecgucle:  weighOrder.VehicleNumber,
		Wdruverid:  weighOrder.DriverId,
		OrderNum:   weighOrder.OrderNo,
		Upid:       0,
		Pmid:       weighOrder.TruckOrderId,
		WGoods:     weighOrder.GoodsName,
		Pihao:      weighOrder.GoodsBatchNo,
		Kou:        weighOrder.OtherWight,
		Wnum:       weighOrder.GoodsNum,
		Wspec:      weighOrder.GoodsSpecification,
		Wunit:      weighOrder.GoodsUnit,
		Wweight:    weighOrder.GoodsWeight,
		Wdeduction: weighOrder.GoodsDeductWeight,
		Wallow:     weighOrder.GoodsExtraWeight,
		Bzdweight:  weighOrder.GoodsBagWeight,
		Wtype:      conf.ORDER_TYPE_MAP[int(weighOrder.OrderType)],
		Wbianhao:   weighOrder.GoodsNo,
		Wghdw:      weighOrder.OriginName,
		Wshdw:      weighOrder.ReceiveName,
		Wjzxgh:     weighOrder.ContainerNo,
		Wzxhdd:     weighOrder.CargotoName,
		Wckname:    weighOrder.WarehouseName,
		Wxianz:     conf.TRUCK_ORDER_IS_LIMIT_LOAD[int(weighOrder.IsWeightLimit)],
		Wsjzz:      weighOrder.LimitTotalLoad,
	}

	if weighOrder.Status >= conf.WEIGH_ORDER_STATUS_WAITE_FINISH {
		responseData.Upid = 1
	}

	return common_struct.DeveiceResponseStruct{
		Status:  "1",
		Message: "获取成功",
		Data:    responseData,
	}
}

func WaitFinish(data *request_entity.WaitFinish) common_struct.DeveiceResponseStruct {
	if data.Upid == "1" {
		if updateErr := wtws_mysql.UpdateOrWeighOrderStatusByWeighOrderNo(conf.WEIGH_ORDER_STATUS_WAITE_FINISH,
			conf.WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM, data.Wno); updateErr != nil {
			return common_struct.DeveiceResponseStruct{
				Status:  "0",
				Message: "接收失败",
			}
		}
	}
	return common_struct.DeveiceResponseStruct{
		Status:  "1",
		Message: "接收成功",
	}
}

func FinishWeighOrder(data *request_entity.FinishWeighOrder) common_struct.DeveiceResponseStruct {
	if data.Over != "1" {
		return common_struct.DeveiceResponseStruct{
			Status:  "1",
			Message: "接收成功",
		}
	}

	var dbWeighOrder *wtws_mysql.OrWeighOrderItem
	var dbTruckOrder *wtws_mysql.OrTruckOrder
	var dbOrder *wtws_mysql.OrOrder
	var getWeighOrderErr, getTruckOrderErr, getOrderErr error

	dbWeighOrder, getWeighOrderErr = wtws_mysql.GetWeighOrderByWeighOrderNo(data.WeighOrderNo)
	if getWeighOrderErr != nil || dbWeighOrder == nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		dbTruckOrder, getTruckOrderErr = wtws_mysql.GetOrTruckOrderById(dbWeighOrder.TruckOrderId)
		wg.Done()
	}()

	go func() {
		dbOrder, getOrderErr = wtws_mysql.GetOrOrderById(dbWeighOrder.OrderId)
		wg.Done()
	}()

	wg.Wait()

	if getTruckOrderErr != nil || getOrderErr != nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
	}

	roughTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.RoughTimeStr, time.Local)
	tareTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.TareTimeStr, time.Local)
	printTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.PrintTimeStr, time.Local)

	tarWight64, _ := strconv.ParseFloat(data.TareWight, 32)
	roughWight64, _ := strconv.ParseFloat(data.RoughWight, 32)
	netWight64, _ := strconv.ParseFloat(data.NetWight, 32)

	updateWeighOrderData := &wtws_mysql.OrWeighOrder{
		Id:             dbWeighOrder.Id,
		Status:         conf.WEIGH_ORDER_STATUS_FINISH,
		TareWight:      float32(tarWight64),
		RoughWight:     float32(roughWight64),
		NetWight:       float32(netWight64),
		RoughTime:      roughTime,
		PrintTime:      printTime,
		TareTime:       tareTime,
		WeightOverTime: time.Now(),
		UpdateTime:     time.Now(),
	}

	updateWeighOrderCols := []string{
		"Status", "TareWight", "RoughWight", "NetWight", "RoughTime", "PrintTime", "TareTime",
		"WeightOverTime", "UpdateTime",
	}

	if updateWeighOrderErr := wtws_mysql.UpdateByWeighOrderId(updateWeighOrderData, updateWeighOrderCols); updateWeighOrderErr != nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
	}

	updateWg := &sync.WaitGroup{}
	updateWg.Add(2)

	var updateTruckOrderErr, updateOrderErr error

	go func() {
		updateTruckOrderCols := []string{"GoodsArranged"}
		updateTruckOrderData := &wtws_mysql.OrTruckOrder{
			Id:            dbTruckOrder.Id,
			GoodsArranged: dbTruckOrder.GoodsArranged + dbWeighOrder.GoodsWeight,
		}
		if updateTruckOrderData.GoodsArranged >= dbTruckOrder.GoodsWeight {
			updateTruckOrderData.Status = conf.TRUCK_ORDER_STATUS_FINISH
			updateTruckOrderCols = append(updateTruckOrderCols, "Status")
		}

		updateTruckOrderErr = wtws_mysql.UpdateByTruckOrderId(updateTruckOrderData, updateTruckOrderCols)

		updateWg.Done()
	}()

	go func() {

		if dbOrder.GoodsMargin <= 0 {
			updateOrderErr = wtws_mysql.UpdateOrOrderById(&wtws_mysql.OrOrder{
				Id:     dbOrder.Id,
				Status: conf.ORDER_STATUS_FINISH,
			}, []string{"Status"})
		}

		updateWg.Done()
	}()

	updateWg.Wait()

	if updateTruckOrderErr != nil || updateOrderErr != nil {
		return common_struct.DeveiceResponseStruct{
			Status:  "0",
			Message: "接收失败",
		}
	}

	return common_struct.DeveiceResponseStruct{
		Status:  "1",
		Message: "处理成功",
	}

}

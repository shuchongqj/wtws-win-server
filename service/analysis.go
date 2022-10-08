package service

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func GetAnalysisDetail(dateType string) common_struct.ResponseStruct {

	var startTimeStr string
	var lastTimeStartStr, lastTimeEndStr string

	endTimeStr := time.Now().Format("2006-01-02")

	if dateType == conf.ANALYSIS_DATE_TYPE_DAY {
		startTime := time.Now().AddDate(0, 0, 0)
		startTimeStr = startTime.Format("2006-01-02")

		lastTimeEndStr = startTime.AddDate(0, 0, -1).
			Format("2006-01-02")
		lastTimeStartStr = lastTimeEndStr
	} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH {
		startTime := time.Now().AddDate(0, 0, 1-time.Now().Day())
		startTimeStr = startTime.Format("2006-01-02")

		lastTimeEndStr = startTime.AddDate(0, 0, -1).
			Format("2006-01-02")
		lastTimeStartStr = startTime.AddDate(0, -1, 0).
			Format("2006-01-02")
	} else {
		startTime := time.Now().AddDate(0, -3, 1-time.Now().Day())
		startTimeStr = startTime.Format("2006-01-02")

		lastTimeEndStr = startTime.AddDate(0, 0, -1).
			Format("2006-01-02")
		lastTimeStartStr = startTime.AddDate(0, -3, 0).
			Format("2006-01-02")
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	var orderList, lastTimeOrderList []wtws_mysql.OrOrder
	var truckOrderList, lastTimeTruckOrderList []wtws_mysql.OrTruckOrder
	var weighOrderList, lastTimeWeighOrderList []wtws_mysql.OrWeighOrder
	//var getOrderErr, getTruckOrderErr, getWeighOrderErr error

	go func() {
		subWg := &sync.WaitGroup{}
		subWg.Add(2)
		go func() {
			orderList, _ = wtws_mysql.GetOrOrdersByInsertTime(startTimeStr, endTimeStr)
			subWg.Done()
		}()
		go func() {
			lastTimeOrderList, _ = wtws_mysql.GetOrOrdersByInsertTime(lastTimeStartStr, lastTimeEndStr)
			subWg.Done()
		}()
		subWg.Wait()

		wg.Done()
	}()

	go func() {
		subWg := &sync.WaitGroup{}
		subWg.Add(2)
		go func() {
			truckOrderList, _ = wtws_mysql.GetOrTruckOrdersByInsertTime(startTimeStr, endTimeStr)
			subWg.Done()
		}()
		go func() {
			lastTimeTruckOrderList, _ = wtws_mysql.GetOrTruckOrdersByInsertTime(lastTimeStartStr, lastTimeEndStr)
			subWg.Done()
		}()
		subWg.Wait()
		wg.Done()
	}()

	go func() {
		subWg := &sync.WaitGroup{}
		subWg.Add(2)
		go func() {
			weighOrderList, _ = wtws_mysql.GetOrWeighOrdersByInsertTime(startTimeStr, endTimeStr)
			subWg.Done()
		}()
		go func() {
			lastTimeWeighOrderList, _ = wtws_mysql.GetOrWeighOrdersByInsertTime(lastTimeStartStr, lastTimeEndStr)
			subWg.Done()
		}()
		subWg.Wait()
		wg.Done()
	}()

	wg.Wait()

	resData := dto.AnalysisDetail{
		PurchaseOrder: dto.OrderStatusMap{
			TruckOrderStatusMap: dto.TruckOrderStatusMap{
				AnalysisDetailItem: dto.AnalysisDetailItem{
					Total:         0,
					LastTotal:     0,
					AddPercentage: 0,
				},
				Wait:    0,
				Checked: 0,
				Reject:  0,
				Failure: 0,
				Finish:  0,
			},
			HasTruck: 0,
		},
		SaleOrder: dto.OrderStatusMap{
			TruckOrderStatusMap: dto.TruckOrderStatusMap{
				AnalysisDetailItem: dto.AnalysisDetailItem{
					Total:         0,
					LastTotal:     0,
					AddPercentage: 0,
				},
				Wait:    0,
				Checked: 0,
				Reject:  0,
				Failure: 0,
				Finish:  0,
			},
			HasTruck: 0,
		},
		SentDirectOrder: dto.OrderStatusMap{
			TruckOrderStatusMap: dto.TruckOrderStatusMap{
				AnalysisDetailItem: dto.AnalysisDetailItem{
					Total:         0,
					LastTotal:     0,
					AddPercentage: 0,
				},
				Wait:    0,
				Checked: 0,
				Reject:  0,
				Failure: 0,
				Finish:  0,
			},
			HasTruck: 0,
		},
		TruckOrder: dto.TruckOrderStatusMap{
			AnalysisDetailItem: dto.AnalysisDetailItem{
				Total:         len(truckOrderList),
				LastTotal:     len(lastTimeTruckOrderList),
				AddPercentage: 0,
			},
			Wait:    0,
			Checked: 0,
			Reject:  0,
			Failure: 0,
			Finish:  0,
		},
		WeighOrder: dto.WeightOrderStatusMap{
			AnalysisDetailItem: dto.AnalysisDetailItem{
				Total:         len(weighOrderList),
				LastTotal:     len(lastTimeWeighOrderList),
				AddPercentage: 0,
			},
			Process:        0,
			Failure:        0,
			WareHouseCheck: 0,
			WaitFinish:     0,
			Finish:         0,
		},
	}

	for _, order := range orderList {
		switch order.OrderType {
		case conf.ORDER_TYPE_PURCHASE:
			resData.PurchaseOrder.Total++
			switch order.Status {
			case conf.ORDER_STATUS_WAIT:
				resData.PurchaseOrder.Wait++
				break
			case conf.ORDER_STATUS_PASS:
				resData.PurchaseOrder.Checked++
				break
			case conf.ORDER_STATUS_REJECT:
				resData.PurchaseOrder.Reject++
				break
			case conf.ORDER_STATUS_FAILURE:
				resData.PurchaseOrder.Failure++
				break
			case conf.ORDER_STATUS_FINISH:
				resData.PurchaseOrder.Finish++
				break
			case conf.ORDER_STATUS_HAS_TRUCK:
				resData.PurchaseOrder.HasTruck++
				break
			}
			break
		case conf.ORDER_TYPE_SALE:
			resData.SaleOrder.Total++
			switch order.Status {
			case conf.ORDER_STATUS_WAIT:
				resData.SaleOrder.Wait++
				break
			case conf.ORDER_STATUS_PASS:
				resData.SaleOrder.Checked++
				break
			case conf.ORDER_STATUS_REJECT:
				resData.SaleOrder.Reject++
				break
			case conf.ORDER_STATUS_FAILURE:
				resData.SaleOrder.Failure++
				break
			case conf.ORDER_STATUS_FINISH:
				resData.SaleOrder.Finish++
				break
			case conf.ORDER_STATUS_HAS_TRUCK:
				resData.SaleOrder.HasTruck++
				break
			}
			break
		case conf.ORDER_TYPE_SENT_DIRECT:
			resData.SentDirectOrder.Total++
			switch order.Status {
			case conf.ORDER_STATUS_WAIT:
				resData.SentDirectOrder.Wait++
				break
			case conf.ORDER_STATUS_PASS:
				resData.SentDirectOrder.Checked++
				break
			case conf.ORDER_STATUS_REJECT:
				resData.SentDirectOrder.Reject++
				break
			case conf.ORDER_STATUS_FAILURE:
				resData.SentDirectOrder.Failure++
				break
			case conf.ORDER_STATUS_FINISH:
				resData.SentDirectOrder.Finish++
				break
			case conf.ORDER_STATUS_HAS_TRUCK:
				resData.SentDirectOrder.HasTruck++
				break
			}
			break
		}
	}

	for _, order := range lastTimeOrderList {
		switch order.OrderType {
		case conf.ORDER_TYPE_PURCHASE:
			resData.PurchaseOrder.LastTotal++
			break
		case conf.ORDER_TYPE_SALE:
			resData.SaleOrder.LastTotal++
			break
		case conf.ORDER_TYPE_SENT_DIRECT:
			resData.SentDirectOrder.LastTotal++
			break
		}
	}

	for _, order := range truckOrderList {
		switch order.Status {
		case conf.TRUCK_ORDER_STATUS_WAIT:
			resData.TruckOrder.Wait++
			break
		case conf.TRUCK_ORDER_STATUS_PASS:
			resData.TruckOrder.Checked++
			break
		case conf.TRUCK_ORDER_STATUS_REJECT:
			resData.TruckOrder.Reject++
			break
		case conf.TRUCK_ORDER_STATUS_FAILURE:
			resData.TruckOrder.Failure++
			break
		case conf.TRUCK_ORDER_STATUS_FINISH:
			resData.TruckOrder.Finish++
			break
		}
	}

	for _, order := range weighOrderList {
		switch order.Status {
		case conf.WEIGH_ORDER_STATUS_PROCESS:
			resData.WeighOrder.Process++
			break
		case conf.WEIGH_ORDER_STATUS_FAILURE:
			resData.WeighOrder.Failure++
			break
		case conf.WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM:
			resData.WeighOrder.WareHouseCheck++
			break
		case conf.WEIGH_ORDER_STATUS_WAITE_FINISH:
			resData.WeighOrder.WaitFinish++
			break
		case conf.WEIGH_ORDER_STATUS_FINISH:
			resData.WeighOrder.Finish++
			break
		}
	}

	var truckOrderDenominator float32 = 1
	var weighOrderDenominator float32 = 1
	var purchaseOrderDenominator float32 = 1
	var saleOrderDenominator float32 = 1
	var sentDirectOrderDenominator float32 = 1

	if resData.PurchaseOrder.LastTotal > 0 {
		purchaseOrderDenominator = float32(resData.PurchaseOrder.LastTotal)
	}

	if resData.SaleOrder.LastTotal > 0 {
		saleOrderDenominator = float32(resData.SaleOrder.LastTotal)
	}

	if resData.SentDirectOrder.LastTotal > 0 {
		sentDirectOrderDenominator = float32(resData.SentDirectOrder.LastTotal)
	}

	if len(lastTimeTruckOrderList) > 0 {
		truckOrderDenominator = float32(len(lastTimeTruckOrderList))
	}

	if len(lastTimeWeighOrderList) > 0 {
		weighOrderDenominator = float32(len(lastTimeWeighOrderList))
	}

	resData.PurchaseOrder.AddPercentage = float32(resData.PurchaseOrder.Total-resData.PurchaseOrder.LastTotal) / purchaseOrderDenominator
	resData.SaleOrder.AddPercentage = float32(resData.SaleOrder.Total-resData.SaleOrder.LastTotal) / saleOrderDenominator
	resData.SentDirectOrder.AddPercentage = float32(resData.SentDirectOrder.Total-resData.SentDirectOrder.LastTotal) / sentDirectOrderDenominator

	resData.TruckOrder.AddPercentage = float32(len(truckOrderList)-len(lastTimeTruckOrderList)) / truckOrderDenominator
	resData.WeighOrder.AddPercentage = float32(len(weighOrderList)-len(lastTimeWeighOrderList)) / weighOrderDenominator

	return common.ResponseStatus(0, "", resData)
}

func GetAnalysisOrderTypeDetail(dateType string, orderType string) common_struct.ResponseStruct {

	var startTimeStr string
	endTimeStr := time.Now().Format("2006-01-02")

	if dateType == conf.ANALYSIS_DATE_TYPE_DAY {
		startTime := time.Now().AddDate(0, 0, -14)
		startTimeStr = startTime.Format("2006-01-02")
	} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH {
		startTime := time.Now().AddDate(0, -12, 0)
		startTimeStr = startTime.Format("2006-01-02")
	} else {
		startTime := time.Now().AddDate(-5, 0, 0)
		startTimeStr = startTime.Format("2006-01-02")
	}

	response := map[string]int{}
	ormList := []orm.Params{}

	switch orderType {
	case conf.ANALYSIS_ORDER_TYPE_PURCHASE, conf.ANALYSIS_ORDER_TYPE_SALE, conf.ANALYSIS_ORDER_TYPE_SENT_DIRECT:
		if dateType == conf.ANALYSIS_DATE_TYPE_DAY {
			ormList, _ = wtws_mysql.GetAnalysisOrOrdersDayTimeGroupCount(startTimeStr, endTimeStr, conf.ANALYSIS_ORDER_TYPE_MAP[orderType])
		} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH {
			ormList, _ = wtws_mysql.GetAnalysisOrOrdersMonthTimeGroupCount(startTimeStr, endTimeStr, conf.ANALYSIS_ORDER_TYPE_MAP[orderType])
		} else {
			ormList, _ = wtws_mysql.GetAnalysisOrOrdersYearTimeGroupCount(startTimeStr, endTimeStr, conf.ANALYSIS_ORDER_TYPE_MAP[orderType])
		}
		break
	case conf.ANALYSIS_ORDER_TYPE_TRUCK:
		if dateType == conf.ANALYSIS_DATE_TYPE_DAY {
			ormList, _ = wtws_mysql.GetAnalysisOrTruckOrdersDayTimeGroupCount(startTimeStr, endTimeStr)
		} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH {
			ormList, _ = wtws_mysql.GetAnalysisOrTruckOrdersMonthTimeGroupCount(startTimeStr, endTimeStr)
		} else {
			ormList, _ = wtws_mysql.GetAnalysisOrTruckOrdersYearTimeGroupCount(startTimeStr, endTimeStr)
		}
		break
	case conf.ANALYSIS_ORDER_TYPE_WEIGH:
		if dateType == conf.ANALYSIS_DATE_TYPE_DAY {
			ormList, _ = wtws_mysql.GetAnalysisOrWeighOrdersDayTimeGroupCount(startTimeStr, endTimeStr)
		} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH {
			ormList, _ = wtws_mysql.GetAnalysisOrWeighOrdersMonthTimeGroupCount(startTimeStr, endTimeStr)
		} else {
			ormList, _ = wtws_mysql.GetAnalysisOrWeighOrdersYearTimeGroupCount(startTimeStr, endTimeStr)
		}
		break
	}

	if len(ormList) != 0 {
		for _, itemOrm := range ormList {
			key := itemOrm["dates"].(string)
			valueStr := itemOrm["count"].(string)
			value, _ := strconv.Atoi(valueStr)
			response[key] = value
		}
	}

	if dateType == conf.ANALYSIS_DATE_TYPE_DAY && len(ormList) < 14 {
		if len(ormList) == 0 {
			now := time.Now()
			for i := 0; i < 14; i++ {
				dateTimeStr := now.Format("01月02日")
				response[dateTimeStr] = 0
				now = now.AddDate(0, 0, -1)
			}
		} else {
			now := time.Now()
			targetStartTimeStr := ormList[0]["dates"].(string)
			dateStr := strconv.Itoa(now.Year()) + "年" + targetStartTimeStr
			targetStartTime, _ := time.ParseInLocation("2006年01月02日", dateStr, time.Local)

			diffDay64 := (now.Unix() - targetStartTime.Unix()) / (60 * 60 * 24)
			diffDay := int(diffDay64)
			for i := 0; i < diffDay; i++ {
				dateTimeStr := now.Format("01月02日")
				response[dateTimeStr] = 0
				now = now.AddDate(0, 0, -1)
			}

			targetTimeInterface := ormList[len(ormList)-1]["dates"]
			targetTimeStr := strconv.Itoa(time.Now().Year()) + "年" + targetTimeInterface.(string)
			targetTime, _ := time.ParseInLocation("2006年01月02日", targetTimeStr, time.Local)
			for i := 0; i < 14-len(ormList)-diffDay; i++ {
				targetTime = targetTime.AddDate(0, 0, -1)
				dateTimeStr := targetTime.Format("01月02日")
				response[dateTimeStr] = 0
			}
		}
	} else if dateType == conf.ANALYSIS_DATE_TYPE_MONTH && len(ormList) < 12 {
		if len(ormList) == 0 {
			now := time.Now()
			for i := 0; i < 12; i++ {
				dateTimeStr := now.Format("2006年01月")
				response[dateTimeStr] = 0
				now = now.AddDate(0, -1, 0)
			}
		} else {
			now := time.Now()
			targetStartTimeStr := ormList[0]["dates"].(string)
			targetStartTime, _ := time.ParseInLocation("2006年01月", targetStartTimeStr, time.Local)

			diffMonth64 := (now.Unix() - targetStartTime.Unix()) / (60 * 60 * 24 * 30)
			diffMonth := int(diffMonth64)
			for i := 0; i < diffMonth; i++ {
				dateTimeStr := now.Format("2006年01月")
				response[dateTimeStr] = 0
				now = now.AddDate(0, -1, 0)
			}

			targetTimeInterface := ormList[len(ormList)-1]["dates"]
			targetTimeStr := targetTimeInterface.(string)
			targetTime, _ := time.ParseInLocation("2006年01月", targetTimeStr, time.Local)
			for i := 0; i < 12-len(ormList)-diffMonth; i++ {
				targetTime = targetTime.AddDate(0, -1, 0)
				dateTimeStr := targetTime.Format("2006年01月")
				response[dateTimeStr] = 0
			}
		}
	} else if dateType == conf.ANALYSIS_DATE_TYPE_YEAR && len(ormList) < 5 {
		if len(ormList) == 0 {
			now := time.Now()
			for i := 0; i < 5; i++ {
				dateTimeStr := now.Format("2006年")
				response[dateTimeStr] = 0
				now = now.AddDate(0, 0, -1)
			}
		} else {
			now := time.Now()
			targetStartTimeStr := ormList[0]["dates"].(string)
			targetStartTime, _ := time.ParseInLocation("2006年", targetStartTimeStr, time.Local)

			diffYear64 := (now.Unix() - targetStartTime.Unix()) / (60 * 60 * 24 * 365)
			diffYear := int(diffYear64)
			for i := 0; i < diffYear; i++ {
				dateTimeStr := now.Format("2006年")
				response[dateTimeStr] = 0
				now = now.AddDate(-1, 0, 0)
			}

			targetTimeInterface := ormList[len(ormList)-1]["dates"]
			targetTimeStr := targetTimeInterface.(string)
			targetTime, _ := time.ParseInLocation("2006年", targetTimeStr, time.Local)
			for i := 0; i < 5-len(ormList)-diffYear; i++ {
				targetTime = targetTime.AddDate(-1, 0, 0)
				dateTimeStr := targetTime.Format("2006年")
				response[dateTimeStr] = 0
			}
		}
	}

	return common.ResponseStatus(0, "", response)
}

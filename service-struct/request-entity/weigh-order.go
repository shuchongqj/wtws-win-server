package request_entity

type WeighOrderList struct {
	PageNum              int      `json:"pageNum" validate:"required"`
	PageSize             int      `json:"pageSize" validate:"required"`
	WeighOrderNo         string   `json:"weighOrderNo"`
	TruckOrderNo         string   `json:"truckOrderNo"`
	OrderNo              string   `json:"orderNo"`
	VehicleNumber        string   `json:"vehicleNumber"`
	OrderType            int      `json:"orderType"`
	Status               int      `json:"status"`
	ReceiveID            int      `json:"receiveID"`
	OriginID             int      `json:"originID"`
	GoodsName            string   `json:"goodsName"`
	GoodsNo              string   `json:"goodsNo"`
	InsertTimes          []string `json:"insertTimes"`
	WarehouseConfirmTime []string `json:"warehouseConfirmTime"`
	FinancialConfirmTime []string `json:"financialConfirmTime"`
}

type AddWeighOrder struct {
	WeighOrderType int `json:"weighOrderType" validate:"required"`
	OriginID       int `json:"originID" validate:"required"`
	ReceiveID      int `json:"receiveID" validate:"required"`
}

type DeleteWeighOrder struct {
	WeighOrderIDs []int `json:"weighOrderIDs"`
}

type UpdateWeighOrder struct {
	WeighOrderID   int     `json:"weighOrderID" validate:"required"`
	WeighOrderNo   string  `json:"weighOrderNo validate:"required""`
	WeighOrderType int     `json:"weighOrderType" validate:"required"`
	ReceiveID      int     `json:"receiveID" validate:"required"`
	ReceiveName    string  `json:"receiveName" validate:"required"`
	OriginID       int     `json:"originID" validate:"required"`
	OriginName     string  `json:"originName" validate:"required"`
	GoodsID        int     `json:"goodsID" validate:"required"`
	GoodsNum       float32 `json:"goodsNum" validate:"required"`
	GoodsWeight    float32 `json:"goodsWeight" validate:"required"`
	GoodsName      string  `json:"goodsName" validate:"required"`
	GoodsNote      string  `json:"goodsNote"`
}

type WarehouseCheckWeighOrder struct {
	WeighOrderID int     `json:"weighOrderID" validate:"required"`
	GoodsBatchNo string  `json:"goodsBatchNo" validate:"required"`
	OtherWight   float32 `json:"otherWight"  validate:"required"`
	TruckOrderID int     `json:"truckOrderID"`
}

type InvalidWeighOrder struct {
	WeighOrderID int `json:"weighOrderID" validate:"required"`
}

type ScanVehicle struct {
	Vehicle       string `json:"vehicle" validate:"required"`
	TransportTime string `json:"transporttime"`
}

type TareWight struct {
	Wvehicle  string `json:"wvehicle" validate:"required"`
	TareWight string `json:"tare_wight" validate:"required"`
	GrossTime string `json:"grosstime" validate:"required"`
	Wno       string `json:"wno" validate:"required"`
}

type CheckWareHouseGoods struct {
	Wno string `json:"wno" validate:"required"`
}

type WaitFinish struct {
	Wno  string `json:"wno" validate:"required"`
	Upid string `json:"upid" validate:"required"`
}

type FinishWeighOrder struct {
	WeighOrderNo string `json:"wno" validate:"required"`
	TareWight    string `json:"tare_wight" validate:"required"`
	RoughWight   string `json:"rough_wight" validate:"required"`
	NetWight     string `json:"net_wight" validate:"required"`
	RoughTimeStr string `json:"grosstime" validate:"required"`
	TareTimeStr  string `json:"skintime" validate:"required"`
	PrintTimeStr string `json:"wgbjssj" validate:"required"`
	Over         string `json:"over" validate:"required"`
}

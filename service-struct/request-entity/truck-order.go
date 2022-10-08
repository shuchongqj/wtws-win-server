package request_entity

type TruckOrderList struct {
	PageNum      int      `json:"pageNum" validate:"required"`
	PageSize     int      `json:"pageSize" validate:"required"`
	TruckOrderNo string   `json:"truckOrderNo"`
	OrderNo      string   `json:"orderNo"`
	OrderType    int      `json:"orderType"`
	Status       int      `json:"status"`
	ReceiveID    int      `json:"receiveID"`
	OriginID     int      `json:"originID"`
	GoodsName    string   `json:"goodsName"`
	GoodsNo      string   `json:"goodsNo"`
	InsertTimes  []string `json:"insertTimes"`
	VerifyTimes  []string `json:"verifyTimes"`
}

type AddTruckOrder struct {
	Orders        []TruckOrders `json:"orders" validate:"required"`
	DriverID      int           `json:"driverID" validate:"required"`
	IsWeightLimit int           `json:"isWeightLimit" validate:"required"`
	CargotoID     int           `json:"cargotoID" validate:"required"`
	PaymentMethod int           `json:"paymentMethod" validate:"required"`
	DriverTime    string        `json:"driverTime" validate:"required"`
	ContainerNo   string        `json:"containerNo"`
	Status        int           `json:"status"`
	VerifierId    int
	VerifierName  string
	VerifierNote  string
}
type TruckOrders struct {
	OrderID           int     `json:"orderID" validate:"required"`
	GoodsLoadQuantity float32 `json:"goodsLoadQuantity" validate:"required"`
}

type AddSentDirectOrder struct {
	OriginID      int     `json:"originID" validate:"required"`
	ReceiveID     int     `json:"receiveID" validate:"required"`
	Goods         []Goods `json:"goods" validate:"required"`
	OrderType     int     `json:"orderType" validate:"required"`
	DriverID      int     `json:"driverID" validate:"required"`
	IsWeightLimit int     `json:"isWeightLimit" validate:"required"`
	DriverTime    string  `json:"driverTime" `
	ContainerNo   string  `json:"containerNo"`
	CargotoID     int     `json:"cargotoID" validate:"required"`
	PaymentMethod int     `json:"paymentMethod" validate:"required"`
}

type DeleteTruckOrder struct {
	TruckOrderIDs []int `json:"truckOrderIDs"`
}

type UpdateTruckOrder struct {
	TruckOrderID   int     `json:"truckOrderID" validate:"required"`
	TruckOrderNo   string  `json:"truckOrderNo validate:"required""`
	TruckOrderType int     `json:"truckOrderType" validate:"required"`
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

type CheckTruckOrder struct {
	TruckOrderID int    `json:"truckOrderID" validate:"required"`
	CheckType    int    `json:"checkType" validate:"required"`
	CheckNote    string `json:"checkNote"`
}

type InvalidTruckOrder struct {
	TruckOrderID int `json:"truckOrderID" validate:"required"`
}

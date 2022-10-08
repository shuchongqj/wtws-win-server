package request_entity

type OrderList struct {
	PageNum     int      `json:"pageNum" validate:"required"`
	PageSize    int      `json:"pageSize" validate:"required"`
	OrderNo     string   `json:"orderNo"`
	OrderType   int      `json:"orderType"`
	OrderStatus int      `json:"orderStatus"`
	ReceiveID   int      `json:"receiveID"`
	OriginID    int      `json:"originID"`
	GoodsName   string   `json:"goodsName"`
	GoodsNo     string   `json:"goodsNo"`
	InsertTimes []string `json:"insertTimes"`
	VerifyTimes []string `json:"verifyTimes"`
}

type AddOrder struct {
	OrderType    int     `json:"orderType" validate:"required"`
	OriginID     int     `json:"originID" validate:"required"`
	ReceiveID    int     `json:"receiveID" validate:"required"`
	Goods        []Goods `json:"goods" validate:"required"`
	Status       int     `json:"status"`
	VerifierId   int
	VerifierName string
	VerifierNote string
}
type Goods struct {
	GoodID             int     `json:"goodID" validate:"required"`
	GoodsName          string  `json:"goodsName" validate:"required"`
	GoodsNo            string  `json:"goodsNo" validate:"required"`
	GoodsSpecification float64 `json:"goodsSpecification" validate:"required"`
	GoodsExtraWeight   float32 `json:"goodsExtraWeight"`
	GoodsNum           float64 `json:"goodsNum" `
	GoodsWeight        float64 `json:"goodsWeight"`
	GoodsNote          string  `json:"goodsNote"`
}

type DeleteOrder struct {
	OrderIDs []int `json:"orderIDs"`
}

type UpdateOrder struct {
	OrderID     int     `json:"orderID" validate:"required"`
	OrderNo     string  `json:"orderNo validate:"required""`
	OrderType   int     `json:"orderType" validate:"required"`
	ReceiveID   int     `json:"receiveID" validate:"required"`
	ReceiveName string  `json:"receiveName" validate:"required"`
	OriginID    int     `json:"originID" validate:"required"`
	OriginName  string  `json:"originName" validate:"required"`
	GoodsID     int     `json:"goodsID" validate:"required"`
	GoodsNum    float32 `json:"goodsNum" validate:"required"`
	GoodsWeight float32 `json:"goodsWeight" validate:"required"`
	GoodsName   string  `json:"goodsName" validate:"required"`
	GoodsNote   string  `json:"goodsNote"`
}

type CheckOrder struct {
	OrderID   int    `json:"orderID" validate:"required"`
	CheckType int    `json:"checkType" validate:"required"`
	CheckNote string `json:"checkNote"`
}

type InvalidOrder struct {
	OrderID int `json:"orderID" validate:"required"`
}

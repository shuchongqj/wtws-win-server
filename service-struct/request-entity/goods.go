package request_entity

type GoodsList struct {
	PageNum    int    `json:"pageNum" validate:"required"`
	PageSize   int    `json:"pageSize" validate:"required"`
	CategoryID int    `json:"categoryID"`
	GoodsName  string `json:"goodsName"`
	GoodsNo    string `json:"goodsNo"`
}

type AddGoods struct {
	GoodsName     string  `json:"goodsName" validate:"required"`
	CategoryID    int     `json:"categoryID" validate:"required"`
	GoodNo        string  `json:"goodsNo"`
	Specification float64 `json:"specification"`
	BagWeight     float64 `json:"bagWeight"`
	DeductWeight  float64 `json:"deductWeight"`
	ExtraWeight   float64 `json:"extraWeight"`
}

type DeleteGoods struct {
	GoodsIDs []int `json:"goodsIDs"`
}

type UpdateGoods struct {
	GoodsID     int     `json:"goodsId" validate:"required"`
	GoodsName   string  `json:"goodsName" validate:"required"`
	Tel         string  `json:"tel" validate:"required"`
	Type        int     `json:"type" validate:"required"`
	ContactName string  `json:"contactName" `
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

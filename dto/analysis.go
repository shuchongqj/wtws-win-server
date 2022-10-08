package dto

type AnalysisDetailItem struct {
	Total         int     `json:"total"`
	LastTotal     int     `json:"lastTotal"`
	AddPercentage float32 `json:"addPercentage"`
}

type TruckOrderStatusMap struct {
	AnalysisDetailItem
	Wait    int `json:"wait"`
	Checked int `json:"checked"`
	Reject  int `json:"reject"`
	Failure int `json:"failure"`
	Finish  int `json:"finish"`
}

type OrderStatusMap struct {
	TruckOrderStatusMap
	HasTruck int `json:"hasTruck"`
}

type WeightOrderStatusMap struct {
	AnalysisDetailItem
	Process        int `json:"process"`
	Failure        int `json:"failure"`
	WareHouseCheck int `json:"wareHouseCheck"`
	WaitFinish     int `json:"waitFinish"`
	Finish         int `json:"finish"`
}

type AnalysisDetail struct {
	PurchaseOrder   OrderStatusMap       `json:"purchaseOrder"`
	SaleOrder       OrderStatusMap       `json:"saleOrder"`
	SentDirectOrder OrderStatusMap       `json:"sentDirectOrder"`
	TruckOrder      TruckOrderStatusMap  `json:"truckOrder"`
	WeighOrder      WeightOrderStatusMap `json:"weighOrder"`
}

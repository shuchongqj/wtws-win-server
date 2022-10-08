package common_struct

type AddTyreData struct {
	StationID      string    `json:"stationID"`
	TyreID         string    `json:"tyreID"`
	TyreStatusText string    `json:"tyreStatusText"`
	TyreNo         string    `json:"tyreNo"`
	ArrivalTime    string    `json:"arrivalTime"`
	Brand          string    `json:"brand"`
	Specification  string    `json:"specification"`
	PatternModel   string    `json:"patternModel"`
	Price          int       `json:"price"`
	InitPattern    float32   `json:"initPattern"`
	PatternNum     int       `json:"patternNum"`
	RetreadBrand   string    `json:"retreadBrand"`
	InitPatterns   []float32 `json:"initPatterns"`
}

type BathAddTyre struct {
	DataList []AddTyreData `json:"dataList"`
}

type TyreEntry struct {
	StationID      string  `json:"stationID"valid:"Required;Length(32)"`
	TyreNo         string  `json:"tyreNo"valid:"Required;MinSize(1)"`
	TyreStatus     int     `json:"tyreStatus"valid:"Required;Range(1,4)"` //1-新胎入库	2-回收胎入库	3-调度胎入库	4-翻新胎入库
	TyreTemplateID string  `json:"tyreTemplateID"valid:"Required"`
	PatternContent float32 `json:"patternContent"valid:"Required"`
	PatternNum     int     `json:"patternNum"valid:"Required;"`
	PatternJSON    string  `json:"patternJson"valid:"Required"`
	ArrivalTime    string  `json:"arrivalTime"`
	TyrePrice      string  `json:"tyrePrice"`
	TyreID         string  `json:"tyreID"`
	RFID           string  `json:"RFID"`
	QrCode         string  `json:"qrCode"`
}

type TyreOut struct {
	OriginalStationID string `json:"originalStationID"`
	StationID         string `json:"stationID"`
	TyreID            string `json:"tyreID"`
	TyreNo            string `json:"tyreNo"`
	RFID              string `json:"RFID"`
}

type ConsumableOut struct {
	OriginalStationID string `json:"originalStationID" valid:"Required"`
	StationID         string `json:"stationID" valid:"Required"`
	ConsumableID      int    `json:"consumableID" valid:"Required"`
}

type imageContent struct {
	Url string `json:"url"`
}

type InsertOperation struct {
	OperationType    int            `json:"operationType"valid:"Required;Range(1,20)"`
	TyreID           string         `json:"tyreID"valid:"Required;MinSize(1)"`
	TyreNo           string         `json:"tyreNo"valid:"Required;MinSize(1)"`
	OperationContent string         `json:"operationContent"valid:"Required;MinSize(1)"`
	VehicleNumber    string         `json:"vehicleNumber"`
	ImagesContent    []imageContent `json:"imagesContent"`
	VehicleID        string         `json:"vehicleID"`
	WheelPosition    string         `json:"wheelPosition"`
	TyrePressure     string         `json:"tyrePressure"`
	CurrentMileage   int            `json:"currentMileage"`
	PatternAVG       float32        `json:"patternAVG"`
	PatternNum       int            `json:"patternNum"`
	Patterns         []float32      `json:"patterns"`
	ToWheelPosition  string         `json:"toWheelPosition"`
	RPatternNum      int            `json:"rPatternNum"`
	RPatterns        []float32      `json:"rPatterns"`
	RPatternAVG      float32        `json:"rPatternAVG"`
}

// GetTyreList ...
type GetTyreList struct {
	StationID      string   `json:"stationID"`
	TyreNo         string   `json:"tyreNo"`
	TyreStatus     string   `json:"tyreStatus"`
	StorageStatus  string   `json:"storageStatus"`
	Brand          string   `json:"brand"`
	IsRetread      int      `json:"isRetread"`
	VehicleNum     string   `json:"vehicleNum"`
	PatternContent string   `json:"patternContent"`
	Times          []string `json:"times"`
	StartTime      int      `json:"startTime"`
	EndTime        int      `json:"endTime"`
	PageNum        int      `json:"pageNum"`
	PageSize       int      `json:"pageSize"`
	IsExport       string   `json:"isExport"`
}

type GetUserByStationIds struct {
	StationIds []string `json:"stationIds"`
}

type PostOutSaleTyre struct {
	TyreID     string `json:"tyreID"valid:"Required"`
	RFID       string `json:"RFID"`
	SaleStatus int8   `json:"saleStatus"`
	IsDelete   int8   `json:"isDelete"`
}

type PostSelectOutConsumable struct {
	ConsumableID int  `json:"consumableID"valid:"Required"`
	SaleStatus   int8 `json:"saleStatus"`
	IsDelete     int8 `json:"isDelete"`
}

type ConfirmOut struct {
	OperateType int8 `json:"operateType"`
}

type InstallSaleOutTyreListItem struct {
	IsTyre       bool    `json:"isTyre"valid:"Required"`
	ConsumableID int     `json:"consumableID"`
	TyreID       string  `json:"tyreID"`
	VehicleNum   string  `json:"vehicleNum"`
	SalePrice    float64 `json:"salePrice"`
}

type InstallSaleOutTyre struct {
	List []InstallSaleOutTyreListItem `json:"list"`
}

type OCRVehicleNumber struct {
	ImageURL string `json:"imageUrl"`
}

type GetTyreQuantity struct {
	Brand         string `json:"brand"`
	PatternModel  string `json:"patternModel"`
	Specification string `json:"specification"`
}

type Inventory struct {
	TyreID string `json:"tyreID"`
}

type SaleTyreList struct {
	TyreID string `json:"tyreID"`
}

type SaleConsumableList struct {
	ConsumableID int `json:"consumableID"`
}

type ParamsInfo struct {
	TyreID string `json:"tyreID" valid:"Required"`
}

type UpdateOrderGoods struct {
	OrderGoodId        int                  `json:"orderGoodId"valid:"Required"`
	ProductID          int                  `json:"productId"valid:"Required"`
	GoodID             int                  `json:"goodId"valid:"Required"`
	IsTyre             bool                 `json:"isTyre"valid:"Required"`
	Number             int                  `json:"number"valid:"Required"`
	InstalledNum       int                  `json:"installedNum"valid:"Required"`
	MaxNeedInstallNum  int                  `json:"maxNeedInstallNum"valid:"Required"`
	CurrentInstallNum  int                  `json:"currentInstallNum"valid:"Required"`
	SaleTyreList       []SaleTyreList       `json:"saleTyreList"`
	SaleConsumableList []SaleConsumableList `json:"saleConsumableList"`
}
type UpdateOrderDetail struct {
	OrderID       int                `json:"orderID"valid:"Required"`
	StoreID       int                `json:"storeID"valid:"Required"`
	VehicleNumber string             `json:"vehicleNumber"valid:"Required"`
	Mobile        string             `json:"mobile"valid:"Required"`
	Goods         []UpdateOrderGoods `json:"goods"valid:"Required"`
}

type GetTyreOperateList struct {
	ParamsInfo ParamsInfo `json:"paramsInfo" valid:"Required"`
}

type EntryConsumable struct {
	StationID               string `json:"stationID" valid:"Required;MinSize(1)"`
	ConsumableEntryMethod   int    `json:"consumableEntryMethod" valid:"Required;Range(1,3)"`
	ConsumableType          int    `json:"consumableType"  valid:"Required"`
	ConsumableNo            string `json:"consumableNo"`
	ConsumableCode          string `json:"consumableCode"`
	PurchasePrice           string `json:"purchasePrice"`
	ArrivalTime             string `json:"arrivalTime"`
	ConsumableBrand         string `json:"consumableBrand"`
	ConsumableSeries        string `json:"consumableSeries"`
	ConsumableSpecification string `json:"consumableSpecification"`
}

type SearchConsumable struct {
	ConsumableNo   string `json:"consumableNo"`
	ConsumableCode string `json:"consumableCode"`
	InStation      int8   `json:"inStation"`
	StationName    string `json:"stationName"`
}

type GetConsumableList struct {
	StationID        string   `json:"stationID"`
	ConsumableNo     string   `json:"consumableNo"`
	ConsumableType   int      `json:"consumableType"`
	ConsumableStatus int      `json:"consumableStatus"`
	Brand            string   `json:"brand"`
	Specification    string   `json:"specification"`
	Series           string   `json:"series"`
	VehicleNumber    string   `json:"vehicleNumber"`
	CreatedTime      []string `json:"createdTime"`
	ArrivalTime      []string `json:"arrivalTime"`
	PageSize         int      `json:"pageSize"`
	PageNum          int      `json:"pageNum"`
}

type UpdateMallUserDetail struct {
	ID       int    `json:"id"`
	Integral int    `json:"integral"`
	Mobile   string `json:"mobile"`
	Username string `json:"username"`
	Type     int    `json:"type"`
}

type DeleteItemByID struct {
	ID     int  `json:"id"`
	Delete bool `json:"delete"`
}

type UpdateIntegralRuleDetail struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Message      string `json:"message"`
	IncreaseType int    `json:"increaseType"`
	Value        int    `json:"value"`
}

type UpdateGood struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	BrandID     int     `json:"brand_id"`
	StoreID     int     `json:"storeId"`
	CategoryID  int     `json:"categoryId"`
	RetailPrice float64 `json:"retailPrice"`
	GoodsUnit   string  `json:"goodsUnit"`
	IsDelete    int     `json:"isDelete"`
	IsHot       int     `json:"is_hot"`
	IsNew       int     `json:"is_new"`
	SellVolume  int     `json:"sell_volume"`

	//AddTime           string `json:"add_time"`
	//AttributeCategory string `json:"attribute_category"`
	//BrandName         string `json:"brandName"`
	//CategoryName      string `json:"categoryName"`
	//CounterPrice      string `json:"counter_price"`
	//GoodsNumber       string `json:"goods_number"`
	//GoodsSn           string `json:"goods_sn"`
	//IsOnSale          string `json:"is_on_sale"`
	//ListPicURL        string `json:"list_pic_url"`
	//PrimaryPicURL     string `json:"primary_pic_url"`
	//PrimaryProductID  string `json:"primary_product_id"`
	//SortOrder         string `json:"sort_order"`
}

type BatchUpdateProductItem struct {
	AttributeID           string            `json:"attributeId"`
	AttributeMap          map[string]string `json:attributeMap"`
	AttributeValue        string            `json:"attributeValue"`
	GoodsID               string            `json:"goods_id"`
	GoodsNumber           string            `json:"goods_number"`
	GoodsSn               string            `json:"goods_sn"`
	GoodsSpecificationIds string            `json:"goods_specification_ids"`
	ID                    string            `json:"id"`
	IsDelete              int               `json:"is_delete"`
	ProductSpecification  string            `json:"productSpecification"`
	ProductDesc           string            `json:"product_desc"`
	RetailPrice           string            `json:"retail_price"`
}

type BatchUpdateProducts struct {
	Products []BatchUpdateProductItem `json:"products"`
}

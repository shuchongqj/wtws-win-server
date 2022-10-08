package common_struct

type ImportTyreError struct {
	ErrField string `json:"errField"`
	ErrMsg   string `json:"errMsg"`
}

type BatchExcelTyre struct {
	StationID      string    `json:"stationID"`
	TyreID         string    `json:"tyreID"`
	TyreStatusText string    `json:"tyreStatusText"`
	TyreNo         string    `json:"tyreNo"`
	Brand          string    `json:"brand"`
	Specification  string    `json:"specification"`
	PatternModel   string    `json:"patternModel"`
	Price          float32   `json:"price"`
	InitPattern    float32   `json:"initPattern"`
	PatternNum     int8      `json:"patternNum"`
	RetreadBrand   string    `json:"retreadBrand"`
	InitPatterns   []float32 `json:"initPatterns"`
	ArrivalTime    string    `json:"arrivalTime"`
	TyreStatus     int8      `json:"tyreStatus"`

	Err ImportTyreError `json:"error"`
}

//PatternMap1 两条花纹的pattern map
type PatternMap1 struct {
	Pattern1 float32 `json:"pattern1"`
	Pattern2 float32 `json:"pattern2"`
}

//PatternMap2 两条花纹的pattern map
type PatternMap2 struct {
	Pattern1 float32 `json:"pattern1"`
	Pattern2 float32 `json:"pattern2"`
	Pattern3 float32 `json:"pattern3"`
}

//PatternMap3 两条花纹的pattern map
type PatternMap3 struct {
	Pattern1 float32 `json:"pattern1"`
	Pattern2 float32 `json:"pattern2"`
	Pattern3 float32 `json:"pattern3"`
	Pattern4 float32 `json:"pattern4"`
}

//PatternMap4 两条花纹的pattern map
type PatternMap4 struct {
	Pattern1 float32 `json:"pattern1"`
	Pattern2 float32 `json:"pattern2"`
	Pattern3 float32 `json:"pattern3"`
	Pattern4 float32 `json:"pattern4"`
	Pattern5 float32 `json:"pattern5"`
}

//PatternMap5 两条花纹的pattern map
type PatternMap5 struct {
	Pattern1 float32 `json:"pattern1"`
	Pattern2 float32 `json:"pattern2"`
	Pattern3 float32 `json:"pattern3"`
	Pattern4 float32 `json:"pattern4"`
	Pattern5 float32 `json:"pattern5"`
	Pattern6 float32 `json:"pattern6"`
}

type InfluxDBGpsStruct struct {
	SN            string  `json:"SN"`
	Type          int     `json:"Type"`
	DeviceUTCTime string  `json:"DeviceUTCTime"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
	Speed         float64 `json:"Speed"`
	Course        int     `json:"Course"`
	DeviceBattery float64 `json:"DeviceBattery"`
	PowerSource   int     `json:"PowerSource"`
}

type InfluxDBTpmsStruct struct {
	SN            string  `json:"SN"`
	Type          int     `json:"Type"`
	SensorID      string  `json:"SensorID"`
	DeviceUTCTime string  `json:"DeviceUTCTime"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
	TirePressure  float64 `json:"TirePressure"`
	Temperature   float64 `json:"Temperature"`
	WheelPlace    string  `json:"WheelPlace"`
	SensorBattery float64 `json:"SensorBattery"`
	TireStatus    string  `json:"TireStatus"`
}

type InfluxDBMileageStruct struct {
	SN         string `json:"SN"`
	Mileages   int    `json:"Mileages"`
	InsertTime string `json:"InsertTime"`
}

type TyreMonthReportItem struct {
	VehicleNumber    string
	WheelPosition    string
	FirstTestTime    string
	FirstTestPattern float64
	LastTestTime     string
	LastTestPattern  float64
}

type SelectSaleOutTyreStruct struct {
	Note        string `json:"note"`
	OperateTime string `json:"operateTime"`
}

type InstallSaleOutTyreOperateContent struct {
	Note          string  `json:"note"`
	SalePrice     float64 `json:"salePrice"`
	VehicleNumber string  `json:"vehicleNumber"`
	OperateTime   string  `json:"operateTime"`
	SaleOrderID   int     `json:"saleOrderID"`
	SaleOrderSn   string  `json:"saleOrderSn"`
	SaleOrderType int     `json:"saleOrderType"`
}

type InstallSaleOutConsumableOperateContent struct {
	InstallSaleOutTyreOperateContent
	ConsumableType int `json:"consumableType"`
}

type CancelSaleOutTyreOperateContent struct {
	Note        string `json:"note"`
	OperateTime string `json:"operateTime"`
}

type CancelSaleOutConsumableOperateContent struct {
	CancelSaleOutTyreOperateContent
	ConsumableType int `json:"consumableType"`
}

type OrderOtherOperateContentStruct struct {
	Note        string `json:"note"`
	OperateTime string `json:"operateTime"`
}

type ConsumableEntryOperateContentStruct struct {
	Note        string `json:"note"`
	OperateTime string `json:"operateTime"`
}

type OrderGoodItem struct {
	GoodName            string               `json:"goodName"`
	GodSpecifition      string               `json:"goodSpecifition"`
	GoodPrice           float64              `json:"goodPrice"`
	IsTyre              bool                 `json:"IsTyre"`
	GoodNum             int                  `json:"goodNum"`
	InstalledNum        int                  `json:"installedNum"`
	CurrentInstalledNum int                  `json:"currentInstalledNum"`
	RemainderNum        int                  `json:"remainderNum"`
	SaleTyreList        []SaleTyreList       `json:"saleTyreList"`
	SaleConsumableList  []SaleConsumableList `json:"saleConsumableList"`
	BonusIntegralRate   float64              `json:"bonusIntegralRate"`
	OperateType         int                  `json:"operateType"`
	Note                string               `json:"note"`
}

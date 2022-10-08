package common_struct

import (
	"github.com/astaxie/beego/orm"
)

var ResponseStructMap = map[int16]ResponseStruct{
	0: {
		Code:    0,
		Message: "请求成功",
		Result:  nil,
	},
	4: {
		Code:    4,
		Message: "敬请期待",
		Result:  nil,
	},
	12: {
		Code:    12,
		Message: "新增成功!",
		Result:  nil,
	},
	14: {
		Code:    14,
		Message: "修改成功!",
		Result:  nil,
	},
	16: {
		Code:    16,
		Message: "删除成功!",
		Result:  nil,
	},
	-99: {
		Code:    -99,
		Message: "参数错误!",
		Result:  nil,
	},
	-1: {
		Code:    -1,
		Message: "请求失败!",
		Result:  nil,
	},
	-2: {
		Code:    -2,
		Message: "账号密码错误，请检查账密重新登陆",
		Result:  nil,
	},
	-3: {
		Code:    -3,
		Message: "token失效，请重新登录!",
		Result:  nil,
	},
	-4: {
		Code:    -4,
		Message: "原密码输入错误，请重新输入",
		Result:  nil,
	},
	-5: {
		Code:    -5,
		Message: "对应用户未找到，请检查请求信息",
		Result:  nil,
	},
	-6: {
		Code:    -6,
		Message: "当前角色已存在",
		Result:  nil,
	},

	-8: {
		Code:    -8,
		Message: "用户权限不足!",
		Result:  nil,
	},

	-9: { //当前状态区别于-8，前端不需要跳转页面
		Code:    -9,
		Message: "用户权限不足!",
		Result:  nil,
	},
	-11: {
		Code:    -11,
		Message: "当前角色存在关联用户，不允许删除!",
		Result:  nil,
	},
	-12: {
		Code:    -12,
		Message: "删除失败！",
		Result:  nil,
	},
	-13: {
		Code:    -13,
		Message: "新增失败!",
		Result:  nil,
	},
	-14: {
		Code:    -14,
		Message: "数据为空!",
		Result:  nil,
	},
	-15: {
		Code:    -15,
		Message: "修改失败!",
		Result:  nil,
	},
	-16: {
		Code:    -16,
		Message: "删除token信息失败，请联系管理员",
		Result:  nil,
	},
	-404: {
		Code:    -404,
		Message: "暂无当前api接口!",
		Result:  nil,
	},
	500: {
		Code:    500,
		Message: "系统异常",
		Result:  nil,
	},
}

type DeveiceResponseStruct struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type BatchImportTyreResultStruct struct {
	Err          []BatchExcelTyre `json:"err"`
	SuccessCount int              `json:"successCount"`
}

// TyreInfoStruct
type TyreInfoStruct struct {
	TyreNo            string  `json:"tyreNo"`
	TyreStatus        int8    `json:"tyreStatus"`
	StorageStatus     int8    `json:"storageStatus"`
	RetreadNum        int8    `json:"retreadNum"`
	TyreTemperature   string  `json:"tyreTemperature"`
	TyrePressure      string  `json:"tyrePressure"`
	WheelPosition     string  `json:"wheelPosition"`
	Price             float64 `json:"price"`
	ValueAddedTax     float64 `json:"valueAddedTax"`
	TaxPrice          float64 `json:"taxPrice"`
	InitPattern       float32 `json:"initPattern"`
	PatternContent    float32 `json:"patternContent"`
	InitPatterns      string  `json:"initPatterns"`
	Patterns          string  `json:"patterns"`
	InitialMileage    int     `json:"initialMileage"`
	CurrentMileage    int     `json:"currentMileage"`
	Brand             string  `json:"brand"`
	Specification     string  `json:"specification"`
	PatternModel      string  `json:"patternModel"`
	WearLimit         float32 `json:"wearLimit"`
	IsRetread         int8    `json:"isRetread"`
	RetreadBrand      string  `json:"retreadBrand"`
	RetreadlPattern   string  `json:"retreadlPattern"`
	SurplusValue      int8    `json:"surplusValue"`
	StationName       string  `json:"stationName"`
	EnterpriseName    string  `json:"enterpriseName"`
	FleetName         string  `json:"fleetName"`
	VehicleNumber     string  `json:"vehicleNumber"`
	CumulativeMileage int     `json:"cumulativeMileage"`
	CumulativePattern float64 `json:"cumulativePattern"`
	KPM               string  `json:"KPM"`
	CPK               string  `json:"CPK"`
	EstimatedMileage  string  `json:"estimatedMileage"`
	RemainMileage     string  `json:"remainMileage"`
}

type StationUserItem struct {
	UserId      string `json:"userId"`
	LoginName   string `json:"loginName"`
	WorkNo      string `json:"workNo"`
	DisplayName string `json:"displayName"`
}

type OperateListStruct struct {
	Count       int          `json:"count"`
	OperateList []orm.Params `json:"list"`
	TyreList    []orm.Params `json:"tyreList"`
}

type ResponseListStruct struct {
	Count int         `json:"count"`
	List  interface{} `json:"list"`
}

type QuantityMap struct {
	Brand         string `json:"brand"`
	PatternModel  string `json:"patternModel"`
	Specification string `json:"specification"`
	Quantity      int    `json:"quantity"`
}

type AnalysisTyreResponse struct {
	TyreNumCount                    int          `json:"tyreNumCount"`
	MonthTyreNum1Count              int          `json:"monthTyreNum1Count"`
	LastMonthTyreNum1Count          int          `json:"lastMonthTyreNum1Count"`
	TyreNumCountInfo                []orm.Params `json:"tyreNumCountInfo"`
	MonthStorageStatusCountInfo     []orm.Params `json:"monthStorageStatusCountInfo"`
	LastMonthStorageStatus1Count    int          `json:"lastMonthStorageStatus1Count"`
	MonthStorageStatus1CountOrMonth []orm.Params `json:"monthStorageStatus1CountOrMonth"`
	MonthNoInstalltyreNumCount      int          `json:"monthNoInstalltyreNumCount"`
	LastMonthNoInstalltyreNumCount  int          `json:"lastMonthNoInstalltyreNumCount"`
	InstallTyrePercentage           string       `json:"installTyrePercentage"`
	StorageStatusPercentage         string       `json:"storageStatusPercentage"`
	NoInstallTyrePercentage         string       `json:"noInstallTyrePercentage"`
}

type AnalysisVehicleResponse struct {
	VehicleCount                       int          `json:"vehicleCount"`
	YearVehicleCount                   int          `json:"yearVehicleCount"`
	MonthVehicleCount                  int          `json:"monthVehicleCount"`
	LastMonthVehicleCount              int          `json:"lastMonthVehicleCount"`
	MonthVehicleCountOrMonth           []orm.Params `json:"monthVehicleCountOrMonth"`
	VehicleTyreNumCount                int          `json:"vehicleTyreNumCount"`
	MonthVehicleTyreNumCount           int          `json:"monthVehicleTyreNumCount"`
	LastMonthVehicleTyreNumCount       int          `json:"lastMonthVehicleTyreNumCount"`
	VehicleInstallTyreNumCount         int          `json:"vehicleInstallTyreNumCount"`
	MonthVehicleTyreNumCountOrMonth    []orm.Params `json:"monthVehicleTyreNumCountOrMonth"`
	LastMothVehicleInstallTyreNumCount int          `json:"lastMothVehicleInstallTyreNumCount"`
	MothVehicleInstallTyreNumCount     int          `json:"mothVehicleInstallTyreNumCount"`
	YearVehicleInstallTyreNumCount     int          `json:"yearVehicleInstallTyreNumCount"`
	VehiclePercentage                  string       `json:"vehiclePercentage"`
	VehicleTyreNumPercentage           string       `json:"vehicleTyreNumPercentage"`
	VehicleInstallTyreNumPercentage    string       `json:"vehicleInstallTyreNumPercentage"`
}

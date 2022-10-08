package request_entity

type DriverList struct {
	PageNum  int `json:"pageNum" validate:"required"`
	PageSize int `json:"pageSize" validate:"required"`
	//DriverAccount string `json:"driverAccount"`
	VehicleNumber string `json:"vehicleNumber"`
	DriverName    string `json:"driverName"`
	IDCardNo      string `json:"idCardNo"`
	Tel           string `json:"tel"`
	BankUserName  string `json:"bankUserName"`
	BankNo        string `json:"bankNo"`
}

type AddDriver struct {
	DriverAccount  string      `json:"driverAccount" validate:"required"`
	DriverName     string      `json:"driverName" validate:"required"`
	VehicleNumber  string      `json:"vehicleNumber" validate:"required"`
	Tel            string      `json:"tel" validate:"required"`
	WorkNo         string      `json:"WorkNo" validate:"required"`
	Gender         int         `json:"gender"`
	IDCardNo       string      `json:"idCardNo"`
	Email          string      `json:"email"`
	BirthDate      string      `json:"birthDate"`
	LimitTotalLoad float64     `json:"limitTotalLoad"`
	Length         string      `json:"length"`
	BankUserName   string      `json:"bankUserName"`
	BankName       string      `json:"bankName"`
	BankNo         string      `json:"bankNo"`
	BagWeight      interface{} `json:"bagWeight"`
	DeductWeight   interface{} `json:"deductWeight"`
	ExtraWeight    interface{} `json:"extraWeight"`
	Specification  interface{} `json:"specification"`
}

type DeleteDriver struct {
	DriverIDs []int `json:"driverIDs"`
}

type UpdateDriver struct {
	DriverID         int         `json:"driverID" validate:"required"`
	DriverName       string      `json:"driverName" validate:"required"`
	VehicleNumber    string      `json:"vehicleNumber" validate:"required"`
	Tel              string      `json:"tel" validate:"required"`
	DriverAccount    string      `json:"driverAccount" `
	WorkNo           string      `json:"WorkNo" `
	UpdateUserInfo   bool        `json:"updateUserInfo"`
	UpdateDriverInfo bool        `json:"updateDriverInfo"`
	Gender           int         `json:"gender"`
	IDCardNo         string      `json:"idCardNo"`
	Email            string      `json:"email"`
	BirthDate        string      `json:"birthDate"`
	LimitTotalLoad   float64     `json:"limitTotalLoad"`
	Length           string      `json:"length"`
	BankUserName     string      `json:"bankUserName"`
	BankName         string      `json:"bankName"`
	BankNo           string      `json:"bankNo"`
	BagWeight        interface{} `json:"bagWeight"`
	DeductWeight     interface{} `json:"deductWeight"`
	ExtraWeight      interface{} `json:"extraWeight"`
	Specification    interface{} `json:"specification"`
}

package request_entity

type StationList struct {
	Name           string `json:"name"`
	ContactPerson  string `json:"contactPerson"`
	EnterpriseName string `json:"enterpriseName"`
	ContactTel     string `json:"contactTel"`
	EnterpryTel    string `json:"enterpryTel"`
	Address        string `json:"address"`
	PageNum        int    `json:"pageNum" validate:"required"`
	PageSize       int    `json:"pageSize" validate:"required"`
}

type UpdateStation struct {
	StationID      int     `json:"stationID"`
	StationName    string  `json:"stationName"`
	StationAddress string  `json:"stationAddress"`
	ContactPerson  string  `json:"contactPerson"`
	ContactTel     string  `json:"contactTel"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	EnterpriseIDs  []int   `json:"enterpriseIDs"`
	Province       string  `json:"province"`
	City           string  `json:"city"`
	Area           string  `json:"area"`
}

type AddStation struct {
	Name          string  `json:"name" validate:"required"`
	Address       string  `json:"address" validate:"required"`
	Province      string  `json:"province"`
	City          string  `json:"city"`
	Area          string  `json:"area"`
	ContactPerson string  `json:"contactPerson"`
	ContactTel    string  `json:"contactTel"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	EnterpriseIDs []int   `json:"enterpriseIDs"`
}

type DeleteStation struct {
	StationIDs []int `json:"stationIDs"`
}

package request_entity

type OriginList struct {
	PageNum     int    `json:"pageNum" validate:"required"`
	PageSize    int    `json:"pageSize" validate:"required"`
	Type        int    `json:"type"`
	OriginName  string `json:"originName"`
	ContactName string `json:"contactName"`
	Tel         string `json:"tel"`
	Address     string `json:"address"`
}

type AddOrigin struct {
	OriginName  string  `json:"originName" validate:"required"`
	Tel         string  `json:"tel" validate:"required"`
	Type        int     `json:"type" validate:"required"`
	ContactName string  `json:"contactName" `
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

type DeleteOrigin struct {
	OriginIDs []int `json:"originIDs"`
}

type UpdateOrigin struct {
	OriginID    int     `json:"originId" validate:"required"`
	OriginName  string  `json:"originName" validate:"required"`
	Tel         string  `json:"tel" validate:"required"`
	Type        int     `json:"type" validate:"required"`
	ContactName string  `json:"contactName" `
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

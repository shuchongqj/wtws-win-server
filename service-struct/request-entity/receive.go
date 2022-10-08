package request_entity

type ReceiveList struct {
	PageNum     int    `json:"pageNum" validate:"required"`
	PageSize    int    `json:"pageSize" validate:"required"`
	Type        int    `json:"type"`
	ReceiveName string `json:"receiveName"`
	ContactName string `json:"contactName"`
	Tel         string `json:"tel"`
	Address     string `json:"address"`
}

type AddReceive struct {
	ReceiveName string  `json:"receiveName" validate:"required"`
	Tel         string  `json:"tel" validate:"required"`
	Type        int     `json:"type" validate:"required"`
	ContactName string  `json:"contactName" `
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

type DeleteReceive struct {
	ReceiveIDs []int `json:"receiveIDs"`
}

type UpdateReceive struct {
	ReceiveID   int     `json:"receiveId" validate:"required"`
	ReceiveName string  `json:"receiveName" validate:"required"`
	Tel         string  `json:"tel" validate:"required"`
	Type        int     `json:"type" validate:"required"`
	ContactName string  `json:"contactName" `
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

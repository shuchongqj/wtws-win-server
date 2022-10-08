package request_entity

type CargotoList struct {
	PageNum     int    `json:"pageNum" validate:"required"`
	PageSize    int    `json:"pageSize" validate:"required"`
	CargotoName string `json:"cargotoName"`
	Code        string `json:"code"`
}

type AddCargoto struct {
	CargotoName string `json:"cargotoName" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

type DeleteCargoto struct {
	CargotoIDs []int `json:"cargotoIDs"`
}

type UpdateCargoto struct {
	CargotoID   int    `json:"cargotoId" validate:"required"`
	CargotoName string `json:"cargotoName" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

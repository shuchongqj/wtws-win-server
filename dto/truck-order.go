package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type TruckOrderList struct {
	List  []wtws_mysql.OrTruckOrder `json:"list"`
	Count int                       `json:"count"`
}

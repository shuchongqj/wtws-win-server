package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type OrderList struct {
	List  []wtws_mysql.OrOrder `json:"list"`
	Count int                  `json:"count"`
}

package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type ReceiveList struct {
	List  []wtws_mysql.OReceive `json:"list"`
	Count int                   `json:"count"`
}

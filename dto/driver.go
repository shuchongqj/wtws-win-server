package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type DriverList struct {
	List  []wtws_mysql.SDriverList `json:"list"`
	Count int                      `json:"count"`
}

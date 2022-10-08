package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type OriginList struct {
	List  []wtws_mysql.OOrigin `json:"list"`
	Count int                  `json:"count"`
}

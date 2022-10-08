package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type GoodsList struct {
	List  []wtws_mysql.GoodListItem `json:"list"`
	Count int                       `json:"count"`
}

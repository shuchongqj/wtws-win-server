package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type CategoryList struct {
	List  []wtws_mysql.GCategory `json:"list"`
	Count int                    `json:"count"`
}

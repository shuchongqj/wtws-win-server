package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type CargotoList struct {
	List  []wtws_mysql.OCargoto `json:"list"`
	Count int                   `json:"count"`
}

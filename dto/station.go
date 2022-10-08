package dto

import wtws_mysql "wtws-server/models/wtws-mysql"

type StationList struct {
	List  []wtws_mysql.StationListStruct `json:"list"`
	Count int                            `json:"count"`
}

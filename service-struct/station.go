package service_struct

// UserListItemStation 用户列表中每个用户
type UserListItemStation struct {
	StationID   int    `json:"stationID"`
	StationName string `json:"stationName"`
}

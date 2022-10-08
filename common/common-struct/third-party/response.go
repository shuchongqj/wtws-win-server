package third_party

type AuthResponseResult struct {
	UserID string `json:"userID"`
}

type AuthResponse struct {
	Code   int16 `json:"code"`
	Result AuthResponseResult
}

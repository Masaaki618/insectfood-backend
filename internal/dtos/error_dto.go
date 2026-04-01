package dtos

// ErrorResponse はエラーレスポンスの共通形式
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail はエラーの詳細
type ErrorDetail struct {
	Message string `json:"message"`
}

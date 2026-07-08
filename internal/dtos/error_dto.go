package dtos

// ErrorResponse はエラーレスポンスの共通形式
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail はエラーの詳細
type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

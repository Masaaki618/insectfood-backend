package dtos

// Response はAPIレスポンスの共通ラッパー
type Response struct {
	Data any `json:"data"`
}

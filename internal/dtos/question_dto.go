package dtos

// QuestionResponse は質問一覧APIのレスポンス（1件分）
type QuestionResponse struct {
	ID       uint   `json:"id"`
	Body     string `json:"body"`
	Category string `json:"category"`
}

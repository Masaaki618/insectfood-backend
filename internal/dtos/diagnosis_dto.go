package dtos

// DiagnosisRequest は診断APIのリクエスト
type DiagnosisRequest struct {
	Scores DiagnosisScores `json:"scores" validate:"required"`
}

// DiagnosisScores はカテゴリ別スコア（各0〜2点）
type DiagnosisScores struct {
	Visual   uint8 `json:"visual"   validate:"max=2"`
	Physical uint8 `json:"physical" validate:"max=2"`
	Mental   uint8 `json:"mental"   validate:"max=2"`
}

// DiagnosisResponse は診断APIのレスポンス
type DiagnosisResponse struct {
	Insect    InsectResponse `json:"insect"`
	AIComment string         `json:"ai_comment"`
}

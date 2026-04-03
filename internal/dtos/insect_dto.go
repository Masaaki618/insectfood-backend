package dtos

// InsectResponse は昆虫一覧・詳細APIのレスポンス（1件分）
type InsectResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Difficulty   uint8   `json:"difficulty"`
	Introduction string  `json:"introduction"`
	Taste        string  `json:"taste"`
	Texture      string  `json:"texture"`
	InsectImg    *string `json:"insect_img"`
}

// InsectDetailResponse は昆虫詳細APIのレスポンス（レーダーチャート・AIコメント付き）
type InsectDetailResponse struct {
	InsectResponse
	AIComment  string             `json:"ai_comment"`
	RadarChart RadarChartResponse `json:"radar_chart"`
}

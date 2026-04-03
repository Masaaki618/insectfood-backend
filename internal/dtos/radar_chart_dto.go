package dtos

// RadarChartResponse はレーダーチャートスコアのレスポンス
type RadarChartResponse struct {
	UmamiScore  uint8 `json:"umami_score"`
	BitterScore uint8 `json:"bitter_score"`
	EguScore    uint8 `json:"egu_score"`
	FlavorScore uint8 `json:"flavor_score"`
	KimoScore   uint8 `json:"kimo_score"`
}

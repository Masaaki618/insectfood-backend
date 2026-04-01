package ai

import "os"

// ClaudeClient はClaude APIへのアクセスを管理するクライアント
type ClaudeClient struct {
	apiKey string
}

// NewClaudeClient は環境変数からAPIキーを取得しClaudeClientを生成する
func NewClaudeClient() *ClaudeClient {
	return &ClaudeClient{
		apiKey: os.Getenv("ANTHROPIC_API_KEY"),
	}
}

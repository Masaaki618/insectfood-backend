package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

const maxTokens = 1024

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

// callAPI はプロンプトをClaude APIに送信してレスポンスのテキストを返す内部メソッド
func (c *ClaudeClient) callAPI(ctx context.Context, prompt string) (string, error) {
	client := anthropic.NewClient(
		option.WithAPIKey(c.apiKey),
	) //nolint:mnd

	message, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		MaxTokens: maxTokens,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
		Model: anthropic.ModelClaudeHaiku4_5,
	})
	if err != nil {
		return "", err
	}
	return message.Content[0].Text, nil
}

// buildInsectCommentPrompt は昆虫詳細用のプロンプトを組み立てる
func (c *ClaudeClient) buildInsectCommentPrompt(insect *models.Insect) string {
	return fmt.Sprintf(`あなたは昆虫食のアドバイザーです。
	以下のルールを必ず守ってください：
	- コメント文のみを出力する。JSON・マークダウン・コードブロック・説明文は一切つけない
	- 80文字以上100文字以内にする
	- 初心者が読んで食べてみたくなる内容にする
	- 味・食感・見た目の特徴を具体的に盛り込む

	【昆虫情報】
	・名前: %s
	・味: %s
	・食感: %s
	・説明: %s

	【良いコメントの例】
	ナッツ系の香ばしい風味が特徴。粉末状なので虫の形が残らず初心者でも安心。スムージーやお菓子に混ぜるだけで手軽にタンパク質を補給できます

	【悪いコメントの例（禁止）】
	「美味しいです」「食べてみてください」などの薄い内容

	コメント文のみを返すこと:`, insect.Name, insect.Taste, insect.Texture, insect.Introduction)
}

// buildDiagnosisResultPrompt はスコアと昆虫リスト用のプロンプトを組み立てる
func (c *ClaudeClient) buildDiagnosisResultPrompt(visual, physical, mental uint8, insects []models.Insect) string {
	insectsJSON, _ := json.Marshal(insects)
	return fmt.Sprintf(`あなたは昆虫食のアドバイザーです。

		重要：コードブロック（` + "```" + `）は絶対に使わないこと。JSONのみを生のテキストで出力すること。

		以下のルールを必ず守ってください：
		- 出力はJSON形式のみ。前後に説明文・マークダウン・コードブロックを一切つけない
		- JSONのキー以外の文字列を出力しない
		- ai_commentは必ず80文字以上100文字以内にする
		- 必ず昆虫リストの中からinsect_idを1つ選ぶ
		- ユーザーのスコアに基づいて最適な昆虫を選ぶ

		【ユーザーの耐性スコア】
		・見た目への耐性（visual）: %d/2点
		・食べる勇気（physical）: %d/2点
		・挑戦する気持ち（mental）: %d/2点

		【選択可能な昆虫リスト】
		%s

		【良いコメントの例】
		「挑戦する気持ちが強いあなたに、白くてまるっとしたシルクワームはいかがでしょう。クリーミーな味わいにきっと驚くはずです。初めての昆虫食にぴったりです」

		【悪いコメントの例（禁止）】
		「おすすめです」「食べてみてください」などの薄い内容

		出力形式（このJSONのみを返すこと）:
		{"insect_id": 1, "ai_comment": "コメント"}`, visual, physical, mental, insectsJSON)
}

// GenerateInsectComment は昆虫情報をもとにAIコメントを生成して返す
func (c *ClaudeClient) GenerateInsectComment(ctx context.Context, insect *models.Insect) (string, error) {
	prompt := c.buildInsectCommentPrompt(insect)

	// APIを呼ぶ（最大3回リトライ）
	var result string
	var err error
	for range 3 {
		result, err = c.callAPI(ctx, prompt)
		if err == nil {
			break
		}
	}
	if err != nil {
		return "", fmt.Errorf("GenerateInsectComment: %w", err)
	}

	return result, nil
}

// GenerateDiagnosisResult はスコアと昆虫リストをもとにAIがおすすめ昆虫IDとコメントを返す
func (c *ClaudeClient) GenerateDiagnosisResult(ctx context.Context, visual, physical, mental uint8, insects []models.Insect) (uint, string, error) {
	prompt := c.buildDiagnosisResultPrompt(visual, physical, mental, insects)

	// APIを呼ぶ（最大3回リトライ）
	var result string
	var err error
	for range 3 {
		result, err = c.callAPI(ctx, prompt)
		if err == nil {
			break
		}
	}
	if err != nil {
		return 0, "", fmt.Errorf("GenerateDiagnosisResult: %w", err)
	}

	// JSONをパースして insect_id と ai_comment を取り出す
	var parsed struct {
		InsectID  uint   `json:"insect_id"`
		AIComment string `json:"ai_comment"`
	}
	if err := json.Unmarshal([]byte(result), &parsed); err != nil {
		return 0, "", fmt.Errorf("GenerateDiagnosisResult JSONパース失敗: %w", err)
	}

	return parsed.InsectID, parsed.AIComment, nil
}

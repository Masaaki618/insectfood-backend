package controllers

import (
	"net/http"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"github.com/gin-gonic/gin"
)

// QuestionController は質問に関するHTTPリクエストを処理するコントローラー
type QuestionController struct {
	service services.IQuestionService
}

// NewQuestionController はQuestionControllerを生成する
func NewQuestionController(service services.IQuestionService) *QuestionController {
	return &QuestionController{service: service}
}

// GetQuestions はカテゴリ別にランダムで6問取得してJSONで返す
func (c *QuestionController) GetQuestions(ctx *gin.Context) {
	questions, err := c.service.GetQuestions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

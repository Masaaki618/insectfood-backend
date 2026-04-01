package controllers

import (
	"net/http"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"github.com/gin-gonic/gin"
)

// DiagnosisController は診断に関するHTTPリクエストを処理するコントローラー
type DiagnosisController struct {
	service services.IDiagnosisService
}

// NewDiagnosisController はDiagnosisControllerを生成する
func NewDiagnosisController(service services.IDiagnosisService) *DiagnosisController {
	return &DiagnosisController{service: service}
}

// Diagnose はスコアを受け取りAIがレコメンドした昆虫とコメントをJSONで返す
func (c *DiagnosisController) Diagnose(ctx *gin.Context) {
	var diagnosisRequest dtos.DiagnosisRequest
	if err := ctx.ShouldBind(&diagnosisRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "invalid request",
			},
		})
		return
	}

	diagnosis, err := c.service.Diagnose(ctx.Request.Context(), diagnosisRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, diagnosis)
}

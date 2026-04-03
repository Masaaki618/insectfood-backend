package controllers

import (
	"net/http"
	"strconv"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"github.com/gin-gonic/gin"
)

// InsectController は昆虫に関するHTTPリクエストを処理するコントローラー
type InsectController struct {
	service services.IInsectService
}

// NewInsectController はInsectControllerを生成する
func NewInsectController(service services.IInsectService) *InsectController {
	return &InsectController{service: service}
}

// GetInsects は昆虫一覧を取得してJSONで返す
func (c *InsectController) GetInsects(ctx *gin.Context) {
	insects, err := c.service.GetInsects(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, insects)
}

// GetInsectByID は指定IDの昆虫詳細を取得してJSONで返す
func (c *InsectController) GetInsectByID(ctx *gin.Context) {
	intStr := ctx.Param("id")
	id, err := strconv.ParseUint(intStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "invalid id",
			},
		})
		return
	}
	insect, err := c.service.GetInsectByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Error: dtos.ErrorDetail{
				Message: "internal server error",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, insect)
}

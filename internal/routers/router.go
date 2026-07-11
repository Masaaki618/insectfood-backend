package routers

import (
	"net/http"

	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router は全エンドポイントのルーティングを管理する
type Router struct {
	db                  *gorm.DB
	insectController    *controllers.InsectController
	questionController  *controllers.QuestionController
	diagnosisController *controllers.DiagnosisController
}

// NewRouter はRouterを生成する
func NewRouter(db *gorm.DB, insectController *controllers.InsectController, questionController *controllers.QuestionController, diagnosisController *controllers.DiagnosisController) *Router {
	return &Router{
		db:                  db,
		insectController:    insectController,
		questionController:  questionController,
		diagnosisController: diagnosisController,
	}
}

// Setup はGinエンジンにルーティングを登録する
func (r *Router) Setup(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/insects", r.insectController.GetInsects)
		v1.GET("/insects/:id", r.insectController.GetInsectByID)
		v1.GET("/questions", r.questionController.GetQuestions)
		v1.POST("/diagnosis", r.diagnosisController.Diagnose)
	}
	engine.GET("/health", func(ctx *gin.Context) {
		sqlDB, err := r.db.DB()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, dtos.ErrorResponse{
				Error: dtos.ErrorDetail{
					Code:    http.StatusServiceUnavailable,
					Message: "service unavailable",
				},
			})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			ctx.JSON(http.StatusServiceUnavailable, dtos.ErrorResponse{
				Error: dtos.ErrorDetail{
					Code:    http.StatusServiceUnavailable,
					Message: "service unavailable",
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

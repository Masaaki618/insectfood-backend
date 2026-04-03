package routers

import (
	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Router は全エンドポイントのルーティングを管理する
type Router struct {
	insectController    *controllers.InsectController
	questionController  *controllers.QuestionController
	diagnosisController *controllers.DiagnosisController
}

// NewRouter はRouterを生成する
func NewRouter(insectController *controllers.InsectController, questionController *controllers.QuestionController, diagnosisController *controllers.DiagnosisController) *Router {
	return &Router{
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
}

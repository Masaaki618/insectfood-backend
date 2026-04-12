package di

import (
	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/infrastructure/ai"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
	"github.com/Masaaki618/insectfood-backend/internal/routers"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"gorm.io/gorm"
)

// NewContainer はDB接続を受け取り各層の依存関係を組み立ててRouterを返す
func NewContainer(db *gorm.DB) *routers.Router {
	// Repository層の生成
	insectRepository := repositories.NewInsectRepository(db)
	questionRepository := repositories.NewQuestionRepository(db)

	// Claude APIクライアントの生成
	claudeClient := ai.NewClaudeClient()

	// Service層の生成
	insectService := services.NewInsectService(insectRepository, claudeClient)
	questionService := services.NewQuestionService(questionRepository)
	diagnosisService := services.NewDiagnosisService(insectRepository, claudeClient)

	// Controller層の生成
	insectController := controllers.NewInsectController(insectService)
	questionController := controllers.NewQuestionController(questionService)
	diagnosisController := controllers.NewDiagnosisController(diagnosisService)

	// Routerの生成
	return routers.NewRouter(insectController, questionController, diagnosisController)
}

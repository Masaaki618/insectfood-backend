package main

import (
	"fmt"
	"os"

	"github.com/Masaaki618/insectfood-backend/internal/controllers"
	"github.com/Masaaki618/insectfood-backend/internal/infrastructure/database"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
	"github.com/Masaaki618/insectfood-backend/internal/routers"
	"github.com/Masaaki618/insectfood-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	insectRepository := repositories.NewInsectRepository(db)
	questionRepository := repositories.NewQuestionRepository(db)
	insectService := services.NewInsectService(insectRepository)
	questionService := services.NewQuestionService(questionRepository)
	diagnosisService := services.NewDiagnosisService(insectRepository)
	insectController := controllers.NewInsectController(insectService)
	questionController := controllers.NewQuestionController(questionService)
	diagnosisController := controllers.NewDiagnosisController(diagnosisService)
	newRouter := routers.NewRouter(insectController, questionController, diagnosisController)
	engine := gin.Default()
	newRouter.Setup(engine)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

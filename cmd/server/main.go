package main

import (
	"fmt"
	"os"

	"github.com/Masaaki618/insectfood-backend/internal/di"
	"github.com/Masaaki618/insectfood-backend/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB接続
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	// 依存関係の組み立て
	router := di.NewContainer(db)

	// Ginエンジンの起動
	engine := gin.Default()
	router.Setup(engine)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

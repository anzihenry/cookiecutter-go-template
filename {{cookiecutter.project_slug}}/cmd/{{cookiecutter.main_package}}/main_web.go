package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"{{cookiecutter.project_slug}}/internal/logger"
)

func main() {
	logger.InitLogger() // 初始化日志
	logger.Logger.Info("Starting web server")

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		logger.Logger.Info("Health check endpoint hit")
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Failed to run web server", zap.Error(err))
	}
}

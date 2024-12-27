package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"{{cookiecutter.project_slug}}/internal/logger"
)

func main() {
	logger.InitLogger() // 初始化日志
	logger.Logger.Info("Starting console application")

	var rootCmd = &cobra.Command{
		Use:   "{{cookiecutter.project_slug}}",
		Short: "A brief description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Logger.Info("Running command", zap.Strings("args", args))
		},
	}

	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Error("Command execution failed", zap.Error(err))
	}
}

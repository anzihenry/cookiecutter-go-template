package main

import (
	"fmt"
	"{{cookiecutter.project_slug}}/pkg/{{cookiecutter.package_name}}"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	{{cookiecutter.package_name}}.SetLogger(logger)

	result := {{cookiecutter.package_name}}.Add(2, 3)
	fmt.Println("Result:", result)
}

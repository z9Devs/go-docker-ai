package main

import (
	"fmt"
	"os"

	"github.com/la-plas-growth/GO-DockerLint-AI/cmd/service/lint"
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	"go.uber.org/zap"
)

func main() {
	// Configure the logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Load environment configuration
	config := env.NewConfiguration()

	// Create the linting service
	service := lint.NewService(config, sugar)

	// Get Dockerfile content from user input or file
	if len(os.Args) < 2 {
		sugar.Fatalf("Usage: %s <path-to-Dockerfile>", os.Args[0])
	}
	dockerfilePath := os.Args[1]

	// Read Dockerfile content
	dockerfileContent, err := os.ReadFile(dockerfilePath)
	if err != nil {
		sugar.Fatalf("Failed to read Dockerfile: %v", err)
	}

	// Analyze the Dockerfile
	lintResponse, err := service.AnalyzeDockerFile(string(dockerfileContent))
	if err != nil {
		sugar.Fatalf("Failed to analyze Dockerfile: %v", err)
	}

	// Display the linting results
	fmt.Println("Linting Results:")
	for _, issue := range lintResponse.Issues {
		fmt.Printf("Line %d [%s]: %s\n", issue.LineNumber, issue.Severity, issue.Message)
	}
}
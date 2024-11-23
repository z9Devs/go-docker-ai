package lint

import (
	"context"
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	"github.com/la-plas-growth/GO-DockerLint-AI/lib"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

// IService defines the interface for the linting service
type IService interface {
	AnalyzeDockerFile(dockerfile string) (*LintResponse, error)
	FetchBestPracticesMarkdown() (string, error)
}

// service implements the IService interface
type service struct {
	configuration *env.Configuration
	logger        *zap.SugaredLogger
	client        *openai.Client
}

// NewService creates a new instance of the linting service
func NewService(configuration *env.Configuration, logger *zap.SugaredLogger) IService {
	if configuration.OpenAIAPIKey == "" {
		logger.Panic("OpenAI API key is required")
	}
	client := openai.NewClient(configuration.OpenAIAPIKey)
	return &service{
		configuration: configuration,
		logger:        logger,
		client:        client,
	}
}

// AnalyzeDockerFile analyzes a Dockerfile using OpenAI and best practices
func (s *service) AnalyzeDockerFile(dockerfile string) (*LintResponse, error) {
	s.logger.Debugf("Analyzing Dockerfile with OpenAI dockerfile path: %s", dockerfile)
	content, err := lib.GetFileContent(dockerfile)
	if err != nil {
		s.logger.Errorf("Failed to read Dockerfile: %v", err)
		return nil, err
	}
	if content == "" {
		s.logger.Errorf("Dockerfile is empty")
		return nil, fmt.Errorf("dockerfile is empty")
	}
	// Fetch best practices content
	bestPractices, err := s.FetchBestPracticesMarkdown()
	if err != nil {
		s.logger.Errorf("Failed to fetch best practices: %v", err)
		return nil, err
	}

	// Clean up best practices to format properly for the prompt
	cleanedBestPractices := strings.ReplaceAll(bestPractices, "\n", " ")

	// Create OpenAI API context
	ctx := context.Background()

	// Construct the prompt for OpenAI
	prompt := fmt.Sprintf(`You are a Dockerfile linter. Use the following best practices as guidance:
	%s

	Analyze the Dockerfile below. For each issue, provide:
	- Line number(s)
	- Severity (info, warning, error)
	- Description

	Dockerfile:
	%s`, cleanedBestPractices, content)

	s.logger.Debugf("prompt: %s", prompt)

	// Call OpenAI API
	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a Dockerfile linter.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})
	if err != nil {
		s.logger.Errorf("Error communicating with OpenAI: %v", err)
		return nil, err
	}

	// Parse the response
	responseContent := resp.Choices[0].Message.Content
	s.logger.Debugf("OpenAI response: %s", responseContent)

	return &LintResponse{Content: responseContent}, nil
}

func (s *service) FetchBestPracticesMarkdown() (string, error) {
	s.logger.Debugf("Fetching best practices from URL: %s", s.configuration.BestPracticesURL)
	resp, err := soup.Get(s.configuration.BestPracticesURL)
	//
	if err != nil {
		s.logger.Errorf("Unexpected status code: %d", err)
		return "", fmt.Errorf("unexpected status code: %d", err)
	}
	doc := soup.HTMLParse(resp)
	//
	return doc.FullText(), nil
}

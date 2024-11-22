package lint

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
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
	restyClient   *resty.Client
}

// NewService creates a new instance of the linting service
func NewService(configuration *env.Configuration, logger *zap.SugaredLogger) IService {
	client := openai.NewClient(configuration.OpenAIAPIKey)
	return &service{
		configuration: configuration,
		logger:        logger,
		client:        client,
		restyClient:   resty.New(),
	}
}

// AnalyzeDockerFile analyzes a Dockerfile using OpenAI and best practices
func (s *service) AnalyzeDockerFile(dockerfile string) (*LintResponse, error) {
	s.logger.Debugf("Analyzing Dockerfile with OpenAI: %s", dockerfile)

	// Fetch best practices content
	bestPractices, err := FetchBestPracticesMarkdown(s.configuration.BestPracticesURL, s.logger)
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
%s`, cleanedBestPractices, dockerfile)

	// Call OpenAI API
	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4",
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

	lines := strings.Split(responseContent, "\n")
	var issues []LintIssue
	for _, line := range lines {
		if strings.HasPrefix(line, "Line") {
			parts := strings.Split(line, ":")
			if len(parts) >= 3 {
				// Extract line number, severity, and message
				lineNumber, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
				severity := strings.TrimSpace(parts[2])
				message := strings.Join(parts[3:], ":")
				issues = append(issues, LintIssue{
					LineNumber: lineNumber,
					Severity:   severity,
					Message:    message,
				})
			} else {
				s.logger.Warnf("Failed to parse line: %s", line)
			}
		}
	}

	return &LintResponse{Issues: issues}, nil
}


func (s *service) FetchBestPracticesMarkdown() (string, error) {
	logger.Debugf("Fetching best practices from URL: %s", s.configuration.BestPracticesURL)

	resp, err := s.restyClient.R().
		Get(s.configuration.BestPracticesURL)
	//
	if err != nil {
		logger.Errorf("Unexpected status code: %d", resp.StatusCode)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("resp %s", result.String())
	}
	
	// Analizza il contenuto Markdown
	content := string(body)
	var sections []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Identifica i titoli principali
		if strings.HasPrefix(line, "#") {
			sections = append(sections, "\n"+line) // Aggiunge una nuova sezione
		} else if len(line) > 0 {
			// Aggiunge paragrafi o contenuti sotto il titolo
			if len(sections) > 0 {
				sections[len(sections)-1] += "\n" + line
			}
		}
	}

	if len(sections) == 0 {
		s.logger.Warn("No relevant content found in the Markdown file.")
		return "", fmt.Errorf("no relevant content found")
	}

	// Combina le sezioni in un risultato leggibile
	result := strings.Join(sections, "\n\n")
	s.logger.Info("Best practices fetched successfully.")
	return result, nil

}
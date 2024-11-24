package lint

import (
	"context"
	"encoding/json"
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

	lintResponse, err := s.analyzeWithChatGPT(cleanedBestPractices, content)
	if err != nil {
		s.logger.Errorf("Failed to analyzeWithChatGPT: %v", err)
		return nil, err
	}

	return lintResponse, nil
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

func (s *service) analyzeWithChatGPT(cleanedBestPractices, content string) (*LintResponse, error) {
	// Create OpenAI API context
	ctx := context.Background()

	// Construct the prompt for OpenAI
	prompt := fmt.Sprintf(`You are a Dockerfile linter. Use only the best practices provided as guidance:
	%s
	
	Analyze the Dockerfile below. For each issue, provide:
	- A brief description of the issue in the field "issue".
	- The severity level of the issue (high, medium, low) in the field "severity".
	- Actionable advice to resolve the issue in the field "advice".
	
	Base your analysis strictly and exclusively on the provided best practices and the content of the Dockerfile below.  
	DO NOT include any suggestions or comments about .dockerignore, as you cannot know if the user has used it.
	
	Dockerfile:
	%s`, cleanedBestPractices, content)

	s.logger.Debugf("prompt: %s", prompt)

	// Define the JSON schema for the response
	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"issues": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"issue":    map[string]string{"type": "string"},
						"severity": map[string]string{"type": "string"},
						"advice":   map[string]string{"type": "string"},
					},
					"required":             []string{"issue", "severity", "advice"},
					"additionalProperties": false,
				},
			},
		},
		"required":             []string{"issues"},
		"additionalProperties": false,
	}

	// Serialize the schema into JSON
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		s.logger.Errorf("Failed to serialize schema: %v", err)
		return nil, fmt.Errorf("failed to serialize schema: %w", err)
	}

	// Define the response format
	responseFormat := &openai.ChatCompletionResponseFormat{
		Type: "json_schema",
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        "DockerfileLintIssues",
			Description: "Linting issues for a Dockerfile",
			Schema:      json.RawMessage(schemaBytes),
			Strict:      true,
		},
	}

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
		ResponseFormat: responseFormat,
	})
	if err != nil {
		s.logger.Errorf("Error communicating with OpenAI: %v", err)
		return nil, err
	}

	// Parse the response
	responseContent := resp.Choices[0].Message.Content
	s.logger.Debugf("OpenAI response: %s", responseContent)

	// Decode the JSON response into the LintResponse structure
	var lintResponse LintResponse
	if err := json.Unmarshal([]byte(responseContent), &lintResponse); err != nil {
		s.logger.Errorf("Failed to parse JSON response: %v", err)
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}
	return &lintResponse, nil
}

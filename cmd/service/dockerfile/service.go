package dockerfile

import (
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

// IService defines the interface for the dockerfile service
type IService interface {
	CreateDockerFile(lang string) (*DockerfileResponse, error)
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

// Create the finale Dockerfile
func (s *service) CreateDockerFile(lang string) (*DockerfileResponse, error) {
	// TODO
	return nil, nil
}

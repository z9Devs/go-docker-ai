package lint

import (
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	"go.uber.org/zap"
)

type IService interface {
	//
	AnalyzeDockerFile(dockerfile string) (*LintResponse, error)
	//FetchBasePractices() error
	//ReadDockerFile() error
	//
}

type service struct {
	configuration *env.Configuration
	logger        *zap.SugaredLogger
}

func NewService(configuration *env.Configuration, logger *zap.SugaredLogger) IService {
	return &service{configuration: configuration, logger: logger}
}

func (s *service) AnalyzeDockerFile(dockerfile string) (*LintResponse, error) {
	s.logger.Debugf("Analyzing dockerfile: %s", dockerfile)
	return nil, nil
}

package project

// IService defines the interface for the dockerfile service
type IService interface {

}

// service implements the IService interface
type service struct {
	configuration *env.Configuration
	logger        *zap.SugaredLogger
}


// NewService creates a new instance of the linting service
func NewService(configuration *env.Configuration, logger *zap.SugaredLogger) IService {
    return &service{
		configuration: configuration,
		logger:        logger,
		//client:        client,
	}
}
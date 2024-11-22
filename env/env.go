package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configuration struct {
	OpenAIAPIKey      string `env:"OPENAI_API_KEY" envDefault:""` 
	LogLevel          string `env:"LOG_LEVEL" envDefault:"info"` 
	BestPracticesURL  string `env:"BEST_PRACTICES_URL" envDefault:"https://github.com/docker/docs/blob/main/content/manuals/build/building/best-practices.md"`
}

func NewConfiguration() *Configuration {

	err := godotenv.Load()
	if err != nil {
	}

	configuration := Configuration{}

	if err := env.Parse(&configuration); err != nil {
		panic(err)
	}

	return &configuration
}

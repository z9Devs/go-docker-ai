package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configuration struct {
    OpenApiKey          string `env:"OPENAI_API_KEY" envDefault:"info"`
	LogLevel            string `env:"LOG_LEVEL" envDefault:"info"`
}

func NewConfiguration() *Configuration {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	configuration := Configuration{
		LogLevel: "info",
	}

	if err := env.Parse(&configuration); err != nil {
		panic(err)
	}

	return &configuration
}

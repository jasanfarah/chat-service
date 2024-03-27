package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"jasanfarah/chat-service/internal/database"
)

// DatabaseConfig holds configuration parameters for the database.

// Configuration holds the entire application configuration.
type Configuration struct {
	Database database.PostgresConfig `yaml:"database"`
	API      struct {
		Version string `yaml:"version"`
	} `yaml:"api"`
}

var AppConfig Configuration

func LoadConfig(env string) Configuration {
	// Set the environment default to 'local' if not provided
	environment := env
	if environment == "" {
		environment = "local"
	}

	viper.SetConfigName("config." + environment + ".yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file:", err)
		panic(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Error("Error unmarshalling config:", err)
		panic(err)
	}

	log.Infof("Initialized application configurations for environment: %s", environment)

	return AppConfig
}

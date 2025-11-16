package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var configuration *Config

func New() (*Config, error) {
	configuration = &Config{}
	if err := Load(configuration); err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	return configuration, nil
}

// Load reads the configuration from file and environment variables.
func Load(object interface{}) error {
	configFolders := []string{
		"./config/",
		"./internal/config/",
		"../../internal/config/", // relative path for debugger
	}

	for _, configFolder := range configFolders {
		// Add Config File Path
		viper.AddConfigPath(configFolder)
	}

	// Config File Name
	viper.SetConfigName("config")
	// Config File Type
	viper.SetConfigType("yaml")

	// Enable environment variable override
	// Environment variables should use underscore instead of dots
	// Example: ASYNQ_REDISADDR will override asynq.redisAddr
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(&object)
}

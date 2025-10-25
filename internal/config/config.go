package config

import (
	"log"

	"github.com/spf13/viper"
)

var configuration *Config

func New() (*Config, error) {
	configuration = &Config{}
	if err := Load(configuration); err == nil {
		log.Fatal(err)
	}

	return configuration, nil
}

// Load reads the configuration from file and environment variables.
func Load(object interface{}) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("./internal/config")
	v.AddConfigPath("./configs")
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// defaults
	v.SetDefault("app.name", "hcm-be")
	v.SetDefault("app.env", "development")
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.readTimeout", "10s")
	v.SetDefault("server.writeTimeout", "15s")
	v.SetDefault("server.idleTimeout", "60s")
	v.SetDefault("server.requestTimeout", "15s")
	v.SetDefault("observability.metricsEnabled", true)
	v.SetDefault("observability.pprofEnabled", true)
	v.SetDefault("database.driver", "memory")
	v.SetDefault("database.dsn", "")
	v.SetDefault("database.maxOpenConnections", 25)
	v.SetDefault("database.maxIdleConnections", 25)
	v.SetDefault("database.maxConnectionLifetime", "5m")
	v.SetDefault("database.maxConnectionIdleTime", "5m")
	v.SetDefault("webhook.apiKey", "your-webhook-api-key")
	v.SetDefault("webhook.hmacSecret", "your-hmac-secret-key")
	v.SetDefault("featureFlag.webhook.enableSignatureValidation", true)
	v.SetDefault("featureFlag.webhook.enableTimestampValidation", true)

	_ = v.ReadInConfig() // ignore if missing

	return viper.Unmarshal(&object)
}

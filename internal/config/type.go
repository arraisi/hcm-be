package config

import (
	"time"
)

type App struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
}
type Server struct {
	Host           string        `yaml:"host"`
	Port           int           `yaml:"port"`
	ReadTimeout    time.Duration `yaml:"readTimeout"`
	WriteTimeout   time.Duration `yaml:"writeTimeout"`
	IdleTimeout    time.Duration `yaml:"idleTimeout"`
	RequestTimeout time.Duration `yaml:"requestTimeout"`
}
type Observability struct {
	MetricsEnabled bool `yaml:"metricsEnabled"`
	PprofEnabled   bool `yaml:"pprofEnabled"`
}
type Database struct {
	Driver                string        `yaml:"driver"` // Supported: memory, postgres, sqlserver
	DSN                   string        `yaml:"dsn"`    // Data Source Name - connection string for the database
	MaxOpenConnections    int           `yaml:"maxOpenConnections"`
	MaxIdleConnections    int           `yaml:"maxIdleConnections"`
	MaxConnectionLifetime time.Duration `yaml:"maxConnectionLifetime"`
	MaxConnectionIdleTime time.Duration `yaml:"maxConnectionIdleTime"`
}
type Webhook struct {
	APIKey     string `yaml:"apiKey"`     // API key for webhook authentication
	HMACSecret string `yaml:"hmacSecret"` // HMAC secret for signature verification
}

type FeatureFlag struct {
	WebhookConfig WebhookFeatureConfig `yaml:"webhook"`
}

type WebhookFeatureConfig struct {
	EnableSignatureValidation        bool `yaml:"enableSignatureValidation"`
	EnableTimestampValidation        bool `yaml:"enableTimestampValidation"`
	EnableDuplicateEventIDValidation bool `yaml:"enableDuplicateEventIDValidation"`
}

type Config struct {
	App           App           `yaml:"app"`
	Server        Server        `yaml:"server"`
	Observability Observability `yaml:"observability"`
	Database      Database      `yaml:"database"`
	Webhook       Webhook       `yaml:"webhook"`
	FeatureFlag   FeatureFlag   `yaml:"featureFlag"`
	Http          HttpConfig    `yaml:"http"`
	Asynq         AsynqConfig   `yaml:"asynq"`
	Condition     Condition     `yaml:"condition"`
}

type HttpConfig struct {
	MockApi     HttpClientConfig `yaml:"mockapi"`
	ApimDIDXApi HttpClientConfig `yaml:"apimDIDXApi"`
}

type HttpClientConfig struct {
	BaseUrl    string        `yaml:"baseUrl"`
	APIKey     string        `yaml:"apiKey"`
	Timeout    time.Duration `yaml:"timeout"`
	RetryCount int           `yaml:"retryCount"`
}

type AsynqConfig struct {
	RedisAddr     string `yaml:"redisAddr"`
	RedisDB       int    `yaml:"redisDB"`
	RedisPassword string `yaml:"redisPassword"`
	Queue         string `yaml:"queue"`
	Concurrency   int    `yaml:"concurrency"`
}

type Condition struct {
	OutletIDs []string `yaml:"outletIDs"`
}

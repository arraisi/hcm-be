package config

import "time"

type App struct {
	Name string
	Env  string
}
type Server struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	RequestTimeout time.Duration
}
type Observability struct {
	MetricsEnabled bool
	PprofEnabled   bool
}
type Database struct {
	Driver                string // Supported: memory, postgres, sqlserver
	DSN                   string // Data Source Name - connection string for the database
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	MaxConnectionIdleTime time.Duration
}
type Webhook struct {
	APIKey     string // API key for webhook authentication
	HMACSecret string // HMAC secret for signature verification
}

type FeatureFlag struct {
	WebhookConfig WebhookFeatureConfig `mapstructure:"webhook"`
}

type WebhookFeatureConfig struct {
	EnableSignatureValidation bool `mapstructure:"enableSignatureValidation"`
	EnableTimestampValidation bool `mapstructure:"enableTimestampValidation"`
}

type Config struct {
	App           App
	Server        Server
	Observability Observability
	Database      Database
	Webhook       Webhook
	FeatureFlag   FeatureFlag `mapstructure:"featureFlag"`
}

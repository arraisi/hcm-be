package config

import (
	"time"

	"github.com/spf13/viper"
)

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

type Config struct {
	App           App
	Server        Server
	Observability Observability
	Database      Database
}

func Load() (*Config, error) {
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

	_ = v.ReadInConfig() // ignore if missing

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

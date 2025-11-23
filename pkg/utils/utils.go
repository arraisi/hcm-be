package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/jmoiron/sqlx"
)

func ToPointer[T any](v T) *T {
	return &v
}

func ToValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func GetTimeUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func ToString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

// ParseDateString converts a "2006-01-02" formatted string to time.Time (UTC).
func ParseDateString(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil // return zero time if empty
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}

func DecodeBase64String(s string) ([]byte, error) {
	if s == "" {
		return nil, nil // gracefully handle empty input
	}
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// JoinSCommaSeparatedString joins string values into comma-separated string
func JoinSCommaSeparatedString(values []string) string {
	if len(values) == 0 {
		return ""
	}
	result := values[0]
	for i := 1; i < len(values); i++ {
		result += "," + values[i]
	}
	return result
}

// OpenDatabase opens a database connection with the given configuration.
func OpenDatabase(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, err
	}

	// configure connection pool from config
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.MaxConnectionLifetime)
	db.SetConnMaxIdleTime(cfg.MaxConnectionIdleTime)

	return db, nil
}

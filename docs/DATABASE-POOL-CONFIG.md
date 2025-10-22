# Database Connection Pool Configuration

## Summary of Changes

### 1. Updated app.Config struct

Added database configuration fields to the `app.Config` struct in `internal/app/app.go`:

```go
type Config struct {
	Name           string
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	RequestTimeout time.Duration
	EnableMetrics  bool
	EnablePprof    bool
	Database       DatabaseConfig  // Added this
}

type DatabaseConfig struct {
	Driver                string
	DSN                   string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	MaxConnectionIdleTime time.Duration
}
```

### 2. Updated database connection setup

Replaced hardcoded values with configuration-driven setup:

**Before:**

```go
db, err := sql.Open("sqlserver" /* cfg.Database.DSN */)
if err != nil {
	return err
}
// optional tuning pool
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(0)
```

**After:**

```go
db, err := sql.Open(cfg.Database.Driver, cfg.Database.DSN)
if err != nil {
	return err
}
defer db.Close()

// configure connection pool from config
db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
db.SetConnMaxLifetime(cfg.Database.MaxConnectionLifetime)
db.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)
```

### 3. Updated main.go

Modified `cmd/server/main.go` to pass database configuration:

```go
if err := app.Run(app.Config{
	Name:           cfg.App.Name,
	Host:           cfg.Server.Host,
	Port:           cfg.Server.Port,
	ReadTimeout:    cfg.Server.ReadTimeout,
	WriteTimeout:   cfg.Server.WriteTimeout,
	IdleTimeout:    cfg.Server.IdleTimeout,
	RequestTimeout: cfg.Server.RequestTimeout,
	EnableMetrics:  cfg.Observability.MetricsEnabled,
	EnablePprof:    cfg.Observability.PprofEnabled,
	Database: app.DatabaseConfig{  // Added this
		Driver:                cfg.Database.Driver,
		DSN:                   cfg.Database.DSN,
		MaxOpenConnections:    cfg.Database.MaxOpenConnections,
		MaxIdleConnections:    cfg.Database.MaxIdleConnections,
		MaxConnectionLifetime: cfg.Database.MaxConnectionLifetime,
		MaxConnectionIdleTime: cfg.Database.MaxConnectionIdleTime,
	},
}); err != nil {
```

### 4. Fixed config path

Updated `internal/config/config.go` to look for config file in the correct location:

```go
v.AddConfigPath("./internal/config")
v.AddConfigPath("./configs")
```

## Current Configuration

The application now uses the following database connection pool settings from `config.yaml`:

```yaml
database:
  driver: sqlserver
  dsn: 'sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=hcmdb&encrypt=disable&trustServerCertificate=true'
  maxOpenConnections: 25 # Maximum number of open connections
  maxIdleConnections: 25 # Maximum number of idle connections
  maxConnectionLifetime: 5m # Maximum time a connection may be reused
  maxConnectionIdleTime: 5m # Maximum time a connection may be idle
```

## Benefits

1. **Configurable**: Connection pool settings can be adjusted without code changes
2. **Environment-specific**: Different values for dev/staging/production
3. **Performance tuning**: Easy to optimize based on load requirements
4. **Monitoring**: Settings are explicit and visible in configuration
5. **Consistent**: All configuration centralized in one place

## Verification

✅ Application builds successfully
✅ Server starts without errors
✅ Database connection established
✅ API endpoints working (`GET /api/v1/users`)
✅ Health checks passing (`GET /healthz`)
✅ Configuration properly loaded from `config.yaml`

The database connection pool is now fully configurable and working correctly with SQL Server!

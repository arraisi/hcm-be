# Docker Deployment Guide

## Build Instructions

### Build the Docker image:
```bash
docker build -t hcm-be:latest .
```

### Build with custom tag:
```bash
docker build -t hcm-be:v1.0.0 .
```

## Run Instructions

### Basic run (requires external database):
```bash
docker run -d \
  --name hcm-backend \
  -p 8080:8080 \
  hcm-be:latest
```

### Run with environment variables for database connection:
```bash
docker run -d \
  --name hcm-backend \
  -p 8080:8080 \
  -e DATABASE_DSN="sqlserver://user:pass@dbhost:1433?database=hcm&encrypt=disable&trustServerCertificate=true" \
  hcm-be:latest
```

### Run with custom config (mount external config file):
```bash
docker run -d \
  --name hcm-backend \
  -p 8080:8080 \
  -v $(pwd)/custom-config.yaml:/app/internal/config/config.yaml:ro \
  hcm-be:latest
```

### Run with networking to existing database container:
```bash
docker run -d \
  --name hcm-backend \
  -p 8080:8080 \
  --network your-db-network \
  -e DATABASE_DSN="sqlserver://user:pass@sqlserver-container:1433?database=hcm&encrypt=disable&trustServerCertificate=true" \
  hcm-be:latest
```

## Health Check

The container includes a health check that calls the `/healthz` endpoint:

```bash
# Check container health status
docker ps

# View health check logs
docker inspect --format='{{json .State.Health}}' hcm-backend
```

## Image Information

- **Base image**: alpine:3.20
- **Working directory**: /app
- **User**: appuser (non-root, uid: 1001)
- **Exposed port**: 8080
- **Config path**: ./internal/config/config.yaml (relative to working directory)
- **Binary**: /app/hcm-be

## Production Considerations

1. **Database Connection**: Ensure the database is accessible from the container
2. **Secrets Management**: Use Docker secrets or external secret management for sensitive data
3. **Logging**: Consider mounting log directories or using centralized logging
4. **Monitoring**: Expose metrics endpoint (/metrics) for monitoring systems
5. **Resource Limits**: Set appropriate CPU and memory limits in production

### Example with resource limits:
```bash
docker run -d \
  --name hcm-backend \
  -p 8080:8080 \
  --memory=512m \
  --cpus=1.0 \
  --restart=unless-stopped \
  hcm-be:latest
```
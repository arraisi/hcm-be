# Deployment Guide for Dokploy

This document provides instructions for deploying the HCM application using Dokploy.

## Overview

The application consists of:
- Backend API service (hcm-be) running on port 8080
- Connects to an existing external SQL Server database
- Docker Compose setup with only the application service

## Prerequisites

- Docker and Docker Compose installed on the target system
- Dokploy instance configured and accessible
- Domain name or subdomain ready for the application
- External SQL Server database accessible from the deployment environment

## Deployment Steps for Dokploy

### 1. Prepare your Dokploy server

1. Ensure your Dokploy server has access to a Docker registry (Docker Hub, GitHub Container Registry, or local registry)
2. Add the domain name in Dokploy settings and ensure it points to your server IP
3. Ensure the deployment server can connect to your external SQL Server database

### 2. Create a new application in Dokploy

1. Go to your Dokploy dashboard
2. Click "Add Application"
3. Select "Docker Compose" type
4. For the source, you can either:
   - Use "Git Repository" if you have this code in a Git repo
   - Use "Raw Compose" and paste the content of `docker-compose.yml`

### 3. Configure the Docker Compose

When using the Raw Compose option, copy the content of your `docker-compose.yml` file:

```yaml
services:
  # Application Service
  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: hcm-app
    environment:
      - APP_DATABASE_DRIVER=sqlserver
      - APP_DATABASE_DSN=${APP_DATABASE_DSN}
      - APP_SERVER_HOST=0.0.0.0
      - APP_SERVER_PORT=8080
      - APP_ENV=production
    ports:
      - "8080:8080"
    networks:
      - hcm-network
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/healthz"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

networks:
  hcm-network:
    driver: bridge
```

### 4. Configure environment variables

In Dokploy, configure the following environment variables:
- `APP_DATABASE_DRIVER` = `sqlserver`
- `APP_DATABASE_DSN` = Your external SQL Server connection string
- `APP_SERVER_HOST` = `0.0.0.0`
- `APP_SERVER_PORT` = `8080`
- `APP_ENV` = `production`

### 5. Configure the domain and SSL

1. In the application settings, go to the "Domains" tab
2. Add your domain (e.g., `hcm.yourdomain.com`)
3. Enable SSL certificate generation (Let's Encrypt) for HTTPS

### 6. Set up the build context (for Git repository source)

If using Git repository source:
1. Point to your repository containing the Dockerfile and docker-compose.yml
2. Set the build context to the root directory of your repository
3. Dokploy will automatically build the application using the Dockerfile

### 7. Deploy and review

1. Click "Deploy" to start the deployment process
2. Monitor the logs to ensure the application starts successfully
3. Wait for the health check to pass
4. Access your application via the configured domain

## Security Considerations

- Use strong credentials for the external database connection
- Consider using Dokploy's secrets management for sensitive database credentials
- Ensure firewall rules allow the application to connect to your external SQL Server

## Ports Configuration

- Application port: 8080 (exposed via domain with SSL)
- Database connection: External to the container (to your existing SQL Server)

## Health Checks

- Application health check: `/healthz` endpoint
- The application will attempt to connect to the external database at startup

## Scaling Considerations

- This setup is configured for single-instance deployment
- Database scaling should be handled separately on your existing SQL Server
- Consider database connection pooling settings for high-traffic scenarios
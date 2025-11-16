# Build stage
FROM golang:1.24-alpine AS builder

# Install git and ca-certificates (needed for module downloads)
RUN apk add --no-cache git ca-certificates

# Set Go toolchain to use latest available
ENV GOTOOLCHAIN=auto

# Set working directory
WORKDIR /src

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies (cached layer if go.mod/go.sum unchanged)
RUN go mod download

# Copy source code
COPY . .

# Build both server and worker binaries
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w" \
    -o hcm-be ./cmd/server

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w" \
    -o hcm-worker ./cmd/worker

# ============================================
# Server runtime stage
# ============================================
FROM alpine:3.20 AS server

# Install ca-certificates for HTTPS calls and create non-root user
RUN apk --no-cache add ca-certificates \
    && addgroup -g 1001 -S appgroup \
    && adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the server binary from builder stage
COPY --from=builder /src/hcm-be .

# Copy config file to the expected relative path
COPY --from=builder /src/internal/config/config.yaml ./internal/config/config.yaml

# Set environment variable for config path (optional, for clarity)
ENV APP_CONFIG=./internal/config/config.yaml

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check using the application's health endpoint
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/healthz || exit 1

# Run the server application
CMD ["./hcm-be"]

# ============================================
# Worker runtime stage
# ============================================
FROM alpine:3.20 AS worker

# Install ca-certificates for HTTPS calls and create non-root user
RUN apk --no-cache add ca-certificates \
    && addgroup -g 1001 -S appgroup \
    && adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the worker binary from builder stage
COPY --from=builder /src/hcm-worker .

# Copy config file to the expected relative path
COPY --from=builder /src/internal/config/config.yaml ./internal/config/config.yaml

# Set environment variable for config path
ENV APP_CONFIG=./internal/config/config.yaml

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Run the worker application
CMD ["./hcm-worker"]
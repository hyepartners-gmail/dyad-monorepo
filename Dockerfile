# ---- Build stage ----
FROM golang:1.24-alpine AS builder

# Install git (required for some Go modules)
RUN apk update && apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Copy the rest of the application
COPY . .

# Build binary with environment variable
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# ---- Runtime stage ----
FROM alpine:latest

WORKDIR /app

# Copy binary from build stage
COPY --from=builder /app/app .

# Environment variable for runtime config
ENV GO_ENV=production

# Expose port for app
EXPOSE 8080

# Healthcheck endpoint
HEALTHCHECK --interval=30s --timeout=3s CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./app"]
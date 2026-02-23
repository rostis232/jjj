# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o /domain-app ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary and migrations
COPY --from=builder /domain-app ./
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Run the application
CMD ["./domain-app"]

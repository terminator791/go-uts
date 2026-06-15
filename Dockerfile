# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the applications
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/migrate ./cmd/migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/seed ./cmd/seed

# Run stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates in case they are needed for external APIs
RUN apk --no-cache add ca-certificates

# Copy the binaries from builder
COPY --from=builder /app/bin/server .
COPY --from=builder /app/bin/migrate .
COPY --from=builder /app/bin/seed .

# Copy the migrations directory so that migration binary can load .sql files
COPY --from=builder /app/migrations ./migrations

# Expose port (default is :8080 in config)
EXPOSE 8080

# Run the server by default
CMD ["./server"]

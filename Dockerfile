# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -ldflags="-w -s" -o server ./api/main.go

# Create data directory
RUN mkdir -p data

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Copy binary and configs from builder
COPY --from=builder /app/server .
COPY --from=builder /app/api/etc ./api/etc
COPY --from=builder /app/data ./data

# Set timezone
ENV TZ=Asia/Shanghai

EXPOSE 8080

CMD ["./server", "-f", "api/etc/config.yaml"]


.PHONY: all build run test clean docker docker-compose help

# 变量定义
APP_NAME := acupofcoffee
API_DIR := api
BUILD_DIR := build
DOCKER_IMAGE := $(APP_NAME)-api

# Go 相关
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOFMT := gofmt
GOLINT := golangci-lint

# 默认目标
all: build

# 安装依赖
deps:
	@echo "📦 Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# 格式化代码
fmt:
	@echo "🎨 Formatting code..."
	$(GOFMT) -s -w .

# 代码检查
lint:
	@echo "🔍 Running linter..."
	$(GOLINT) run ./...

# 运行测试
test:
	@echo "🧪 Running tests..."
	$(GOTEST) -v -cover ./...

# 构建 API 服务
build:
	@echo "🔨 Building API server..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GOBUILD) -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./$(API_DIR)/main.go

# 开发模式运行
run:
	@echo "🚀 Starting API server in development mode..."
	$(GOCMD) run ./$(API_DIR)/main.go -f ./$(API_DIR)/etc/config.yaml

# 清理构建产物
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -rf ./logs

# Docker 构建
docker:
	@echo "🐳 Building Docker image..."
	docker build -t $(DOCKER_IMAGE):latest -f deploy/docker/Dockerfile .

# Docker Compose 启动
docker-up:
	@echo "🐳 Starting services with Docker Compose..."
	cd deploy/docker && docker-compose up -d

# Docker Compose 停止
docker-down:
	@echo "🐳 Stopping services..."
	cd deploy/docker && docker-compose down

# Docker Compose 日志
docker-logs:
	@echo "📋 Showing logs..."
	cd deploy/docker && docker-compose logs -f

# 数据库迁移（开发用）
migrate:
	@echo "📊 Running database migrations..."
	$(GOCMD) run ./$(API_DIR)/main.go migrate

# 生成 API 文档
docs:
	@echo "📚 Generating API documentation..."
	@command -v swag >/dev/null 2>&1 || { echo "Installing swag..."; go install github.com/swaggo/swag/cmd/swag@latest; }
	swag init -g ./$(API_DIR)/main.go -o ./docs

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make deps        - Install dependencies"
	@echo "  make fmt         - Format code"
	@echo "  make lint        - Run linter"
	@echo "  make test        - Run tests"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run in development mode"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make docker      - Build Docker image"
	@echo "  make docker-up   - Start with Docker Compose"
	@echo "  make docker-down - Stop Docker Compose services"
	@echo "  make docker-logs - Show Docker Compose logs"
	@echo "  make docs        - Generate API documentation"
	@echo "  make help        - Show this help"


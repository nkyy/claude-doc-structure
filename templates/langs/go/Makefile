# Go Project Makefile

# Application
APP_NAME := {name}
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build settings
BINARY_DIR := bin
BINARY_NAME := $(APP_NAME)
MAIN_PATH := ./cmd/$(APP_NAME)
BUILD_FLAGS := -ldflags="-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -w -s"

# Go settings
GO := go
GO_VERSION := $(shell $(GO) version | cut -d' ' -f3)
GOPATH := $(shell $(GO) env GOPATH)
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)

# Tools
GOLANGCI_LINT := golangci-lint
GOTESTSUM := gotestsum

# Colors for output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
PURPLE := \033[0;35m
CYAN := \033[0;36m
WHITE := \033[0;37m
NC := \033[0m # No Color

.PHONY: help
help: ## Show this help message
	@echo '$(CYAN)Available commands:$(NC)'
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make $(CYAN)<target>$(NC)\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  $(CYAN)%-15s$(NC) %s\n", $$1, $$2 } /^##@/ { printf "\n$(PURPLE)%s$(NC)\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
.PHONY: dev
dev: ## Run the application in development mode
	@echo "$(GREEN)Running $(APP_NAME) in development mode...$(NC)"
	$(GO) run $(MAIN_PATH) server

.PHONY: build
build: clean ## Build the application binary
	@echo "$(GREEN)Building $(APP_NAME) for $(GOOS)/$(GOARCH)...$(NC)"
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build complete: $(BINARY_DIR)/$(BINARY_NAME)$(NC)"

.PHONY: build-all
build-all: clean ## Build for all supported platforms
	@echo "$(GREEN)Building for all platforms...$(NC)"
	@mkdir -p $(BINARY_DIR)
	GOOS=linux GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=arm64 $(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 $(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "$(GREEN)Cross-platform build complete$(NC)"

.PHONY: install
install: ## Install the binary to GOPATH/bin
	@echo "$(GREEN)Installing $(APP_NAME)...$(NC)"
	$(GO) install $(BUILD_FLAGS) $(MAIN_PATH)

.PHONY: run
run: build ## Build and run the application
	@echo "$(GREEN)Running $(APP_NAME)...$(NC)"
	./$(BINARY_DIR)/$(BINARY_NAME)

##@ Testing
.PHONY: test
test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	$(GO) test -v -race ./...

.PHONY: test-cover
test-cover: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GO) test -v -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

.PHONY: test-short
test-short: ## Run short tests
	@echo "$(GREEN)Running short tests...$(NC)"
	$(GO) test -short -v ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	@echo "$(GREEN)Running benchmarks...$(NC)"
	$(GO) test -bench=. -benchmem ./...

##@ Code Quality
.PHONY: fmt
fmt: ## Format code
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GO) fmt ./...

.PHONY: vet
vet: ## Run go vet
	@echo "$(GREEN)Running go vet...$(NC)"
	$(GO) vet ./...

.PHONY: lint
lint: ## Run golangci-lint
	@echo "$(GREEN)Running golangci-lint...$(NC)"
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: ## Run golangci-lint with auto-fix
	@echo "$(GREEN)Running golangci-lint with auto-fix...$(NC)"
	$(GOLANGCI_LINT) run --fix

.PHONY: check
check: fmt vet lint test ## Run all checks (format, vet, lint, test)

##@ Dependencies
.PHONY: deps
deps: ## Download dependencies
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	$(GO) mod download

.PHONY: deps-update
deps-update: ## Update dependencies
	@echo "$(GREEN)Updating dependencies...$(NC)"
	$(GO) get -u ./...
	$(GO) mod tidy

.PHONY: deps-verify
deps-verify: ## Verify dependencies
	@echo "$(GREEN)Verifying dependencies...$(NC)"
	$(GO) mod verify

.PHONY: tidy
tidy: ## Tidy up dependencies
	@echo "$(GREEN)Tidying dependencies...$(NC)"
	$(GO) mod tidy

##@ Database
.PHONY: migrate-up
migrate-up: ## Run database migrations up
	@echo "$(GREEN)Running database migrations up...$(NC)"
	./$(BINARY_DIR)/$(BINARY_NAME) migrate up

.PHONY: migrate-down
migrate-down: ## Run database migrations down
	@echo "$(GREEN)Running database migrations down...$(NC)"
	./$(BINARY_DIR)/$(BINARY_NAME) migrate down

.PHONY: migrate-create
migrate-create: ## Create a new migration (usage: make migrate-create name=create_users)
	@echo "$(GREEN)Creating migration: $(name)...$(NC)"
	./$(BINARY_DIR)/$(BINARY_NAME) migrate create $(name)

##@ Docker
.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "$(GREEN)Building Docker image...$(NC)"
	docker build -t $(APP_NAME):$(VERSION) -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run: ## Run Docker container
	@echo "$(GREEN)Running Docker container...$(NC)"
	docker run --rm -p 8080:8080 --env-file .env $(APP_NAME):latest

.PHONY: docker-push
docker-push: docker-build ## Push Docker image
	@echo "$(GREEN)Pushing Docker image...$(NC)"
	docker push $(APP_NAME):$(VERSION)
	docker push $(APP_NAME):latest

##@ Utilities
.PHONY: clean
clean: ## Clean build artifacts
	@echo "$(GREEN)Cleaning build artifacts...$(NC)"
	rm -rf $(BINARY_DIR)
	rm -f coverage.out coverage.html

.PHONY: version
version: ## Show version information
	@echo "$(CYAN)Application:$(NC) $(APP_NAME)"
	@echo "$(CYAN)Version:$(NC) $(VERSION)"
	@echo "$(CYAN)Go Version:$(NC) $(GO_VERSION)"
	@echo "$(CYAN)Git Commit:$(NC) $(GIT_COMMIT)"
	@echo "$(CYAN)Build Time:$(NC) $(BUILD_TIME)"
	@echo "$(CYAN)OS/Arch:$(NC) $(GOOS)/$(GOARCH)"

.PHONY: tools
tools: ## Install development tools
	@echo "$(GREEN)Installing development tools...$(NC)"
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install gotest.tools/gotestsum@latest

.PHONY: generate
generate: ## Run go generate
	@echo "$(GREEN)Running go generate...$(NC)"
	$(GO) generate ./...

# Default target
.DEFAULT_GOAL := help
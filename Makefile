# Makefile for Claude Doc Structure (Go version)
# Provides convenient shortcuts for common operations

.PHONY: help build install clean test run dev release

# Default target
help:
	@echo "Claude Doc Structure (Go) - Available Commands:"
	@echo ""
	@echo "Build & Installation:"
	@echo "  make build       Build the binary"
	@echo "  make install     Install the binary to GOPATH/bin"
	@echo "  make clean       Clean up build artifacts"
	@echo ""
	@echo "Development:"
	@echo "  make dev         Run in development mode"
	@echo "  make test        Run tests"
	@echo "  make run CMD=... Run with arguments (e.g., make run CMD='init myproject')"
	@echo ""
	@echo "Release:"
	@echo "  make release     Build for multiple platforms"
	@echo ""
	@echo "Documentation Operations:"
	@echo "  make init        Initialize documentation structure in current directory"
	@echo "  make validate    Validate documentation structure"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make run CMD='init myproject'"
	@echo "  make run CMD='split README.md --by-headers'"
	@echo ""

# Build the binary
build:
	@echo "Building claude-docs..."
	go build -o bin/claude-docs .
	@echo "Build complete! Binary: bin/claude-docs"

# Install to GOPATH/bin
install: build
	@echo "Installing claude-docs..."
	go install .
	@echo "Installation complete! You can now use 'claude-docs' command."

# Clean up
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf bin/
	go clean
	@echo "Clean complete."

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Development mode (with hot reload if available)
dev:
	@echo "Running in development mode..."
	go run . --help

# Run with arguments
run:
ifndef CMD
	@echo "Error: CMD parameter is required"
	@echo "Usage: make run CMD='command arguments'"
	@echo "Example: make run CMD='init myproject'"
	@exit 1
endif
	@echo "Running: claude-docs $(CMD)"
	go run . $(CMD)

# Initialize documentation structure
init:
	@echo "Initializing Claude documentation structure..."
	go run . init

# Validate structure
validate:
	@echo "Validating documentation structure..."
	go run . validate

# Build for multiple platforms
release:
	@echo "Building release binaries..."
	@mkdir -p bin
	
	# Linux amd64
	GOOS=linux GOARCH=amd64 go build -o bin/claude-docs-linux-amd64 .
	
	# Linux arm64
	GOOS=linux GOARCH=arm64 go build -o bin/claude-docs-linux-arm64 .
	
	# macOS amd64 (Intel)
	GOOS=darwin GOARCH=amd64 go build -o bin/claude-docs-darwin-amd64 .
	
	# macOS arm64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build -o bin/claude-docs-darwin-arm64 .
	
	# Windows amd64
	GOOS=windows GOARCH=amd64 go build -o bin/claude-docs-windows-amd64.exe .
	
	@echo "Release binaries built in bin/ directory:"
	@ls -la bin/

# Quick setup
quick-setup: build init
	@echo "Quick setup complete!"
	@echo "Binary built and documentation structure initialized."
	@echo "Next steps:"
	@echo "1. Edit CLAUDE.md with your project details"
	@echo "2. Run './bin/claude-docs validate' to check your structure"
	@echo "3. Use './bin/claude-docs template api' to generate templates"

# Development dependencies
deps:
	@echo "Installing development dependencies..."
	go mod tidy
	go mod download

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Generate documentation
docs:
	@echo "Generating documentation..."
	@./bin/claude-docs --help > docs/cli-help.txt
	@echo "Documentation generated in docs/"
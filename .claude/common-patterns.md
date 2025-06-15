# Common Command Patterns

## Development Workflow

### Quick Development Cycle
```bash
# Build and test in one command
make build && ./bin/claude-docs --help

# Run with arguments during development
make run CMD="init test-project"
make run CMD="split README.md --by-headers"
make run CMD="validate"
```

### Cross-Platform Building
```bash
# Build for current platform
make build

# Build for all platforms (release)
make release

# Clean up build artifacts
make clean
```

## CLI Usage Patterns

### Project Initialization
```bash
# Initialize documentation structure
claude-docs init

# Initialize with specific project name
claude-docs init my-project

# Validate existing structure
claude-docs validate
```

### Document Processing
```bash
# Split large documents
claude-docs split README.md --by-headers
claude-docs split large-doc.md --by-lines 100
claude-docs split huge-file.md --by-size 50KB

# Merge multiple documents
claude-docs merge docs/ --output combined.md
claude-docs merge specs/ --output requirements.md
```

### Template Generation
```bash
# Generate API documentation template
claude-docs template api

# Generate screen specification template
claude-docs template screens

# List available templates
claude-docs template --list
```

## Git Workflow Patterns

### Release Process
```bash
# Create and push release tag
git tag v0.2.0
git push origin v0.2.0

# This automatically triggers:
# 1. GitHub Actions build
# 2. Multi-platform binary creation
# 3. GitHub Release
# 4. Homebrew formula update
```

### Development Workflow
```bash
# Feature development
git checkout -b feature/new-command
# ... make changes ...
git add .
git commit -m "Add new command"
git push origin feature/new-command

# Create PR and merge
# Release when ready with tag push
```

## Testing Patterns

### Manual Testing
```bash
# Test build process
make clean && make build

# Test CLI commands
./bin/claude-docs --help
./bin/claude-docs init test-dir
./bin/claude-docs validate

# Test cross-platform builds
make release
ls -la bin/
```

### Automated Testing (Future)
```bash
# Run test suite
make test

# Run with coverage
go test -cover ./...

# Lint code
make lint
```

## Homebrew Management

### Installation Testing
```bash
# Install from tap
brew install nkyy/claude-doc/claude-doc

# Test installation
claude-docs --help

# Uninstall for testing
brew uninstall claude-doc

# Reinstall specific version
brew reinstall claude-doc
```

### Formula Updates
```bash
# Manual formula update (if needed)
brew tap nkyy/claude-doc
cd $(brew --repository nkyy/claude-doc)
# Edit Formula/claude-doc.rb
git add Formula/claude-doc.rb
git commit -m "Update to vX.Y.Z"
git push
```

## Documentation Maintenance

### Template Updates
```bash
# Update templates
edit templates/specs/api.md
edit templates/specs/screens.md

# Test template generation
claude-docs template api
claude-docs template screens
```

### Example Projects
```bash
# Test example projects
cd templates/examples/small-project
claude-docs validate

cd ../large-project
claude-docs validate
```

## Debugging Patterns

### Build Issues
```bash
# Check Go environment
go version
go env

# Clean build
make clean
go mod tidy
make build

# Verbose build
go build -v -o bin/claude-docs .
```

### Runtime Issues
```bash
# Check binary info
file ./bin/claude-docs
ldd ./bin/claude-docs  # Linux
otool -L ./bin/claude-docs  # macOS

# Debug specific commands
./bin/claude-docs init --verbose
./bin/claude-docs split file.md --debug
```

## Performance Monitoring

### Binary Size Optimization
```bash
# Check binary size
ls -lh bin/claude-docs*

# Build with optimizations
go build -ldflags="-s -w" -o bin/claude-docs .

# Strip debug info (if needed)
strip bin/claude-docs
```

### Execution Time
```bash
# Time command execution
time ./bin/claude-docs init
time ./bin/claude-docs split large-file.md

# Profile memory usage (development)
go build -o bin/claude-docs-debug .
# Use with profiling tools
```

## IDE Integration Patterns

### VS Code Development
```bash
# Quick build and test
Cmd+Shift+P -> "Tasks: Run Task" -> "build"

# Debug configuration in launch.json
{
  "type": "go",
  "request": "launch",
  "mode": "debug",
  "program": "${workspaceFolder}",
  "args": ["init", "test-project"]
}
```

### Command Line Shortcuts
```bash
# Add to shell profile (.bashrc, .zshrc)
alias cds='claude-docs'
alias cds-build='cd ~/path/to/claude-doc-structure && make build'
alias cds-test='cd ~/path/to/claude-doc-structure && make run CMD="validate"'
```
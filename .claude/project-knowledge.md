# Project Technical Knowledge

## Architecture Insights

### CLI Design Patterns
- **Cobra Framework**: Used for subcommand structure (`cmd/` directory)
- **Single Binary Philosophy**: All functionality compiled into one executable
- **Internal Package Structure**: Clean separation of concerns in `internal/`

### Key Technical Decisions

#### Go Migration Rationale
- **Problem**: Python dependency management complexity
- **Solution**: Go single binary compilation
- **Benefits**: Zero runtime dependencies, faster startup, easier distribution

#### Build System Design
```go
// Cross-platform build targets in Makefile
GOOS=darwin GOARCH=amd64  // macOS Intel
GOOS=darwin GOARCH=arm64  // macOS Apple Silicon  
GOOS=linux GOARCH=amd64   // Linux x64
GOOS=windows GOARCH=amd64 // Windows x64
```

#### Document Processing Architecture
- **Splitter**: `internal/splitter/` - Multi-method document splitting
- **Merger**: `internal/merger/` - Smart document merging with TOC
- **Template System**: `templates/` - User-facing documentation templates

### Performance Optimizations
1. **Single Binary Distribution**: No installation dependencies
2. **Fast Startup**: Go compiled binary vs interpreted Python
3. **Memory Efficient**: No runtime interpretation overhead

### Code Patterns & Conventions

#### CLI Command Structure
```go
// cmd/root.go - Base command setup
var rootCmd = &cobra.Command{
    Use:   "claude-docs",
    Short: "Claude Documentation Tool",
}

// Subcommands in separate files
// cmd/init.go, cmd/split.go, cmd/merge.go, etc.
```

#### Error Handling Pattern
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

#### File Processing Pattern
- Always check file existence before processing
- Use `filepath.Walk` for directory traversal
- Implement proper cleanup with `defer`

### Testing Strategy
- **Unit Tests**: `go test ./...`
- **Integration Tests**: CLI command testing
- **Manual Testing**: Cross-platform binary testing

### Deployment & Distribution

#### GitHub Actions Workflow
1. **Trigger**: Git tag push (`v*`)
2. **Build**: Multi-platform binaries
3. **Release**: GitHub Releases with assets
4. **Homebrew**: Automatic formula update

#### Homebrew Integration
```ruby
# Formula structure
class ClaudeDoc < Formula
  desc "Documentation toolkit optimized for Claude Code"
  url "https://github.com/.../releases/download/..."
  sha256 "..."
  
  def install
    bin.install "claude-docs-darwin-amd64" => "claude-docs"
  end
end
```

## Critical Implementation Details

### Template System
- Uses `{name}` placeholder format
- Templates stored in `templates/` directory
- Support for nested template structures

### Document Splitting Logic
- **By Headers**: Markdown header-based splitting
- **By Lines**: Fixed line count chunks
- **By Size**: File size-based splitting

### Makefile Design Philosophy
- **Convenience**: Simple `make build`, `make test` commands
- **Development**: `make run CMD="..."` for testing
- **Cross-Platform**: `make release` for all targets

## Common Gotchas & Solutions

### Build Issues
- **Problem**: `GOOS`/`GOARCH` environment variables persist
- **Solution**: Explicit setting in Makefile targets

### Path Handling
- **Problem**: Cross-platform path separators
- **Solution**: Use `filepath.Join()` consistently

### Binary Naming
- **Problem**: Windows `.exe` extension requirement
- **Solution**: Conditional naming in build scripts
# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Doc Structure is a documentation toolkit that provides templates, tools, and best practices for organizing project context to maximize Claude Code's effectiveness. It transforms documentation chaos into Claude-optimized clarity through systematic organization and automation.

## Architecture & Technology Stack

**Core Technologies:**
- Go 1.21+ (single binary, zero dependencies at runtime)
- Bash scripting for automation
- Markdown for all documentation templates

**Key Components:**
- CLI tool (`main.go` + `cmd/`) - Unified command interface in Go
- Document processing (`internal/splitter`, `internal/merger`) - Split/merge utilities
- Template system (`templates/`) - User-facing documentation templates
- Build system (`Makefile`) - Cross-platform binary generation

## Project Structure

```
claude-doc-structure/
├── main.go                    # CLI entry point
├── go.mod                     # Go module definition
├── cmd/                       # CLI commands (cobra-based)
│   ├── root.go               # Root command setup
│   ├── init.go               # Project initialization
│   ├── split.go              # Document splitting
│   ├── merge.go              # Document merging
│   ├── template.go           # Template generation
│   └── validate.go           # Structure validation
├── internal/                  # Internal packages
│   ├── splitter/             # Document splitting logic
│   └── merger/               # Document merging logic
├── templates/                 # User-facing templates
│   ├── .claude/              # Claude-specific assets
│   ├── specs/                # Documentation templates
│   └── examples/             # Reference implementations
└── Makefile                  # Build automation
```

## Common Development Commands

**Build & Installation:**
```bash
make build                     # Build single binary
make install                   # Install to GOPATH/bin
make release                   # Build for all platforms
```

**Development:**
```bash
make dev                       # Run in development mode
make run CMD="init myproject"  # Run with arguments
make test                      # Run test suite
make fmt                       # Format code
make lint                      # Code quality checks
```

**CLI Usage:**
```bash
./bin/claude-docs --help      # Show help
./bin/claude-docs init         # Initialize project
./bin/claude-docs validate     # Validate structure
./bin/claude-docs split README.md --by-headers
./bin/claude-docs merge docs/ --output combined.md
```

## Key Implementation Details

**CLI Architecture:**
- `cmd/root.go` - Cobra-based CLI with subcommands
- `cmd/init.go` - Project initialization with template generation
- `cmd/split.go` - Document splitting with multiple methods
- `cmd/merge.go` - Document merging with Claude optimization
- `cmd/template.go` - Template generation system
- `cmd/validate.go` - Structure validation logic

**Document Processing:**
- `internal/splitter/splitter.go` - Multi-method document splitting
- `internal/merger/merger.go` - Smart document merging with TOC
- Supports splitting by headers, lines, or file size
- Merging includes Claude-specific optimizations

**Build System:**
- Single binary compilation with zero runtime dependencies
- Cross-platform builds for Linux, macOS, Windows (x64/ARM64)
- Template placeholders use `{name}` format

## Current Development Focus

**Go Migration (Completed):**
The project has been migrated from Python to Go for better environment independence:
- ✅ Zero runtime dependencies - single binary deployment
- ✅ Cross-platform builds for all major platforms  
- ✅ Improved performance and faster startup times
- ✅ Maintained all original functionality and CLI interface

**Next Steps:**
- Add comprehensive test suite
- Implement GitHub Actions for automated releases
- Add configuration file support (.claude-docs.yaml)
- Enhanced template customization options

**Documentation Maintenance:**
- **IMPORTANT**: When making improvements or changes to the project, always update the README.md to reflect:
  - New features and capabilities
  - Changed installation instructions
  - Updated command examples
  - Modified project structure
- Keep README.md synchronized with actual functionality
- Update badges and version information as needed
- Ensure all code examples in README.md are tested and working

## Testing Strategy

**Manual Testing:**
1. Build and test: `make build && ./bin/claude-docs --help`
2. Test commands: `make run CMD="init test-project"`
3. Test document processing: Use sample markdown files
4. Test cross-platform builds: `make release`

**Automated Testing:**
- `go test ./...` - Run all tests
- `make test` - Convenient test execution
- `make lint` - Code quality checks with golangci-lint

## Dependencies

**Runtime:** None (single binary, zero dependencies)
**Development:** 
- Go 1.21+ for building
- golangci-lint for code quality (optional)
- Make for build automation

This tool achieves true environment independence with single binary distribution.
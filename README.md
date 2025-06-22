# Claude Doc Structure 📋

A comprehensive documentation toolkit for efficient collaboration with Claude Code. Streamline your development workflow with organized project context, automated tools, and reusable templates.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Go](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![CLI Tool](https://img.shields.io/badge/CLI-Tool-brightgreen.svg)](https://github.com/your-username/claude-doc-structure)
[![Zero Dependencies](https://img.shields.io/badge/Dependencies-Zero-green.svg)]()

## 🎯 Purpose

This template provides a structured approach to organizing project documentation that maximizes Claude Code's effectiveness. By following these patterns, you'll:

- Enable Claude Code to understand your project context faster
- Reduce repetitive explanations through organized documentation
- Maintain consistency across development sessions
- Scale documentation efficiently as your project grows
- Support multiple programming languages with tailored configurations
- Leverage adaptive-trading project insights for real-world optimization

## 🚀 Quick Start

### Option 1: Binary Download (Recommended)

1. **Download the binary** for your platform from [Releases](https://github.com/your-username/claude-doc-structure/releases)

2. **Make it executable and add to PATH:**
   ```bash
   chmod +x claude-docs
   sudo mv claude-docs /usr/local/bin/
   ```

3. **Start using immediately:**
   ```bash
   claude-docs init my-project
   claude-docs template api users
   claude-docs validate
   ```

### Option 2: Build from Source

1. **Clone and build:**
   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cd claude-doc-structure
   make build
   ```

2. **Use the binary:**
   ```bash
   ./bin/claude-docs init my-project
   ```

### Option 3: Manual Setup

1. **Copy the structure** to your project:
   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cp -r claude-doc-structure/templates/* ./
   ```

2. **Customize CLAUDE.md** with your project details and start collaborating!

## 📁 Structure Overview

```
your-project/
├── CLAUDE.md              # Main project context for Claude Code
├── .claude/               # Claude Code optimization files
│   ├── context.md         # Project background and constraints
│   ├── project-knowledge.md # Technical insights and patterns
│   ├── project-improvements.md # Development history
│   ├── common-patterns.md # Command patterns and workflows
│   ├── debug-log.md       # Issue resolution history
│   ├── project-specifics.md # Project-specific information
│   ├── ai-collaboration.md # AI collaboration best practices
│   └── code-patterns.md   # Common code patterns and solutions
└── specs/                 # Detailed specifications
    ├── api.md            # API documentation (basic)
    ├── api-spec.md       # Comprehensive API specification
    ├── screens.md        # UI/UX specifications (basic)
    └── ui-spec.md        # Detailed UI specification

claude-doc-structure/ (this repo)
├── main.go               # CLI entry point
├── go.mod               # Go module definition
├── cmd/                 # CLI commands
├── internal/            # Document processing logic
├── bin/claude-docs      # Built binary (after make build)
└── Makefile            # Build automation
```

## 📖 Documentation Approach

### Claude-Optimized Structure

This project implements **systematic knowledge management** for AI-assisted development, inspired by [effective Claude knowledge management practices](https://zenn.dev/driller/articles/2a23ef94f1d603).

**Core Philosophy:**
- **Living Documentation**: Documents that evolve with your project
- **Context Separation**: Organized information for quick AI access
- **Pattern Recognition**: Reusable knowledge for consistent development

**Structure Benefits:**
- Improved Claude Code understanding of project context
- Faster onboarding for new team members
- Systematic capture of development learnings
- Reduced repetitive explanations across sessions

### Documentation Layers

**Layer 1: Quick Context (`CLAUDE.md`)**
- Main project overview and current status
- References to detailed context files
- Key files and immediate project state

**Layer 2: Detailed Context (`.claude/` directory)**
- `context.md`: Project background, constraints, business context
- `project-knowledge.md`: Technical architecture, patterns, decisions
- `project-improvements.md`: Development history, lessons learned
- `common-patterns.md`: Frequently used commands and workflows
- `debug-log.md`: Issue resolution history and troubleshooting

**Layer 3: Specifications (`specs/` directory)**
- `api.md` / `api-spec.md`: API documentation and endpoint specifications
- `screens.md` / `ui-spec.md`: UI/UX specifications and screen descriptions  
- Feature specifications and requirements

**Language-Specific Templates (`templates/langs/`)**
- Python: `requirements.txt`, `.env.example`, `pyproject.toml`
- Node.js: `package.json`, `.env.example`, TypeScript configs
- Go: `go.mod`, `Makefile`, `.env.example`

### Best Practices

**For Claude Code Integration:**
- Use clear, descriptive headers
- Include file paths with line numbers: `src/api/users.js:42`
- Maintain consistent terminology throughout docs
- Keep context under 200KB per file for optimal performance

**Documentation Writing:**
- Lead with purpose and context
- Use examples over abstract descriptions
- Include "what changed" and "why" in updates
- Reference actual code locations

## 🛠️ Tools & Commands

### Unified CLI Tool

The `claude-docs` command provides all functionality in one place:

```bash
# Project initialization
claude-docs init [project-name]           # Create documentation structure
claude-docs validate [directory]          # Validate documentation structure

# Document management
claude-docs split <file> [options]        # Split large documents
claude-docs merge <directory> [options]   # Merge multiple documents

# Template generation
claude-docs template <type> [name]        # Generate documentation templates
#   Types: api, screen, feature, python, nodejs, go
```

### CLI Options & Features

**Document Splitting:**
```bash
claude-docs split large-doc.md --by-headers --max-sections 8
claude-docs split large-doc.md --by-size --max-size-kb 50
claude-docs split large-doc.md --by-lines --lines-per-file 200
```

**Document Merging:**
```bash
claude-docs merge specs/ --output combined.md
claude-docs merge docs/ --recursive --exclude "*.draft.md"
claude-docs merge . --pattern "*.md" --no-claude-optimization
```

**Cross-Platform Builds:**
```bash
make release    # Build for Linux, macOS, Windows (x64 & ARM64)
```

### Template System

Generate professional documentation templates:

```bash
# API documentation template
claude-docs template api user-management

# Screen specification template
claude-docs template screen dashboard

# Feature specification template
claude-docs template feature authentication
```

## 🌟 Examples & Workflows

### Common Workflows

**1. Starting a New Project:**
```bash
# Initialize structure
claude-docs init my-awesome-app

# Generate initial templates
claude-docs template api authentication
claude-docs template screen login
claude-docs template feature user-management

# Generate language-specific configurations
claude-docs template python my-api      # Creates Python-specific files
claude-docs template nodejs my-frontend # Creates Node.js setup
claude-docs template go my-service      # Creates Go project files

# Validate everything is set up correctly
claude-docs validate
```

**2. Managing Large Documentation:**
```bash
# Split a large specification into manageable chunks
claude-docs split specs/massive-api-spec.md --by-headers --max-sections 8

# Later, merge them for a comprehensive review
claude-docs merge specs/ --output complete-api-docs.md
```

**3. Team Collaboration:**
```bash
# Each team member works on their section
claude-docs merge specs/ --exclude work-in-progress.md --output team-review.md

# Validate the team's documentation structure
claude-docs validate
```

### Small Project Example
```markdown
# CLAUDE.md for TodoApp

## Project Overview
Simple React todo application with local storage.

## Key Files
- `src/App.js:1` - Main component
- `src/components/TodoList.js:15` - Todo list logic
- `src/utils/storage.js:8` - Local storage utilities

## Current Focus
Adding user authentication with Firebase.
```

### Generated Template Example
After running `claude-docs template api users`, you get:
```markdown
# Users API Endpoint

## Overview
Brief description of what this endpoint does.

## HTTP Method and URL
```
GET/POST/PUT/DELETE /api/users
```

## Parameters
### Path Parameters
- `id` (string): User identifier

### Query Parameters
- `limit` (number, optional): Results per page
- `offset` (number, optional): Pagination offset
...
```

### Large Project Example
See `examples/large-project/` for a comprehensive multi-service application structure.

## 🚀 Installation

### Binary Installation (Recommended)
```bash
# Download from GitHub Releases
curl -L https://github.com/your-username/claude-doc-structure/releases/latest/download/claude-docs-linux-amd64 -o claude-docs
chmod +x claude-docs
sudo mv claude-docs /usr/local/bin/
```

### Build from Source
```bash
git clone https://github.com/your-username/claude-doc-structure.git
cd claude-doc-structure
make build
# Binary will be in ./bin/claude-docs
```

### Requirements
- **Runtime**: None (single binary, zero dependencies)
- **Build**: Go 1.21+ (only if building from source)

## 🆘 Troubleshooting

**Command not found after installation:**
```bash
# Make sure the binary is executable and in PATH
chmod +x claude-docs
which claude-docs  # Should show the path

# Or use full path
./bin/claude-docs --help
```

**Permission denied errors:**
```bash
# Make sure you have write permissions to /usr/local/bin/
sudo mv claude-docs /usr/local/bin/
```

**For development:**
```bash
# Build and test
make build
./bin/claude-docs --help

# Run tests
make test

# Format and lint code
make fmt
make lint
```

## ✨ New Features in Latest Version

### Enhanced Templates
- **Language-Specific Support**: Python, Node.js, and Go project templates
- **Advanced AI Collaboration**: Best practices for Claude Code integration
- **Project-Specific Documentation**: Tailored templates for unique project needs
- **Comprehensive Specifications**: Detailed API and UI specification templates

### Template Types Available
```bash
# Documentation templates
claude-docs template api [name]       # Basic API documentation
claude-docs template screen [name]    # UI/screen specifications  
claude-docs template feature [name]   # Feature specifications

# Language-specific templates (coming soon)
claude-docs template python [name]    # Python project setup
claude-docs template nodejs [name]    # Node.js project setup
claude-docs template go [name]        # Go project setup
```

### Inspired by Real Projects
These templates incorporate patterns learned from production projects like adaptive-trading, ensuring practical applicability and real-world effectiveness.

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please see [CONTRIBUTING.md](.github/CONTRIBUTING.md) for detailed guidelines.

## 📝 License

MIT License - see [LICENSE](LICENSE) for details.

## 🔗 Related Resources

- [Claude Code Documentation](https://docs.anthropic.com/claude-code)
- [Best Practices for AI-Assisted Development](https://github.com/anthropics/claude-code)
- [Claude Knowledge Management (Zenn Article)](https://zenn.dev/driller/articles/2a23ef94f1d603) - **Inspiration for this project's documentation optimization approach**
- [Markdown Guide](https://www.markdownguide.org/)

## 📊 Features Summary

| Feature | Description | Command |
|---------|-------------|---------|
| 🚀 **Project Init** | Create optimal documentation structure | `claude-docs init` |
| ✂️ **Document Split** | Break large docs into manageable pieces | `claude-docs split` |
| 🔄 **Document Merge** | Combine multiple docs for review | `claude-docs merge` |
| 📝 **Template Gen** | Generate professional doc templates | `claude-docs template` |
| ✅ **Structure Validation** | Verify documentation best practices | `claude-docs validate` |
| 🏗️ **Cross-Platform** | Single binary for Linux, macOS, Windows | `make release` |
| ⚡ **Zero Dependencies** | No runtime dependencies needed | Download & run |
| 🎯 **Claude Optimized** | Built specifically for Claude Code workflows | All commands |

---

*Made with ❤️ for developers working with Claude Code*

*Turn your documentation chaos into Claude-optimized clarity!* ✨
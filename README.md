# Claude Doc Structure 📋

A comprehensive documentation toolkit for efficient collaboration with Claude Code. Streamline your development workflow with organized project context, automated tools, and reusable templates.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Python](https://img.shields.io/badge/Python-3.8%2B-blue.svg)](https://www.python.org/)
[![CLI Tool](https://img.shields.io/badge/CLI-Tool-brightgreen.svg)](https://github.com/your-username/claude-doc-structure)

## 🎯 Purpose

This template provides a structured approach to organizing project documentation that maximizes Claude Code's effectiveness. By following these patterns, you'll:

- Enable Claude Code to understand your project context faster
- Reduce repetitive explanations through organized documentation
- Maintain consistency across development sessions
- Scale documentation efficiently as your project grows

## 🚀 Quick Start

### Option 1: Use the CLI Tool (Recommended)

1. **Install the tool:**
   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cd claude-doc-structure
   make install
   ```

2. **Initialize your project:**
   ```bash
   claude-docs init my-project
   ```

3. **Start using the tools:**
   ```bash
   # Generate templates
   claude-docs template api users
   
   # Split large documents
   claude-docs split README.md --by-headers
   
   # Merge multiple documents
   claude-docs merge specs/ --output combined.md
   
   # Validate structure
   claude-docs validate
   ```

### Option 2: Manual Setup

1. **Copy the structure** to your project:
   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cp -r claude-doc-structure/.claude ./
   cp claude-doc-structure/CLAUDE.md ./
   ```

2. **Customize CLAUDE.md** with your project details and start collaborating!

## 📁 Structure Overview

```
your-project/
├── CLAUDE.md              # Main project context for Claude Code
├── specs/                 # Detailed specifications
│   ├── api.md            # API documentation
│   └── screens.md        # UI/UX specifications
├── .claude/              # Claude Code assets
│   ├── prompts/          # Reusable prompt templates
│   └── templates/        # Documentation templates
├── scripts/              # Documentation utilities
│   ├── split_docs.py     # Split large docs into smaller files
│   └── merge_docs.py     # Merge docs for comprehensive view
├── claude-docs           # Unified CLI tool
├── Makefile             # Convenient shortcuts
└── setup.py             # Package configuration
```

## 📖 Documentation Approach

### Staged Documentation Strategy

**Stage 1: Single File (Small Projects)**
- Use only `CLAUDE.md`
- Keep everything in one place for quick reference
- Perfect for projects under 10,000 lines

**Stage 2: Split by Domain (Medium Projects)**
- Create `specs/api.md`, `specs/screens.md`
- Reference from `CLAUDE.md` using relative paths
- Ideal for projects with multiple components

**Stage 3: Hierarchical Structure (Large Projects)**
- Organize by features and modules
- Use `scripts/split_docs.py` for automated organization
- Maintain cross-references between documents

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
```

### Makefile Shortcuts

For even easier usage, use the provided Makefile:

```bash
make help                                  # Show all available commands
make init                                  # Initialize documentation structure
make split FILE=large-doc.md              # Split a document
make merge DIR=specs/                     # Merge documents
make template TYPE=api NAME=users         # Generate template
make validate                             # Validate structure
```

### Document Management Scripts

**Direct script usage (if you prefer):**
```bash
# Split large documentation
python scripts/split_docs.py specs/large-spec.md --by-headers --max-sections 5

# Merge for comprehensive review
python scripts/merge_docs.py specs/ --output combined.md
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

### Quick Install
```bash
git clone https://github.com/your-username/claude-doc-structure.git
cd claude-doc-structure
make install
```

### Manual Install
```bash
pip install -e .
```

### Requirements
- Python 3.8+
- No external dependencies (uses standard library only)

## 🆘 Troubleshooting

**Command not found after installation:**
```bash
# Make sure the script is executable
chmod +x claude-docs

# Add to PATH or use full path
export PATH=$PATH:$(pwd)
```

**Permission denied errors:**
```bash
# Run with appropriate permissions
sudo make install
```

**For development:**
```bash
# Install in development mode
pip install -e ".[dev]"

# Run tests
make test

# Run linting
make lint
```

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
- [Markdown Guide](https://www.markdownguide.org/)

## 📊 Features Summary

| Feature | Description | Command |
|---------|-------------|---------|
| 🚀 **Project Init** | Create optimal documentation structure | `claude-docs init` |
| ✂️ **Document Split** | Break large docs into manageable pieces | `claude-docs split` |
| 🔄 **Document Merge** | Combine multiple docs for review | `claude-docs merge` |
| 📝 **Template Gen** | Generate professional doc templates | `claude-docs template` |
| ✅ **Structure Validation** | Verify documentation best practices | `claude-docs validate` |
| 🛠️ **Makefile Support** | Convenient shortcuts for all operations | `make help` |
| 📦 **Easy Install** | Simple installation process | `make install` |

---

*Made with ❤️ for developers working with Claude Code*

*Turn your documentation chaos into Claude-optimized clarity!* ✨
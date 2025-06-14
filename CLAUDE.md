# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Doc Structure is a documentation toolkit that provides templates, tools, and best practices for organizing project context to maximize Claude Code's effectiveness. It transforms documentation chaos into Claude-optimized clarity through systematic organization and automation.

## Architecture & Technology Stack

**Core Technologies:**
- Python 3.8+ (standard library only, no external dependencies)
- Bash scripting for automation
- Markdown for all documentation templates

**Key Components:**
- CLI tool (`tools/cli/claude-docs`) - Unified command interface
- Document processing scripts (`tools/scripts/`) - Split/merge utilities
- Template system (`templates/`) - User-facing documentation templates
- Development tools (`tools/`) - Testing and automation

## Project Structure

```
claude-doc-structure/
├── templates/                  # User-facing templates (copied to projects)
│   ├── .claude/               # Claude-specific assets
│   ├── specs/                 # Documentation templates
│   ├── examples/              # Reference implementations
│   └── CLAUDE.md              # Generic project template
├── tools/                     # Development tools
│   ├── cli/claude-docs        # Main CLI application
│   ├── scripts/               # Document processing utilities
│   └── tests/                 # Test suite
├── setup.py & pyproject.toml  # Package configuration
└── Makefile                   # Development automation
```

## Common Development Commands

**Installation & Setup:**
```bash
make install                   # Install CLI tool locally
pip install -e .              # Install in development mode
pip install -e ".[dev]"       # Install with dev dependencies
```

**CLI Testing:**
```bash
tools/cli/claude-docs --help          # Test CLI functionality
tools/cli/claude-docs init test-proj  # Test project initialization
tools/cli/claude-docs validate        # Test structure validation
```

**Script Testing:**
```bash
python tools/scripts/split_docs.py --help
python tools/scripts/merge_docs.py --help
```

**Quality Checks:**
```bash
make test                      # Run test suite
make lint                      # Code quality checks (black, flake8, mypy)
make clean                     # Clean build artifacts
```

## Key Implementation Details

**CLI Architecture:**
- `tools/cli/claude-docs:31` - Project initialization using templates/
- `tools/cli/claude-docs:175` - Document splitting functionality
- `tools/cli/claude-docs:215` - Template generation system
- `tools/cli/claude-docs:483` - Structure validation logic

**Document Processing:**
- `tools/scripts/split_docs.py:24` - DocumentSplitter class for large files
- `tools/scripts/split_docs.py:45` - Header-based splitting algorithm
- `tools/scripts/merge_docs.py:26` - DocumentMerger for combining files

**Template System:**
- Templates in `templates/` are copied to user projects
- CLI references `templates/.claude/` for Claude-specific assets
- Template placeholders use `{PROJECT_NAME}` format

## Current Development Focus

**Path Migration (In Progress):**
The project was recently restructured to separate user templates from development tools. Key updates needed:
- CLI script paths: Update template source paths from legacy structure
- Setup configuration: Ensure `setup.py` points to correct script location
- Documentation: Update README and help text for new structure

**Configuration Notes:**
- `setup.py:55` - References correct CLI script path: `tools/cli/claude-docs`
- `pyproject.toml:47` - Uses entry point for CLI installation
- `Makefile` - Provides convenient development shortcuts

## Testing Strategy

**Manual Testing:**
1. Test CLI commands: `tools/cli/claude-docs [command]`
2. Test document processing: Use sample markdown files
3. Test template generation: Verify output structure
4. Test Makefile shortcuts: `make [command]`

**Automated Testing:**
- `tools/tests/test_cli.py` - CLI functionality tests
- Use `pytest` for test execution
- Code quality tools: black, flake8, mypy (optional)

## Dependencies

**Runtime:** None (uses Python standard library only)
**Development:** pytest, black, flake8, mypy (optional via `pip install -e ".[dev]"`)

This tool is designed to be dependency-free for easy adoption across different environments.
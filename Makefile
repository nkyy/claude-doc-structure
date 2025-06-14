# Makefile for Claude Doc Structure
# Provides convenient shortcuts for common operations

.PHONY: help install clean test split merge validate init template

# Default target
help:
	@echo "Claude Doc Structure - Available Commands:"
	@echo ""
	@echo "Setup & Installation:"
	@echo "  make install     Install the claude-docs tool locally"
	@echo "  make clean       Clean up build artifacts"
	@echo ""
	@echo "Documentation Operations:"
	@echo "  make init        Initialize documentation structure in current directory"
	@echo "  make split FILE=<file> [OPTIONS]  Split a large document"
	@echo "  make merge DIR=<dir> [OPTIONS]    Merge multiple documents"
	@echo "  make validate    Validate documentation structure"
	@echo "  make template TYPE=<type> [NAME] Generate documentation template"
	@echo ""
	@echo "Development:"
	@echo "  make test        Run tests (if available)"
	@echo "  make lint        Run code quality checks"
	@echo ""
	@echo "Examples:"
	@echo "  make init"
	@echo "  make split FILE=README.md"
	@echo "  make merge DIR=specs/"
	@echo "  make template TYPE=api NAME=users"
	@echo ""

# Installation
install:
	@echo "Installing claude-docs tool..."
	pip install -e .
	@echo "Installation complete! You can now use 'claude-docs' command."

# Clean up
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf build/
	rm -rf dist/
	rm -rf *.egg-info/
	find . -type d -name __pycache__ -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	@echo "Clean complete."

# Initialize documentation structure
init:
	@echo "Initializing Claude documentation structure..."
	./claude-docs init

# Split document
split:
ifndef FILE
	@echo "Error: FILE parameter is required"
	@echo "Usage: make split FILE=path/to/document.md [OPTIONS='--by-headers --max-sections 5']"
	@exit 1
endif
	@echo "Splitting document: $(FILE)"
	./claude-docs split "$(FILE)" $(OPTIONS)

# Merge documents
merge:
ifndef DIR
	@echo "Error: DIR parameter is required"
	@echo "Usage: make merge DIR=path/to/directory [OPTIONS='--output combined.md']"
	@exit 1
endif
	@echo "Merging documents from: $(DIR)"
	./claude-docs merge "$(DIR)" $(OPTIONS)

# Validate structure
validate:
	@echo "Validating documentation structure..."
	./claude-docs validate

# Generate template
template:
ifndef TYPE
	@echo "Error: TYPE parameter is required"
	@echo "Usage: make template TYPE=api|screen|feature [NAME=template-name]"
	@exit 1
endif
	@echo "Generating $(TYPE) template..."
ifdef NAME
	./claude-docs template $(TYPE) $(NAME)
else
	./claude-docs template $(TYPE)
endif

# Development tasks
test:
	@echo "Running tests..."
	@if [ -d "tests" ]; then \
		python -m pytest tests/ -v; \
	else \
		echo "No tests directory found."; \
	fi

lint:
	@echo "Running code quality checks..."
	@if command -v black >/dev/null 2>&1; then \
		echo "Running Black formatter..."; \
		black --check scripts/ claude-docs; \
	fi
	@if command -v flake8 >/dev/null 2>&1; then \
		echo "Running Flake8..."; \
		flake8 scripts/ claude-docs; \
	fi
	@if command -v mypy >/dev/null 2>&1; then \
		echo "Running MyPy type checker..."; \
		mypy scripts/ claude-docs; \
	fi

# Quick operations for common workflows
quick-setup: init
	@echo "Quick setup complete!"
	@echo "Next steps:"
	@echo "1. Edit CLAUDE.md with your project details"
	@echo "2. Run 'make validate' to check your structure"
	@echo "3. Use 'make template TYPE=api' to generate templates"

# Example workflows
example-split:
	@echo "Example: Splitting a large README into sections..."
	@if [ -f "README.md" ]; then \
		./claude-docs split README.md --by-headers --max-sections 5 --prefix readme-; \
	else \
		echo "No README.md found to demonstrate splitting."; \
	fi

example-merge:
	@echo "Example: Merging specs directory..."
	@if [ -d "specs" ]; then \
		./claude-docs merge specs/ --output combined-specs.md; \
	else \
		echo "No specs/ directory found to demonstrate merging."; \
	fi

# Help for specific commands
help-split:
	@echo "Split Command Options:"
	@echo "  FILE=<path>              Input markdown file (required)"
	@echo "  OPTIONS='--by-headers'   Split by markdown headers (default)"
	@echo "  OPTIONS='--by-lines'     Split by line count"
	@echo "  OPTIONS='--by-size'      Split by file size"
	@echo "  OPTIONS='--max-sections N'  Maximum sections (default: 10)"
	@echo "  OPTIONS='--header-level N'   Header level to split on (default: 2)"
	@echo "  OPTIONS='--prefix text'      Filename prefix"
	@echo ""
	@echo "Examples:"
	@echo "  make split FILE=large-doc.md"
	@echo "  make split FILE=api.md OPTIONS='--by-headers --max-sections 8'"
	@echo "  make split FILE=guide.md OPTIONS='--by-size --max-size-kb 50'"

help-merge:
	@echo "Merge Command Options:"
	@echo "  DIR=<path>               Input directory (required)"
	@echo "  OPTIONS='--output file'  Output filename (default: merged-docs.md)"
	@echo "  OPTIONS='--pattern *.md' File pattern (default: *.md)"
	@echo "  OPTIONS='--recursive'    Search subdirectories"
	@echo "  OPTIONS='--exclude file' Files to exclude"
	@echo ""
	@echo "Examples:"
	@echo "  make merge DIR=docs/"
	@echo "  make merge DIR=specs/ OPTIONS='--output api-docs.md'"
	@echo "  make merge DIR=. OPTIONS='--recursive --exclude README.md'"

help-template:
	@echo "Template Command Options:"
	@echo "  TYPE=api                 Generate API endpoint template"
	@echo "  TYPE=screen              Generate screen specification template"
	@echo "  TYPE=feature             Generate feature specification template"
	@echo "  NAME=<name>              Template name (optional)"
	@echo ""
	@echo "Examples:"
	@echo "  make template TYPE=api NAME=users"
	@echo "  make template TYPE=screen NAME=dashboard"
	@echo "  make template TYPE=feature NAME=authentication"
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize documentation structure",
	Long:  "Initialize Claude documentation structure in current directory.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := getProjectName(args)
		initProject(projectName)
	},
}

func getProjectName(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	
	cwd, err := os.Getwd()
	if err != nil {
		return "project"
	}
	
	return filepath.Base(cwd)
}

func initProject(projectName string) {
	fmt.Printf("Initializing Claude documentation structure for '%s'...\n", projectName)
	
	// Create enhanced structure with .claude/ optimization
	directories := []string{
		"specs",
		".claude",
	}
	
	for _, dir := range directories {
		err := os.MkdirAll(dir, 0755)
		checkError(err)
		fmt.Printf("Created directory: %s\n", dir)
	}
	
	// Load and customize CLAUDE.md template
	claudeMdContent := loadTemplate("CLAUDE.md", projectName)
	
	claudeMdPath := "CLAUDE.md"
	if _, err := os.Stat(claudeMdPath); os.IsNotExist(err) {
		err := os.WriteFile(claudeMdPath, []byte(claudeMdContent), 0644)
		checkError(err)
		fmt.Printf("Created: %s\n", claudeMdPath)
	} else {
		fmt.Println("CLAUDE.md already exists, skipping...")
	}
	
	// Create all templates with .claude/ structure
	templateFiles := map[string]string{
		"specs/api.md":                    loadTemplate("specs/api.md", projectName),
		"specs/screens.md":                loadTemplate("specs/screens.md", projectName),
		".claude/context.md":              loadTemplate(".claude/context.md", projectName),
		".claude/project-knowledge.md":    loadTemplate(".claude/project-knowledge.md", projectName),
		".claude/project-improvements.md": loadTemplate(".claude/project-improvements.md", projectName),
		".claude/common-patterns.md":      loadTemplate(".claude/common-patterns.md", projectName),
		".claude/debug-log.md":            loadTemplate(".claude/debug-log.md", projectName),
	}
	
	for filePath, content := range templateFiles {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err := os.WriteFile(filePath, []byte(content), 0644)
			checkError(err)
			fmt.Printf("Created: %s\n", filePath)
		}
	}
	
	fmt.Println("\nClaude-optimized documentation structure initialized successfully!")
	fmt.Println("Next steps:")
	fmt.Println("1. Edit CLAUDE.md with your project details")
	fmt.Println("2. Update .claude/context.md with project background and constraints")
	fmt.Println("3. Fill in .claude/project-knowledge.md with technical insights")
	fmt.Println("4. Update specs/api.md and specs/screens.md")
	fmt.Println("5. Start collaborating with Claude Code for enhanced AI assistance!")
	fmt.Println("\nThis structure follows best practices from:")
	fmt.Println("https://zenn.dev/driller/articles/2a23ef94f1d603")
}

func loadTemplate(templatePath, projectName string) string {
	// For now, use built-in templates
	// In the future, this could read from external files or embedded templates
	return createBuiltinTemplate(templatePath, projectName)
}

func createBuiltinTemplate(templatePath, projectName string) string {
	switch templatePath {
	case "specs/api.md":
		return `# API Documentation

## Endpoints

### GET /api/endpoint
Description of the endpoint.

**Parameters:**
- param1 (string): Description
- param2 (number): Description

**Response:**
` + "```json" + `
{
  "example": "response"
}
` + "```" + `

**Example:**
` + "```bash" + `
curl -X GET "http://localhost:3000/api/endpoint?param1=value"
` + "```" + `
`
	case "specs/screens.md":
		return `# Screen Specifications

## Screen Name

### Purpose
What this screen is for and when it's used.

### Layout
Description of the visual layout and components.

### User Interactions
- Action 1: What happens
- Action 2: What happens

### API Calls
- Endpoint used: Purpose
- Data displayed: Source

### Navigation
- From: Where users come from
- To: Where users can go next
`
	case ".claude/context.md":
		return strings.ReplaceAll(`# Project Context

## Project Overview
{project_name} - Brief description of your project and its purpose.

## Core Mission
Define the main goal and value proposition of your project.

## Key Constraints & Requirements

### Technical Constraints
- **Runtime Environment**: List required environments (Node.js, Python, Go, etc.)
- **Platform Support**: Target platforms (Web, Mobile, Desktop, etc.)
- **Performance Requirements**: Critical performance metrics
- **Security Requirements**: Authentication, data protection, compliance needs

### Design Principles
- **User Experience**: Core UX principles and guidelines
- **Code Quality**: Standards for maintainability and scalability
- **Architecture**: Key architectural decisions and patterns
- **Integration**: External systems and API requirements

### Business Context
- **Target Users**: Primary user personas and use cases
- **Market Requirements**: Business constraints and competitive factors
- **Timeline**: Key milestones and delivery expectations
- **Success Metrics**: How success is measured

### Current Development Phase
- 🔄 **Active Development**: Current focus areas
- 📋 **Backlog**: Planned features and improvements
- ⚠️ **Known Issues**: Critical issues requiring attention
- ✅ **Completed**: Recent achievements and milestones

## Success Metrics
- Performance targets and benchmarks
- User satisfaction and engagement goals
- Technical debt and code quality metrics
- Business objectives and KPIs`, "{project_name}", projectName)

	case ".claude/project-knowledge.md":
		return strings.ReplaceAll(`# Project Technical Knowledge

## Architecture Insights

### Core Architecture Patterns
- **Design Pattern**: [MVC/MVP/MVVM/etc.] - Rationale and implementation
- **Data Flow**: How data moves through the system
- **State Management**: Approach to managing application state
- **Error Handling**: Consistent error handling patterns

### Key Technical Decisions

#### Technology Stack Rationale
- **Frontend**: [Framework/Library] - Why chosen, benefits, trade-offs
- **Backend**: [Framework/Language] - Performance and maintainability considerations
- **Database**: [Database type] - Data structure and query patterns
- **Infrastructure**: [Hosting/Services] - Scalability and cost considerations

#### Code Organization
` + "```" + `
{project_name}/
├── src/
│   ├── components/     # Reusable UI components
│   ├── services/       # Business logic and API calls
│   ├── utils/          # Helper functions and utilities
│   └── types/          # Type definitions and interfaces
├── tests/              # Test files and test utilities
└── docs/               # Documentation and guides
` + "```" + `

### Performance Optimizations
1. **Critical Path Optimization**: Key performance bottlenecks addressed
2. **Caching Strategy**: What is cached and how
3. **Load Time Improvements**: Techniques used to improve initial load
4. **Memory Management**: Approaches to prevent memory leaks

### Code Patterns & Conventions

#### Naming Conventions
- **Variables**: camelCase for variables, PascalCase for classes
- **Functions**: Descriptive names following verb-noun pattern
- **Files**: Consistent naming scheme for different file types
- **APIs**: RESTful conventions or GraphQL patterns

#### Common Patterns
` + "```javascript" + `
// Example: Error handling pattern
try {
  const result = await apiCall();
  return handleSuccess(result);
} catch (error) {
  return handleError(error);
}

// Example: Component structure pattern
const ComponentName = ({ prop1, prop2 }) => {
  // Hook declarations
  // Event handlers
  // Render logic
};
` + "```" + `

### Testing Strategy
- **Unit Tests**: Testing individual functions and components
- **Integration Tests**: Testing component interactions
- **E2E Tests**: Testing complete user workflows
- **Performance Tests**: Load testing and benchmarking

## Critical Implementation Details

### Security Considerations
- **Authentication**: User authentication and session management
- **Authorization**: Role-based access control
- **Data Validation**: Input sanitization and validation
- **Secure Communication**: HTTPS, API security, data encryption

### Common Gotchas & Solutions

#### Development Pitfalls
- **Problem**: Common issue developers encounter
- **Solution**: Established solution or workaround
- **Prevention**: How to avoid the issue in the future`, "{project_name}", projectName)

	case ".claude/project-improvements.md":
		return `# Project Improvement History

## Major Milestones

### YYYY-MM-DD: Project Initialization
**Achievement**: Initial project setup and foundation
- ✅ Project structure established
- ✅ Core dependencies configured
- ✅ Development environment setup
- **Impact**: Ready for feature development

## Performance Improvements

### [Performance Area] Optimization
**Before**: Description of previous state with metrics
**After**: Description of improved state with metrics
**Techniques Used**: 
- Specific optimization technique 1
- Specific optimization technique 2
**Benefit**: Impact on development workflow or user experience

## Lessons Learned

### Technical Lessons
- **What Worked Well**: Successful approaches and decisions
- **What Could Be Improved**: Areas for future enhancement
- **Unexpected Challenges**: Problems that weren't anticipated
- **Best Practices Developed**: New practices established during development

## Future Improvement Areas

### High Priority
- [ ] Specific improvement needed with business impact
- [ ] Technical debt that should be addressed
- [ ] Performance bottleneck to resolve

### Medium Priority
- [ ] Enhancement that would improve developer experience
- [ ] Code refactoring for better maintainability
- [ ] Additional feature that users have requested

## Metrics & Success Indicators

### Technical Metrics
- **Code Quality**: Test coverage, lint score, complexity metrics
- **Performance**: Load times, response times, bundle sizes
- **Reliability**: Uptime, error rates, crash reports

### Business Metrics
- **User Engagement**: Usage statistics, retention rates
- **Performance**: Key business KPIs and goals
- **Cost Efficiency**: Infrastructure costs, development velocity`

	case ".claude/common-patterns.md":
		return strings.ReplaceAll(`# Common Command Patterns

## Development Workflow

### Quick Development Cycle
` + "```bash" + `
# Start development server
npm run dev
# or
yarn dev
# or
python manage.py runserver

# Run tests during development
npm run test:watch
# or
pytest --watch

# Build for production
npm run build
# or
python -m build
` + "```" + `

### Code Quality Checks
` + "```bash" + `
# Linting and formatting
npm run lint
npm run format
# or
flake8 . && black .

# Type checking
npm run type-check
# or
mypy .
` + "```" + `

## Testing Patterns

### Test Execution
` + "```bash" + `
# Run all tests
npm test
# or
pytest

# Run specific test file
npm test -- ComponentName.test.js
# or
pytest tests/test_specific.py

# Run tests with coverage
npm run test:coverage
# or
pytest --cov=src
` + "```" + `

## API Development

### API Testing
` + "```bash" + `
# Test API endpoints
curl -X GET "http://localhost:3000/api/users"
curl -X POST "http://localhost:3000/api/users" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'

# Using HTTPie (alternative to curl)
http GET localhost:3000/api/users
http POST localhost:3000/api/users name="John Doe" email="john@example.com"
` + "```" + `

## Environment Management

### Environment Variables
` + "```bash" + `
# Load environment variables
source .env
# or
export $(cat .env | xargs)

# Check environment configuration
npm run env:check
# or
python -c "import os; print(os.environ)"
` + "```" + `

## IDE Integration Patterns

### Command Line Shortcuts
` + "```bash" + `
# Add to shell profile (.bashrc, .zshrc)
alias dev='npm run dev'
alias test='npm run test'
alias build='npm run build'

# Project-specific shortcuts
alias {project_name}-dev='cd ~/projects/{project_name} && npm run dev'
alias {project_name}-test='cd ~/projects/{project_name} && npm test'
` + "```", "{project_name}", projectName)

	case ".claude/debug-log.md":
		return `# Debug Log & Issue Resolution

## Critical Issues Resolved

### YYYY-MM-DD: [Issue Title]
**Problem**: Detailed description of the issue
**Root Cause**: What was causing the problem
**Investigation Steps**:
` + "```bash" + `
# Commands used to investigate
command1 --flag
command2 --debug
# Output analysis
` + "```" + `
**Solution**: How the issue was resolved
**Prevention**: Steps taken to prevent recurrence

## Common Development Issues

### Environment Setup Problems
**Issue**: Development environment not working correctly
**Symptoms**:
- Command not found errors
- Missing dependencies
- Permission issues
- Configuration conflicts

**Debugging Steps**:
` + "```bash" + `
# Check installed versions
node --version
npm --version
python --version

# Verify PATH
echo $PATH

# Check permissions
ls -la
whoami
` + "```" + `

**Solutions**:
- Environment variable configuration
- Dependency installation order
- Permission adjustments
- Configuration file updates

## Performance Monitoring

### Key Metrics to Track
- Response time percentiles
- Error rates by endpoint
- Database query performance
- Memory and CPU usage
- Disk space utilization

## Issue Templates

### Bug Report Template
` + "```markdown" + `
**Environment:**
- OS: [Operating System]
- Runtime: [Node.js/Python/Go version]
- Browser: [if applicable]
- Deployment: [local/staging/production]

**Steps to Reproduce:**
1. Step 1
2. Step 2
3. Step 3

**Expected Behavior:**
What should happen

**Actual Behavior:**
What actually happened

**Error Messages:**
` + "```" + `
Paste error messages here
` + "```" + `

**Additional Context:**
Any additional information
` + "```" + `

### Performance Issue Template
` + "```markdown" + `
**Performance Issue:**
- Endpoint/Feature: [specific area]
- Current Performance: [metrics]
- Expected Performance: [target metrics]
- Impact: [user/business impact]

**Investigation:**
` + "```" + `
Paste investigation commands and output
` + "```" + `

**Potential Solutions:**
- Solution 1: [description and effort]
- Solution 2: [description and effort]
` + "```" + ``

	case "CLAUDE.md":
		return strings.ReplaceAll(`# {project_name}

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Context Access

**Essential Files for AI Context:**
- ` + "`.claude/context.md`" + ` - Project background, constraints, and requirements
- ` + "`.claude/project-knowledge.md`" + ` - Technical architecture and patterns  
- ` + "`.claude/project-improvements.md`" + ` - Development history and lessons learned
- ` + "`.claude/common-patterns.md`" + ` - Frequently used command patterns
- ` + "`.claude/debug-log.md`" + ` - Critical issues and troubleshooting

## Project Overview

{project_name} - Brief description of your project and its purpose. This project follows Claude Code optimization best practices for enhanced AI-assisted development.

## Architecture & Technology Stack

**Core Technologies:**
- List your main technologies here
- Framework versions
- Key dependencies

**Key Components:**
- Component 1: Description and location
- Component 2: Description and location
- Component 3: Description and location

## Current Development Status

**✅ Completed:**
- Initial project setup and foundation
- Core feature implementation
- Basic testing infrastructure

**🔄 Active Development:**
- Feature enhancements
- Performance optimizations
- User experience improvements

**📋 Next Priority:**
- Planned features and improvements
- Technical debt resolution
- Documentation updates

## Key Files & Components

- ` + "`src/main.js:1`" + ` - Main application entry point
- ` + "`src/components/App.js:15`" + ` - Main application component
- ` + "`src/services/api.js:8`" + ` - API service layer
- ` + "`src/utils/helpers.js:12`" + ` - Utility functions
- ` + "`tests/integration/api.test.js:25`" + ` - API integration tests

## Documentation Maintenance

**Important:** When making changes to the project:
1. Update ` + "`.claude/project-improvements.md`" + ` with changes and lessons learned
2. Document any issues in ` + "`.claude/debug-log.md`" + ` with solutions
3. Update ` + "`.claude/project-knowledge.md`" + ` with new technical insights
4. Keep ` + "`.claude/common-patterns.md`" + ` current with new workflow patterns
5. Maintain this CLAUDE.md file with structural changes

This project structure is optimized for Claude Code AI assistance following best practices from [Zenn article on Claude knowledge management](https://zenn.dev/driller/articles/2a23ef94f1d603).`, "{project_name}", projectName)

	default:
		return fmt.Sprintf("# %s\n\nTemplate content for %s", projectName, templatePath)
	}
}
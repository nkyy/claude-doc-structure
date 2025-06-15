# {project_name}

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Context Access

**Essential Files for AI Context:**
- `.claude/context.md` - Project background, constraints, and requirements
- `.claude/project-knowledge.md` - Technical architecture and patterns  
- `.claude/project-improvements.md` - Development history and lessons learned
- `.claude/common-patterns.md` - Frequently used command patterns
- `.claude/debug-log.md` - Critical issues and troubleshooting

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

## Project Structure

```
{project_name}/
├── src/                    # Source code
├── tests/                  # Test files
├── docs/                   # Documentation
├── CLAUDE.md              # This file - main project context
├── .claude/               # Claude Code optimization files
│   ├── context.md         # Project background and constraints
│   ├── project-knowledge.md # Technical insights and patterns
│   ├── project-improvements.md # Development history
│   ├── common-patterns.md # Command patterns and workflows
│   └── debug-log.md       # Issue resolution history
└── specs/                 # Detailed specifications
    ├── api.md             # API documentation
    └── screens.md         # UI/screen specifications
```

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

## Common Development Commands

**Development:**
```bash
npm run dev              # Start development server
npm run build           # Build for production
npm run test            # Run test suite
npm run lint            # Code quality checks
```

**Database:**
```bash
npm run db:migrate      # Run database migrations
npm run db:seed         # Seed with sample data
npm run db:reset        # Reset database
```

**Deployment:**
```bash
npm run deploy:staging  # Deploy to staging
npm run deploy:prod     # Deploy to production
```

## Key Files & Components

- `src/main.js:1` - Main application entry point
- `src/components/App.js:15` - Main application component
- `src/services/api.js:8` - API service layer
- `src/utils/helpers.js:12` - Utility functions
- `tests/integration/api.test.js:25` - API integration tests

## Dependencies

**Production Dependencies:**
- dependency-name: Purpose and version
- framework-name: Core functionality

**Development Dependencies:**
- testing-framework: Testing infrastructure
- build-tool: Build and bundling

## Testing Strategy

**Test Coverage:**
- Unit Tests: Individual component testing
- Integration Tests: Component interaction testing
- E2E Tests: Complete user workflow testing

**Test Commands:**
```bash
npm test                # Run all tests
npm run test:unit       # Unit tests only
npm run test:e2e        # End-to-end tests
npm run test:coverage   # Generate coverage reports
```

## Documentation Maintenance

**Important:** When making changes to the project:
1. Update `.claude/project-improvements.md` with changes and lessons learned
2. Document any issues in `.claude/debug-log.md` with solutions
3. Update `.claude/project-knowledge.md` with new technical insights
4. Keep `.claude/common-patterns.md` current with new workflow patterns
5. Maintain this CLAUDE.md file with structural changes

## Usage Guidelines

**Getting Started:**
1. Clone the repository
2. Install dependencies: `npm install`
3. Set up environment: Copy `.env.example` to `.env`
4. Run migrations: `npm run db:migrate`
5. Start development: `npm run dev`

**Development Workflow:**
1. Create feature branch from main
2. Implement changes with tests
3. Run quality checks: `npm run lint && npm test`
4. Submit pull request with clear description
5. Deploy after code review and approval

This project structure is optimized for Claude Code AI assistance following best practices from [Zenn article on Claude knowledge management](https://zenn.dev/driller/articles/2a23ef94f1d603).
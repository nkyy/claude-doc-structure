# Development Prompts for Claude Doc Structure

## Adding New Features

```
I need to add a new feature to the claude-docs tool: ${FEATURE_DESCRIPTION}

Current codebase structure:
- bin/claude-docs: Main CLI script with all commands
- scripts/: Individual utility scripts (split_docs.py, merge_docs.py)
- .claude/templates/: Template system for generating documentation

Please:
1. Analyze the existing code structure in bin/claude-docs
2. Implement the new feature following the existing patterns
3. Add appropriate help text and argument parsing
4. Test the implementation
5. Update the README.md if needed

Make sure to maintain consistency with existing command structure and error handling.
```

## Code Refactoring

```
I want to refactor the claude-docs codebase to improve maintainability:

Current issues:
- Single large file (bin/claude-docs) contains all functionality
- Repeated code patterns
- ${SPECIFIC_ISSUE}

Please:
1. Break down the monolithic script into logical modules
2. Create a proper Python package structure
3. Maintain backward compatibility
4. Add proper error handling and logging
5. Keep the CLI interface unchanged

Focus on clean code principles and maintainability.
```

## Template Updates

```
I need to update the documentation templates in .claude/templates/:

Template type: ${TEMPLATE_TYPE}
Required changes: ${CHANGES_DESCRIPTION}

Please:
1. Review the current template structure
2. Update the template content while maintaining formatting
3. Ensure placeholder variables work correctly
4. Test template generation
5. Update any related documentation

Keep templates professional and consistent with the project style.
```

## Bug Fixes

```
There's a bug in the claude-docs tool:

Issue: ${BUG_DESCRIPTION}
Expected behavior: ${EXPECTED_BEHAVIOR}
Actual behavior: ${ACTUAL_BEHAVIOR}
Steps to reproduce: ${REPRODUCTION_STEPS}

Please:
1. Investigate the root cause
2. Implement a fix
3. Add error handling to prevent similar issues
4. Test the fix thoroughly
5. Consider if documentation needs updates

Make sure the fix doesn't break existing functionality.
```

## Documentation Improvements

```
I want to improve the project documentation:

Area: ${DOCUMENTATION_AREA}
Current issues: ${CURRENT_ISSUES}
Target audience: ${TARGET_AUDIENCE}

Please:
1. Review existing documentation
2. Identify gaps and inconsistencies
3. Improve clarity and completeness
4. Add practical examples
5. Ensure consistency across all docs

Focus on making the documentation helpful for both new users and contributors.
```
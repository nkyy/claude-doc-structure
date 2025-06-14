# Refactoring Prompts for Claude Code

Collection of reusable prompts for common code refactoring tasks. Replace variables in `${VARIABLE_NAME}` format with your specific values.

## Code Structure Refactoring

### Extract Function/Method
```
I need to extract a function from ${SOURCE_FILE} at ${LINE_RANGE}.

Extraction details:
- Current code location: ${SOURCE_FILE}:${LINE_RANGE}
- Purpose of extracted function: ${FUNCTION_PURPOSE}
- Proposed function name: ${FUNCTION_NAME}
- Parameters needed: ${FUNCTION_PARAMETERS}
- Return type: ${RETURN_TYPE}

Requirements:
- Follow the existing code style and naming conventions
- Maintain all existing functionality
- Update all call sites
- Add appropriate documentation
- Include proper error handling if needed

Please examine the code and create a clean extraction that improves code organization while maintaining all existing behavior.
```

### Extract Class/Component
```
I need to extract a ${CLASS_TYPE} from ${SOURCE_FILE} for ${EXTRACTION_PURPOSE}.

Current situation:
- Source file: ${SOURCE_FILE}
- Code to extract: ${CODE_DESCRIPTION}
- New ${CLASS_TYPE} name: ${NEW_CLASS_NAME}
- Responsibilities: ${CLASS_RESPONSIBILITIES}

Requirements:
- Follow the existing architectural patterns
- Maintain proper separation of concerns
- Update all dependencies and imports
- Preserve existing functionality
- Follow the project's naming conventions

Please examine the existing code structure and create a well-designed extraction that improves the overall architecture.
```

### Consolidate Duplicate Code
```
I have identified duplicate code in the following locations:
${DUPLICATE_LOCATIONS}

Duplication details:
- Common functionality: ${COMMON_FUNCTIONALITY}
- Variations between instances: ${DIFFERENCES}
- Proposed consolidation approach: ${CONSOLIDATION_STRATEGY}

Requirements:
- Create a reusable solution that handles all variations
- Update all existing usage sites
- Maintain backward compatibility
- Follow the project's patterns for shared utilities
- Add appropriate tests for the consolidated code

Please examine the duplicate code and create an elegant solution that eliminates duplication while handling all the variations.
```

## Architecture Refactoring

### Improve Separation of Concerns
```
I need to improve the separation of concerns in ${MODULE_NAME} which currently handles ${MIXED_RESPONSIBILITIES}.

Current issues:
- ${CURRENT_PROBLEMS}
- Mixed responsibilities: ${RESPONSIBILITY_OVERLAP}
- Tight coupling between: ${COUPLING_ISSUES}

Refactoring goals:
- Separate ${RESPONSIBILITY_A} from ${RESPONSIBILITY_B}
- Create clear interfaces between components
- Improve testability and maintainability
- Follow ${ARCHITECTURAL_PATTERN} pattern

Please examine the current structure and propose a refactoring approach that creates clear separation of concerns while maintaining all existing functionality.
```

### Introduce Design Pattern
```
I need to refactor ${COMPONENT_NAME} to use the ${DESIGN_PATTERN} pattern.

Current situation:
- Current implementation: ${CURRENT_APPROACH}
- Problems with current approach: ${CURRENT_PROBLEMS}
- Benefits of ${DESIGN_PATTERN}: ${PATTERN_BENEFITS}

Pattern requirements:
- ${PATTERN_SPECIFIC_REQUIREMENTS}
- Integration with existing code: ${INTEGRATION_REQUIREMENTS}
- Backward compatibility: ${COMPATIBILITY_REQUIREMENTS}

Please examine the current implementation and refactor it to properly implement the ${DESIGN_PATTERN} pattern while maintaining all existing functionality.
```

### Refactor to Dependency Injection
```
I need to refactor ${COMPONENT_NAME} to use dependency injection instead of ${CURRENT_DEPENDENCY_APPROACH}.

Current dependencies:
- ${CURRENT_DEPENDENCIES}
- Hard-coded dependencies: ${HARD_CODED_DEPS}
- Testing difficulties: ${TESTING_ISSUES}

DI requirements:
- Injectable dependencies: ${INJECTABLE_DEPS}
- Configuration approach: ${DI_CONFIGURATION}
- Container/framework: ${DI_FRAMEWORK}

Please examine the current code and refactor it to use proper dependency injection while maintaining all existing functionality and improving testability.
```

## Performance Refactoring

### Optimize Algorithm Performance
```
I need to optimize the performance of ${FUNCTION_NAME} in ${FILE_PATH}.

Current performance issues:
- ${PERFORMANCE_PROBLEMS}
- Current time complexity: ${CURRENT_COMPLEXITY}
- Target time complexity: ${TARGET_COMPLEXITY}
- Performance bottlenecks: ${BOTTLENECKS}

Optimization requirements:
- Maintain existing functionality and API
- Improve performance for ${PERFORMANCE_SCENARIOS}
- Consider memory usage implications
- Preserve existing error handling

Please analyze the current algorithm and implement optimizations that achieve the target performance while maintaining all existing behavior.
```

### Reduce Memory Usage
```
I need to reduce memory usage in ${COMPONENT_NAME} which currently has ${MEMORY_ISSUES}.

Memory problems:
- ${SPECIFIC_MEMORY_ISSUES}
- Memory leaks: ${LEAK_LOCATIONS}
- Inefficient data structures: ${INEFFICIENT_STRUCTURES}
- Large object retention: ${RETENTION_ISSUES}

Optimization goals:
- Reduce memory footprint by ${TARGET_REDUCTION}
- Fix memory leaks
- Optimize data structures
- Improve garbage collection

Please examine the current implementation and refactor it to use memory more efficiently while maintaining all existing functionality.
```

### Add Caching/Memoization
```
I need to add caching to ${COMPONENT_NAME} to improve performance for ${PERFORMANCE_SCENARIO}.

Caching requirements:
- Cache candidates: ${CACHEABLE_OPERATIONS}
- Cache invalidation strategy: ${INVALIDATION_STRATEGY}
- Cache size limits: ${CACHE_LIMITS}
- Cache storage: ${CACHE_STORAGE}

Implementation considerations:
- Thread safety: ${THREAD_SAFETY_REQUIREMENTS}
- Memory constraints: ${MEMORY_CONSTRAINTS}
- Cache hit rate optimization: ${HIT_RATE_GOALS}

Please examine the current code and implement an effective caching strategy that improves performance while maintaining correctness.
```

## Code Quality Refactoring

### Improve Error Handling
```
I need to improve error handling in ${COMPONENT_NAME} which currently has ${ERROR_HANDLING_ISSUES}.

Current problems:
- ${SPECIFIC_ERROR_ISSUES}
- Inconsistent error handling: ${INCONSISTENCIES}
- Missing error cases: ${MISSING_ERROR_CASES}
- Poor error messages: ${MESSAGE_ISSUES}

Improvement goals:
- Consistent error handling pattern: ${TARGET_PATTERN}
- Comprehensive error coverage: ${ERROR_COVERAGE}
- Clear error messages: ${MESSAGE_REQUIREMENTS}
- Proper error propagation: ${PROPAGATION_STRATEGY}

Please examine the current error handling and refactor it to provide robust, consistent error handling throughout the component.
```

### Add Input Validation
```
I need to add comprehensive input validation to ${FUNCTION_NAME} in ${FILE_PATH}.

Validation requirements:
- Input parameters: ${INPUT_PARAMETERS}
- Validation rules: ${VALIDATION_RULES}
- Error handling: ${VALIDATION_ERROR_HANDLING}
- Performance considerations: ${VALIDATION_PERFORMANCE}

Current state:
- Existing validation: ${CURRENT_VALIDATION}
- Missing validation: ${MISSING_VALIDATION}
- Validation inconsistencies: ${VALIDATION_ISSUES}

Please examine the current function and add comprehensive input validation that follows the project's validation patterns while maintaining performance.
```

### Improve Code Readability
```
I need to improve the readability of ${COMPONENT_NAME} in ${FILE_PATH}.

Readability issues:
- ${READABILITY_PROBLEMS}
- Complex logic: ${COMPLEX_AREAS}
- Poor naming: ${NAMING_ISSUES}
- Lack of documentation: ${DOCUMENTATION_GAPS}

Improvement goals:
- Simplify complex logic
- Improve variable and function names
- Add clear comments and documentation
- Break down large functions/methods
- Follow the project's style guidelines

Please examine the current code and refactor it to improve readability while maintaining all existing functionality.
```

## Database Refactoring

### Optimize Database Queries
```
I need to optimize database queries in ${COMPONENT_NAME} that are causing ${PERFORMANCE_ISSUES}.

Current query problems:
- Slow queries: ${SLOW_QUERIES}
- N+1 query problems: ${N_PLUS_ONE_ISSUES}
- Missing indexes: ${INDEX_OPPORTUNITIES}
- Inefficient joins: ${JOIN_ISSUES}

Optimization goals:
- Reduce query execution time by ${TARGET_IMPROVEMENT}
- Eliminate N+1 queries
- Add appropriate indexes
- Optimize join strategies

Please examine the current database access patterns and optimize the queries while maintaining all existing functionality.
```

### Refactor Database Schema
```
I need to refactor the database schema for ${SCHEMA_AREA} to address ${SCHEMA_ISSUES}.

Current schema problems:
- ${SPECIFIC_SCHEMA_ISSUES}
- Normalization issues: ${NORMALIZATION_PROBLEMS}
- Performance bottlenecks: ${SCHEMA_PERFORMANCE_ISSUES}
- Data integrity concerns: ${INTEGRITY_ISSUES}

Refactoring goals:
- ${SCHEMA_IMPROVEMENT_GOALS}
- Migration strategy: ${MIGRATION_APPROACH}
- Backward compatibility: ${COMPATIBILITY_REQUIREMENTS}

Please examine the current schema and propose a refactoring approach that addresses the issues while maintaining data integrity and system functionality.
```

## API Refactoring

### Improve API Design
```
I need to refactor the ${API_NAME} API to improve ${API_ISSUES}.

Current API problems:
- ${SPECIFIC_API_ISSUES}
- Inconsistent endpoints: ${INCONSISTENCY_ISSUES}
- Poor response formats: ${RESPONSE_ISSUES}
- Missing error handling: ${ERROR_HANDLING_GAPS}

Refactoring goals:
- Consistent API design: ${CONSISTENCY_GOALS}
- Better response formats: ${RESPONSE_IMPROVEMENTS}
- Comprehensive error handling: ${ERROR_HANDLING_GOALS}
- Backward compatibility: ${COMPATIBILITY_STRATEGY}

Please examine the current API and refactor it to follow REST/GraphQL best practices while maintaining backward compatibility where possible.
```

### Add API Versioning
```
I need to add versioning to the ${API_NAME} API to support ${VERSIONING_GOALS}.

Current situation:
- Existing API endpoints: ${CURRENT_ENDPOINTS}
- Breaking changes needed: ${BREAKING_CHANGES}
- Backward compatibility requirements: ${COMPATIBILITY_NEEDS}

Versioning strategy:
- Versioning approach: ${VERSIONING_APPROACH}
- Version migration plan: ${MIGRATION_PLAN}
- Deprecation timeline: ${DEPRECATION_TIMELINE}

Please examine the current API and implement a versioning strategy that allows for future evolution while maintaining support for existing clients.
```

## Testing Refactoring

### Improve Test Coverage
```
I need to improve test coverage for ${COMPONENT_NAME} which currently has ${CURRENT_COVERAGE}% coverage.

Coverage gaps:
- Untested functions: ${UNTESTED_FUNCTIONS}
- Missing edge cases: ${MISSING_EDGE_CASES}
- Insufficient error testing: ${ERROR_TEST_GAPS}
- Integration test gaps: ${INTEGRATION_GAPS}

Testing goals:
- Target coverage: ${TARGET_COVERAGE}%
- Test quality improvements: ${QUALITY_GOALS}
- Test maintainability: ${MAINTAINABILITY_GOALS}

Please examine the current tests and add comprehensive test coverage that follows the project's testing patterns.
```

### Refactor Test Structure
```
I need to refactor the test structure for ${TEST_AREA} to improve ${TEST_ISSUES}.

Current test problems:
- ${SPECIFIC_TEST_ISSUES}
- Test duplication: ${DUPLICATION_ISSUES}
- Poor test organization: ${ORGANIZATION_ISSUES}
- Brittle tests: ${BRITTLENESS_ISSUES}

Refactoring goals:
- Better test organization: ${ORGANIZATION_GOALS}
- Reduce test duplication: ${DUPLICATION_REDUCTION}
- Improve test maintainability: ${MAINTAINABILITY_IMPROVEMENTS}
- Follow testing best practices: ${BEST_PRACTICES}

Please examine the current test structure and refactor it to create more maintainable and reliable tests.
```

## Variables Reference

Common variables used in these prompts:

- `${SOURCE_FILE}` - File containing code to refactor
- `${LINE_RANGE}` - Specific lines to refactor (e.g., "25-45")
- `${COMPONENT_NAME}` - Name of component being refactored
- `${FUNCTION_NAME}` - Name of function being refactored
- `${DESIGN_PATTERN}` - Design pattern to implement
- `${PERFORMANCE_ISSUES}` - Specific performance problems
- `${CURRENT_APPROACH}` - How it's currently implemented
- `${TARGET_IMPROVEMENT}` - Goal for improvement
- `${REFACTORING_GOALS}` - What the refactoring should achieve

## Usage Examples

### Example 1: Extract a Complex Function
```
I need to extract a function from src/services/userService.js at lines 45-75.

Extraction details:
- Current code location: src/services/userService.js:45-75
- Purpose of extracted function: Validate and format user profile data
- Proposed function name: validateAndFormatProfile
- Parameters needed: profileData, validationRules
- Return type: { isValid: boolean, formattedData: object, errors: string[] }

Requirements:
- Follow the existing code style and naming conventions
- Maintain all existing functionality
- Update all call sites
- Add appropriate documentation
- Include proper error handling if needed

Please examine the code and create a clean extraction that improves code organization while maintaining all existing behavior.
```

This approach helps Claude Code understand your refactoring goals and implement improvements that align with your project's architecture and standards.
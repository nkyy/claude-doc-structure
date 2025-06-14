# Backend Development Prompts for Claude Code

Collection of reusable prompts for common backend development tasks. Replace variables in `${VARIABLE_NAME}` format with your specific values.

## API Development

### Create REST API Endpoint
```
I need to create a new REST API endpoint for ${ENDPOINT_PURPOSE}.

Endpoint details:
- Path: ${API_PATH}
- Method: ${HTTP_METHOD}
- Expected request: ${REQUEST_FORMAT}
- Expected response: ${RESPONSE_FORMAT}
- Authentication: ${AUTH_REQUIREMENTS}

Requirements:
- Follow the existing API patterns in ${API_DIRECTORY}
- Include proper input validation
- Add appropriate error handling
- Follow the project's response format standards
- Include proper HTTP status codes
- Add rate limiting if needed

Please examine the existing API endpoints to understand the current patterns and create an endpoint that fits the established architecture.
```

### Add GraphQL Resolver
```
I need to create a GraphQL resolver for ${RESOLVER_PURPOSE}.

Schema details:
- Type: ${GRAPHQL_TYPE}
- Fields: ${FIELD_DEFINITIONS}
- Arguments: ${RESOLVER_ARGUMENTS}
- Return type: ${RETURN_TYPE}

Requirements:
- Follow the existing GraphQL patterns in ${SCHEMA_DIRECTORY}
- Include proper field resolvers
- Add authentication/authorization checks
- Handle errors appropriately
- Follow the project's GraphQL conventions

Please examine the existing resolvers and schema to understand the current patterns and create a resolver that integrates seamlessly.
```

### API Authentication & Authorization
```
I need to implement ${AUTH_TYPE} authentication for the ${ENDPOINT_NAME} endpoint.

Authentication requirements:
- ${AUTH_DETAILS}
- Permissions needed: ${PERMISSION_REQUIREMENTS}
- Token validation: ${TOKEN_VALIDATION}
- Error handling: ${AUTH_ERROR_HANDLING}

Current auth system: ${CURRENT_AUTH_SYSTEM}
Auth middleware location: ${AUTH_MIDDLEWARE_PATH}

Please examine the existing authentication patterns and implement authentication that follows the established security approach.
```

## Database Operations

### Create Database Model
```
I need to create a database model for ${MODEL_NAME} with the following structure:

Fields:
${FIELD_DEFINITIONS}

Relationships:
${RELATIONSHIP_DEFINITIONS}

Requirements:
- Follow the existing model patterns in ${MODELS_DIRECTORY}
- Include proper data validation
- Add appropriate indexes
- Follow the project's naming conventions
- Include timestamps if standard in the project

ORM/Database: ${DATABASE_SYSTEM}

Please examine the existing models to understand the current patterns and create a model that fits the established data architecture.
```

### Create Database Migration
```
I need to create a database migration for ${MIGRATION_PURPOSE}.

Migration details:
- ${MIGRATION_ACTIONS}
- Affected tables: ${TABLE_NAMES}
- Data transformations: ${DATA_CHANGES}
- Rollback strategy: ${ROLLBACK_PLAN}

Migration system: ${MIGRATION_SYSTEM}
Migration directory: ${MIGRATIONS_DIRECTORY}

Please examine the existing migrations to understand the current patterns and create a migration that follows the established approach.
```

### Add Database Queries
```
I need to implement database queries for ${QUERY_PURPOSE}.

Query requirements:
- ${QUERY_SPECIFICATIONS}
- Performance considerations: ${PERFORMANCE_REQUIREMENTS}
- Pagination: ${PAGINATION_NEEDS}
- Filtering: ${FILTER_OPTIONS}
- Sorting: ${SORT_OPTIONS}

Database/ORM: ${DATABASE_SYSTEM}
Query location: ${QUERY_DIRECTORY}

Please examine the existing query patterns and implement queries that follow the established database access patterns.
```

## Service Layer

### Create Service Class
```
I need to create a service class for ${SERVICE_PURPOSE}.

Service responsibilities:
- ${SERVICE_FUNCTIONS}
- External integrations: ${EXTERNAL_APIS}
- Business logic: ${BUSINESS_RULES}
- Error handling: ${ERROR_SCENARIOS}

Requirements:
- Follow the existing service patterns in ${SERVICES_DIRECTORY}
- Include proper dependency injection
- Add comprehensive error handling
- Follow the project's architecture patterns
- Include proper logging

Please examine the existing services to understand the current patterns and create a service that fits the established architecture.
```

### Add Business Logic
```
I need to implement business logic for ${BUSINESS_FEATURE}.

Business requirements:
- ${BUSINESS_RULES}
- Validation rules: ${VALIDATION_REQUIREMENTS}
- Workflow steps: ${WORKFLOW_STEPS}
- Edge cases: ${EDGE_CASES}

Integration points:
- ${INTEGRATION_REQUIREMENTS}

Please examine the existing business logic patterns and implement logic that follows the established approach while meeting the specified requirements.
```

## External API Integration

### Add Third-Party API Integration
```
I need to integrate with ${API_NAME} for ${INTEGRATION_PURPOSE}.

API details:
- Base URL: ${API_BASE_URL}
- Authentication: ${API_AUTH_METHOD}
- Endpoints needed: ${API_ENDPOINTS}
- Rate limits: ${RATE_LIMIT_INFO}
- Error handling: ${API_ERROR_HANDLING}

Requirements:
- Follow the existing API integration patterns
- Include proper error handling and retries
- Add request/response logging
- Handle rate limiting appropriately
- Include proper timeout handling

Please examine the existing external API integrations to understand the current patterns and create an integration that follows the established approach.
```

### Add Webhook Handler
```
I need to create a webhook handler for ${WEBHOOK_SOURCE} that processes ${WEBHOOK_EVENTS}.

Webhook details:
- Expected payload: ${WEBHOOK_PAYLOAD}
- Verification method: ${WEBHOOK_VERIFICATION}
- Processing requirements: ${PROCESSING_LOGIC}
- Response format: ${RESPONSE_FORMAT}

Requirements:
- Follow the existing webhook patterns in ${WEBHOOKS_DIRECTORY}
- Include proper payload validation
- Add idempotency handling
- Include comprehensive error handling
- Add proper logging and monitoring

Please examine the existing webhook handlers to understand the current patterns and create a handler that fits the established architecture.
```

## Background Jobs & Queues

### Create Background Job
```
I need to create a background job for ${JOB_PURPOSE}.

Job details:
- Job parameters: ${JOB_PARAMETERS}
- Processing logic: ${JOB_LOGIC}
- Retry strategy: ${RETRY_CONFIGURATION}
- Error handling: ${ERROR_HANDLING}
- Scheduling: ${SCHEDULE_REQUIREMENTS}

Queue system: ${QUEUE_SYSTEM}
Job directory: ${JOBS_DIRECTORY}

Requirements:
- Follow the existing job patterns
- Include proper error handling and retries
- Add comprehensive logging
- Handle job failures gracefully
- Include progress tracking if needed

Please examine the existing background jobs to understand the current patterns and create a job that follows the established approach.
```

### Add Queue Configuration
```
I need to configure queues for ${QUEUE_PURPOSE}.

Queue requirements:
- Queue types: ${QUEUE_TYPES}
- Priority levels: ${PRIORITY_CONFIGURATION}
- Concurrency settings: ${CONCURRENCY_REQUIREMENTS}
- Retry policies: ${RETRY_POLICIES}
- Dead letter handling: ${DEAD_LETTER_CONFIGURATION}

Queue system: ${QUEUE_SYSTEM}

Please examine the existing queue configurations and create queue setup that follows the established patterns while meeting the specified requirements.
```

## Testing

### Add Unit Tests
```
I need comprehensive unit tests for ${COMPONENT_NAME} at ${FILE_PATH}.

Testing requirements:
- Test all public methods/functions
- Test error scenarios and edge cases
- Test integration points with mocks
- Follow the existing testing patterns
- Achieve ${COVERAGE_TARGET}% code coverage

Testing framework: ${TEST_FRAMEWORK}
Test directory: ${TEST_DIRECTORY}

Please examine the existing test files to understand the testing approach and create thorough tests that match the project's testing standards.
```

### Add Integration Tests
```
I need integration tests for ${FEATURE_NAME} that test ${INTEGRATION_SCOPE}.

Test scenarios:
- ${TEST_SCENARIOS}
- Database interactions: ${DATABASE_TESTS}
- External API calls: ${API_TESTS}
- End-to-end workflows: ${WORKFLOW_TESTS}

Testing setup:
- Test database: ${TEST_DB_CONFIG}
- Test environment: ${TEST_ENV_CONFIG}
- Mock services: ${MOCK_REQUIREMENTS}

Please examine the existing integration tests to understand the current patterns and create comprehensive tests that follow the established approach.
```

## Error Handling & Logging

### Implement Error Handling
```
I need to implement comprehensive error handling for ${COMPONENT_NAME}.

Error handling requirements:
- Handle ${ERROR_TYPES}
- Error response format: ${ERROR_FORMAT}
- Logging requirements: ${LOGGING_SPECIFICATIONS}
- Error recovery: ${RECOVERY_STRATEGIES}
- Error reporting: ${ERROR_REPORTING}

Current error handling approach: ${CURRENT_ERROR_SYSTEM}

Please examine the existing error handling patterns and implement error handling that follows the established approach while meeting the specified requirements.
```

### Add Structured Logging
```
I need to add structured logging to ${COMPONENT_NAME} for ${LOGGING_PURPOSE}.

Logging requirements:
- Log levels: ${LOG_LEVELS}
- Log format: ${LOG_FORMAT}
- Contextual information: ${LOG_CONTEXT}
- Performance metrics: ${METRICS_LOGGING}
- Security considerations: ${SECURITY_LOGGING}

Logging system: ${LOGGING_SYSTEM}

Please examine the existing logging patterns and implement logging that follows the established approach while providing the required visibility.
```

## Security

### Add Input Validation
```
I need to implement input validation for ${ENDPOINT_NAME} endpoint.

Validation requirements:
- Input fields: ${INPUT_FIELDS}
- Validation rules: ${VALIDATION_RULES}
- Sanitization needs: ${SANITIZATION_REQUIREMENTS}
- Error messaging: ${VALIDATION_ERROR_FORMAT}

Current validation approach: ${CURRENT_VALIDATION_SYSTEM}

Please examine the existing validation patterns and implement validation that follows the established security approach.
```

### Implement Security Middleware
```
I need to implement security middleware for ${SECURITY_PURPOSE}.

Security requirements:
- ${SECURITY_FEATURES}
- Headers to set: ${SECURITY_HEADERS}
- Rate limiting: ${RATE_LIMIT_CONFIG}
- Input sanitization: ${SANITIZATION_RULES}

Current security setup: ${CURRENT_SECURITY_MIDDLEWARE}

Please examine the existing security middleware and implement security measures that integrate with the established security approach.
```

## Performance Optimization

### Add Caching
```
I need to implement caching for ${CACHING_TARGET} to improve performance.

Caching requirements:
- Cache keys: ${CACHE_KEY_STRATEGY}
- TTL settings: ${CACHE_TTL}
- Cache invalidation: ${INVALIDATION_STRATEGY}
- Cache storage: ${CACHE_STORAGE_SYSTEM}

Current caching approach: ${CURRENT_CACHING_SYSTEM}

Please examine the existing caching patterns and implement caching that follows the established approach while meeting the performance requirements.
```

### Database Query Optimization
```
I need to optimize database queries for ${PERFORMANCE_TARGET}.

Current performance issues: ${PERFORMANCE_PROBLEMS}
Query locations: ${QUERY_FILES}

Optimization goals:
- ${OPTIMIZATION_TARGETS}
- Indexing strategy: ${INDEX_REQUIREMENTS}
- Query restructuring: ${QUERY_CHANGES}

Please analyze the current queries and implement optimizations that follow the established database patterns while achieving the performance goals.
```

## Deployment & DevOps

### Add Health Checks
```
I need to implement health checks for ${SERVICE_NAME}.

Health check requirements:
- Database connectivity: ${DB_HEALTH_CHECK}
- External services: ${EXTERNAL_SERVICE_CHECKS}
- System resources: ${RESOURCE_CHECKS}
- Custom checks: ${CUSTOM_HEALTH_CHECKS}

Health check endpoint: ${HEALTH_ENDPOINT}

Please examine any existing health checks and implement comprehensive health monitoring that follows the established patterns.
```

### Add Environment Configuration
```
I need to add environment configuration for ${CONFIGURATION_PURPOSE}.

Configuration requirements:
- Environment variables: ${ENV_VARIABLES}
- Default values: ${DEFAULT_VALUES}
- Validation: ${CONFIG_VALIDATION}
- Documentation: ${CONFIG_DOCUMENTATION}

Current config system: ${CURRENT_CONFIG_SYSTEM}

Please examine the existing configuration patterns and implement configuration that follows the established approach.
```

## Variables Reference

Common variables used in these prompts:

- `${ENDPOINT_NAME}` - Name of the API endpoint
- `${API_PATH}` - URL path for the endpoint
- `${HTTP_METHOD}` - GET, POST, PUT, DELETE, etc.
- `${MODEL_NAME}` - Database model name
- `${SERVICE_NAME}` - Name of the service
- `${DATABASE_SYSTEM}` - PostgreSQL, MySQL, MongoDB, etc.
- `${QUEUE_SYSTEM}` - Redis, RabbitMQ, AWS SQS, etc.
- `${TEST_FRAMEWORK}` - Jest, Mocha, PyTest, etc.
- `${AUTH_TYPE}` - JWT, OAuth, API Key, etc.
- `${INTEGRATION_PURPOSE}` - Purpose of the integration
- `${BUSINESS_FEATURE}` - Name of the business feature

## Usage Examples

### Example 1: Create User Registration Endpoint
```
I need to create a new REST API endpoint for user registration.

Endpoint details:
- Path: /api/v1/auth/register
- Method: POST
- Expected request: { "email": "string", "password": "string", "name": "string" }
- Expected response: { "user": {...}, "token": "string" }
- Authentication: None (public endpoint)

Requirements:
- Follow the existing API patterns in src/controllers/
- Include proper input validation
- Add appropriate error handling
- Follow the project's response format standards
- Include proper HTTP status codes
- Add rate limiting if needed

Please examine the existing API endpoints to understand the current patterns and create an endpoint that fits the established architecture.
```

This structured approach helps Claude Code understand your backend architecture and create code that integrates seamlessly with your existing patterns.
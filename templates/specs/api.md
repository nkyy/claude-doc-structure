# API Documentation Template

This template provides a structure for documenting REST APIs and GraphQL schemas in a way that's optimized for Claude Code understanding.

## API Overview

### Purpose
[Describe what this API accomplishes and its role in the system]

### Base Configuration
- **Base URL:** `https://api.example.com/v1`
- **Authentication:** Bearer token (JWT)
- **Content Type:** `application/json`
- **Rate Limiting:** 1000 requests/hour per API key

### Architecture
[Describe the API architecture, middleware stack, and key design decisions]

## Authentication

### Token-Based Authentication
```http
Authorization: Bearer <jwt_token>
```

**Implementation Location:** `src/middleware/auth.js:15-35`

### API Key Authentication (Admin endpoints)
```http
X-API-Key: <api_key>
```

**Implementation Location:** `src/middleware/admin-auth.js:8-20`

## Core Resources

### Users Resource

**Base Path:** `/users`

#### List Users
```http
GET /users?page=1&limit=20&sort=created_at
```

**Response:**
```json
{
  "data": [
    {
      "id": "usr_123",
      "email": "user@example.com",
      "created_at": "2024-01-15T10:30:00Z",
      "profile": {
        "name": "John Doe",
        "avatar_url": "https://cdn.example.com/avatars/123.jpg"
      }
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 150,
    "has_next": true
  }
}
```

**Implementation:** `src/controllers/users.js:25-60`

#### Create User
```http
POST /users
Content-Type: application/json

{
  "email": "newuser@example.com",
  "password": "securePassword123",
  "profile": {
    "name": "New User"
  }
}
```

**Validation Rules:**
- Email: Required, valid email format, unique
- Password: Required, minimum 8 characters, must contain number and special character
- Name: Required, 2-50 characters

**Implementation:** `src/controllers/users.js:85-120`

#### Get User by ID
```http
GET /users/{id}
```

**Path Parameters:**
- `id` (string, required): User ID in format `usr_xxx`

**Implementation:** `src/controllers/users.js:135-155`

#### Update User
```http
PUT /users/{id}
Content-Type: application/json

{
  "profile": {
    "name": "Updated Name",
    "bio": "New bio text"
  }
}
```

**Implementation:** `src/controllers/users.js:170-200`

#### Delete User
```http
DELETE /users/{id}
```

**Implementation:** `src/controllers/users.js:215-235`
**Soft Delete Logic:** `src/models/user.js:45-55`

## Error Handling

### Standard Error Format
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "The provided data is invalid",
    "details": [
      {
        "field": "email",
        "message": "Email address is already in use"
      }
    ],
    "request_id": "req_789"
  }
}
```

### Common Error Codes
- `VALIDATION_ERROR` (400): Invalid input data
- `UNAUTHORIZED` (401): Missing or invalid authentication
- `FORBIDDEN` (403): Insufficient permissions
- `NOT_FOUND` (404): Resource not found
- `RATE_LIMITED` (429): Too many requests
- `INTERNAL_ERROR` (500): Server error

**Error Handler Implementation:** `src/middleware/error-handler.js:10-45`

## Data Models

### User Model
```typescript
interface User {
  id: string;           // Format: usr_xxxxx
  email: string;        // Unique, validated
  password_hash: string; // Bcrypt hash
  profile: {
    name: string;
    bio?: string;
    avatar_url?: string;
  };
  created_at: string;   // ISO 8601 timestamp
  updated_at: string;   // ISO 8601 timestamp
  deleted_at?: string;  // Soft delete timestamp
}
```

**Database Schema:** `migrations/001_create_users.sql:15-30`
**Model Implementation:** `src/models/user.js:1-80`

## Middleware Stack

1. **Request Logging** (`src/middleware/logging.js:5-20`)
2. **CORS Handling** (`src/middleware/cors.js:8-25`)
3. **Rate Limiting** (`src/middleware/rate-limit.js:12-30`)
4. **Authentication** (`src/middleware/auth.js:15-35`)
5. **Request Validation** (`src/middleware/validation.js:20-50`)
6. **Error Handling** (`src/middleware/error-handler.js:10-45`)

## Database Integration

### Connection Configuration
**File:** `src/config/database.js:5-25`

### Query Patterns
- **List Queries:** Use pagination with cursor-based approach
- **Search Queries:** Full-text search with PostgreSQL
- **Joins:** Minimize N+1 queries with eager loading

### Migration Strategy
**Directory:** `migrations/`
**Runner:** `src/scripts/migrate.js:1-50`

## Testing

### Unit Tests
**Location:** `tests/unit/controllers/`
**Framework:** Jest
**Coverage:** 90%+ for controller methods

### Integration Tests
**Location:** `tests/integration/api/`
**Framework:** Supertest + Jest
**Database:** Test database with fixtures

### API Documentation Tests
**Location:** `tests/docs/`
**Purpose:** Validate that documentation examples work

**Test Runner:** `package.json:scripts.test-api`

## Deployment

### Environment Configuration
- **Development:** `config/development.json`
- **Staging:** `config/staging.json`  
- **Production:** `config/production.json`

### Health Checks
```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "version": "1.2.3",
  "database": "connected",
  "cache": "connected"
}
```

**Implementation:** `src/controllers/health.js:5-25`

## Performance Considerations

### Caching Strategy
- **Redis Cache:** `src/services/cache.js:10-40`
- **Cache Keys:** Follow pattern `api:resource:id:version`
- **TTL:** 300 seconds for user data, 3600 for static data

### Database Optimization
- **Indexes:** Defined in `migrations/` files
- **Query Optimization:** Use EXPLAIN ANALYZE for slow queries
- **Connection Pooling:** Max 20 connections per instance

## Security

### Input Validation
**Implementation:** `src/middleware/validation.js:20-50`
**Rules:** Defined in `src/schemas/` directory

### SQL Injection Prevention
- Use parameterized queries only
- **ORM:** Sequelize with prepared statements
- **Validation:** All inputs validated before database queries

### Rate Limiting
**Implementation:** `src/middleware/rate-limit.js:12-30`
**Limits:** 
- Anonymous: 100 requests/hour
- Authenticated: 1000 requests/hour
- Admin: 5000 requests/hour

## Monitoring

### Logging
**Implementation:** `src/utils/logger.js:5-30`
**Format:** Structured JSON logs
**Levels:** ERROR, WARN, INFO, DEBUG

### Metrics
**Collection:** Prometheus metrics
**Dashboards:** Grafana dashboards in `monitoring/`
**Alerts:** Defined in `alerts/api-alerts.yml`

## Related Documentation

- Database Schema: `specs/database.md`
- Frontend Integration: `specs/frontend-api.md`
- Deployment Guide: `docs/deployment.md`
- Security Guidelines: `docs/security.md`

## Changelog

### Version 1.2.3 (2024-01-15)
- Added user profile photo upload
- Improved error message clarity
- Fixed pagination edge case

### Version 1.2.2 (2024-01-10)
- Added rate limiting per endpoint
- Enhanced authentication middleware
- Updated API documentation

### Version 1.2.1 (2024-01-05)
- Fixed user deletion bug
- Added request validation
- Improved database performance
# API Specification - {name}

## Overview

**Base URL:** `{base_url}`  
**Version:** `{api_version}`  
**Authentication:** `{auth_type}`

## Authentication

### {auth_method}
```http
{auth_header}: {auth_value}
```

**Example:**
```bash
curl -H "{auth_header}: {auth_token}" {base_url}/api/endpoint
```

## Endpoints

### {endpoint_category}

#### GET {endpoint_path}
**Description:** {endpoint_description}

**Parameters:**
- `{param1}` ({param1_type}, {param1_required}): {param1_description}
- `{param2}` ({param2_type}, {param2_required}): {param2_description}
- `{param3}` ({param3_type}, {param3_required}): {param3_description}

**Query Parameters:**
- `limit` (integer, optional): Number of items to return (default: 10, max: 100)
- `offset` (integer, optional): Number of items to skip (default: 0)

**Response:**
```json
{
  "status": "success",
  "data": {
    "{response_field1}": "{response_value1}",
    "{response_field2}": {response_value2},
    "{response_field3}": [
      {
        "{nested_field1}": "{nested_value1}",
        "{nested_field2}": {nested_value2}
      }
    ]
  },
  "meta": {
    "total": 100,
    "count": 10,
    "limit": 10,
    "offset": 0
  }
}
```

**Error Response:**
```json
{
  "status": "error",
  "error": {
    "code": "{error_code}",
    "message": "{error_message}",
    "details": "{error_details}"
  }
}
```

**Status Codes:**
- `200 OK`: Success
- `400 Bad Request`: Invalid parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Access denied
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

**Example:**
```bash
curl -X GET "{base_url}{endpoint_path}?{param1}={example_value1}" \
  -H "{auth_header}: {auth_token}" \
  -H "Content-Type: application/json"
```

#### POST {endpoint_path}
**Description:** {post_endpoint_description}

**Request Body:**
```json
{
  "{request_field1}": "{request_value1}",
  "{request_field2}": {request_value2},
  "{request_field3}": {
    "{nested_request_field}": "{nested_request_value}"
  }
}
```

**Validation Rules:**
- `{request_field1}`: {validation_rule1}
- `{request_field2}`: {validation_rule2}
- `{request_field3}`: {validation_rule3}

**Response:**
```json
{
  "status": "success",
  "data": {
    "id": "{generated_id}",
    "{created_field1}": "{created_value1}",
    "{created_field2}": {created_value2},
    "created_at": "{timestamp}",
    "updated_at": "{timestamp}"
  }
}
```

**Example:**
```bash
curl -X POST "{base_url}{endpoint_path}" \
  -H "{auth_header}: {auth_token}" \
  -H "Content-Type: application/json" \
  -d '{
    "{request_field1}": "{example_request_value1}",
    "{request_field2}": {example_request_value2}
  }'
```

#### PUT {endpoint_path}/{id}
**Description:** {put_endpoint_description}

**Path Parameters:**
- `id` (string, required): {id_description}

**Request Body:**
```json
{
  "{update_field1}": "{update_value1}",
  "{update_field2}": {update_value2}
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "id": "{id}",
    "{updated_field1}": "{updated_value1}",
    "{updated_field2}": {updated_value2},
    "updated_at": "{timestamp}"
  }
}
```

**Example:**
```bash
curl -X PUT "{base_url}{endpoint_path}/{example_id}" \
  -H "{auth_header}: {auth_token}" \
  -H "Content-Type: application/json" \
  -d '{
    "{update_field1}": "{example_update_value1}"
  }'
```

#### DELETE {endpoint_path}/{id}
**Description:** {delete_endpoint_description}

**Path Parameters:**
- `id` (string, required): {delete_id_description}

**Response:**
```json
{
  "status": "success",
  "message": "{delete_success_message}"
}
```

**Example:**
```bash
curl -X DELETE "{base_url}{endpoint_path}/{example_id}" \
  -H "{auth_header}: {auth_token}"
```

## Data Models

### {model_name}
```json
{
  "id": "string",
  "{model_field1}": "string",
  "{model_field2}": "integer",
  "{model_field3}": "boolean",
  "{model_field4}": "array",
  "{model_field5}": "object",
  "created_at": "string (ISO 8601 datetime)",
  "updated_at": "string (ISO 8601 datetime)"
}
```

**Field Descriptions:**
- `id`: Unique identifier
- `{model_field1}`: {field1_description}
- `{model_field2}`: {field2_description}
- `{model_field3}`: {field3_description}
- `{model_field4}`: {field4_description}
- `{model_field5}`: {field5_description}

## Error Handling

### Standard Error Format
```json
{
  "status": "error",
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message",
    "details": "Additional error details",
    "field_errors": {
      "field_name": ["Field specific error message"]
    }
  }
}
```

### Common Error Codes
- `VALIDATION_ERROR`: Request validation failed
- `AUTHENTICATION_ERROR`: Authentication failed
- `AUTHORIZATION_ERROR`: Access denied
- `NOT_FOUND`: Resource not found
- `RATE_LIMIT_EXCEEDED`: Rate limit exceeded
- `INTERNAL_ERROR`: Internal server error

## Rate Limiting

**Rate Limit:** {rate_limit_requests} requests per {rate_limit_window}

**Headers:**
- `X-RateLimit-Limit`: Maximum requests allowed
- `X-RateLimit-Remaining`: Remaining requests in current window
- `X-RateLimit-Reset`: Unix timestamp when the rate limit resets

**Example Response Headers:**
```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1609459200
```

## Pagination

**Query Parameters:**
- `limit`: Number of items per page (default: 10, max: 100)
- `offset`: Number of items to skip (default: 0)
- `page`: Page number (alternative to offset)

**Response Meta:**
```json
{
  "meta": {
    "total": 100,
    "count": 10,
    "limit": 10,
    "offset": 0,
    "page": 1,
    "pages": 10,
    "has_next": true,
    "has_prev": false
  }
}
```

## Filtering and Sorting

**Filtering:**
- `filter[{field_name}]`: Filter by field value
- `filter[{field_name}][eq]`: Exact match
- `filter[{field_name}][like]`: Partial match
- `filter[{field_name}][gt]`: Greater than
- `filter[{field_name}][lt]`: Less than

**Sorting:**
- `sort`: Field to sort by
- `order`: Sort direction (`asc` or `desc`)

**Example:**
```
GET /api/items?filter[status]=active&sort=created_at&order=desc
```

## Webhook Events

### {webhook_event_name}
**Trigger:** {webhook_trigger}

**Payload:**
```json
{
  "event": "{webhook_event_name}",
  "timestamp": "{iso_timestamp}",
  "data": {
    "{webhook_field1}": "{webhook_value1}",
    "{webhook_field2}": {webhook_value2}
  }
}
```

## SDKs and Libraries

### {language1}
```{lang1_syntax}
# Installation
{lang1_install_command}

# Usage
{lang1_usage_example}
```

### {language2}
```{lang2_syntax}
// Installation
{lang2_install_command}

// Usage
{lang2_usage_example}
```

## Changelog

### Version {version1}
- {change1}
- {change2}
- {change3}

### Version {version2}
- {change4}
- {change5}
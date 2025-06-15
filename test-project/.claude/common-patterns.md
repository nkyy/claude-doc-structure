# Common Command Patterns

## Development Workflow

### Quick Development Cycle
```bash
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
```

### Code Quality Checks
```bash
# Linting and formatting
npm run lint
npm run format
# or
flake8 . && black .

# Type checking
npm run type-check
# or
mypy .
```

## Testing Patterns

### Test Execution
```bash
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
```

## API Development

### API Testing
```bash
# Test API endpoints
curl -X GET "http://localhost:3000/api/users"
curl -X POST "http://localhost:3000/api/users" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'

# Using HTTPie (alternative to curl)
http GET localhost:3000/api/users
http POST localhost:3000/api/users name="John Doe" email="john@example.com"
```

## Environment Management

### Environment Variables
```bash
# Load environment variables
source .env
# or
export $(cat .env | xargs)

# Check environment configuration
npm run env:check
# or
python -c "import os; print(os.environ)"
```

## IDE Integration Patterns

### Command Line Shortcuts
```bash
# Add to shell profile (.bashrc, .zshrc)
alias dev='npm run dev'
alias test='npm run test'
alias build='npm run build'

# Project-specific shortcuts
alias test-app-dev='cd ~/projects/test-app && npm run dev'
alias test-app-test='cd ~/projects/test-app && npm test'
```
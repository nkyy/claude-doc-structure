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
# or
go run main.go

# Run tests during development
npm run test:watch
# or
pytest --watch
# or
go test -v ./...

# Build for production
npm run build
# or
python -m build
# or
go build -o bin/app
```

### Code Quality Checks
```bash
# Linting and formatting
npm run lint
npm run format
# or
flake8 . && black .
# or
golangci-lint run && gofmt -w .

# Type checking
npm run type-check
# or
mypy .
# or
go vet ./...

# Security scanning
npm audit
# or
safety check
# or
gosec ./...
```

## Testing Patterns

### Test Execution
```bash
# Run all tests
npm test
# or
pytest
# or
go test ./...

# Run specific test file
npm test -- ComponentName.test.js
# or
pytest tests/test_specific.py
# or
go test ./pkg/specific

# Run tests with coverage
npm run test:coverage
# or
pytest --cov=src
# or
go test -cover ./...
```

### Test Data Management
```bash
# Reset test database
npm run db:reset:test
# or
python manage.py migrate --run-syncdb
# or
go run scripts/reset-test-db.go

# Seed test data
npm run db:seed:test
# or
python manage.py loaddata fixtures/test_data.json
# or
go run scripts/seed-test-data.go
```

## Database Operations

### Migrations
```bash
# Create new migration
npm run db:migrate:create
# or
python manage.py makemigrations
# or
migrate create -ext sql -dir migrations migration_name

# Run migrations
npm run db:migrate
# or
python manage.py migrate
# or
migrate -path migrations -database "connection_string" up
```

### Database Maintenance
```bash
# Backup database
npm run db:backup
# or
python manage.py dumpdata > backup.json
# or
pg_dump dbname > backup.sql

# Restore database
npm run db:restore backup.sql
# or
python manage.py loaddata backup.json
# or
psql dbname < backup.sql
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

# Using Postman or Insomnia collections
postman collection run api-tests.json
insomnia export spec.json
```

### API Documentation
```bash
# Generate API documentation
npm run docs:api
# or
python manage.py generate_docs
# or
swag init # for Swagger with Go

# Serve documentation locally
npm run docs:serve
# or
mkdocs serve
# or
redoc-cli serve openapi.json
```

## Deployment Patterns

### Local Deployment
```bash
# Build for production
npm run build
# or
python setup.py build
# or
go build -ldflags="-s -w" -o bin/app

# Start production server
npm start
# or
gunicorn app:app
# or
./bin/app
```

### Container Operations
```bash
# Build Docker image
docker build -t {project_name}:latest .

# Run container locally
docker run -p 3000:3000 {project_name}:latest

# Docker Compose operations
docker-compose up -d
docker-compose down
docker-compose logs -f
```

### Cloud Deployment
```bash
# Deploy to Heroku
git push heroku main
heroku logs --tail
heroku ps:scale web=1

# Deploy to AWS/GCP/Azure
aws ecs update-service --cluster prod --service {project_name}
gcloud app deploy
az webapp deployment source config-zip
```

## Monitoring & Debugging

### Log Analysis
```bash
# View application logs
npm run logs
# or
tail -f logs/app.log
# or
journalctl -u {project_name} -f

# Search logs for errors
grep -i error logs/app.log
# or
journalctl -u {project_name} | grep ERROR
# or
kubectl logs -f deployment/{project_name}
```

### Performance Monitoring
```bash
# Monitor system resources
top
htop
# or
docker stats
# or
kubectl top pods

# Profile application performance
npm run profile
# or
python -m cProfile app.py
# or
go tool pprof http://localhost:6060/debug/pprof/profile
```

## Environment Management

### Environment Variables
```bash
# Load environment variables
source .env
# or
export $(cat .env | xargs)
# or
set -a && source .env && set +a

# Check environment configuration
npm run env:check
# or
python -c "import os; print(os.environ)"
# or
printenv | grep APP_
```

### Dependency Updates
```bash
# Update dependencies
npm update
npm audit fix
# or
pip install --upgrade -r requirements.txt
pip-audit
# or
go get -u ./...
go mod tidy
```

## IDE Integration Patterns

### VS Code Tasks
```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "dev",
      "type": "shell",
      "command": "npm run dev",
      "group": "build"
    },
    {
      "label": "test",
      "type": "shell",
      "command": "npm test",
      "group": "test"
    }
  ]
}
```

### Command Line Shortcuts
```bash
# Add to shell profile (.bashrc, .zshrc)
alias dev='npm run dev'
alias test='npm run test'
alias build='npm run build'
alias deploy='npm run deploy'

# Project-specific shortcuts
alias {project_name}-dev='cd ~/projects/{project_name} && npm run dev'
alias {project_name}-test='cd ~/projects/{project_name} && npm test'
```
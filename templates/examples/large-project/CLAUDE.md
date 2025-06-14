# E-Commerce Platform - Large Project Example

## Project Overview

A comprehensive multi-service e-commerce platform demonstrating the Claude Doc Structure approach for large, complex projects. This example shows how to organize documentation across multiple files while maintaining clear cross-references and context.

## Architecture & Technology Stack

**Frontend Services:**
- React 18.x + TypeScript (Customer Web App)
- Next.js 13.x (Admin Dashboard)
- React Native 0.72.x (Mobile App)
- Redux Toolkit (State Management)

**Backend Services:**
- Node.js + Express (API Gateway)
- Node.js + Fastify (Product Service)
- Python + FastAPI (Recommendation Service)
- Go + Gin (Payment Service)
- Java + Spring Boot (Order Service)

**Infrastructure:**
- PostgreSQL (Primary Database)
- Redis (Caching & Sessions)
- Elasticsearch (Search & Analytics)
- Docker + Kubernetes (Containerization)
- AWS (Cloud Infrastructure)
- RabbitMQ (Message Queue)

## Project Structure

```
ecommerce-platform/
├── docs/                          # Comprehensive documentation
│   ├── README.md                 # Project overview and getting started
│   ├── CLAUDE.md                 # This file - main context for Claude Code
│   ├── architecture/             # System architecture docs
│   │   ├── overview.md           # High-level architecture
│   │   ├── services.md           # Service breakdown
│   │   ├── data-flow.md         # Data flow diagrams
│   │   └── security.md          # Security architecture
│   ├── api/                      # API documentation  
│   │   ├── gateway.md           # API Gateway specifications
│   │   ├── products.md          # Product service API
│   │   ├── users.md             # User management API
│   │   ├── orders.md            # Order service API
│   │   └── payments.md          # Payment service API
│   ├── deployment/               # Deployment and infrastructure
│   │   ├── kubernetes.md        # K8s deployment specs
│   │   ├── aws-setup.md         # AWS infrastructure setup
│   │   ├── monitoring.md        # Monitoring and alerting
│   │   └── ci-cd.md             # CI/CD pipeline
│   └── development/              # Development workflows
│       ├── setup.md             # Local development setup
│       ├── testing.md           # Testing strategies
│       ├── contributing.md      # Contribution guidelines
│       └── troubleshooting.md   # Common issues and solutions
├── services/                     # Microservices
│   ├── api-gateway/             # API Gateway service
│   ├── product-service/         # Product management
│   ├── user-service/            # User authentication & profiles
│   ├── order-service/           # Order processing
│   ├── payment-service/         # Payment processing
│   ├── recommendation-service/   # ML-based recommendations
│   └── notification-service/    # Email/SMS notifications
├── web-apps/                    # Frontend applications
│   ├── customer-web/           # Customer-facing website
│   ├── admin-dashboard/        # Admin management portal
│   └── mobile-app/             # React Native mobile app
├── infrastructure/              # Infrastructure as Code
│   ├── terraform/              # AWS infrastructure
│   ├── kubernetes/             # K8s manifests
│   └── docker/                 # Docker configurations
├── scripts/                    # Automation scripts
│   ├── setup/                  # Environment setup
│   ├── deployment/            # Deployment automation
│   └── maintenance/           # Maintenance tasks
└── tests/                     # Cross-service integration tests
    ├── e2e/                   # End-to-end tests
    ├── performance/           # Load testing
    └── security/              # Security testing
```

## Current Development Focus

**Q2 2024: Microservices Migration**
- 🔄 Breaking down monolithic API into microservices
- 🔄 Implementing event-driven architecture with RabbitMQ
- ⏳ Setting up service mesh with Istio

**Q3 2024: Performance & Scale**
- ⏳ Implementing caching strategies across services
- ⏳ Database sharding for product catalog
- ⏳ CDN setup for static assets

**Q4 2024: Advanced Features**
- ⏳ ML-powered recommendation engine
- ⏳ Real-time inventory management
- ⏳ Mobile app feature parity

## Key Documentation Files

### Architecture Documentation
- `docs/architecture/overview.md:1` - System architecture and service interactions
- `docs/architecture/services.md:15` - Detailed service specifications
- `docs/architecture/data-flow.md:8` - Request/response flow diagrams
- `docs/architecture/security.md:22` - Authentication, authorization, and security measures

### API Documentation
- `docs/api/gateway.md:1` - API Gateway routing and middleware
- `docs/api/products.md:45` - Product CRUD operations and search
- `docs/api/users.md:30` - User registration, login, profile management
- `docs/api/orders.md:60` - Order lifecycle and state management
- `docs/api/payments.md:25` - Payment processing and webhook handling

### Service Implementation
- `services/api-gateway/src/routes/index.js:1` - Main routing configuration
- `services/product-service/src/controllers/productController.js:20` - Product business logic
- `services/user-service/src/auth/middleware.js:15` - Authentication middleware
- `services/order-service/src/workflows/orderProcessing.js:35` - Order workflow engine
- `services/payment-service/internal/handlers/stripe.go:10` - Stripe integration

### Frontend Applications
- `web-apps/customer-web/src/App.tsx:1` - Main customer application
- `web-apps/admin-dashboard/pages/index.tsx:1` - Admin dashboard home
- `web-apps/mobile-app/src/navigation/AppNavigator.tsx:20` - Mobile app navigation

### Infrastructure
- `infrastructure/terraform/main.tf:1` - AWS infrastructure definition
- `infrastructure/kubernetes/services/product-service.yaml:1` - Product service K8s manifest
- `scripts/deployment/deploy.sh:25` - Automated deployment script

## Service Dependencies & Communication

### Service Mesh Architecture
```
API Gateway → [Product Service, User Service, Order Service, Payment Service]
Order Service → Product Service (inventory check)
Order Service → Payment Service (payment processing)
Order Service → User Service (user validation)
Recommendation Service → Product Service (product data)
Notification Service ← All Services (events via RabbitMQ)
```

### Database Relationships
- **PostgreSQL Main**: Users, Orders, Products (normalized)
- **Redis**: Sessions, Cart Data, API Rate Limiting
- **Elasticsearch**: Product Search, Analytics, Logs

## Development Workflow

### Local Development Setup
```bash
# Start infrastructure services
docker-compose up -d postgres redis rabbitmq

# Start backend services
cd services && npm run dev:all

# Start frontend applications  
cd web-apps/customer-web && npm run dev
cd web-apps/admin-dashboard && npm run dev
```

### Testing Strategy
```bash
# Unit tests per service
npm run test:unit

# Integration tests
npm run test:integration

# End-to-end tests
npm run test:e2e

# Performance tests
npm run test:load
```

## Recent Architectural Decisions

**2024-06-01: Microservices Migration**
- Split monolithic API into domain-driven microservices
- Implemented eventual consistency with event sourcing
- Choice: RabbitMQ over Kafka for message queuing (smaller scale, simpler ops)

**2024-05-15: Database Strategy**
- PostgreSQL for transactional data (ACID compliance)
- Redis for caching and session management
- Elasticsearch for search and analytics
- Decision: No MongoDB to reduce operational complexity

**2024-05-01: Frontend Architecture**
- React for customer web (SEO flexibility)
- Next.js for admin dashboard (SSR benefits)
- React Native for mobile (code sharing)
- Redux Toolkit for consistent state management

## Cross-Cutting Concerns

### Security
- JWT-based authentication across all services
- API rate limiting and request validation
- HTTPS enforcement and certificate management
- Regular security audits and dependency updates

### Monitoring & Observability
- Centralized logging with ELK stack
- Distributed tracing with Jaeger
- Metrics collection with Prometheus + Grafana
- Alerting for critical system events

### Performance Considerations
- Database query optimization and indexing
- API response caching strategies
- Image optimization and CDN usage
- Connection pooling and load balancing

This large project example demonstrates how to organize complex documentation while maintaining clear navigation and comprehensive context for Claude Code collaboration.
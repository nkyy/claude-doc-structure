# Specification Writing Guidelines

This guide explains how to write effective specifications that maximize Claude Code's understanding and enable efficient development collaboration.

## When to Create Specifications

### Single File Approach (Small Projects)
- Project under 10,000 lines of code
- Simple architecture with few modules
- Single developer or small team
- **Use:** Keep everything in `CLAUDE.md`

### Domain Split Approach (Medium Projects)
- Project between 10,000-50,000 lines
- Multiple distinct components (API, UI, database)
- Team of 2-5 developers
- **Use:** Create `specs/api.md`, `specs/screens.md`, etc.

### Hierarchical Approach (Large Projects)
- Project over 50,000 lines
- Multiple services or applications
- Large development team
- **Use:** Full `specs/` directory with sub-categories

## Specification Structure

### Essential Sections

1. **Purpose & Context**
   - What problem does this solve?
   - How does it fit into the larger system?
   - Who are the stakeholders?

2. **Requirements**
   - Functional requirements (what it must do)
   - Non-functional requirements (performance, security)
   - Constraints and assumptions

3. **Design Overview**
   - High-level architecture
   - Key components and their relationships
   - Data flow diagrams (as text or ASCII art)

4. **Implementation Details**
   - File locations and structure
   - Key algorithms or business logic
   - Integration points with other systems

5. **Validation Criteria**
   - How to test the implementation
   - Success metrics
   - Edge cases to consider

### Claude Code Optimization

**Use Specific File References:**
```markdown
## Implementation
The user authentication logic is implemented in:
- `src/auth/login.js:25-45` - Main login function
- `src/auth/middleware.js:12` - JWT validation
- `config/auth.yaml:8` - Authentication configuration
```

**Include Context Breadcrumbs:**
```markdown
## Related Components
This feature interacts with:
- User Management System (see `specs/users.md`)
- Payment Processing (see `specs/billing.md:payment-flow`)
- Notification Service (external API)
```

**Maintain Consistent Terminology:**
- Define technical terms once and use consistently
- Create a glossary for domain-specific language
- Use the same naming conventions as the codebase

## Writing Best Practices

### For Claude Code Understanding

1. **Lead with Context**
   ```markdown
   # User Profile Management

   ## Context
   Users need to update their profile information after registration.
   This feature is part of the account management system and integrates
   with the billing module for subscription changes.
   ```

2. **Use Examples Over Abstractions**
   ```markdown
   ## API Response Format
   ```json
   {
     "user": {
       "id": "usr_123",
       "email": "user@example.com",
       "profile": {
         "name": "John Doe",
         "avatar_url": "https://cdn.example.com/avatars/123.jpg"
       }
     }
   }
   ```
   ```

3. **Include "What Changed" Context**
   ```markdown
   ## Recent Changes
   - 2024-01-15: Added avatar upload functionality
   - 2024-01-10: Migrated from REST to GraphQL (see migration guide)
   - 2023-12-20: Added email verification requirement
   ```

### For Team Collaboration

1. **Decision Documentation**
   ```markdown
   ## Technical Decisions
   
   **Database Choice: PostgreSQL**
   - Reason: Complex relational data requirements
   - Alternative considered: MongoDB (rejected due to ACID needs)
   - Decision date: 2024-01-10
   - Decision maker: Tech Lead (Jane Doe)
   ```

2. **Status Tracking**
   ```markdown
   ## Implementation Status
   - [x] Database schema design
   - [x] API endpoint implementation
   - [ ] Frontend integration (in progress)
   - [ ] Testing and validation
   - [ ] Documentation updates
   ```

## Common Anti-Patterns

### Avoid These Mistakes

❌ **Vague Descriptions**
```markdown
## User Management
The system handles users and their data.
```

✅ **Specific and Actionable**
```markdown
## User Management
Provides CRUD operations for user accounts including:
- Registration with email verification
- Profile updates with avatar upload
- Account deletion with data retention policy
```

❌ **Missing Context**
```markdown
## Database Changes
Add new column to users table.
```

✅ **Full Context**
```markdown
## Database Changes
Add `last_login_at` timestamp column to `users` table to support:
- Inactive account detection (see `specs/cleanup.md`)
- Security audit logging (see `specs/security.md:audit-trail`)
- User engagement analytics (see `specs/analytics.md`)
```

❌ **Outdated Information**
```markdown
## API Endpoints
Uses REST API at `/api/v1/users`
```

✅ **Current and Versioned**
```markdown
## API Endpoints
Current: GraphQL at `/graphql` (see `specs/api.md:graphql-schema`)
Legacy: REST at `/api/v1/users` (deprecated, removal planned for Q2 2024)
```

## Template Usage

Use the provided templates as starting points:

- `api.md` - For REST/GraphQL API documentation
- `screens.md` - For UI/UX specifications
- `.claude/templates/api-endpoint.md` - For individual endpoint docs
- `.claude/templates/screen-spec.md` - For individual screen specs

## Maintenance

### Keep Specifications Current

1. **Update During Development**
   - Modify specs when implementation differs from design
   - Add actual file paths and line numbers
   - Document any architectural decisions made during coding

2. **Regular Review Cycles**
   - Monthly spec review meetings
   - Update status indicators
   - Remove obsolete sections

3. **Cross-Reference Validation**
   - Ensure all file references are current
   - Verify links to other specifications work
   - Check that terminology remains consistent

Remember: The goal is to create documentation that makes Claude Code's assistance more effective and reduces the need for repetitive explanations across development sessions.
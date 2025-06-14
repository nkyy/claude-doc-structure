# Frontend Development Prompts for Claude Code

Collection of reusable prompts for common frontend development tasks. Replace variables in `${VARIABLE_NAME}` format with your specific values.

## Component Development

### Create New Component
```
I need to create a new React component called ${COMPONENT_NAME} that ${COMPONENT_PURPOSE}.

Context:
- This component will be used in ${USAGE_CONTEXT}
- It should follow the existing component patterns in src/components/
- The project uses ${TECH_STACK} (e.g., React + TypeScript + Tailwind CSS)

Requirements:
- ${SPECIFIC_REQUIREMENTS}
- Follow the existing code style and patterns
- Include proper TypeScript types
- Add appropriate accessibility attributes
- Include error handling where needed

Please examine the existing components in src/components/ to understand the current patterns and create a component that fits the established architecture.
```

### Refactor Existing Component
```
I need to refactor the ${COMPONENT_NAME} component located at ${FILE_PATH}.

Goals:
- ${REFACTORING_GOALS}
- Improve ${SPECIFIC_AREAS} (e.g., performance, readability, maintainability)
- Maintain existing functionality and API

Please:
1. First examine the current component implementation
2. Identify areas for improvement
3. Propose the refactoring approach
4. Implement the changes while preserving existing behavior
5. Update any related tests if they exist
```

### Component Testing
```
I need comprehensive tests for the ${COMPONENT_NAME} component at ${FILE_PATH}.

Testing requirements:
- Unit tests for all component functionality
- Test user interactions and state changes
- Test error scenarios and edge cases
- Follow the existing testing patterns in the project

Please examine the existing test files to understand the testing approach and create thorough tests that match the project's testing standards.
```

## State Management

### Add State Management
```
I need to implement state management for ${FEATURE_NAME} using ${STATE_SOLUTION} (e.g., Zustand, Redux, Context API).

Requirements:
- Manage ${STATE_DESCRIPTION}
- Handle ${SPECIFIC_ACTIONS}
- Include proper TypeScript types
- Follow the existing state management patterns in the project

Current state management approach: ${CURRENT_APPROACH}
Files to examine: ${RELEVANT_FILES}

Please create the state management solution that integrates seamlessly with the existing architecture.
```

### Connect Component to State
```
I need to connect the ${COMPONENT_NAME} component at ${FILE_PATH} to the ${STORE_NAME} state.

The component should:
- ${STATE_INTERACTIONS}
- Handle loading and error states appropriately
- Follow the existing patterns for state connections

Please examine the existing state connections in the project and implement the integration following the established patterns.
```

## API Integration

### Add API Integration
```
I need to integrate ${API_ENDPOINT} with the ${COMPONENT_NAME} component.

API Details:
- Endpoint: ${API_ENDPOINT}
- Method: ${HTTP_METHOD}
- Expected data: ${DATA_STRUCTURE}
- Authentication: ${AUTH_REQUIREMENTS}

Requirements:
- Handle loading, success, and error states
- Follow the existing API integration patterns
- Include proper error handling and user feedback
- Add appropriate TypeScript types

Please examine the existing API integrations in src/api/ or src/hooks/ to understand the current patterns and implement the integration consistently.
```

### Create Custom Hook
```
I need a custom hook called ${HOOK_NAME} that ${HOOK_PURPOSE}.

The hook should:
- ${HOOK_FUNCTIONALITY}
- Return ${RETURN_VALUES}
- Handle ${ERROR_SCENARIOS}
- Follow the existing custom hook patterns

Please examine the existing hooks in src/hooks/ to understand the current patterns and create a hook that fits the established architecture.
```

## Styling and UI

### Implement Design System Component
```
I need to implement a ${COMPONENT_TYPE} component that follows our design system.

Design specifications:
- ${VISUAL_REQUIREMENTS}
- Variants: ${COMPONENT_VARIANTS}
- States: ${COMPONENT_STATES}
- Responsive behavior: ${RESPONSIVE_REQUIREMENTS}

Styling approach: ${STYLING_METHOD} (e.g., Tailwind CSS, Styled Components, CSS Modules)

Please examine the existing UI components to understand the design system patterns and create a component that maintains consistency with the established visual language.
```

### Responsive Design Implementation
```
I need to make ${COMPONENT_NAME} at ${FILE_PATH} fully responsive.

Breakpoints to support:
- Mobile: ${MOBILE_SPECS}
- Tablet: ${TABLET_SPECS}  
- Desktop: ${DESKTOP_SPECS}

Current styling approach: ${STYLING_METHOD}

Please examine the existing responsive patterns in the project and implement responsive design that follows the established approach.
```

## Performance Optimization

### Optimize Component Performance
```
The ${COMPONENT_NAME} component at ${FILE_PATH} has performance issues with ${PERFORMANCE_PROBLEM}.

Please:
1. Analyze the current implementation for performance bottlenecks
2. Identify specific optimization opportunities
3. Implement optimizations such as:
   - ${OPTIMIZATION_TECHNIQUES} (e.g., memoization, virtualization, lazy loading)
4. Maintain existing functionality while improving performance
5. Add performance monitoring if appropriate

Please examine the component and suggest the most effective optimization approach.
```

### Bundle Size Optimization
```
I need to optimize the bundle size for ${FEATURE_AREA} by implementing code splitting and lazy loading.

Current bundle size issues: ${SIZE_ISSUES}
Target improvements: ${TARGET_GOALS}

Please:
1. Analyze the current bundle structure
2. Identify opportunities for code splitting
3. Implement lazy loading for appropriate components
4. Ensure proper loading states and error boundaries
5. Follow the existing code splitting patterns if any exist
```

## Form Handling

### Create Form with Validation
```
I need to create a form for ${FORM_PURPOSE} with the following fields:
${FIELD_LIST}

Requirements:
- Client-side validation using ${VALIDATION_LIBRARY}
- Server-side error handling
- Proper accessibility attributes
- Loading states during submission
- Follow existing form patterns in the project

Please examine the existing forms to understand the current patterns and create a form that integrates with the established form handling approach.
```

### Form State Management
```
I need to implement form state management for the ${FORM_NAME} form using ${FORM_LIBRARY} (e.g., React Hook Form, Formik).

Form requirements:
- ${FORM_FIELDS}
- Validation rules: ${VALIDATION_RULES}
- Submission handling: ${SUBMISSION_LOGIC}

Please examine the existing form implementations to understand the current approach and create form state management that follows the established patterns.
```

## Error Handling and Loading States

### Implement Error Boundaries
```
I need to implement error boundaries for ${COMPONENT_AREA} to handle ${ERROR_SCENARIOS}.

Requirements:
- Catch and handle errors gracefully
- Provide user-friendly error messages
- Include error reporting/logging
- Follow the existing error handling patterns

Please examine the current error handling approach and implement error boundaries that integrate with the existing error management system.
```

### Add Loading States
```
I need to add comprehensive loading states to ${COMPONENT_NAME} for ${ASYNC_OPERATIONS}.

Loading state requirements:
- Skeleton loading for ${SKELETON_AREAS}
- Progress indicators for ${PROGRESS_AREAS}
- Disabled states during operations
- Follow existing loading state patterns

Please examine the existing loading state implementations and create loading states that maintain consistency with the established UX patterns.
```

## Accessibility

### Improve Component Accessibility
```
I need to improve the accessibility of ${COMPONENT_NAME} at ${FILE_PATH} to meet WCAG ${WCAG_LEVEL} standards.

Current accessibility issues: ${ACCESSIBILITY_ISSUES}

Requirements:
- Proper ARIA labels and roles
- Keyboard navigation support
- Screen reader compatibility
- Focus management
- Color contrast compliance

Please audit the component for accessibility issues and implement improvements following WCAG guidelines and the project's accessibility standards.
```

## Testing

### Add E2E Tests
```
I need to create end-to-end tests for the ${FEATURE_NAME} user flow.

User flow:
${USER_FLOW_STEPS}

Testing framework: ${E2E_FRAMEWORK} (e.g., Playwright, Cypress)

Please examine the existing E2E test patterns and create comprehensive tests that cover the critical user paths and edge cases for this feature.
```

### Visual Regression Testing
```
I need to set up visual regression testing for ${COMPONENT_AREA} to catch unintended UI changes.

Components to test: ${COMPONENTS_LIST}
Testing scenarios: ${TEST_SCENARIOS}

Please implement visual regression tests that integrate with the existing testing pipeline and provide reliable detection of visual changes.
```

## Variables Reference

Common variables used in these prompts:

- `${COMPONENT_NAME}` - Name of the component
- `${FILE_PATH}` - Path to the file
- `${TECH_STACK}` - Technology stack (React, Vue, Angular, etc.)
- `${STYLING_METHOD}` - CSS framework or styling approach
- `${STATE_SOLUTION}` - State management solution
- `${API_ENDPOINT}` - API endpoint URL
- `${FEATURE_NAME}` - Name of the feature being worked on
- `${USAGE_CONTEXT}` - Where/how the component will be used
- `${REQUIREMENTS}` - Specific requirements for the task

## Usage Examples

### Example 1: Create a User Profile Component
```
I need to create a new React component called UserProfile that displays user information and allows editing.

Context:
- This component will be used in the dashboard and profile pages
- It should follow the existing component patterns in src/components/
- The project uses React + TypeScript + Tailwind CSS

Requirements:
- Display user avatar, name, email, and bio
- Include edit mode with form validation
- Handle loading and error states
- Follow the existing code style and patterns
- Include proper TypeScript types
- Add appropriate accessibility attributes
- Include error handling where needed

Please examine the existing components in src/components/ to understand the current patterns and create a component that fits the established architecture.
```

This will help Claude Code understand the context and create a component that fits your project's patterns and requirements.
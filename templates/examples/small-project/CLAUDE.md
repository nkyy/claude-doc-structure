# Todo App - Small Project Example

## Project Overview

A simple React-based todo application demonstrating the Claude Doc Structure approach for small projects. This example shows how to effectively document a straightforward project using a single CLAUDE.md file.

## Architecture & Technology Stack

**Core Technologies:**
- React 18.x (Frontend framework)
- Local Storage API (Data persistence)
- CSS Modules (Styling)
- Vite (Build tool)

**Key Components:**
- TodoApp: Main application component
- TodoList: List display and management
- TodoItem: Individual todo item component
- AddTodo: New todo creation form

## Project Structure

```
todo-app/
├── src/
│   ├── components/
│   │   ├── TodoApp.jsx        # Main app component
│   │   ├── TodoList.jsx       # Todo list container
│   │   ├── TodoItem.jsx       # Individual todo item
│   │   └── AddTodo.jsx        # Add new todo form
│   ├── utils/
│   │   ├── storage.js         # Local storage utilities
│   │   └── helpers.js         # Helper functions
│   ├── styles/
│   │   └── app.css           # Main styles
│   └── main.jsx              # Application entry point
├── public/
│   └── index.html            # HTML template
├── package.json              # Dependencies and scripts
├── vite.config.js           # Vite configuration
└── CLAUDE.md                # This file - project context
```

## Current Development Focus

**Phase 1: Core Functionality (Completed)**
- ✅ Basic todo CRUD operations
- ✅ Local storage persistence
- ✅ Responsive design

**Phase 2: Enhanced Features (In Progress)**
- 🔄 Add due dates to todos
- 🔄 Categories/tags system
- ⏳ Priority levels

**Phase 3: Advanced Features (Planned)**
- ⏳ Data export/import
- ⏳ Keyboard shortcuts
- ⏳ Dark mode theme

## Key Files & Implementation Details

### Core Components
- `src/components/TodoApp.jsx:1` - Main application state management
- `src/components/TodoList.jsx:15` - Todo rendering and filtering logic  
- `src/components/TodoItem.jsx:8` - Individual todo interactions (complete, edit, delete)
- `src/components/AddTodo.jsx:5` - Form handling and validation

### Utilities
- `src/utils/storage.js:10` - Local storage read/write operations
- `src/utils/helpers.js:3` - Date formatting and ID generation

### Configuration
- `vite.config.js:1` - Build configuration and dev server setup
- `package.json:5` - Dependencies: react, react-dom, vite

## Development Workflow

### Running the Project
```bash
npm install          # Install dependencies
npm run dev         # Start development server
npm run build       # Build for production
npm run preview     # Preview production build
```

### Adding New Features
1. Create component in appropriate directory
2. Update imports in parent components
3. Add any necessary utility functions
4. Update this CLAUDE.md with new file references

## Recent Changes & Decisions

**2024-06-14: Initial Setup**
- Chose Vite over Create React App for faster builds
- Implemented local storage for simplicity over external database
- Used CSS Modules for scoped styling

**Current Issues:**
- Need to add proper error handling for storage operations
- Consider adding PropTypes or TypeScript for better type safety
- Mobile touch interactions could be improved

## Testing Strategy

**Manual Testing Checklist:**
- [ ] Add new todo
- [ ] Mark todo as complete/incomplete  
- [ ] Edit existing todo
- [ ] Delete todo
- [ ] Persist data across browser sessions
- [ ] Handle empty states gracefully

**Future: Automated Testing**
- Unit tests for utility functions
- Component testing with React Testing Library
- E2E tests with Playwright

This small project example demonstrates how a single CLAUDE.md file can effectively capture all necessary context for Claude Code collaboration on simple projects.
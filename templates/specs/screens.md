# Screen Specifications Template

This template provides a structure for documenting user interface screens and user experience flows in a way that's optimized for Claude Code understanding.

## UI/UX Overview

### Design System
- **Framework:** React + TypeScript
- **Styling:** Tailwind CSS + Custom Components
- **State Management:** Zustand
- **Component Library:** Custom design system (see `src/components/ui/`)

### Responsive Breakpoints
- **Mobile:** 320px - 767px
- **Tablet:** 768px - 1023px  
- **Desktop:** 1024px+

**Implementation:** `src/styles/breakpoints.css:5-15`

## Core Screens

### Login Screen

**Route:** `/login`
**Component:** `src/pages/LoginPage.tsx:15-80`

#### Purpose
Allow existing users to authenticate and access their account.

#### Layout Structure
```
┌─────────────────────────────────┐
│           App Header            │
├─────────────────────────────────┤
│                                 │
│        [Company Logo]           │
│                                 │
│     Welcome Back Message        │
│                                 │
│    ┌─────────────────────┐     │
│    │   Email Input       │     │
│    └─────────────────────┘     │
│                                 │
│    ┌─────────────────────┐     │
│    │   Password Input    │     │
│    └─────────────────────┘     │
│                                 │
│         [Login Button]          │
│                                 │
│    Forgot Password Link         │
│    Create Account Link          │
│                                 │
├─────────────────────────────────┤
│           App Footer            │
└─────────────────────────────────┘
```

#### Form Fields
1. **Email Address**
   - Type: Email input with validation
   - Required: Yes
   - Validation: Valid email format
   - Error messages in `src/utils/validation.ts:25-35`

2. **Password**
   - Type: Password input with show/hide toggle
   - Required: Yes
   - Validation: Minimum 8 characters
   - Show/hide component: `src/components/PasswordInput.tsx:20-45`

#### User Interactions
1. **Form Submission**
   - Validation occurs on blur and submit
   - Loading state shows spinner on button
   - Success redirects to dashboard
   - Errors display inline below respective fields

2. **Forgot Password Flow**
   - Links to `/forgot-password` route
   - Component: `src/pages/ForgotPasswordPage.tsx:10-55`

3. **Create Account Flow**
   - Links to `/register` route
   - Component: `src/pages/RegisterPage.tsx:15-120`

#### State Management
```typescript
interface LoginState {
  email: string;
  password: string;
  isLoading: boolean;
  errors: {
    email?: string;
    password?: string;
    general?: string;
  };
}
```

**Store Implementation:** `src/stores/authStore.ts:15-45`

#### Error Handling
- **Invalid Credentials:** "Invalid email or password"
- **Account Locked:** "Account temporarily locked. Try again in 15 minutes"
- **Network Error:** "Connection failed. Please check your internet and try again"

**Error Display Component:** `src/components/ErrorMessage.tsx:8-25`

### Dashboard Screen

**Route:** `/dashboard`
**Component:** `src/pages/DashboardPage.tsx:20-150`

#### Purpose
Provide users with an overview of their account status and quick access to key features.

#### Layout Structure
```
┌─────────────────────────────────────────────────────────┐
│                    Top Navigation                       │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐    │
│  │   Metric 1  │  │   Metric 2  │  │   Metric 3  │    │
│  │   Card      │  │   Card      │  │   Card      │    │
│  └─────────────┘  └─────────────┘  └─────────────┘    │
│                                                         │
│  ┌─────────────────────────────────────────────────┐   │
│  │              Recent Activity Table              │   │
│  │                                                 │   │
│  │  Date | Action | Status | Details               │   │
│  │  -----|--------|--------|--------               │   │
│  │  ...  |  ...   |  ...   |  ...                  │   │
│  └─────────────────────────────────────────────────┘   │
│                                                         │
│  ┌─────────────┐           ┌─────────────┐             │
│  │   Quick     │           │   Recent    │             │
│  │   Actions   │           │   Updates   │             │
│  │   Widget    │           │   Widget    │             │
│  └─────────────┘           └─────────────┘             │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

#### Components Breakdown

1. **Metric Cards**
   - Component: `src/components/dashboard/MetricCard.tsx:10-35`
   - Data source: `src/hooks/useDashboardMetrics.ts:15-40`
   - Updates every 30 seconds via polling

2. **Recent Activity Table**
   - Component: `src/components/dashboard/ActivityTable.tsx:15-80`
   - Pagination: 10 items per page
   - Data source: `src/api/activities.ts:20-45`

3. **Quick Actions Widget**
   - Component: `src/components/dashboard/QuickActions.tsx:8-30`
   - Actions: Create, Import, Export, Settings

4. **Recent Updates Widget**
   - Component: `src/components/dashboard/UpdatesWidget.tsx:12-40`
   - Shows last 5 system updates or announcements

#### Data Loading
```typescript
interface DashboardData {
  metrics: {
    totalUsers: number;
    activeUsers: number;
    revenue: number;
  };
  recentActivity: Activity[];
  quickActions: QuickAction[];
  updates: SystemUpdate[];
}
```

**Data Fetching:** `src/hooks/useDashboardData.ts:10-55`
**Loading States:** Skeleton components in `src/components/ui/Skeleton.tsx`

### User Profile Screen

**Route:** `/profile`
**Component:** `src/pages/ProfilePage.tsx:25-200`

#### Purpose
Allow users to view and edit their personal information and account settings.

#### Layout Structure (Mobile-First)

**Mobile Layout:**
```
┌─────────────────────┐
│    Back Button      │
├─────────────────────┤
│                     │
│   [Profile Photo]   │
│                     │
│    Upload Button    │
│                     │
├─────────────────────┤
│                     │
│   Personal Info     │
│   ┌─────────────┐   │
│   │    Name     │   │
│   └─────────────┘   │
│   ┌─────────────┐   │
│   │    Email    │   │
│   └─────────────┘   │
│   ┌─────────────┐   │
│   │     Bio     │   │
│   └─────────────┘   │
│                     │
├─────────────────────┤
│                     │
│   Account Settings  │
│   • Change Password │
│   • Notifications   │
│   • Privacy         │
│   • Delete Account  │
│                     │
├─────────────────────┤
│   [Save Changes]    │
└─────────────────────┘
```

**Desktop Layout:**
```
┌─────────────────────────────────────────────┐
│               Page Header                   │
├─────────────────────────────────────────────┤
│                                             │
│  ┌─────────────┐    ┌─────────────────┐    │
│  │             │    │                 │    │
│  │  [Profile   │    │  Personal Info  │    │
│  │   Photo]    │    │                 │    │
│  │             │    │  Name: [Input]  │    │
│  │  [Upload]   │    │  Email:[Input]  │    │
│  │             │    │  Bio:  [Text]   │    │
│  └─────────────┘    │                 │    │
│                     │  [Save Button]  │    │
│                     └─────────────────┘    │
│                                             │
│  ┌─────────────────────────────────────┐   │
│  │         Account Settings            │   │
│  │                                     │   │
│  │  Security   Notifications  Privacy  │   │
│  │  [Tab]      [Tab]          [Tab]    │   │
│  │                                     │   │
│  │  [Tab Content Area]                 │   │
│  │                                     │   │
│  └─────────────────────────────────────┘   │
│                                             │
└─────────────────────────────────────────────┘
```

#### Form Sections

1. **Profile Photo Upload**
   - Component: `src/components/profile/PhotoUpload.tsx:15-60`
   - File types: JPG, PNG, WebP
   - Max size: 5MB
   - Crop functionality: `src/utils/imageCrop.ts:10-45`

2. **Personal Information**
   - Name: Text input, 2-50 characters
   - Email: Email input with verification flow
   - Bio: Textarea, 0-500 characters
   - Form validation: `src/schemas/profileSchema.ts:8-25`

3. **Security Settings**
   - Change password form
   - Two-factor authentication toggle
   - Active sessions management
   - Component: `src/components/profile/SecuritySettings.tsx:20-120`

#### State Management
```typescript
interface ProfileState {
  user: User;
  isEditing: boolean;
  isLoading: boolean;
  uploadProgress: number;
  errors: Record<string, string>;
}
```

**Store:** `src/stores/profileStore.ts:20-80`

## Shared Components

### Navigation Components

1. **Top Navigation**
   - Component: `src/components/navigation/TopNav.tsx:15-85`
   - Responsive: Hamburger menu on mobile
   - User menu dropdown on desktop

2. **Sidebar Navigation**  
   - Component: `src/components/navigation/Sidebar.tsx:20-120`
   - Collapsible on desktop
   - Hidden on mobile (use top nav instead)

### Form Components

1. **Input Components**
   - Text Input: `src/components/ui/Input.tsx:10-40`
   - Password Input: `src/components/ui/PasswordInput.tsx:15-50`
   - File Upload: `src/components/ui/FileUpload.tsx:20-80`

2. **Validation Components**
   - Error Message: `src/components/ui/ErrorMessage.tsx:8-25`
   - Field Wrapper: `src/components/ui/FieldWrapper.tsx:12-35`

### Loading States

1. **Skeleton Components**
   - Card Skeleton: `src/components/ui/CardSkeleton.tsx:5-20`
   - Table Skeleton: `src/components/ui/TableSkeleton.tsx:8-30`
   - Form Skeleton: `src/components/ui/FormSkeleton.tsx:10-40`

2. **Loading Indicators**
   - Spinner: `src/components/ui/Spinner.tsx:5-15`
   - Progress Bar: `src/components/ui/ProgressBar.tsx:8-25`

## Accessibility

### WCAG Compliance
- **Level:** AA compliance target
- **Screen Reader:** Full ARIA label support
- **Keyboard Navigation:** All interactive elements accessible via keyboard
- **Color Contrast:** Minimum 4.5:1 ratio

### Implementation
- Focus management: `src/hooks/useFocusManagement.ts:10-35`
- ARIA labels: `src/utils/accessibility.ts:5-25`
- Screen reader announcements: `src/components/ui/Announcer.tsx:8-20`

## Testing Strategy

### Unit Tests
- Component tests: `src/components/**/*.test.tsx`
- Hook tests: `src/hooks/**/*.test.ts`
- Utility tests: `src/utils/**/*.test.ts`

### Integration Tests
- User flow tests: `tests/integration/flows/`
- API integration: `tests/integration/api/`
- Form submission: `tests/integration/forms/`

### E2E Tests
- Playwright tests: `e2e/tests/`
- Critical user paths: Login, dashboard, profile update
- Cross-browser testing: Chrome, Firefox, Safari

**Test Configuration:** `playwright.config.ts:10-50`

## Performance Optimization

### Code Splitting
- Route-based splitting: `src/router/routes.tsx:15-40`
- Component lazy loading: `src/components/lazy/`
- Bundle analysis: `npm run analyze`

### Image Optimization
- Next.js Image component for optimization
- WebP format with fallbacks
- Responsive image sizes
- Implementation: `src/components/ui/OptimizedImage.tsx:10-40`

### Caching Strategy
- Static assets: 1 year cache
- API responses: 5 minutes cache
- User data: Session storage
- Configuration: `src/utils/cache.ts:15-45`

## Design Tokens

### Colors
```css
:root {
  --color-primary: #3b82f6;
  --color-primary-dark: #1d4ed8;
  --color-secondary: #6b7280;
  --color-success: #10b981;
  --color-warning: #f59e0b;
  --color-error: #ef4444;
}
```

**File:** `src/styles/tokens.css:1-30`

### Typography
- **Headings:** Inter font family
- **Body:** Inter font family  
- **Code:** JetBrains Mono
- **Sizes:** 12px, 14px, 16px, 18px, 20px, 24px, 32px, 48px

**File:** `src/styles/typography.css:5-40`

### Spacing
- **Scale:** 4px base unit (4, 8, 12, 16, 20, 24, 32, 40, 48, 64, 80, 96)
- **Grid:** 24px grid system
- **Breakpoints:** 768px, 1024px, 1280px

**File:** `src/styles/spacing.css:1-20`

## Related Documentation

- API Integration: `specs/api.md`
- Component Library: `docs/components.md` 
- Design System: `docs/design-system.md`
- Accessibility Guide: `docs/accessibility.md`

## Changelog

### Version 2.1.0 (2024-01-15)
- Added profile photo upload functionality
- Improved mobile responsive design
- Enhanced form validation UX

### Version 2.0.5 (2024-01-10) 
- Fixed dashboard loading performance
- Updated navigation accessibility
- Added skeleton loading states

### Version 2.0.4 (2024-01-05)
- Enhanced error message clarity
- Fixed mobile navigation bugs
- Improved form field focus management
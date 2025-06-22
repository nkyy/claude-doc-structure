package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template <type> [name]",
	Short: "Generate documentation templates",
	Long:  "Generate documentation templates for API endpoints, screens, features, or language-specific configurations.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		templateType := args[0]
		var name string
		if len(args) > 1 {
			name = args[1]
		}
		generateTemplate(templateType, name)
	},
}

func generateTemplate(templateType, name string) {
	templates := map[string]map[string]string{
		"api": {
			"filename": "api-endpoint.md",
			"content": `# {name} API Endpoint

## Overview
Brief description of what this endpoint does.

## HTTP Method and URL
` + "```" + `
GET/POST/PUT/DELETE /api/{name}
` + "```" + `

## Parameters

### Path Parameters
- ` + "`id`" + ` (string): Description

### Query Parameters
- ` + "`param1`" + ` (string, optional): Description
- ` + "`param2`" + ` (number, required): Description

### Request Body
` + "```json" + `
{
  "field1": "value",
  "field2": 123
}
` + "```" + `

## Response

### Success Response (200 OK)
` + "```json" + `
{
  "success": true,
  "data": {
    "result": "value"
  }
}
` + "```" + `

### Error Responses
- ` + "`400 Bad Request`" + `: Invalid parameters
- ` + "`404 Not Found`" + `: Resource not found
- ` + "`500 Internal Server Error`" + `: Server error

## Examples

### Request
` + "```bash" + `
curl -X GET "http://localhost:3000/api/{name}?param1=value" \
  -H "Content-Type: application/json"
` + "```" + `

### Response
` + "```json" + `
{
  "success": true,
  "data": []
}
` + "```" + `

## Notes
Additional implementation notes or considerations.
`,
		},
		"screen": {
			"filename": "screen-spec.md",
			"content": `# {name} Screen Specification

## Overview
Brief description of the screen's purpose and functionality.

## User Stories
- As a [user type], I want to [goal] so that [benefit]
- As a [user type], I want to [goal] so that [benefit]

## Layout & Components

### Header Section
- Component 1: Description and behavior
- Component 2: Description and behavior

### Main Content
- Component 1: Description and behavior
- Component 2: Description and behavior

### Footer/Actions
- Button 1: What it does
- Button 2: What it does

## User Interactions

### Primary Actions
1. Action 1: Step-by-step description
2. Action 2: Step-by-step description

### Secondary Actions
- Action A: Description
- Action B: Description

## Data Requirements

### API Calls
- ` + "`GET /api/endpoint`" + `: Purpose and when called
- ` + "`POST /api/endpoint`" + `: Purpose and when called

### State Management
- State 1: Description and initial value
- State 2: Description and initial value

## Navigation

### Entry Points
- From Screen A: Via action/button
- From Screen B: Via navigation

### Exit Points
- To Screen C: Via action/button
- To Screen D: Via navigation

## Validation Rules
- Field 1: Validation requirements
- Field 2: Validation requirements

## Error Handling
- Error Type 1: How it's displayed/handled
- Error Type 2: How it's displayed/handled

## Responsive Behavior
Description of how the screen adapts to different screen sizes.

## Accessibility
- ARIA labels and roles
- Keyboard navigation
- Screen reader considerations

## Notes
Additional implementation notes, design decisions, or technical considerations.
`,
		},
		"feature": {
			"filename": "feature-spec.md",
			"content": `# {name} Feature Specification

## Overview
High-level description of the feature and its business value.

## Requirements

### Functional Requirements
1. Requirement 1: Detailed description
2. Requirement 2: Detailed description
3. Requirement 3: Detailed description

### Non-Functional Requirements
- Performance: Expected response times, throughput
- Security: Authentication, authorization, data protection
- Usability: User experience considerations
- Compatibility: Browser/platform support

## User Stories
- As a [user type], I want to [goal] so that [benefit]
- As a [user type], I want to [goal] so that [benefit]
- As a [user type], I want to [goal] so that [benefit]

## Technical Design

### Architecture Overview
Description of how the feature fits into the overall system architecture.

### Components
- Component 1: Responsibility and interfaces
- Component 2: Responsibility and interfaces
- Component 3: Responsibility and interfaces

### Data Model
` + "```" + `
Entity 1:
- field1: type, description
- field2: type, description

Entity 2:
- field1: type, description
- field2: type, description
` + "```" + `

### API Design
- ` + "`GET /api/{name}`" + `: List/retrieve resources
- ` + "`POST /api/{name}`" + `: Create new resource
- ` + "`PUT /api/{name}/:id`" + `: Update existing resource
- ` + "`DELETE /api/{name}/:id`" + `: Delete resource

## Implementation Plan

### Phase 1: Core Functionality
- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

### Phase 2: Enhanced Features
- [ ] Task 1: Description
- [ ] Task 2: Description

### Phase 3: Polish & Optimization
- [ ] Task 1: Description
- [ ] Task 2: Description

## Testing Strategy

### Unit Tests
- Component 1: Test scenarios
- Component 2: Test scenarios

### Integration Tests
- API endpoints: Test scenarios
- Database operations: Test scenarios

### User Acceptance Tests
- User Story 1: Test scenarios
- User Story 2: Test scenarios

## Deployment Considerations
- Database migrations
- Configuration changes
- Feature flags
- Rollback procedures

## Success Metrics
- Metric 1: Target value and measurement method
- Metric 2: Target value and measurement method

## Risks & Mitigation
- Risk 1: Description and mitigation strategy
- Risk 2: Description and mitigation strategy

## Future Considerations
Ideas for future enhancements or related features.
`,
		},
	}
	
	template, exists := templates[templateType]
	if !exists {
		fmt.Printf("Unknown template type: %s\n", templateType)
		fmt.Printf("Available templates: %s\n", strings.Join(getKeys(templates), ", "))
		return
	}
	
	if name == "" {
		name = fmt.Sprintf("example-%s", templateType)
	}
	
	filename := strings.Replace(template["filename"], templateType+"-", name+"-", 1)
	content := strings.ReplaceAll(template["content"], "{name}", name)
	
	// Create templates directory if it doesn't exist
	templatesDir := ".claude/templates"
	err := os.MkdirAll(templatesDir, 0755)
	checkError(err)
	
	filePath := filepath.Join(templatesDir, filename)
	err = os.WriteFile(filePath, []byte(content), 0644)
	checkError(err)
	
	fmt.Printf("Generated template: %s\n", filePath)
}

func getKeys(m map[string]map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
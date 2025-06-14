package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize documentation structure",
	Long:  "Initialize Claude documentation structure in current directory.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := getProjectName(args)
		initProject(projectName)
	},
}

func getProjectName(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	
	cwd, err := os.Getwd()
	if err != nil {
		return "project"
	}
	
	return filepath.Base(cwd)
}

func initProject(projectName string) {
	fmt.Printf("Initializing Claude documentation structure for '%s'...\n", projectName)
	
	// Create basic structure
	directories := []string{
		"specs",
		".claude/prompts",
		".claude/templates",
	}
	
	for _, dir := range directories {
		err := os.MkdirAll(dir, 0755)
		checkError(err)
		fmt.Printf("Created directory: %s\n", dir)
	}
	
	// Create CLAUDE.md template
	claudeMdContent := fmt.Sprintf(`# %s

## Project Overview

Brief description of your project and its purpose.

## Architecture & Technology Stack

**Core Technologies:**
- List your main technologies here
- Framework versions
- Key dependencies

**Key Components:**
- Component 1: Description
- Component 2: Description

## Project Structure

` + "```" + `
%s/
├── src/                    # Source code
├── docs/                   # Documentation
├── tests/                  # Test files
├── CLAUDE.md              # This file - project context for Claude Code
└── specs/                 # Detailed specifications
    ├── api.md             # API documentation
    └── screens.md         # UI/screen specifications
` + "```" + `

## Current Development Focus

Describe what you're currently working on or planning to implement.

## Key Files

- ` + "`src/main.js:1`" + ` - Main application entry point
- ` + "`src/components/App.js:15`" + ` - Main application component
- ` + "`src/utils/helpers.js:8`" + ` - Utility functions

## Dependencies

List key dependencies and their purposes:
- dependency-name: Purpose and version

## Usage Guidelines

Instructions for running, building, and testing the project.
`, projectName, projectName)
	
	claudeMdPath := "CLAUDE.md"
	if _, err := os.Stat(claudeMdPath); os.IsNotExist(err) {
		err := os.WriteFile(claudeMdPath, []byte(claudeMdContent), 0644)
		checkError(err)
		fmt.Printf("Created: %s\n", claudeMdPath)
	} else {
		fmt.Println("CLAUDE.md already exists, skipping...")
	}
	
	// Create specs templates
	apiTemplate := `# API Documentation

## Endpoints

### GET /api/endpoint
Description of the endpoint.

**Parameters:**
- param1 (string): Description
- param2 (number): Description

**Response:**
` + "```json" + `
{
  "example": "response"
}
` + "```" + `

**Example:**
` + "```bash" + `
curl -X GET "http://localhost:3000/api/endpoint?param1=value"
` + "```" + `
`
	
	screensTemplate := `# Screen Specifications

## Screen Name

### Purpose
What this screen is for and when it's used.

### Layout
Description of the visual layout and components.

### User Interactions
- Action 1: What happens
- Action 2: What happens

### API Calls
- Endpoint used: Purpose
- Data displayed: Source

### Navigation
- From: Where users come from
- To: Where users can go next
`
	
	specFiles := map[string]string{
		"specs/api.md":     apiTemplate,
		"specs/screens.md": screensTemplate,
	}
	
	for filePath, content := range specFiles {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err := os.WriteFile(filePath, []byte(content), 0644)
			checkError(err)
			fmt.Printf("Created: %s\n", filePath)
		}
	}
	
	fmt.Println("\nDocumentation structure initialized successfully!")
	fmt.Println("Next steps:")
	fmt.Println("1. Edit CLAUDE.md with your project details")
	fmt.Println("2. Update specs/api.md and specs/screens.md")
	fmt.Println("3. Start collaborating with Claude Code!")
}
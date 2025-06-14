package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [directory]",
	Short: "Validate documentation structure",
	Long:  "Validate Claude documentation structure in the specified directory (default: current directory).",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}
		validateStructure(directory)
	},
}

func validateStructure(directory string) {
	fmt.Printf("Validating documentation structure in: %s\n", directory)
	
	var issues []string
	var recommendations []string
	
	// Check for CLAUDE.md
	claudeMdPath := filepath.Join(directory, "CLAUDE.md")
	if _, err := os.Stat(claudeMdPath); os.IsNotExist(err) {
		issues = append(issues, "Missing CLAUDE.md file (main project context)")
	} else {
		// Check CLAUDE.md content
		content, err := os.ReadFile(claudeMdPath)
		if err == nil {
			contentStr := string(content)
			if len(contentStr) < 200 {
				recommendations = append(recommendations, "CLAUDE.md seems quite short - consider adding more project context")
			}
			if !strings.Contains(contentStr, "Project Overview") {
				recommendations = append(recommendations, "Consider adding a 'Project Overview' section to CLAUDE.md")
			}
		}
	}
	
	// Check for specs directory
	specsDir := filepath.Join(directory, "specs")
	if _, err := os.Stat(specsDir); os.IsNotExist(err) {
		recommendations = append(recommendations, "Consider creating a 'specs/' directory for detailed specifications")
	} else {
		// Check for markdown files in specs
		specFiles, err := filepath.Glob(filepath.Join(specsDir, "*.md"))
		if err == nil && len(specFiles) == 0 {
			recommendations = append(recommendations, "specs/ directory exists but contains no markdown files")
		}
	}
	
	// Check for .claude directory
	claudeDir := filepath.Join(directory, ".claude")
	if _, err := os.Stat(claudeDir); os.IsNotExist(err) {
		recommendations = append(recommendations, "Consider creating a '.claude/' directory for Claude-specific assets")
	} else {
		promptsDir := filepath.Join(claudeDir, "prompts")
		templatesDir := filepath.Join(claudeDir, "templates")
		
		if _, err := os.Stat(promptsDir); os.IsNotExist(err) {
			recommendations = append(recommendations, "Consider creating '.claude/prompts/' for reusable prompts")
		}
		if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
			recommendations = append(recommendations, "Consider creating '.claude/templates/' for documentation templates")
		}
	}
	
	// Report results
	if len(issues) > 0 {
		fmt.Println("\n❌ Issues found:")
		for _, issue := range issues {
			fmt.Printf("  - %s\n", issue)
		}
	}
	
	if len(recommendations) > 0 {
		fmt.Println("\n💡 Recommendations:")
		for _, rec := range recommendations {
			fmt.Printf("  - %s\n", rec)
		}
	}
	
	if len(issues) == 0 && len(recommendations) == 0 {
		fmt.Println("\n✅ Documentation structure looks good!")
	}
	
	// Count markdown files
	markdownCount := 0
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			markdownCount++
		}
		return nil
	})
	
	fmt.Printf("\nScanned %d markdown files\n", markdownCount)
}
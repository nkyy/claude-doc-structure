package cmd

import (
	"github.com/claude-code/claude-doc-structure/internal/merger"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <input-directory>",
	Short: "Merge multiple documents",
	Long:  "Merge multiple documents into a single file for easier processing.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputDir := args[0]
		
		output, _ := cmd.Flags().GetString("output")
		pattern, _ := cmd.Flags().GetString("pattern")
		exclude, _ := cmd.Flags().GetStringSlice("exclude")
		recursive, _ := cmd.Flags().GetBool("recursive")
		noTOC, _ := cmd.Flags().GetBool("no-toc")
		noDividers, _ := cmd.Flags().GetBool("no-dividers")
		noStructure, _ := cmd.Flags().GetBool("no-structure")
		noSummary, _ := cmd.Flags().GetBool("no-summary")
		noClaudeOptimization, _ := cmd.Flags().GetBool("no-claude-optimization")
		
		// Create merger
		m := merger.New(inputDir, output)
		m.Pattern = pattern
		m.Exclude = exclude
		m.Recursive = recursive
		m.AddTOC = !noTOC
		m.AddDividers = !noDividers
		m.PreserveStructure = !noStructure
		m.AddSummary = !noSummary
		m.OptimizeForClaude = !noClaudeOptimization
		
		err := m.Merge()
		checkError(err)
	},
}

func init() {
	mergeCmd.Flags().String("output", "merged-docs.md", "Output filename")
	mergeCmd.Flags().String("pattern", "*.md", "File pattern")
	mergeCmd.Flags().StringSlice("exclude", []string{}, "Files to exclude")
	mergeCmd.Flags().Bool("recursive", false, "Search recursively")
	mergeCmd.Flags().Bool("no-toc", false, "Skip table of contents")
	mergeCmd.Flags().Bool("no-dividers", false, "Skip section dividers")
	mergeCmd.Flags().Bool("no-structure", false, "Skip link processing")
	mergeCmd.Flags().Bool("no-summary", false, "Skip summary section")
	mergeCmd.Flags().Bool("no-claude-optimization", false, "Skip Claude optimization")
}
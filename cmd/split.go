package cmd

import (
	"github.com/claude-code/claude-doc-structure/internal/splitter"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split <input-file>",
	Short: "Split large documents",
	Long:  "Split large documents into smaller, manageable sections.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		
		outputDir, _ := cmd.Flags().GetString("output-dir")
		prefix, _ := cmd.Flags().GetString("prefix")
		maxSections, _ := cmd.Flags().GetInt("max-sections")
		headerLevel, _ := cmd.Flags().GetInt("header-level")
		linesPerFile, _ := cmd.Flags().GetInt("lines-per-file")
		maxSizeKB, _ := cmd.Flags().GetInt64("max-size-kb")
		noNavigation, _ := cmd.Flags().GetBool("no-navigation")
		
		byLines, _ := cmd.Flags().GetBool("by-lines")
		bySize, _ := cmd.Flags().GetBool("by-size")
		
		// Create splitter
		s := splitter.New(inputFile, outputDir, prefix)
		s.MaxSections = maxSections
		s.HeaderLevel = headerLevel
		s.LinesPerFile = linesPerFile
		s.MaxSizeKB = maxSizeKB
		s.AddNavigation = !noNavigation
		
		// Determine split method
		method := splitter.ByHeaders // default
		if byLines {
			method = splitter.ByLines
		} else if bySize {
			method = splitter.BySize
		}
		
		err := s.Split(method)
		checkError(err)
	},
}

func init() {
	splitCmd.Flags().String("output-dir", "", "Output directory")
	splitCmd.Flags().String("prefix", "", "Filename prefix")
	splitCmd.Flags().Bool("by-headers", false, "Split by headers (default)")
	splitCmd.Flags().Bool("by-lines", false, "Split by line count")
	splitCmd.Flags().Bool("by-size", false, "Split by file size")
	splitCmd.Flags().Int("max-sections", 10, "Maximum sections")
	splitCmd.Flags().Int("header-level", 2, "Header level to split on")
	splitCmd.Flags().Int("lines-per-file", 200, "Lines per file")
	splitCmd.Flags().Int64("max-size-kb", 100, "Max file size in KB")
	splitCmd.Flags().Bool("no-navigation", false, "Skip navigation links")
}
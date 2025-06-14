package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "claude-docs",
	Short: "Claude Documentation Tool",
	Long: `Claude Documentation Tool - Unified CLI for Claude Code documentation management.

A toolkit that provides templates, tools, and best practices for organizing 
project context to maximize Claude Code's effectiveness.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(splitCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(validateCmd)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
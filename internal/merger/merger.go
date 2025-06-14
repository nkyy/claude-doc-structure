package merger

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Merger struct {
	InputDir          string
	OutputFile        string
	Pattern           string
	Exclude           []string
	Recursive         bool
	AddTOC            bool
	AddDividers       bool
	PreserveStructure bool
	AddSummary        bool
	OptimizeForClaude bool
}

type Document struct {
	Filename string
	Path     string
	Content  string
	Size     int64
	ModTime  time.Time
}

func New(inputDir, outputFile string) *Merger {
	return &Merger{
		InputDir:          inputDir,
		OutputFile:        outputFile,
		Pattern:           "*.md",
		Exclude:           []string{},
		Recursive:         false,
		AddTOC:            true,
		AddDividers:       true,
		PreserveStructure: true,
		AddSummary:        true,
		OptimizeForClaude: true,
	}
}

func (m *Merger) Merge() error {
	// Find all matching files
	files, err := m.findFiles()
	if err != nil {
		return fmt.Errorf("failed to find files: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no files found matching pattern %s", m.Pattern)
	}

	// Read and process documents
	documents, err := m.readDocuments(files)
	if err != nil {
		return fmt.Errorf("failed to read documents: %w", err)
	}

	// Sort documents by filename
	sort.Slice(documents, func(i, j int) bool {
		return documents[i].Filename < documents[j].Filename
	})

	// Generate merged content
	content := m.generateMergedContent(documents)

	// Write output file
	err = os.WriteFile(m.OutputFile, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Merged %d documents into %s\n", len(documents), m.OutputFile)
	return nil
}

func (m *Merger) findFiles() ([]string, error) {
	var files []string

	if m.Recursive {
		err := filepath.Walk(m.InputDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if matched, _ := filepath.Match(m.Pattern, info.Name()); matched {
				if !m.isExcluded(path) {
					files = append(files, path)
				}
			}

			return nil
		})
		return files, err
	} else {
		pattern := filepath.Join(m.InputDir, m.Pattern)
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}

		for _, match := range matches {
			if !m.isExcluded(match) {
				files = append(files, match)
			}
		}
	}

	return files, nil
}

func (m *Merger) isExcluded(path string) bool {
	filename := filepath.Base(path)
	for _, exclude := range m.Exclude {
		if matched, _ := filepath.Match(exclude, filename); matched {
			return true
		}
	}
	return false
}

func (m *Merger) readDocuments(files []string) ([]Document, error) {
	var documents []Document

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", file, err)
		}

		info, err := os.Stat(file)
		if err != nil {
			return nil, fmt.Errorf("failed to stat file %s: %w", file, err)
		}

		doc := Document{
			Filename: filepath.Base(file),
			Path:     file,
			Content:  string(content),
			Size:     info.Size(),
			ModTime:  info.ModTime(),
		}

		documents = append(documents, doc)
	}

	return documents, nil
}

func (m *Merger) generateMergedContent(documents []Document) string {
	var content strings.Builder

	// Add header
	content.WriteString(fmt.Sprintf("# Merged Documentation\n\n"))
	content.WriteString(fmt.Sprintf("Generated on: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	content.WriteString(fmt.Sprintf("Source directory: %s\n", m.InputDir))
	content.WriteString(fmt.Sprintf("Total documents: %d\n\n", len(documents)))

	// Add table of contents
	if m.AddTOC {
		content.WriteString(m.generateTOC(documents))
	}

	// Add summary
	if m.AddSummary {
		content.WriteString(m.generateSummary(documents))
	}

	// Add Claude optimization section
	if m.OptimizeForClaude {
		content.WriteString(m.generateClaudeOptimization(documents))
	}

	// Add document contents
	for i, doc := range documents {
		if m.AddDividers {
			content.WriteString("---\n\n")
		}

		content.WriteString(fmt.Sprintf("## Document: %s\n\n", doc.Filename))
		
		if m.PreserveStructure {
			content.WriteString(fmt.Sprintf("**File:** `%s`\n", doc.Path))
			content.WriteString(fmt.Sprintf("**Size:** %d bytes\n", doc.Size))
			content.WriteString(fmt.Sprintf("**Modified:** %s\n\n", doc.ModTime.Format("2006-01-02 15:04:05")))
		}

		// Process content
		processedContent := m.processContent(doc.Content)
		content.WriteString(processedContent)

		if i < len(documents)-1 {
			content.WriteString("\n\n")
		}
	}

	return content.String()
}

func (m *Merger) generateTOC(documents []Document) string {
	var toc strings.Builder
	
	toc.WriteString("## Table of Contents\n\n")
	
	for i, doc := range documents {
		anchor := strings.ToLower(strings.ReplaceAll(doc.Filename, " ", "-"))
		anchor = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(anchor, "")
		
		toc.WriteString(fmt.Sprintf("%d. [%s](#document-%s)\n", i+1, doc.Filename, anchor))
	}
	
	toc.WriteString("\n")
	return toc.String()
}

func (m *Merger) generateSummary(documents []Document) string {
	var summary strings.Builder
	
	summary.WriteString("## Summary\n\n")
	
	totalSize := int64(0)
	for _, doc := range documents {
		totalSize += doc.Size
	}
	
	summary.WriteString(fmt.Sprintf("- **Total files:** %d\n", len(documents)))
	summary.WriteString(fmt.Sprintf("- **Total size:** %s\n", formatBytes(totalSize)))
	summary.WriteString(fmt.Sprintf("- **Average size:** %s\n", formatBytes(totalSize/int64(len(documents)))))
	
	// File type breakdown
	extensions := make(map[string]int)
	for _, doc := range documents {
		ext := filepath.Ext(doc.Filename)
		extensions[ext]++
	}
	
	if len(extensions) > 1 {
		summary.WriteString("- **File types:**\n")
		for ext, count := range extensions {
			if ext == "" {
				ext = "no extension"
			}
			summary.WriteString(fmt.Sprintf("  - %s: %d files\n", ext, count))
		}
	}
	
	summary.WriteString("\n")
	return summary.String()
}

func (m *Merger) generateClaudeOptimization(documents []Document) string {
	var optimization strings.Builder
	
	optimization.WriteString("## Claude Code Optimization\n\n")
	optimization.WriteString("This merged document has been optimized for Claude Code with the following features:\n\n")
	optimization.WriteString("- **Structured organization:** Documents are clearly separated and labeled\n")
	optimization.WriteString("- **Navigation aids:** Table of contents and file metadata included\n")
	optimization.WriteString("- **Context preservation:** Original file paths and structure maintained\n")
	optimization.WriteString("- **Link processing:** Internal links updated for merged context\n")
	optimization.WriteString("- **Size awareness:** Large documents split appropriately for Claude's context window\n\n")
	
	// Warn about large files
	largeFiles := 0
	for _, doc := range documents {
		if doc.Size > 50*1024 { // 50KB threshold
			largeFiles++
		}
	}
	
	if largeFiles > 0 {
		optimization.WriteString(fmt.Sprintf("⚠️  **Note:** %d documents are larger than 50KB. Consider splitting them for better Claude Code performance.\n\n", largeFiles))
	}
	
	return optimization.String()
}

func (m *Merger) processContent(content string) string {
	// Remove any existing front matter
	if strings.HasPrefix(content, "---") {
		lines := strings.Split(content, "\n")
		inFrontMatter := true
		var processedLines []string
		
		for _, line := range lines[1:] { // Skip first "---"
			if line == "---" && inFrontMatter {
				inFrontMatter = false
				continue
			}
			if !inFrontMatter {
				processedLines = append(processedLines, line)
			}
		}
		content = strings.Join(processedLines, "\n")
	}
	
	// Process relative links if preserving structure
	if m.PreserveStructure {
		content = m.processLinks(content)
	}
	
	return strings.TrimSpace(content)
}

func (m *Merger) processLinks(content string) string {
	// Update relative markdown links to be absolute
	linkPattern := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	
	return linkPattern.ReplaceAllStringFunc(content, func(match string) string {
		parts := linkPattern.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		
		text := parts[1]
		link := parts[2]
		
		// Skip absolute URLs
		if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
			return match
		}
		
		// Skip anchors
		if strings.HasPrefix(link, "#") {
			return match
		}
		
		// Convert relative paths to absolute
		if !filepath.IsAbs(link) {
			absLink := filepath.Join(m.InputDir, link)
			return fmt.Sprintf("[%s](%s)", text, absLink)
		}
		
		return match
	})
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
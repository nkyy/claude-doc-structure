package splitter

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Splitter struct {
	InputFile   string
	OutputDir   string
	Prefix      string
	MaxSections int
	HeaderLevel int
	MaxSizeKB   int64
	LinesPerFile int
	AddNavigation bool
}

type SplitMethod int

const (
	ByHeaders SplitMethod = iota
	ByLines
	BySize
)

func New(inputFile, outputDir, prefix string) *Splitter {
	return &Splitter{
		InputFile:     inputFile,
		OutputDir:     outputDir,
		Prefix:        prefix,
		MaxSections:   10,
		HeaderLevel:   2,
		MaxSizeKB:     100,
		LinesPerFile:  200,
		AddNavigation: true,
	}
}

func (s *Splitter) Split(method SplitMethod) error {
	// Read input file
	content, err := os.ReadFile(s.InputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Create output directory
	if s.OutputDir == "" {
		s.OutputDir = filepath.Dir(s.InputFile)
	}
	
	err = os.MkdirAll(s.OutputDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	switch method {
	case ByHeaders:
		return s.splitByHeaders(string(content))
	case ByLines:
		return s.splitByLines(string(content))
	case BySize:
		return s.splitBySize(string(content))
	default:
		return fmt.Errorf("unknown split method")
	}
}

func (s *Splitter) splitByHeaders(content string) error {
	lines := strings.Split(content, "\n")
	headerPattern := regexp.MustCompile(fmt.Sprintf(`^#{%d}\s+(.+)`, s.HeaderLevel))
	
	var sections []Section
	var currentSection Section
	var currentLines []string
	
	for _, line := range lines {
		if matches := headerPattern.FindStringSubmatch(line); matches != nil {
			// Save previous section
			if currentSection.Title != "" {
				currentSection.Content = strings.Join(currentLines, "\n")
				sections = append(sections, currentSection)
			}
			
			// Start new section
			currentSection = Section{
				Title:    matches[1],
				Filename: s.generateFilename(matches[1]),
			}
			currentLines = []string{line}
		} else {
			currentLines = append(currentLines, line)
		}
	}
	
	// Add last section
	if currentSection.Title != "" {
		currentSection.Content = strings.Join(currentLines, "\n")
		sections = append(sections, currentSection)
	}
	
	// Limit sections if needed
	if len(sections) > s.MaxSections {
		sections = sections[:s.MaxSections]
	}
	
	return s.writeSections(sections)
}

func (s *Splitter) splitByLines(content string) error {
	lines := strings.Split(content, "\n")
	var sections []Section
	
	for i := 0; i < len(lines); i += s.LinesPerFile {
		end := i + s.LinesPerFile
		if end > len(lines) {
			end = len(lines)
		}
		
		sectionLines := lines[i:end]
		title := fmt.Sprintf("Part %d", (i/s.LinesPerFile)+1)
		
		section := Section{
			Title:    title,
			Filename: s.generateFilename(title),
			Content:  strings.Join(sectionLines, "\n"),
		}
		
		sections = append(sections, section)
	}
	
	return s.writeSections(sections)
}

func (s *Splitter) splitBySize(content string) error {
	maxSizeBytes := s.MaxSizeKB * 1024
	var sections []Section
	
	currentSize := int64(0)
	var currentLines []string
	partNum := 1
	
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		lineSize := int64(len(line) + 1) // +1 for newline
		
		if currentSize+lineSize > maxSizeBytes && len(currentLines) > 0 {
			// Create section
			title := fmt.Sprintf("Part %d", partNum)
			section := Section{
				Title:    title,
				Filename: s.generateFilename(title),
				Content:  strings.Join(currentLines, "\n"),
			}
			sections = append(sections, section)
			
			// Reset for next section
			currentLines = []string{line}
			currentSize = lineSize
			partNum++
		} else {
			currentLines = append(currentLines, line)
			currentSize += lineSize
		}
	}
	
	// Add last section
	if len(currentLines) > 0 {
		title := fmt.Sprintf("Part %d", partNum)
		section := Section{
			Title:    title,
			Filename: s.generateFilename(title),
			Content:  strings.Join(currentLines, "\n"),
		}
		sections = append(sections, section)
	}
	
	return s.writeSections(sections)
}

func (s *Splitter) generateFilename(title string) string {
	// Clean title for filename
	cleaned := strings.ToLower(title)
	cleaned = regexp.MustCompile(`[^a-z0-9\s-]`).ReplaceAllString(cleaned, "")
	cleaned = regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, "-")
	cleaned = strings.Trim(cleaned, "-")
	
	if s.Prefix != "" {
		return fmt.Sprintf("%s%s.md", s.Prefix, cleaned)
	}
	return fmt.Sprintf("%s.md", cleaned)
}

func (s *Splitter) writeSections(sections []Section) error {
	for i, section := range sections {
		content := section.Content
		
		// Add navigation if enabled
		if s.AddNavigation {
			content = s.addNavigation(content, sections, i)
		}
		
		filePath := filepath.Join(s.OutputDir, section.Filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write section %s: %w", section.Filename, err)
		}
		
		fmt.Printf("Created: %s\n", filePath)
	}
	
	return nil
}

func (s *Splitter) addNavigation(content string, sections []Section, currentIndex int) string {
	var nav strings.Builder
	
	nav.WriteString("---\n")
	nav.WriteString("## Navigation\n\n")
	
	// Previous
	if currentIndex > 0 {
		nav.WriteString(fmt.Sprintf("← [Previous: %s](%s)\n", 
			sections[currentIndex-1].Title, sections[currentIndex-1].Filename))
	}
	
	// Next
	if currentIndex < len(sections)-1 {
		nav.WriteString(fmt.Sprintf("→ [Next: %s](%s)\n", 
			sections[currentIndex+1].Title, sections[currentIndex+1].Filename))
	}
	
	nav.WriteString("\n---\n\n")
	
	return nav.String() + content
}

type Section struct {
	Title    string
	Filename string
	Content  string
}
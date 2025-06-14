#!/usr/bin/env python3
"""
Document Splitting Utility for Claude Code Projects

This script splits large markdown documents into smaller, more manageable files
that are easier for Claude Code to process and understand.

Usage:
    python split_docs.py input.md [options]

Examples:
    python split_docs.py specs/large-spec.md --by-headers --max-sections 5
    python split_docs.py README.md --output-dir docs/ --prefix section-
"""

import argparse
import re
import os
import sys
from pathlib import Path
from typing import List, Dict, Tuple, Optional


class DocumentSplitter:
    """Splits markdown documents into smaller files based on various criteria."""
    
    def __init__(self, input_file: str, output_dir: str = None, prefix: str = ""):
        self.input_file = Path(input_file)
        self.output_dir = Path(output_dir) if output_dir else self.input_file.parent
        self.prefix = prefix
        self.sections: List[Dict] = []
        
    def read_input_file(self) -> str:
        """Read the input markdown file."""
        try:
            with open(self.input_file, 'r', encoding='utf-8') as f:
                return f.read()
        except FileNotFoundError:
            print(f"Error: File '{self.input_file}' not found.")
            sys.exit(1)
        except Exception as e:
            print(f"Error reading file: {e}")
            sys.exit(1)
    
    def split_by_headers(self, content: str, max_sections: int = 10, 
                        header_level: int = 2) -> List[Dict]:
        """Split document by markdown headers."""
        header_pattern = rf'^#{{{header_level},{header_level}}}\s+(.+)$'
        sections = []
        
        lines = content.split('\n')
        current_section = {'title': '', 'content': '', 'header_level': 0}
        
        for line in lines:
            header_match = re.match(header_pattern, line, re.MULTILINE)
            
            if header_match:
                # Save previous section if it has content
                if current_section['content'].strip():
                    sections.append(current_section.copy())
                
                # Start new section
                current_section = {
                    'title': header_match.group(1).strip(),
                    'content': line + '\n',
                    'header_level': header_level
                }
            else:
                current_section['content'] += line + '\n'
        
        # Add the last section
        if current_section['content'].strip():
            sections.append(current_section)
        
        # Limit sections if max_sections is specified
        if max_sections and len(sections) > max_sections:
            print(f"Warning: Document has {len(sections)} sections, "
                  f"limiting to {max_sections}")
            sections = sections[:max_sections]
        
        return sections
    
    def split_by_lines(self, content: str, lines_per_file: int = 200) -> List[Dict]:
        """Split document by number of lines."""
        lines = content.split('\n')
        sections = []
        
        for i in range(0, len(lines), lines_per_file):
            section_lines = lines[i:i + lines_per_file]
            section_content = '\n'.join(section_lines)
            
            # Try to find a title from the first header in the section
            title = f"Section {i // lines_per_file + 1}"
            for line in section_lines:
                if line.startswith('#'):
                    title = line.strip('#').strip()
                    break
            
            sections.append({
                'title': title,
                'content': section_content,
                'header_level': 1
            })
        
        return sections
    
    def split_by_size(self, content: str, max_size_kb: int = 100) -> List[Dict]:
        """Split document by file size in KB."""
        max_size_bytes = max_size_kb * 1024
        sections = []
        
        # Split by paragraphs first
        paragraphs = content.split('\n\n')
        current_section = {'title': '', 'content': '', 'size': 0}
        section_num = 1
        
        for paragraph in paragraphs:
            paragraph_bytes = len(paragraph.encode('utf-8'))
            
            if (current_section['size'] + paragraph_bytes > max_size_bytes and 
                current_section['content']):
                
                # Find title for current section
                if not current_section['title']:
                    current_section['title'] = f"Section {section_num}"
                    section_num += 1
                
                sections.append(current_section.copy())
                current_section = {'title': '', 'content': '', 'size': 0}
            
            current_section['content'] += paragraph + '\n\n'
            current_section['size'] += paragraph_bytes
            
            # Try to extract title from headers
            if paragraph.startswith('#') and not current_section['title']:
                current_section['title'] = paragraph.split('\n')[0].strip('#').strip()
        
        # Add the last section
        if current_section['content'].strip():
            if not current_section['title']:
                current_section['title'] = f"Section {section_num}"
            sections.append(current_section)
        
        return sections
    
    def generate_filename(self, title: str, index: int) -> str:
        """Generate a safe filename from section title."""
        # Clean the title for use as filename
        safe_title = re.sub(r'[^\w\s-]', '', title)
        safe_title = re.sub(r'[-\s]+', '-', safe_title)
        safe_title = safe_title.strip('-').lower()
        
        # Ensure filename isn't too long
        if len(safe_title) > 50:
            safe_title = safe_title[:50]
        
        # Add prefix and index
        filename = f"{self.prefix}{index:02d}-{safe_title}.md"
        return filename
    
    def create_index_file(self, sections: List[Dict]) -> None:
        """Create an index file with links to all sections."""
        index_content = f"# {self.input_file.stem} - Split Documentation\n\n"
        index_content += "This document has been split into the following sections:\n\n"
        
        for i, section in enumerate(sections, 1):
            filename = self.generate_filename(section['title'], i)
            index_content += f"{i}. [{section['title']}]({filename})\n"
        
        index_content += f"\n---\n\n"
        index_content += f"*Split from: {self.input_file.name}*\n"
        index_content += f"*Generated on: {__import__('datetime').datetime.now().strftime('%Y-%m-%d %H:%M:%S')}*\n"
        
        index_file = self.output_dir / f"{self.prefix}index.md"
        with open(index_file, 'w', encoding='utf-8') as f:
            f.write(index_content)
        
        print(f"Created index file: {index_file}")
    
    def add_navigation_links(self, sections: List[Dict]) -> None:
        """Add navigation links to each section."""
        for i, section in enumerate(sections):
            nav_links = "\n\n---\n\n"
            
            # Previous link
            if i > 0:
                prev_filename = self.generate_filename(sections[i-1]['title'], i)
                nav_links += f"← Previous: [{sections[i-1]['title']}]({prev_filename}) | "
            
            # Index link
            nav_links += f"[📑 Index]({self.prefix}index.md)"
            
            # Next link
            if i < len(sections) - 1:
                next_filename = self.generate_filename(sections[i+1]['title'], i+2)
                nav_links += f" | Next: [{sections[i+1]['title']}]({next_filename}) →"
            
            section['content'] += nav_links
    
    def write_sections(self, sections: List[Dict], add_navigation: bool = True) -> None:
        """Write sections to individual files."""
        # Create output directory if it doesn't exist
        self.output_dir.mkdir(parents=True, exist_ok=True)
        
        if add_navigation:
            self.add_navigation_links(sections)
        
        for i, section in enumerate(sections, 1):
            filename = self.generate_filename(section['title'], i)
            filepath = self.output_dir / filename
            
            # Add metadata header
            content = f"# {section['title']}\n\n"
            if section['content'].startswith('#'):
                # Remove the first header line since we're adding our own
                content_lines = section['content'].split('\n')
                content += '\n'.join(content_lines[1:])
            else:
                content += section['content']
            
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            
            print(f"Created: {filepath}")
    
    def split(self, method: str = 'headers', **kwargs) -> None:
        """Main splitting method."""
        print(f"Reading input file: {self.input_file}")
        content = self.read_input_file()
        
        if method == 'headers':
            sections = self.split_by_headers(
                content, 
                kwargs.get('max_sections', 10),
                kwargs.get('header_level', 2)
            )
        elif method == 'lines':
            sections = self.split_by_lines(
                content, 
                kwargs.get('lines_per_file', 200)
            )
        elif method == 'size':
            sections = self.split_by_size(
                content,
                kwargs.get('max_size_kb', 100)
            )
        else:
            print(f"Error: Unknown split method '{method}'")
            sys.exit(1)
        
        if not sections:
            print("No sections found to split.")
            return
        
        print(f"Found {len(sections)} sections to create")
        
        # Write sections
        self.write_sections(sections, kwargs.get('add_navigation', True))
        
        # Create index file
        self.create_index_file(sections)
        
        print(f"\nSplit complete! Created {len(sections)} files in {self.output_dir}")


def main():
    parser = argparse.ArgumentParser(
        description="Split large markdown documents into smaller files for Claude Code"
    )
    
    parser.add_argument(
        'input_file',
        help='Input markdown file to split'
    )
    
    parser.add_argument(
        '--output-dir',
        help='Output directory for split files (default: same as input file)'
    )
    
    parser.add_argument(
        '--prefix',
        default='',
        help='Prefix for generated filenames'
    )
    
    # Splitting methods
    split_group = parser.add_mutually_exclusive_group()
    split_group.add_argument(
        '--by-headers',
        action='store_true',
        help='Split by markdown headers (default)'
    )
    split_group.add_argument(
        '--by-lines',
        action='store_true',
        help='Split by number of lines'
    )
    split_group.add_argument(
        '--by-size',
        action='store_true',
        help='Split by file size'
    )
    
    # Method-specific options
    parser.add_argument(
        '--max-sections',
        type=int,
        default=10,
        help='Maximum number of sections when splitting by headers (default: 10)'
    )
    
    parser.add_argument(
        '--header-level',
        type=int,
        default=2,
        help='Header level to split on (default: 2 for ##)'
    )
    
    parser.add_argument(
        '--lines-per-file',
        type=int,
        default=200,
        help='Lines per file when splitting by lines (default: 200)'
    )
    
    parser.add_argument(
        '--max-size-kb',
        type=int,
        default=100,
        help='Maximum file size in KB when splitting by size (default: 100)'
    )
    
    parser.add_argument(
        '--no-navigation',
        action='store_true',
        help='Don\'t add navigation links between sections'
    )
    
    args = parser.parse_args()
    
    # Determine split method
    if args.by_lines:
        method = 'lines'
    elif args.by_size:
        method = 'size'
    else:
        method = 'headers'  # default
    
    # Create splitter and run
    splitter = DocumentSplitter(
        args.input_file,
        args.output_dir,
        args.prefix
    )
    
    splitter.split(
        method=method,
        max_sections=args.max_sections,
        header_level=args.header_level,
        lines_per_file=args.lines_per_file,
        max_size_kb=args.max_size_kb,
        add_navigation=not args.no_navigation
    )


if __name__ == '__main__':
    main()
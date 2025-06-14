#!/usr/bin/env python3
"""
Document Merging Utility for Claude Code Projects

This script merges multiple markdown documents into a single comprehensive file
for review, export, or comprehensive analysis by Claude Code.

Usage:
    python merge_docs.py input_directory [options]

Examples:
    python merge_docs.py specs/ --output combined.md
    python merge_docs.py docs/ --pattern "*.md" --exclude index.md
    python merge_docs.py . --recursive --output full-docs.md
"""

import argparse
import os
import re
import sys
from pathlib import Path
from typing import List, Dict, Optional, Set
import glob


class DocumentMerger:
    """Merges multiple markdown documents into a single file."""
    
    def __init__(self, input_dir: str, output_file: str = "merged-docs.md"):
        self.input_dir = Path(input_dir)
        self.output_file = Path(output_file)
        self.merged_content = ""
        self.files_processed: List[Path] = []
        self.table_of_contents: List[Dict] = []
        
    def find_markdown_files(self, pattern: str = "*.md", recursive: bool = False,
                           exclude: List[str] = None) -> List[Path]:
        """Find all markdown files matching the criteria."""
        exclude = exclude or []
        exclude_set = set(exclude)
        
        if recursive:
            # Use rglob for recursive search
            files = list(self.input_dir.rglob(pattern))
        else:
            # Use glob for non-recursive search
            files = list(self.input_dir.glob(pattern))
        
        # Filter out excluded files
        filtered_files = []
        for file_path in files:
            if file_path.name not in exclude_set:
                # Skip if it's the output file
                if file_path.resolve() != self.output_file.resolve():
                    filtered_files.append(file_path)
        
        # Sort files for consistent ordering
        return sorted(filtered_files)
    
    def extract_title_from_content(self, content: str, filename: str) -> str:
        """Extract title from markdown content."""
        lines = content.strip().split('\n')
        
        # Look for the first header
        for line in lines:
            if line.startswith('#'):
                # Remove markdown header symbols and clean up
                title = re.sub(r'^#+\s*', '', line).strip()
                if title:
                    return title
        
        # Fallback to filename without extension
        return filename.replace('.md', '').replace('-', ' ').replace('_', ' ').title()
    
    def process_internal_links(self, content: str, current_file: Path) -> str:
        """Process internal markdown links to work in merged document."""
        # Pattern for markdown links: [text](link)
        link_pattern = r'\[([^\]]+)\]\(([^)]+)\)'
        
        def replace_link(match):
            text = match.group(1)
            link = match.group(2)
            
            # Skip external links (http/https)
            if link.startswith(('http://', 'https://')):
                return match.group(0)
            
            # Skip anchors
            if link.startswith('#'):
                return match.group(0)
            
            # Handle relative file links
            if link.endswith('.md'):
                # Convert to anchor link within the merged document
                link_path = Path(link)
                if not link_path.is_absolute():
                    link_path = current_file.parent / link_path
                
                # Create anchor from filename
                anchor = link_path.stem.lower().replace(' ', '-').replace('_', '-')
                return f'[{text}](#{anchor})'
            
            return match.group(0)
        
        return re.sub(link_pattern, replace_link, content)
    
    def create_section_divider(self, title: str, filename: str) -> str:
        """Create a section divider for the merged document."""
        divider = f"\n\n{'='*80}\n"
        divider += f"# {title}\n"
        divider += f"*Source: {filename}*\n"
        divider += f"{'='*80}\n\n"
        return divider
    
    def generate_table_of_contents(self) -> str:
        """Generate table of contents for the merged document."""
        toc = "# Table of Contents\n\n"
        
        for i, entry in enumerate(self.table_of_contents, 1):
            anchor = entry['filename'].replace('.md', '').lower().replace(' ', '-').replace('_', '-')
            toc += f"{i}. [{entry['title']}](#{anchor})\n"
        
        toc += "\n---\n\n"
        return toc
    
    def extract_headers_for_toc(self, content: str) -> List[Dict]:
        """Extract headers from content for detailed table of contents."""
        headers = []
        lines = content.split('\n')
        
        for line in lines:
            if line.startswith('#'):
                level = len(line) - len(line.lstrip('#'))
                title = line.strip('#').strip()
                if title:
                    headers.append({
                        'level': level,
                        'title': title,
                        'anchor': title.lower().replace(' ', '-').replace('.', '').replace(',', '')
                    })
        
        return headers
    
    def add_metadata_header(self) -> str:
        """Add metadata header to the merged document."""
        from datetime import datetime
        
        header = "<!-- Merged Documentation -->\n"
        header += f"<!-- Generated on: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')} -->\n"
        header += f"<!-- Source directory: {self.input_dir} -->\n"
        header += f"<!-- Files merged: {len(self.files_processed)} -->\n\n"
        return header
    
    def merge_files(self, files: List[Path], add_toc: bool = True,
                   add_dividers: bool = True, preserve_structure: bool = True) -> None:
        """Merge multiple markdown files into one."""
        print(f"Merging {len(files)} files...")
        
        # Add metadata header
        self.merged_content = self.add_metadata_header()
        
        # Process each file
        for file_path in files:
            try:
                print(f"Processing: {file_path}")
                
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                
                if not content.strip():
                    print(f"  Skipping empty file: {file_path}")
                    continue
                
                # Extract title
                title = self.extract_title_from_content(content, file_path.name)
                
                # Store for TOC
                self.table_of_contents.append({
                    'title': title,
                    'filename': file_path.name,
                    'path': str(file_path)
                })
                
                # Process internal links
                if preserve_structure:
                    content = self.process_internal_links(content, file_path)
                
                # Add section divider
                if add_dividers:
                    self.merged_content += self.create_section_divider(title, file_path.name)
                
                # Add content
                self.merged_content += content
                
                # Add spacing between sections
                if not content.endswith('\n\n'):
                    self.merged_content += '\n\n'
                
                self.files_processed.append(file_path)
                
            except Exception as e:
                print(f"  Error processing {file_path}: {e}")
                continue
        
        # Add table of contents at the beginning
        if add_toc and self.table_of_contents:
            toc = self.generate_table_of_contents()
            # Insert TOC after metadata header
            parts = self.merged_content.split('\n\n', 1)
            if len(parts) == 2:
                self.merged_content = parts[0] + '\n\n' + toc + parts[1]
            else:
                self.merged_content = toc + self.merged_content
    
    def add_summary_section(self) -> None:
        """Add a summary section at the end of the document."""
        summary = f"\n\n{'='*80}\n"
        summary += "# Document Summary\n"
        summary += f"{'='*80}\n\n"
        summary += f"**Total files merged:** {len(self.files_processed)}\n"
        summary += f"**Source directory:** {self.input_dir}\n"
        summary += f"**Generated on:** {__import__('datetime').datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n\n"
        
        summary += "## Files included:\n\n"
        for i, file_path in enumerate(self.files_processed, 1):
            summary += f"{i}. {file_path.name} ({file_path})\n"
        
        summary += f"\n**Total word count:** {len(self.merged_content.split())}\n"
        summary += f"**Total character count:** {len(self.merged_content)}\n"
        
        self.merged_content += summary
    
    def optimize_for_claude(self) -> None:
        """Optimize the merged document for Claude Code processing."""
        # Add Claude Code specific markers
        claude_header = "<!-- Optimized for Claude Code -->\n"
        claude_header += "<!-- This is a comprehensive merged document -->\n"
        claude_header += "<!-- Use Ctrl+F to search for specific sections -->\n\n"
        
        # Insert at the beginning (after existing metadata)
        lines = self.merged_content.split('\n')
        if lines[0].startswith('<!--'):
            # Find the end of existing metadata
            insert_index = 0
            for i, line in enumerate(lines):
                if line.startswith('<!--') or line.strip() == '':
                    insert_index = i + 1
                else:
                    break
            
            lines.insert(insert_index, claude_header)
            self.merged_content = '\n'.join(lines)
        else:
            self.merged_content = claude_header + self.merged_content
        
        # Add file navigation comments
        self.add_file_markers()
    
    def add_file_markers(self) -> None:
        """Add HTML comments to mark file boundaries for easier navigation."""
        # This is already handled by section dividers, but we can add HTML comments too
        for entry in self.table_of_contents:
            filename = entry['filename']
            marker = f"<!-- START: {filename} -->"
            end_marker = f"<!-- END: {filename} -->"
            
            # Find the section and add markers
            section_pattern = f"# {re.escape(entry['title'])}\n\\*Source: {re.escape(filename)}\\*"
            match = re.search(section_pattern, self.merged_content)
            
            if match:
                # Add start marker before the section
                start_pos = match.start()
                self.merged_content = (
                    self.merged_content[:start_pos] + 
                    marker + '\n' + 
                    self.merged_content[start_pos:]
                )
    
    def write_output(self, add_summary: bool = True) -> None:
        """Write the merged content to the output file."""
        if add_summary:
            self.add_summary_section()
        
        # Create output directory if it doesn't exist
        self.output_file.parent.mkdir(parents=True, exist_ok=True)
        
        try:
            with open(self.output_file, 'w', encoding='utf-8') as f:
                f.write(self.merged_content)
            
            print(f"\nMerged document written to: {self.output_file}")
            print(f"Total size: {len(self.merged_content)} characters")
            print(f"Files processed: {len(self.files_processed)}")
            
        except Exception as e:
            print(f"Error writing output file: {e}")
            sys.exit(1)
    
    def merge(self, pattern: str = "*.md", recursive: bool = False,
             exclude: List[str] = None, add_toc: bool = True,
             add_dividers: bool = True, preserve_structure: bool = True,
             add_summary: bool = True, optimize_claude: bool = True) -> None:
        """Main merge method."""
        print(f"Searching for files in: {self.input_dir}")
        
        # Find files to merge
        files = self.find_markdown_files(pattern, recursive, exclude)
        
        if not files:
            print("No markdown files found to merge.")
            return
        
        print(f"Found {len(files)} files to merge:")
        for file_path in files:
            print(f"  - {file_path}")
        
        # Merge files
        self.merge_files(files, add_toc, add_dividers, preserve_structure)
        
        # Optimize for Claude Code if requested
        if optimize_claude:
            self.optimize_for_claude()
        
        # Write output
        self.write_output(add_summary)


def main():
    parser = argparse.ArgumentParser(
        description="Merge multiple markdown documents into a single file"
    )
    
    parser.add_argument(
        'input_directory',
        help='Directory containing markdown files to merge'
    )
    
    parser.add_argument(
        '--output',
        default='merged-docs.md',
        help='Output filename (default: merged-docs.md)'
    )
    
    parser.add_argument(
        '--pattern',
        default='*.md',
        help='File pattern to match (default: *.md)'
    )
    
    parser.add_argument(
        '--exclude',
        nargs='*',
        default=[],
        help='Files to exclude from merging'
    )
    
    parser.add_argument(
        '--recursive',
        action='store_true',
        help='Search subdirectories recursively'
    )
    
    parser.add_argument(
        '--no-toc',
        action='store_true',
        help='Don\'t generate table of contents'
    )
    
    parser.add_argument(
        '--no-dividers',
        action='store_true',
        help='Don\'t add section dividers between files'
    )
    
    parser.add_argument(
        '--no-structure',
        action='store_true',
        help='Don\'t preserve internal link structure'
    )
    
    parser.add_argument(
        '--no-summary',
        action='store_true',
        help='Don\'t add summary section at the end'
    )
    
    parser.add_argument(
        '--no-claude-optimization',
        action='store_true',
        help='Don\'t optimize for Claude Code processing'
    )
    
    args = parser.parse_args()
    
    # Validate input directory
    if not Path(args.input_directory).exists():
        print(f"Error: Directory '{args.input_directory}' does not exist.")
        sys.exit(1)
    
    # Create merger and run
    merger = DocumentMerger(args.input_directory, args.output)
    
    merger.merge(
        pattern=args.pattern,
        recursive=args.recursive,
        exclude=args.exclude,
        add_toc=not args.no_toc,
        add_dividers=not args.no_dividers,
        preserve_structure=not args.no_structure,
        add_summary=not args.no_summary,
        optimize_claude=not args.no_claude_optimization
    )


if __name__ == '__main__':
    main()
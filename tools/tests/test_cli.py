#!/usr/bin/env python3
"""
Tests for the claude-docs CLI tool
"""

import os
import tempfile
import subprocess
import shutil
from pathlib import Path
import pytest


class TestCLI:
    """Test cases for the claude-docs CLI tool"""
    
    @pytest.fixture
    def temp_dir(self):
        """Create a temporary directory for tests"""
        temp_dir = tempfile.mkdtemp()
        yield temp_dir
        shutil.rmtree(temp_dir)
    
    @pytest.fixture
    def cli_path(self):
        """Get the path to the claude-docs CLI script"""
        return Path(__file__).parent.parent / "bin" / "claude-docs"
    
    def run_cli(self, cli_path, args, cwd=None):
        """Helper to run CLI commands"""
        cmd = [str(cli_path)] + args
        result = subprocess.run(
            cmd,
            capture_output=True,
            text=True,
            cwd=cwd
        )
        return result
    
    def test_cli_help(self, cli_path):
        """Test that CLI shows help"""
        result = self.run_cli(cli_path, ["--help"])
        assert result.returncode == 0
        assert "Claude Documentation Tool" in result.stdout
        assert "init" in result.stdout
        assert "split" in result.stdout
        assert "merge" in result.stdout
        assert "template" in result.stdout
        assert "validate" in result.stdout
    
    def test_init_command(self, cli_path, temp_dir):
        """Test project initialization"""
        result = self.run_cli(cli_path, ["init", "test-project"], cwd=temp_dir)
        assert result.returncode == 0
        
        # Check that required files were created
        assert (Path(temp_dir) / "CLAUDE.md").exists()
        assert (Path(temp_dir) / "specs" / "api.md").exists()
        assert (Path(temp_dir) / "specs" / "screens.md").exists()
        assert (Path(temp_dir) / ".claude" / "prompts").exists()
        assert (Path(temp_dir) / ".claude" / "templates").exists()
    
    def test_validate_command(self, cli_path, temp_dir):
        """Test structure validation"""
        # First initialize a project
        self.run_cli(cli_path, ["init", "test-project"], cwd=temp_dir)
        
        # Then validate it
        result = self.run_cli(cli_path, ["validate"], cwd=temp_dir)
        assert result.returncode == 0
        assert "Documentation structure looks good" in result.stdout
    
    def test_template_generation(self, cli_path, temp_dir):
        """Test template generation"""
        # Initialize project first
        self.run_cli(cli_path, ["init", "test-project"], cwd=temp_dir)
        
        # Generate API template
        result = self.run_cli(cli_path, ["template", "api", "users"], cwd=temp_dir)
        assert result.returncode == 0
        assert "Generated template:" in result.stdout
        
        # Check that template file was created
        template_file = Path(temp_dir) / ".claude" / "templates" / "users-endpoint.md"
        assert template_file.exists()
        
        # Check template content
        with open(template_file, 'r') as f:
            content = f.read()
            assert "users API Endpoint" in content
            assert "/api/users" in content
    
    def test_split_command_help(self, cli_path):
        """Test split command help"""
        result = self.run_cli(cli_path, ["split", "--help"])
        assert result.returncode == 0
        assert "Split large documents" in result.stdout
        assert "--by-headers" in result.stdout
        assert "--max-sections" in result.stdout
    
    def test_merge_command_help(self, cli_path):
        """Test merge command help"""
        result = self.run_cli(cli_path, ["merge", "--help"])
        assert result.returncode == 0
        assert "Merge multiple documents" in result.stdout
        assert "--output" in result.stdout
        assert "--recursive" in result.stdout


if __name__ == "__main__":
    pytest.main([__file__])
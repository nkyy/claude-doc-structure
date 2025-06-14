#!/usr/bin/env python3
"""
Setup script for Claude Doc Structure
"""

from setuptools import setup, find_packages
from pathlib import Path

# Read the README file
readme_path = Path(__file__).parent / "README.md"
if readme_path.exists():
    with open(readme_path, "r", encoding="utf-8") as f:
        long_description = f.read()
else:
    long_description = "A documentation structure template for efficient collaboration with Claude Code."

setup(
    name="claude-doc-structure",
    version="1.0.0",
    author="Claude Code Community",
    author_email="noreply@anthropic.com",
    description="Documentation structure template for Claude Code projects",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/your-username/claude-doc-structure",
    py_modules=[],
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Topic :: Documentation",
        "Topic :: Software Development :: Documentation",
        "Topic :: Text Processing :: Markup",
    ],
    python_requires=">=3.8",
    install_requires=[
        # No external dependencies - uses only standard library
    ],
    extras_require={
        "dev": [
            "pytest>=6.0",
            "black>=21.0",
            "flake8>=3.9",
            "mypy>=0.910",
        ],
    },
    scripts=[
        "tools/cli/claude-docs",
    ],
    include_package_data=False,
    zip_safe=False,
)
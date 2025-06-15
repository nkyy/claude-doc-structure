# Debug Log & Issue Resolution

## Critical Issues Resolved

### 2024-06-15: Homebrew Formula Binary Naming Issue
**Problem**: `claude-docs` command not found after Homebrew installation
**Root Cause**: Binary installed as `claude-docs` but user expected `claude-doc`
**Investigation**:
```bash
brew --prefix  # /usr/local
ls -la /usr/local/bin/claude-doc*  # Found claude-docs symlink
which claude-docs  # /usr/local/var/pyenv/shims/claude-docs (wrong!)
```
**Solution**: pyenv shim was intercepting command
**Prevention**: Document correct binary name and potential PATH conflicts

### 2024-06-15: Python Import Error in Go Binary
**Problem**: `Error importing required modules: No module named 'split_docs'`
**Root Cause**: Old Python binary cached in pyenv shims
**Investigation**:
```bash
file /usr/local/bin/claude-docs  # Mach-O 64-bit executable (correct)
which claude-docs  # pyenv shim (incorrect)
```
**Solution**: Remove pyenv shim or use direct path
**Prevention**: Clear documentation about installation and PATH management

### 2024-06-15: GitHub Actions Homebrew Update Failure
**Problem**: Formula directory not found during automated update
**Root Cause**: Empty repository with no Formula directory structure
**Investigation**:
```bash
# GitHub Actions log showed:
# Cloning into 'tap'...
# warning: You appear to have cloned an empty repository.
# /home/runner/work/_temp/.../sh: line 11: Formula/claude-doc.rb: No such file or directory
```
**Solution**: Created proper Homebrew tap repository structure
**Prevention**: Initialize tap repository with proper directory structure before automation

## Build System Issues

### Cross-Platform Binary Paths
**Issue**: Makefile inconsistency between `dist/` and `bin/` directories
**Fix**: Standardized on `bin/` directory for all build outputs
**Impact**: GitHub Actions and local builds now consistent

### Go Module Dependencies
**Monitoring**: Watch for dependency bloat
**Current**: Minimal dependencies (cobra for CLI)
**Strategy**: Evaluate each new dependency for necessity

## Installation Issues

### Environment Conflicts
**Common Issue**: Multiple installation methods conflict
**Symptoms**:
- Command not found despite successful installation
- Wrong version executing
- Python error in Go binary

**Debugging Steps**:
```bash
# Check installation
which claude-docs
file $(which claude-docs)

# Check Homebrew installation
brew list claude-doc
brew --cellar claude-doc

# Check PATH priority
echo $PATH | tr ':' '\n' | grep -E "(homebrew|pyenv)"
```

## Performance Monitoring

### Binary Size Tracking
- **v0.1.0**: ~5MB (baseline)
- **v0.1.1**: ~5MB (stable)
- **Target**: Keep under 10MB

### Memory Usage
- **Startup**: <10MB heap
- **Document Processing**: Scales with file size
- **Target**: Efficient streaming for large files

## Testing Gaps Identified

### Manual Testing Coverage
- ✅ Basic CLI commands
- ✅ Cross-platform builds
- ❌ Large file processing
- ❌ Error condition handling
- ❌ Template generation edge cases

### Automated Testing Needs
- Unit tests for splitter/merger logic
- Integration tests for CLI commands
- Performance regression tests
- Cross-platform compatibility tests

## Monitoring & Alerting

### GitHub Actions Health
**Check**: Weekly review of failed builds
**Metrics**: Build time, success rate, artifact size
**Alerts**: Failed releases, Homebrew update failures

### User Issue Patterns
**Common Problems**:
1. Installation PATH conflicts
2. Binary naming confusion
3. Missing documentation

**Solutions**:
1. Improved installation documentation
2. Consistent naming across platforms
3. Troubleshooting guide

## Future Debug Strategies

### Logging Implementation
**Needed**: Structured logging for troubleshooting
```go
// Add debug flag to commands
var debugFlag bool
cmd.Flags().BoolVar(&debugFlag, "debug", false, "Enable debug logging")
```

### Error Context Enhancement
**Needed**: Better error messages with context
```go
// Instead of generic errors
return fmt.Errorf("failed to process file: %w", err)

// Provide actionable context
return fmt.Errorf("failed to process file %s (size: %d bytes): %w", filename, size, err)
```

### Performance Profiling
**Tool**: Go built-in profiling
**Usage**: `go build -race` for race condition detection
**Monitoring**: Memory leaks in long-running operations

## Issue Templates

### Bug Report Template
```
**Environment:**
- OS: [macOS/Linux/Windows]
- Architecture: [amd64/arm64]
- Installation: [Homebrew/Manual]
- Version: [claude-docs --version]

**Command Executed:**
[Exact command that failed]

**Expected Behavior:**
[What should happen]

**Actual Behavior:**
[What actually happened]

**Debug Info:**
[Output of debug commands]
```

### Performance Issue Template
```
**File Details:**
- Size: [file size]
- Type: [markdown/text/other]
- Line count: [approximate]

**Performance Issue:**
- Operation: [split/merge/validate]
- Duration: [how long it took]
- Expected: [how long it should take]

**System Info:**
- RAM: [available memory]
- CPU: [processor info]
- Disk: [available space]
```
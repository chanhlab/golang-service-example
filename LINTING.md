# Code Linting Setup

This project uses [golangci-lint](https://golangci-lint.run/) for code quality and style enforcement.

## Installation

golangci-lint is already configured in the project. If you need to install it separately:

```bash
# macOS (using Homebrew)
brew install golangci-lint

# Linux/macOS (using Go)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Or download binary from releases
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
```

## Usage

The project provides several make targets for linting:

### Basic Linting
```bash
make lint
```
Runs golangci-lint with default configuration to check for code issues.

### Auto-fix Issues
```bash
make lint-fix
```
Automatically fixes issues that can be resolved automatically (formatting, imports, etc.).

### Verbose Output
```bash
make lint-verbose
```
Runs linting with verbose output to see detailed information about the process.

## Enabled Linters

The project uses golangci-lint default configuration which includes:

### Core Linters (Always Enabled)
- **errcheck**: Checks for unchecked errors
- **gosimple**: Simplifies Go code
- **govet**: Examines Go source code and reports suspicious constructs
- **ineffassign**: Detects unused variable assignments
- **staticcheck**: Advanced static analysis
- **typecheck**: Type checks Go code
- **unused**: Finds unused constants, variables, functions and types

### Additional Linters
The default configuration also enables other useful linters for code quality.

## Fixed Issues

During setup, the following issues were identified and resolved:

1. **Unchecked errors** (`errcheck`):
   - Fixed `m.Close()` calls in migration functions
   - Fixed `connection.Close()` call in worker

2. **Error message capitalization** (`staticcheck`):
   - Fixed error messages to follow Go conventions (lowercase, no capitalization)
   - Updated error messages in credential service

3. **Deprecated API usage** (`staticcheck`):
   - Replaced `grpc.Dial` with `grpc.NewClient` in worker code

## Configuration

The project currently uses golangci-lint's default configuration. For custom configuration, create a `.golangci.yml` file in the project root.

Example configuration based on the [Ollama project](https://github.com/ollama/ollama/blob/main/.golangci.yaml):

```yaml
linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - typecheck
    - unused
    - bodyclose
    - gofmt
    - goimports
    - misspell
    - gocritic
    - revive
    - gosec

linters-settings:
  goimports:
    local-prefixes: github.com/chanhlab/golang-service-example
  
  staticcheck:
    checks: ["all", "-SA1019"]

issues:
  exclude-rules:
    # Skip generated files
    - path: "generated/.*\\.pb\\.go$"
      linters:
        - all
    
    # Skip some linters for test files
    - path: "_test\\.go$"
      linters:
        - goconst
        - gosec

run:
  timeout: 5m
```

## Integration with CI/CD

To integrate with CI/CD pipelines, add the following to your workflow:

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linting
make lint
```

## IDE Integration

### VS Code
Install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) which includes golangci-lint integration.

### GoLand/IntelliJ
GoLand has built-in support for golangci-lint. Enable it in Settings > Go > Linter.

### Vim/Neovim
Use plugins like `ale` or `vim-go` with golangci-lint integration.

## Best Practices

1. **Run linting before commits**: Always run `make lint` before committing code
2. **Fix issues incrementally**: Use `make lint-fix` for auto-fixable issues
3. **Keep code clean**: Address linting issues as they appear, don't accumulate technical debt
4. **Understand the rules**: Learn what each linter checks for to write better code from the start

## Troubleshooting

### Common Issues

1. **"command not found: golangci-lint"**
   - Install golangci-lint using one of the methods above

2. **Configuration errors**
   - Ensure `.golangci.yml` syntax is correct
   - Check golangci-lint version compatibility

3. **Too many issues**
   - Start by running `make lint-fix` to auto-fix simple issues
   - Address issues incrementally, file by file

### Getting Help

- [golangci-lint documentation](https://golangci-lint.run/)
- [List of all available linters](https://golangci-lint.run/usage/linters/)
- [Configuration options](https://golangci-lint.run/usage/configuration/)
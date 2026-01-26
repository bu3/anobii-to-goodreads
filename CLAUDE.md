# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI tool that converts Anobii book export files to Goodreads import format. The tool processes CSV files containing book data and transforms them between the two platforms' data schemas.

## Key Commands

### Building and Running
- `go run cmd/main.go -in <input_file> -out <output_file>` - Run the converter directly
- `go build cmd/main.go` - Build the binary

### Testing
- `go test ./...` - Run all tests (uses Ginkgo/Gomega testing framework)
- `go test ./pkg/...` - Run package tests only
- `go test ./acceptance/...` - Run acceptance tests only
- `go test -coverprofile=cover.out ./...` - Run tests with coverage

### Code Quality
- `go fmt ./...` - Format all Go files
- `go vet ./...` - Run Go static analysis
- `go mod tidy` - Clean up module dependencies

## Architecture

The project follows a layered architecture with clear separation of concerns:

### Core Components
- **cmd/main.go** - CLI entry point handling flags and orchestrating the conversion
- **pkg/convert/** - Main conversion orchestrator that coordinates all other components
- **pkg/providers/** - Data models and file I/O for different platforms (anobii, goodreads)
- **pkg/mapping/** - Business logic for transforming Anobii records to Goodreads format
- **pkg/file/** - File system operations abstraction

### Data Flow
1. Anobii CSV file is read via `anobii.AnobiiFile.Read()`
2. Data is transformed through `mapping.AnobiiToGoodReadsMapper`
3. Goodreads format is written via `goodreads.GoodReadsFile.Write()`

### Key Transformation Logic
The mapper handles:
- Reading status conversion (finished → read, being read → reading, etc.)
- Date parsing from Anobii status strings (supports YYYY-MM-DD and YYYY formats)
- Direct field mapping (ISBN, Title, Author, Rating)

## Testing Framework

Uses Ginkgo v2 with Gomega for BDD-style testing. Test structure includes:
- Unit tests for individual packages
- Acceptance tests for end-to-end functionality
- Test coverage configuration via `.testcoverage.yml` with 70% minimum coverage
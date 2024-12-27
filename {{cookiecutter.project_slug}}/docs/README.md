# {{ cookiecutter.project_name }} Documentation

## Overview
{{ cookiecutter.project_name }} is a Go project scaffold generated using Cookiecutter. This project supports multiple types:
- Library
- Console Application
- Web Server
- gRPC Server

## Getting Started

### Requirements
- Go 1.20 or newer
- `make` utility

### Project Structure
- `cmd/`: Main application entry point.
- `pkg/`: Reusable library code.
- `internal/`: Internal-only packages, not meant to be used externally.
- `examples/`: Usage examples (for library projects).
- `proto/`: Protocol buffer definitions (for gRPC projects).
- `docs/`: Documentation.

## Usage

### Building the Project
Run the following command to build the project:
```bash
make build

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`devlog` is a CLI tool for tracking daily developer activities. It stores timestamped log entries in a JSON file located at `~/.devlog/logs.json`.

## Build and Run

```bash
# Build the binary
go build -o devlog .

# Run directly with go
go run main.go [command]

# Install to $GOPATH/bin
go install
```

## Architecture

The codebase follows a clean layered architecture:

- **main.go**: Entry point that delegates to the cmd package
- **cmd/**: Cobra command definitions (one file per command)
  - `root.go`: Root command setup
  - `add.go`: Add new log entries
  - `today.go`: Display today's entries
  - `list.go`: Display all entries
- **internal/storage/**: Data persistence layer
  - `json.go`: Handles reading/writing entries to `~/.devlog/logs.json`
  - Defines the `Entry` struct with `Message` and `Date` fields

### Command Pattern

All commands follow the Cobra pattern:
1. Define a `cobra.Command` struct with Use, Short, Long, and RunE
2. Register the command in `init()` via `rootCmd.AddCommand()`
3. Commands interact with storage through `storage.LoadEntries()` and `storage.SaveEntry()`

### Data Storage

The storage layer abstracts JSON file operations:
- `dataFilePath()`: Returns `~/.devlog/logs.json`, creating directory if needed
- `LoadEntries()`: Reads all entries from JSON file
- `SaveEntry(message string)`: Appends new entry with current timestamp

Date filtering happens in the command layer, not in storage.

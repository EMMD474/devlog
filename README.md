# devlog

A simple CLI tool for tracking your daily development activities. Keep a log of what you worked on each day, review your progress, and maintain a record of your development journey.

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/emmd474/devlog.git
cd devlog

# Build the binary
go build -o devlog .

# Move to a location in your PATH (optional)
sudo mv devlog /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/emmd474/devlog@latest
```

## Usage

### Adding Log Entries

Add a new entry to today's log:

```bash
devlog add "Fixed authentication bug in login module"
devlog add "Implemented user profile API endpoint"
devlog add "Reviewed PRs and updated documentation"
```

### Viewing Today's Entries

Display all entries logged today:

```bash
devlog today
```

Output:
```
- Fixed authentication bug in login module
- Implemented user profile API endpoint
- Reviewed PRs and updated documentation
```

### Viewing All Entries

Display all log entries with timestamps:

```bash
devlog list
```

Output:
```
[08 Jan 26 14:30 UTC] Fixed authentication bug in login module
[08 Jan 26 15:45 UTC] Implemented user profile API endpoint
[08 Jan 26 16:20 UTC] Reviewed PRs and updated documentation
```

## Daily Workflow

Here's a typical daily workflow with devlog:

1. **Start your day**: Check what you did yesterday
   ```bash
   devlog list
   ```

2. **Log as you work**: Add entries throughout the day
   ```bash
   devlog add "Started work on payment gateway integration"
   devlog add "Fixed CSS layout issue on dashboard"
   ```

3. **End of day review**: See what you accomplished
   ```bash
   devlog today
   ```

## Data Storage

All log entries are stored in `~/.devlog/logs.json` as JSON data. This file is automatically created on first use.

## Examples

```bash
# Log a bug fix
devlog add "Fixed memory leak in background worker"

# Log a feature implementation
devlog add "Added dark mode toggle to settings page"

# Log learning and research
devlog add "Researched Redis caching strategies"

# Log code review activities
devlog add "Reviewed 3 PRs, approved 2, requested changes on 1"

# View today's accomplishments
devlog today
```

## License

MIT

# Cobra-CMD

A CLI application for processing and partitioning JIRA tickets using the Cobra framework.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Process Command](#process-command)
  - [Partition Command](#partition-command)
- [Examples](#examples)
- [Development](#development)

## Installation

### Prerequisites
- Go 1.24.1 or later

### Setup
1. Clone the repository:
   ```
   git clone https://github.com/user/cobra-cmd.git
   cd cobra-cmd
   ```

2. Build the application:
   ```
   go build -o cobra-cmd
   ```

3. (Optional) Add the binary to your PATH:
   ```
   chmod +x cobra-cmd
   mv cobra-cmd /usr/local/bin/
   ```

## Usage

### Process Command

Process JIRA tickets with optional version specification.

```
cobra-cmd process [flags]
```

Flags:
- `--tickets string`: Comma-separated list of JIRA tickets to process
- `--version string`: Version for processing

### Partition Command

Partition JIRA tickets into a specified file or default file.

```
cobra-cmd partition [flags]
```

Flags:
- `--tickets string`: Comma-separated list of JIRA tickets to partition
- `--filename string`: Output filename for partitioned tickets (default "default_partition.txt")

## Examples

### Default Command (Process)

Running the application without any subcommand will execute the process command:

```
cobra-cmd
```

### Process Command Examples

Process specific tickets:
```
cobra-cmd process --tickets "JIRA-123,JIRA-456"
```

Process tickets with a version:
```
cobra-cmd process --tickets "JIRA-123,JIRA-456" --version "v1.0.0"
```

### Partition Command Examples

Partition tickets to the default file:
```
cobra-cmd partition --tickets "JIRA-123,JIRA-456"
```

Partition tickets to a custom file:
```
cobra-cmd partition --tickets "JIRA-123,JIRA-456" --filename "sprint_tickets.txt"
```

## Development

To add new commands to the application:

1. Use Cobra to generate a new command:
   ```
   cobra add [command-name]
   ```

2. Implement your command logic in the generated file.

3. Add the command to the root command in `cmd/root.go`.
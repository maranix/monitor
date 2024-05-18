# Monitor

Monitor is a CLI tool for reloading commands/services on filesystem changes using GoLang. It leverages the `cobra` and `fsnotify` packages to watch for changes in a specified directory and execute a command when changes are detected.

## Features

- Watches a specified directory for file changes.
- Debounces events to prevent multiple executions from rapid subsequent changes.
- Ignores extra `CHMOD` events on Linux to avoid duplicate command executions.

## Installation

To install Monitor, you need to have Go installed on your machine. Then, you can build the project using the following commands:

```bash
git clone https://github.com/yourusername/monitor.git
cd monitor
go build -o monitor .
```

## Usage

After building the project, you can run the monitor command with the following options:

```bash
./monitor <directory> <command/service to run>
```

## Flags

- `--verbose`, `-v` gives out the verbose output.

## Commands

- `version`: prints the cli version

## Example

To watch the current directory and output `hello` everytime a modification is made within the directory:

```bash
./monitor ./ "echo hello"

or

go run main.go ./ "echo hello"
```

## Contributions

Contributions are welcome! Please feel free to submit a Pull Request for anything you might like to see or have any issues with.

## License

This project is Licensed under [MIT License](https://github.com/maranix/monitor/blob/main/LICENSE)

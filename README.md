<p align="center">
    <img src="https://github.com/maranix/monitor/assets/59292838/4d8598f2-a345-4ab3-b533-e07a1d3e9610" alt="Monitor Logo" width="300" height="300"/>
</p>

<h1 align="center">Monitor</h1>
Monitor is a CLI tool for reloading commands/services on filesystem changes. It leverages the `fsnotify` package to watch for changes in a specified directory and executes command when changes are detected.

> [!WARNING]
> ### Project Status
> Monitor is currently under active development and is a part of my daily development workflow.
>
> As I am still learning GoLang, the APIs, features, and other aspects of the tool are subject to regular changes. This project serves as both a practical tool and a learning experience.
>
> I welcome any suggestions on features, code practices, and improvements. If you have any feedback, please feel free to open an issue. I'd love to learn more and discuss potential enhancements!

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

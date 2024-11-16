# Go Text CLI

A Go-powered CLI tool that counts anything in your text with a single command.

## Features

- Count total characters (including spaces)
- Count characters excluding spaces
- Count total words in text
- Find longest and shortest words in text (with character count)
- Estimate reading time
- Detailed help command

## Installation

First, clone the repository:

```bash
git clone https://github.com/erdembaran/go-text-cli.git
cd go-text-cli
```

Make sure you have Go installed on your system. Then build the project:

```bash
go build
```

## Usage

The tool provides several commands for text analysis:

### Count All Characters

```bash
go run main.go -all "Your text here"
```

### Count Characters (Excluding Spaces)

```bash
go run main.go -nospace "Your text here"
```

### Count Words

```bash
go run main.go -words "Your text here"
```

### Show Help

```bash
go run main.go -help
```

## Examples

```bash
# Count all characters in a text
go run main.go -all "Hello World"
> Total characters (including spaces): 11

# Count characters excluding spaces
go run main.go -nospace "Hello World"
> Total characters (excluding spaces): 10

# Count total words
go run main.go -words "Hello World"
> Total words: 2
```

## Command Reference

| Command    | Description                           | Example                                 |
| ---------- | ------------------------------------- | --------------------------------------- |
| `-all`     | Count all characters including spaces | `go run main.go -all "Hello World"`     |
| `-nospace` | Count characters excluding spaces     | `go run main.go -nospace "Hello World"` |
| `-words`   | Count total words in text             | `go run main.go -words "Hello World"`   |
| `-help`    | Show help message                     | `go run main.go -help`                  |

## Requirements

- Go 1.23.1 or higher

## Contributing

    1. Fork the repository
    2. Create your feature branch (`git checkout -b feature/amazing-feature`)
    3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
    4. Push to the branch (`git push origin feature/amazing-feature`)
    5. Open a Pull Request

## Author

Baran Erdem

## Project Status

This project is actively maintained. Feel free to create issues for bugs or feature requests.

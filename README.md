# Chat PDF Generator

A CLI tool that generates styled PDF documents from chat logs. The tool creates professional-looking PDFs with headers, footers, and formatted chat entries.

## Features

- Header with logo and title
- Footer with page numbers and hyperlink
- Formatted chat entries with timestamps
- Support for emojis and Unicode characters
- Automatic pagination
- Clean, modular code structure

## Requirements

- Go 1.16 or higher
- gofpdf library

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Place your logo file as `logo.png` in the project directory
2. Run the program:
   ```bash
   go run main.go
   ```

The program will generate a `chat_log.pdf` file with sample chat entries.

## Customization

You can modify the following aspects of the PDF:
- Header title and logo
- Footer URL and styling
- Chat entry formatting
- Font styles and sizes
- Page layout and margins

## License

MIT License 
# Gemini CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/Pradipbabar/gemini-cli)](https://goreportcard.com/report/github.com/Pradipbabar/gemini-cli)
[![Go Reference](https://pkg.go.dev/badge/github.com/Pradipbabar/gemini-cli.svg)](https://pkg.go.dev/github.com/Pradipbabar/gemini-cli)
[![License](https://img.shields.io/github/license/Pradipbabar/gemini-cli)](LICENSE)

A powerful command-line interface for Google's Gemini AI API, built with Go. This CLI tool enables seamless interaction with Gemini AI for chat conversations and content generation directly from your terminal.

## 🚀 Features

- 💬 Interactive chat conversations with Gemini AI
- 📝 Generate explanations and content from text inputs
- 🔑 Simple API key configuration and management
- 💾 Save responses to files
- 📤 JSON output format support
- 🔒 Secure API key handling

## 📋 Prerequisites

- Go 1.21 or higher
- Google Cloud account with Gemini API access
- Gemini API key

## ⚡ Quick Start

### Configuration

Set up your Gemini API key:

```bash
gemini-cli config -k YOUR_API_KEY
```

### Basic Usage

Chat with Gemini:
```bash
gemini-cli chat -p "What is artificial intelligence?"
```

Generate content with JSON output:
```bash
gemini-cli chat -p "Explain quantum computing" -o json
```

## 🛠️ Commands

### `config`
Configure your Gemini API key:
```bash
gemini-cli config -k YOUR_API_KEY
```

### `chat`
Start a chat conversation:
```bash
gemini-cli chat -p "Your prompt" [-s output.txt] [-o json]
```
Options:
- `-p, --prompt`: Input prompt text
- `-s, --save`: Save response to file
- `-o, --output`: Output format (text/json)

### `explain`
Generate explanations from file input:
```bash
gemini-cli explain -f input.txt [-s output.txt] [-o json]
```
Options:
- `-f, --file`: Input file path
- `-s, --save`: Save response to file
- `-o, --output`: Output format (text/json)

## 🏗️ Project Structure

```
gemini-cli/
├── cmd/            # Command implementations
├── pkg/            # Core packages
│   ├── chat.go     # Chat functionality
│   ├── client.go   # Gemini API client
│   ├── config.go   # Configuration management
│   ├── explain.go  # Explanation generation
│   └── models.go   # Data models
├── go.mod
├── go.sum
├── main.go         # Entry point
└── README.md
```

## 🔧 Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Pradipbabar/gemini-cli.git
   cd gemini-cli
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o gemini-cli
   ```

4. Run tests:
   ```bash
   go test ./...
   ```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## 📝 Environment Variables

- `GEMINI_API_KEY`: Your Gemini API key (set automatically by the config command)

## 🔒 Security

- API keys are stored securely in the user's environment
- HTTPS is used for all API communications
- Input validation is performed on all commands

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Google's Gemini AI team for providing the API
- The Go community for excellent tools and libraries
- Contributors to this project

## ⚠️ Disclaimer

This is not an official Google product. This CLI tool is maintained independently and is subject to the Gemini API's terms of service and usage limits.

## 📞 Support

For support, please:
1. Check the [issues](https://github.com/Pradipbabar/gemini-cli/issues) page
2. Create a new issue if your problem isn't already listed
3. Provide as much context as possible when reporting issues

---

Made with ❤️ by [Pradip Babar](https://github.com/Pradipbabar)

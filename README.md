

# Gemini AI CLI

Gemini AI CLI is a command-line interface (CLI) application written in Go, powered by the Cobra package. It allows you to interact with the Gemini AI API for various tasks such as configuring access keys, engaging in chat conversations, and generating explanations.

## Installation

```bash
go install github.com/Pradipbabar/gimini-cli
```

## Commands

### 1. Configure

Set up your Gemini AI access key.

```bash
gimini-cli configure -k YOUR_ACCESS_KEY
```

### 2. Chat

Engage in chat with Gemini AI.

```bash
gimini-cli chat -p "Your prompt text" -s output.txt
```

- `-p`: The prompt text.
- `-s`: (Optional) Save the response to the specified file.

### 3. Explain

Generate explanations using a file as input.

```bash
gimini-cli explain -f input.txt -s output.txt
```

- `-f`: The file path with the input data for the prompt.
- `-s`: (Optional) Save the response to the specified file.

## Usage

1. Configure the access key using the `configure` command.
2. Engage in chat by providing a prompt with the `chat` command.
3. Generate explanations using a file as input with the `explain` command.

## Environment Variable

The access key is stored in the environment variable `GEMINI_AI_ACCESS_KEY` after configuration.

## Examples

Configure access key:

```bash
gimini-cli configure -k YOUR_ACCESS_KEY
```

Chat with Gemini AI and save the response:

```bash
gimini-cli chat -p "Tell me a story" -s output.txt
```

Generate explanations using a file as input:

```bash
gimini-cli explain -f input.txt -s output.txt
```



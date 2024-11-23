# GO-DockerLint-AI

**GO-DockerLint-AI** is a tool written in Go designed to analyze and lint Dockerfiles, providing suggestions and corrections to improve the quality and efficiency of Docker images.

## Features

- Static analysis of Dockerfiles.
- Suggestions for optimization and fixing common mistakes.
- Multi-platform support (Linux, Windows, ARM, and x86_64).
- No external dependencies: standalone binary.
- Integration with OpenAI for advanced suggestions.

## System Requirements

- **Go** 1.20 or later.
- Operating System:
  - Linux
  - Windows ( experimental compatibility)
  - macOS (experimental compatibility)
- Docker (optional for testing).
- OpenAI API Key (required for functionality).

## Setting Up OpenAI Key

Before running the tool, export your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=your_openai_key
```

Make sure to replace `your_openai_key` with your actual OpenAI API key.

## Installation

### Clone the Repository

```bash
git clone https://github.com/la-plas-growth/GO-DockerLint-AI.git
cd GO-DockerLint-AI
```

### Build

#### For Linux (amd64):
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-linux-amd64 .
```

#### For Windows (amd64):
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-windows-amd64 .
```

You can also use the `build.sh` script to generate binaries for all supported platforms:

```bash
sh build.sh
```

### Using the Binary

1. Download the precompiled binary from the [releases](https://github.com/la-plas-growth/GO-DockerLint-AI/releases) section.
2. Add the binary to your `$PATH` or execute it directly.

## Usage

Run the following command to analyze a Dockerfile:

```bash
./go-dockerlint-ai-linux-amd64 dockerlint /path/to/Dockerfile
```

Example output:
```
[INFO] Dockerfile analysis completed.
[WARN] Use more recent base images.
[SUGGESTION] Combine RUN commands to reduce image layers.
```

### Available Commands

| Command             | Description                               |
|---------------------|-------------------------------------------|
| `dockerlint <file>`       | Lint the specified Dockerfile.            |
| `version`           | Show the tool version.                   |
| `help`              | Display available commands.              |

## Development

If you want to contribute or customize the project, make sure you have Go installed. Use the following commands to start the project in development mode:

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run tests:
   ```bash
   go test ./...
   ```

3. Start the tool:
   ```bash
   go run main.go dockerlint /path/to/Dockerfile
   ```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a branch for your feature/bugfix:
   ```bash
   git checkout -b my-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Description of the change"
   ```
4. Submit a pull request.

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions or issues, open an [issue](https://github.com/la-plas-growth/GO-DockerLint-AI/issues) or contact:

- **Author**: [la-plas-growth](https://github.com/la-plas-growth)

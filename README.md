# GO-Docker-AI

**GO-Docker-AI** is a tool written in Go designed to analyze and lint Dockerfiles, providing suggestions and corrections to improve the quality and efficiency of Docker images.

## Features

- Static analysis of Dockerfiles.
- Suggestions for optimization and fixing common mistakes.
- NOT YET: Multi-platform support, FOR NOW WE ONLY SUPPORT "linux amd64"
- No external dependencies: standalone binary.
- Integration with OpenAI for advanced suggestions.

## New features v0.2.0
- Creation of base Dockerfile based on best practices.
  

## System Requirements

- **Go** 1.20 or later.
- Operating System:
  - Linux
  - NOT YET: Windows
  - NOT YET: macOS
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
git clone https://github.com/la-plas-growth/go-docker-ai.git
cd go-docker-ai
```

### Build

#### For Linux (amd64):
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o go-docker-ai-linux-amd64 .
```

#### For Linux (arm64):
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -a -o go-docker-ai-linux-arm64 .
```

#### For Windows (amd64) - NOT TESTED YET:
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -o go-Docker-ai-windows-amd64 .
```

You can also use the `build.sh` script to generate binaries for all supported platforms:

```bash
sh build.sh
```

### Using the Binary

1. Download the precompiled binary from the [releases](https://github.com/la-plas-growth/go-docker-ai/releases) section.
2. Add the binary to your `$PATH` or execute it directly.

## Usage

Run the following command to analyze a Dockerfile:

```bash
./go-docker-ai-linux-amd64 dockerfile lint -f /path/to/Dockerfile
```

Example output:
```
{
   issues": [
          {
                  "issue": "Using a large base image without necessity.",
                  "severity": "high",
                  "advice": "Consider using a leaner base image, such as `alpine` without glibc dependencies, to minimize the size and security risks."
          },
          {
                  "issue": "Base image version is mutable and not pinned.",
                  "severity": "medium",
                  "advice": "It's advisable to pin the base image version using a specific digest instead of just `alpine-3.9_glibc-2.29` to avoid potential issues with future updates."
          },
          {
                  "issue": "No explicit update before package installation.",
                  "severity": "medium",
                  "advice": "Run `apk update` before installing packages to ensure you're getting the latest available versions of packages."       
          }
   ]
}                 
```

Run the followinf command to create a base Dockerfile:

```bash
./go-docker-ai-linux-amd64 dockerfile create -t golang
```

Output: the file will be created in the folder where the command is been executed

### Available Commands

| Command               | Description                               |
|---------------------  |-------------------------------------------|
| `dockerfile create -t <lang> -p <path>`| Create a base dockerfile in any language  |
| `dockerifle lint -f <file>`   | Lint the specified Dockerfile.            |
| `version`             | Show the tool version.                    |
| `help`                | Display available commands.               |

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
   go run main.go <command> <arg>
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

For questions or issues, open an [issue](https://github.com/la-plas-growth/GO-Docker-AI/issues) or contact:

- **Author**: [la-plas-growth](https://github.com/la-plas-growth)
- **Author**: [Allan-Nava](https://github.com/Allan-Nava)

# micbun.io-backend

Backend for micbun.io, using gRPC, and Redis.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Requirements

- Go 1.18 or higher
- Redis server
- gRPC

## Installation

1. Clone the repository:

```bash
git clone https://github.com/MicBun/micbun.io-backend.git
cd micbun.io-backend
```

2. Install dependencies:

```bash
go mod tidy
```

## Configuration

1. Copy the example environment file and update the values as needed:

```bash
cp .env.example .env
```

2. Ensure Redis server is running and update the Redis connection details in the `.env` file.

## Running the Server

1. Start the gRPC server:

```bash
go run cmd/server/main.go
```

2. The server should now be running and accessible via the configured gRPC port.

## Testing

1. Run tests using the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

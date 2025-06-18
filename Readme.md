# GoBank - Simple Banking API

A lightweight banking API built with Go that provides basic banking operations through RESTful endpoints.

## Features

- Account management (Create, Read, Delete operations)
- JSON-based API responses
- Error handling middleware
- Gorilla Mux router for efficient routing

## Project Structure

```
gobank/
  ├── api.go      # API handlers and server implementation
  ├── types.go    # Data structures and types
  ├── main.go     # Application entry point
  ├── go.mod      # Go module file
  ├── go.sum      # Go module checksum
  └── Makefile    # Build automation
```

## API Endpoints

### Account Management

#### GET /account
Retrieves account information.

**Response:**
```json
{
    "id": integer,
    "firstName": string,
    "lastName": string,
    "number": integer,
    "balance": integer
}
```

#### POST /account
Creates a new account.

#### DELETE /account
Deletes an existing account.

## Getting Started

### Prerequisites

- Go 1.x or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone github.com/tobiadiks/gobank
cd gobank
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run .
```

The server will start on the configured port (check main.go for the port configuration).

## Development

The project uses the following Go packages:
- `github.com/gorilla/mux` for routing
- Standard Go libraries for JSON handling and HTTP server

## Error Handling

The API includes a custom error handling middleware that returns errors in the following format:

```json
{
    "error": "error message"
}
```

## Contributing

Feel free to submit issues and enhancement requests.

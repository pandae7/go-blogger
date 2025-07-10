# Go Blogger

A simple blogging platform built with Go for the backend and a demo gRPC client inside the `cmd` package.

## Features

- CRUD operations for blog posts via gRPC
- gRPC server implementation
- Demo Go client in `cmd/client`

## Prerequisites

- [Go](https://golang.org/dl/) 1.18+
- [protoc](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go) and [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc) plugins

## Installation

1. **Clone the repository**
    ```bash
    git clone https://github.com/pandae7/go-blogger.git
    cd go_blogger
    ```

2. **Install Go dependencies**
    ```bash
    go mod tidy
    ```

3. **Generate gRPC code (if needed)**
    ```bash
    protoc --go_out=. --go-grpc_out=. proto/*.proto
    ```

## Running the Application

### Start the gRPC server

```bash
go run cmd/server/main.go
```
The server listens on `localhost:8080` (or as configured).

### Run the demo client

```bash
go run cmd/client/main.go
```
The client will connect to the running gRPC server and demonstrate API usage.

## Testing

```bash
go test ./...
```

## gRPC Methods

- `CreateBlogPost` — Create a new blog post
- `GetBlogPost` — Get a post by ID
- `UpdateBlogPost` — Update a post by ID
- `DeleteBlogPost` — Delete a post by ID

> **Note:** There is currently no method to fetch all posts.

## License

MIT

### References

For gRPC Server & Client - 
https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05

For Folder Structure - 
https://github.com/pandae7/go-ai-stream

For Interfaces, methods and Tests - 
https://opendev.org/starlingx/app-node-interface-metrics-exporter/commit/423e49980e0f1e28e94c5ceeb6fa128854f8aeae 

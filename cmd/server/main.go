package main

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	defaultPort = "8080"
	defaultHost = "localhost"
)

func main() {
	port := defaultPort
	host := defaultHost

	// Create network listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Infof("Starting server on %s:%s", host, port)

	// Create a new gRPC server instance
	newServer := grpc.NewServer()

		go func() {
		log.Printf("gRPC server starting on port %s", port)
		log.Printf("Server ready to accept connections...")
		
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}

func printServerInfo(port string) {
	fmt.Println("===========================================")
	fmt.Println("          gRPC Blog Service")
	fmt.Println("===========================================")
	fmt.Printf("Server Address: localhost:%s\n", port)
	fmt.Println("Available Methods:")
	fmt.Println("  - CreatePost")
	fmt.Println("  - GetPost")
	fmt.Println("  - UpdatePost")
	fmt.Println("  - DeletePost")
	fmt.Println("")
	fmt.Println("Test with grpcurl:")
	fmt.Printf("  grpcurl -plaintext localhost:%s list\n", port)
	fmt.Printf("  grpcurl -plaintext localhost:%s blog.v1.BlogService/GetPost\n", port)
	fmt.Println("===========================================")
}



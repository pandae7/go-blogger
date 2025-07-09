package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pandae7/go-blogger/internal/server"
	storage "github.com/pandae7/go-blogger/internal/storage"
	pb "github.com/pandae7/go-blogger/proto/blog"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	storage := storage.NewBlogStorage()

	// Create a new gRPC server instance
	newServer := grpc.NewServer()

	// creating a default blog service server for now
	blogserver := server.NewBlogServiceServer(storage)

	// register blog service server
	pb.RegisterBlogServiceServer(newServer, blogserver)
	// Print server information
	printServerInfo(host, port)

	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	newServer.GracefulStop()
	log.Println("Server stopped")
}

func printServerInfo(host string, port string) {
	fmt.Println("===========================================")
	fmt.Println("          gRPC Blog Service")
	fmt.Println("===========================================")
	fmt.Printf("Server Address: %s:%s\n", host, port)
	fmt.Println("Available Methods:")
	fmt.Println("  - CreateBlogPost")
	fmt.Println("  - GetBlogPost")
	fmt.Println("  - UpdateBlogPost")
	fmt.Println("  - DeleteBlogPost")
	fmt.Println("===========================================")
}

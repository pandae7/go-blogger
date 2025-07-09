package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/pandae7/go-blogger/proto/blog"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Creating a new blog post...")
	createBlogReq := &pb.CreateBlogPostRequest{
		Title:           "My First Blog Post",
		Content:         "This is the content of my first blog post.",
		Author:          "Aman Pandae",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"trending", "topic", "cloud"},
	}

	createBlogResp, err := client.CreateBlogPost(ctx, createBlogReq)
	if err != nil {
		log.Fatal("Failed to create blog post: %v", err)
	}

	if !createBlogResp.Success {
		log.Fatal("Failed to create blog post: %s", createBlogResp.Message)
	}

	fmt.Printf("Blog post created successfully with ID: %s\n", createBlogResp.Post.PostId)
	printBlogPostDetails(createBlogResp.Post)

	fmt.Println("Fetch created Blog Post...")
	getBlogReq := &pb.GetBlogPostRequest{
		PostId: createBlogResp.Post.PostId,
	}

	getBlogResp, err := client.GetBlogPost(ctx, getBlogReq)
	if err != nil {
		log.Fatal("Failed to fetch blog post: %v", err)
	}

	if !getBlogResp.Success {
		log.Fatal("Failed to fetch blog post: %s", getBlogResp.Message)
	}
	fmt.Printf("Blog post fetched successfully with ID: %s\n", getBlogResp.Post.PostId)
	printBlogPostDetails(getBlogResp.Post)

	fmt.Println("Updating the blog post...")
	updateBlogReq := &pb.UpdateBlogPostRequest{
		PostId:  createBlogResp.Post.PostId,
		Title:   "Updated Blog Post Title",
		Content: "This is the updated content of my blog post.",
		Tags:    []string{"news", "trends", "views", "updates"},
	}
	updateBlogResp, err := client.UpdateBlogPost(ctx, updateBlogReq)
	if err != nil {
		log.Fatal("Failed to update blog post: %v", err)
	}

	if !updateBlogResp.Success {
		log.Fatal("Failed to update blog post: %s", updateBlogResp.Message)
	}

	fmt.Printf("Blog post updated successfully with ID: %s\n", updateBlogResp.Post.PostId)
	printBlogPostDetails(updateBlogResp.Post)

	fmt.Println("Deleting the blog post...")
	deleteBlogReq := &pb.DeleteBlogPostRequest{
		PostId: createBlogResp.Post.PostId,
	}
	deleteBlogResp, err := client.DeleteBlogPost(ctx, deleteBlogReq)
	if err != nil {
		log.Fatal("Failed to delete blog post: %v", err)
	}
	if !deleteBlogResp.Success {
		log.Fatal("Failed to delete blog post: %s", deleteBlogResp.Message)
	}
	fmt.Printf("Blog post deleted successfully with ID: %s\n", deleteBlogReq.PostId)

}

func printBlogPostDetails(post *pb.BlogPost) {
	fmt.Println("******Blog Post Details:******")
	fmt.Printf("ID: %s\n", post.PostId)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Content: %s\n", post.Content)
	fmt.Printf("Author: %s\n", post.Author)
	fmt.Printf("Publication Date: %s\n", post.PublicationDate.AsTime().Format(time.RFC3339))
	fmt.Printf("Tags: %v\n", post.Tags)
	fmt.Printf("Updated At: %s\n", post.UpdatedAt.AsTime().Format(time.RFC3339))
	fmt.Println("*******************************")
}

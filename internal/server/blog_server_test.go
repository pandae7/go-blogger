package server

import (
	"context"
	"errors"
	"testing"
	"time"

	models "github.com/pandae7/go-blogger/internal/models"
	pb "github.com/pandae7/go-blogger/proto/blog"
)

// Mock storage for testing
type mockBlogStorage struct {
	CreatePostFunc func(ctx context.Context, post *models.BlogPost) error
	GetPostFunc    func(ctx context.Context, postID string) (*models.BlogPost, error)
	UpdatePostFunc func(ctx context.Context, req *models.UpdateBlogPostRequest) (*models.BlogPost, error)
	DeletePostFunc func(ctx context.Context, postID string) error
}

func (m *mockBlogStorage) CreatePost(ctx context.Context, post *models.BlogPost) error {
	return m.CreatePostFunc(ctx, post)
}
func (m *mockBlogStorage) GetPost(ctx context.Context, postID string) (*models.BlogPost, error) {
	return m.GetPostFunc(ctx, postID)
}
func (m *mockBlogStorage) UpdatePost(ctx context.Context, req *models.UpdateBlogPostRequest) (*models.BlogPost, error) {
	return m.UpdatePostFunc(ctx, req)
}
func (m *mockBlogStorage) DeletePost(ctx context.Context, postID string) error {
	return m.DeletePostFunc(ctx, postID)
}

func TestCreateBlogPost_Success(t *testing.T) {
	mockStorage := &mockBlogStorage{
		CreatePostFunc: func(ctx context.Context, post *models.BlogPost) error {
			return nil
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.CreateBlogPostRequest{
		Title:   "Blog Test",
		Content: "Test Blog Content",
		Author:  "NotAman",
		Tags:    []string{"test1", "test2"},
	}
	resp, err := server.CreateBlogPost(context.Background(), req)
	if err != nil || !resp.Success {
		t.Errorf("expected success, got error: %v, resp: %+v", err, resp)
	}
}

func TestCreateBlogPost_InvalidRequest(t *testing.T) {
	mockStorage := &mockBlogStorage{}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.CreateBlogPostRequest{
		Title:   "",
		Content: "",
		Author:  "",
	}
	resp, err := server.CreateBlogPost(context.Background(), req)
	if err == nil || resp.Success {
		t.Errorf("expected error for invalid request, got: %v, resp: %+v", err, resp)
	}
}

func TestGetBlogPost_Success(t *testing.T) {
	mockStorage := &mockBlogStorage{
		GetPostFunc: func(ctx context.Context, postID string) (*models.BlogPost, error) {
			return &models.BlogPost{
				PostId:          postID,
				Title:           "Title",
				Content:         "Content",
				Author:          "Author",
				PublicationDate: time.Now(),
				Tags:            []string{"go"},
				UpdatedAt:       time.Now(),
			}, nil
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.GetBlogPostRequest{PostId: "123"}
	resp, err := server.GetBlogPost(context.Background(), req)
	if err != nil || !resp.Success {
		t.Errorf("expected success, got error: %v, resp: %+v", err, resp)
	}
}

func TestGetBlogPost_NotFound(t *testing.T) {
	mockStorage := &mockBlogStorage{
		GetPostFunc: func(ctx context.Context, postID string) (*models.BlogPost, error) {
			return nil, errors.New("not found")
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.GetBlogPostRequest{PostId: "not-exist"}
	resp, err := server.GetBlogPost(context.Background(), req)
	if err == nil || resp.Success {
		t.Errorf("expected error for not found, got: %v, resp: %+v", err, resp)
	}
}

func TestUpdateBlogPost_Success(t *testing.T) {
	mockStorage := &mockBlogStorage{
		UpdatePostFunc: func(ctx context.Context, req *models.UpdateBlogPostRequest) (*models.BlogPost, error) {
			return &models.BlogPost{
				PostId:          req.PostId,
				Title:           req.Title,
				Content:         req.Content,
				Tags:            req.Tags,
				Author:          "Author",
				PublicationDate: time.Now(),
				UpdatedAt:       req.UpdatedAt,
			}, nil
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.UpdateBlogPostRequest{
		PostId:  "123",
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"go", "update"},
	}
	resp, err := server.UpdateBlogPost(context.Background(), req)
	if err != nil || !resp.Success {
		t.Errorf("expected success, got error: %v, resp: %+v", err, resp)
	}
}

func TestUpdateBlogPost_InvalidRequest(t *testing.T) {
	mockStorage := &mockBlogStorage{}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.UpdateBlogPostRequest{
		PostId: "",
	}
	resp, err := server.UpdateBlogPost(context.Background(), req)
	if err == nil || resp.Success {
		t.Errorf("expected error for invalid update request, got: %v, resp: %+v", err, resp)
	}
}

func TestDeleteBlogPost_Success(t *testing.T) {
	mockStorage := &mockBlogStorage{
		DeletePostFunc: func(ctx context.Context, postID string) error {
			return nil
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.DeleteBlogPostRequest{PostId: "123"}
	resp, err := server.DeleteBlogPost(context.Background(), req)
	if err != nil || !resp.Success {
		t.Errorf("expected success, got error: %v, resp: %+v", err, resp)
	}
}

func TestDeleteBlogPost_Failure(t *testing.T) {
	mockStorage := &mockBlogStorage{
		DeletePostFunc: func(ctx context.Context, postID string) error {
			return errors.New("delete failed")
		},
	}
	server := NewBlogServiceServer(mockStorage)
	req := &pb.DeleteBlogPostRequest{PostId: "fail"}
	resp, err := server.DeleteBlogPost(context.Background(), req)
	if err == nil || resp.Success {
		t.Errorf("expected error for delete failure, got: %v, resp: %+v", err, resp)
	}
}

func TestModelToProtobuf(t *testing.T) {
	server := NewBlogServiceServer(nil)
	now := time.Now()
	post := &models.BlogPost{
		PostId:          "id",
		Title:           "title",
		Content:         "content",
		Author:          "author",
		PublicationDate: now,
		UpdatedAt:       now,
		Tags:            []string{"go"},
	}
	pbPost := server.modelToProtobuf(post)
	if pbPost.PostId != post.PostId || pbPost.Title != post.Title || pbPost.Author != post.Author {
		t.Errorf("modelToProtobuf did not map fields correctly")
	}
	if pbPost.PublicationDate.AsTime().Unix() != now.Unix() {
		t.Errorf("PublicationDate not mapped correctly")
	}
}

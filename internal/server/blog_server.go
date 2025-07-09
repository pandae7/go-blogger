package server

import (
	"context"
	"time"

	"github.com/google/uuid"
	models "github.com/pandae7/go-blogger/internal/models"
	storage "github.com/pandae7/go-blogger/internal/storage"
	pb "github.com/pandae7/go-blogger/proto/blog"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type BlogServiceServer struct {
	pb.UnimplementedBlogServiceServer
	storage storage.BlogStorage
}

func NewBlogServiceServer(storage storage.BlogStorage) *BlogServiceServer {
	return &BlogServiceServer{
		storage: storage,
	}
}

func (s *BlogServiceServer) CreateBlogPost(ctx context.Context, req *pb.CreateBlogPostRequest) (*pb.CreateBlogPostResponse, error) {
	log.Infof("Creating new post with title: %s", req.GetTitle())

	if err := s.validateCreatePostRequest(req); err != nil {
		log.Errorf("Invalid request: %v", err)
		return &pb.CreateBlogPostResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	// check PublishedDate
	publicationDate := req.GetPublicationDate()
	if publicationDate == nil {
		log.Warnf("Publication date is not set, using current time")
		publicationDate = timestamppb.Now()
	}

	post := &models.BlogPost{
		PostId:          uuid.New().String(),
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		PublicationDate: publicationDate.AsTime(),
		Tags:            req.GetTags(),
		UpdatedAt:       time.Now(),
	}

	if err := s.storage.CreatePost(ctx, post); err != nil {
		log.Errorf("Failed to create post: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create post: %v", err)
	}

	log.Infof("Post created successfully with ID: %s", post.PostId)
	return &pb.CreateBlogPostResponse{
		Post:    s.modelToProtobuf(post),
		Success: true,
		Message: "Post created successfully",
	}, nil
}

func (s *BlogServiceServer) GetBlogPost(ctx context.Context, req *pb.GetBlogPostRequest) (*pb.GetBlogPostResponse, error) {
	log.Infof("Retrieving post with ID: %s", req.GetPostId())

	post, err := s.storage.GetPost(ctx, req.GetPostId())
	if err != nil {
		return &pb.GetBlogPostResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.GetBlogPostResponse{
		Post:    s.modelToProtobuf(post),
		Success: true,
		Message: "Post retrieved successfully",
	}, nil
}

func (s *BlogServiceServer) UpdateBlogPost(ctx context.Context, req *pb.UpdateBlogPostRequest) (*pb.UpdateBlogPostResponse, error) {
	log.Infof("Updating post with ID: %s", req.GetPostId())

	if err := s.validateUpdatePostRequest(req); err != nil {
		return &pb.UpdateBlogPostResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	updateReq := &models.UpdateBlogPostRequest{
		PostId:    req.GetPostId(),
		Title:     req.GetTitle(),
		Content:   req.GetContent(),
		Tags:      req.GetTags(),
		UpdatedAt: time.Now(),
	}

	updatedPost, err := s.storage.UpdatePost(ctx, updateReq)
	if err != nil {
		return &pb.UpdateBlogPostResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.UpdateBlogPostResponse{
		Post:    s.modelToProtobuf(updatedPost),
		Success: true,
		Message: "Post updated successfully",
	}, nil
}

func (s *BlogServiceServer) DeleteBlogPost(ctx context.Context, req *pb.DeleteBlogPostRequest) (*pb.DeleteBlogPostResponse, error) {
	log.Infof("Deleting post with ID: %s", req.GetPostId())

	if err := s.storage.DeletePost(ctx, req.GetPostId()); err != nil {
		return &pb.DeleteBlogPostResponse{
			Success: false,
			Message: "Failed to delete post: " + err.Error(),
		}, err
	}

	log.Infof("Post deleted successfully with ID: %s", req.GetPostId())
	return &pb.DeleteBlogPostResponse{
		Success: true,
		Message: "Post deleted successfully",
	}, nil
}

func (s *BlogServiceServer) validateCreatePostRequest(req *pb.CreateBlogPostRequest) error {
	if req.GetTitle() == "" {
		return status.Error(codes.InvalidArgument, "Post title cannot be empty")
	}
	if req.GetContent() == "" {
		return status.Error(codes.InvalidArgument, "Post content cannot be empty")
	}
	if req.GetAuthor() == "" {
		return status.Error(codes.InvalidArgument, "Post author cannot be empty")
	}
	return nil
}

func (s *BlogServiceServer) validateUpdatePostRequest(req *pb.UpdateBlogPostRequest) error {
	if req.GetPostId() == "" {
		return status.Error(codes.InvalidArgument, "Post ID cannot be empty")
	}
	if req.GetTitle() == "" && req.GetContent() == "" && len(req.GetTags()) == 0 {
		return status.Error(codes.InvalidArgument, "At least one field (title, content, tags) must be provided for update")
	}
	return nil
}

func (s *BlogServiceServer) modelToProtobuf(post *models.BlogPost) *pb.BlogPost {
	return &pb.BlogPost{
		PostId:          post.PostId,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: timestamppb.New(post.PublicationDate),
		UpdatedAt:       timestamppb.New(post.UpdatedAt),
		Tags:            post.Tags,
	}
}

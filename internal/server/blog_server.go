package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)


type BlogServiceServer struct {
	pb.UnimplementedBlogServiceServer
	storage storage.BlogStorage
}

func (s *BlogServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	log.Info("Creating new post with title: %s", req.GetTitle())

	if err := s.validateCreatePostRequest(req); err != nil {
		log.Error("Invalid request: %v", err)
		return pb.CreateBlogPostResponse{
			success: false,
			message: err.Error(),
		}, err
	}

	// check PublishedDate
	publicationDate := req.GetPublicationDate()
	if publicationDate == nil {
		log.Warn("Publication date is not set, using current time")
		publicationDate = timestamppb.Now()
	}

	post := &models.Post{
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		PublicationDate: publicationDate.AsTime(),
		Tags:            req.GetTags(),
		updatedAt:       time.Now(),
	}

	if err := s.storage.CreatePost(ctx, post); err != nil {
		log.Error("Failed to create post: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create post: %v", err)
	}

	log.Info("Post created successfully with ID: %s", post.PostID)
	return &pb.CreateBlogPostResponse{
		Post: s.modelToProtobuf(post),
		Success: true,
		Message: "Post created successfully",
	}, nil
}


func (s *BlogServiceServer) validateCreatePostRequest(req *pb.CreatePostRequest) error {
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

func (s *BlogServiceServer) modelToProtobuf(post *models.Post) *pb.Post {
	return &pb.Post{
		PostId:          post.PostID,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: timestamppb.New(post.PublicationDate),
		UpdatedAt:       timestamppb.New(post.UpdatedAt),
		Tags:            post.Tags,
	}
}
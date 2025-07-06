package storage

import (
	"context"
	"time"
	"sync"
	"github.com/pandae7/go-blogger/internal/models" 
)

// BlogStorage defines the interface for blog-related storage operations.
type BlogStorage interface {
	// CreatePost creates a new blog post in the storage.
	CreatePost(ctx context.Context, post *models.BlogPost) error

	// GetPost retrieves a blog post by its ID.
	GetPost(ctx context.Context, postId string) (*models.BlogPost, error)

	// UpdatePost updates an existing blog post.
	UpdatePost(ctx context.Context, post *models.UpdateBlogPostRequest) (*models.BlogPost, error)

	// DeletePost deletes a blog post by its ID.
	DeletePost(ctx context.Context, postId string) error	

}

// AuthorStorage defines the interface for author-related storage operations.
type AuthorStorage interface {
	// CreateAuthor creates a new author in the storage.
	CreateAuthor(ctx context.Context, author *models.Author) error

	// GetAuthor retrieves an author by their ID.
	GetAuthor(ctx context.Context, authorId string) (*models.Author, error)

	// UpdateAuthor updates an existing author's information.
	UpdateAuthor(ctx context.Context, author *models.Author) (*models.Author, error)

	// DeleteAuthor deletes an author by their ID.
	DeleteAuthor(ctx context.Context, authorId string) error
}

type BlogStorageImpl struct {
	// In Memory storage
	posts map[string]*models.BlogPost
	
	// mu protects concurrent access to the posts map
	mu sync.RWMutex
	
	// createdAt tracks when the Blogs storage was created.
	createdAt time.Time
}

type AuthorStorageImpl struct {
	// In Memory storage
	authors map[string]*models.author

	// mu protects concurrent access to the authors map
	mu sync.RWMutex

	// createdAt tracks when the Author storage was created.
	createdAt time.Time
}

func NewBlogStorage() *BlogStorageImpl {
	return &BlogStorageImpl{
		posts: make(map[string]*models.BlogPost),
		createdAt: time.Now(),
	}
}

func NewAuthorStorage() *AuthorStorageImpl {
	return &AuthorStorageImpl{
		authors: make(map[string]*models.Author),
		createdAt: time.Now(),
	}
}
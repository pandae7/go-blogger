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

func (s *BlogStorageImpl) CreatePost(ctx context.Context, post *models.BlogPost) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if post already exists
	if _, exists := s.posts[post.PostId]; exists {
		return models.ErrDuplicatePost
	}

	now = time.Now()
	// Set the publication date if not provided
	if post.PublicationDate.IsZero() {
		post.PublicationDate = now
	}
	// Set the updated at time
	post.UpdatedAt = now
	
	// Add the post to the storage
	s.posts[post.PostId] = post
	return nil
}

func (s *BlogStorageImpl) GetPost(ctx context.Context, postId string) (*models.BlogPost, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retrieve the post by ID
	post, exists := s.posts[postId]
	if !exists {
		return nil, models.ErrPostNotFound
	}
	return post, nil
}

func (s *BlogStorageImpl) UpdatePost(ctx context.Context, post *models.UpdateBlogPostRequest) (*models.BlogPost, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Retrieve the existing post
	existingPost, exists := s.posts[post.PostId]
	if !exists {
		return nil, models.ErrPostNotFound
	}

	// Update fields if provided
	if post.Title != "" {
		existingPost.Title = post.Title
	}
	if post.Content != "" {
		existingPost.Content = post.Content
	}
	if len(post.Tags) > 0 {
		existingPost.Tags = post.Tags
	}
	existingPost.UpdatedAt = time.Now()

	return existingPost, nil
}

func (s *BlogStorageImpl) DeletePost(ctx context.Context, postId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the post exists
	if _, exists := s.posts[postId]; !exists {
		return models.ErrPostNotFound
	}

	// Delete the post
	delete(s.posts, postId)
	return nil
}

func (s *AuthorStorageImpl) CreateAuthor(ctx context.Context, author *models.Author) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if author already exists
	if _, exists := s.authors[author.AuthorId]; exists {
		return models.ErrDuplicateAuthor
	}

	now := time.Now()
	// Set created and updated times
	author.CreatedAt = now
	author.UpdatedAt = now

	// Add the author to the storage
	s.authors[author.AuthorId] = author
	return nil
}

func (s *AuthorStorageImpl) GetAuthor(ctx context.Context, authorId string) (*models.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retrieve the author by ID
	author, exists := s.authors[authorId]
	if !exists {
		return nil, models.ErrAuthorNotFound
	}
	return author, nil
}	

func (s *AuthorStorageImpl) UpdateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Retrieve the existing author
	existingAuthor, exists := s.authors[author.AuthorId]
	if !exists {
		return nil, models.ErrAuthorNotFound
	}

	// Update fields if provided
	if author.Name != "" {
		existingAuthor.Name = author.Name
	}
	if author.Email != "" {
		existingAuthor.Email = author.Email
	}
	if author.ProfileName != "" {
		existingAuthor.ProfileName = author.ProfileName
	}
	existingAuthor.UpdatedAt = time.Now()

	return existingAuthor, nil
}

func (s *AuthorStorageImpl) DeleteAuthor(ctx context.Context, authorId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the author exists
	if _, exists := s.authors[authorId]; !exists {
		return models.ErrAuthorNotFound
	}

	// Delete the author
	delete(s.authors, authorId)
	return nil
}
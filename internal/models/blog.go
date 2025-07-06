package models

import "time"

type BlogPost struct {
	PostId      string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	publicationDate time.Time `json:"publication_date"`
	UpdatedAt  time.Time `json:"updated_at"`
	Tags        []string `json:"tags"`
}

type Author struct {
	AuthorId   string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfileName string `json:"profile_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateBlogPostRequest struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Author      string   `json:"author"`
	PublicationDate time.Time `json:"publication_date,omitempty"`
	Tags        []string `json:"tags"`
}

type UpdateBlogPostRequest struct {
	PostId      string   `json:"id"`
	Title       string   `json:"title,omitempty"`
	Content     string   `json:"content,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CreateBlogPostResponse struct {
	Post        *BlogPost `json:"post"`
	Success     bool      `json:"success"`
	Message     string    `json:"message,omitempty"`
}
	
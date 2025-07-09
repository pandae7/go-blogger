package models

import "errors"

// Error constants
var (
	ErrPostNotFound    = errors.New("post not found")
	ErrAuthorNotFound  = errors.New("author not found")
	ErrTagNotFound     = errors.New("tag not found")
	ErrInvalidPostID   = errors.New("invalid post ID")
	ErrInvalidAuthorID = errors.New("invalid author ID")
	ErrInvalidTagID    = errors.New("invalid tag ID")
	ErrEmptyTitle      = errors.New("post title cannot be empty")
	ErrEmptyContent    = errors.New("post content cannot be empty")
	ErrEmptyAuthor     = errors.New("post author cannot be empty")
	ErrDuplicatePost   = errors.New("post with this ID already exists")
)

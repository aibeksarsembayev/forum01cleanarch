package models

import (
	"context"
	"time"
)

// Post model domain ...
type Post struct {
	PostID       int       `json:"post_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	CategoryName string    `json:"category_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// PostRequestDTO ...
type PostRequestDTO struct {
	PostID       int       `json:"post_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	CategoryName string    `json:"category_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	VoteLike     int       `json:"vote_like"`
	VoteDislike  int       `json:"vote_dislike"`
}

// PostCreateRequestDTO ...
type PostCreateRequestDTO struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	CategoryName string `json:"category_name"`
}

// PostResponceDTO ...
type PostResponceDTO struct {
	PostID       int       `json:"post_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	CategoryName string    `json:"category_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// PostUsecase represents posts's usecases
type PostUsecase interface {
	Create(ctx context.Context, post *Post) (id int, err error)
	Update(ctx context.Context, post *Post) (err error)
	GetAll(ctx context.Context) (*[]PostRequestDTO, error)
	GetByID(ctx context.Context, id int) (post *PostRequestDTO, err error)
	GetByCategory(ctx context.Context, category string) (*[]PostRequestDTO, error)
	GetByUserID(ctx context.Context, id int) (*[]PostRequestDTO, error)
	GetLikedByUserID(ctx context.Context, id int) (*[]PostRequestDTO, error)
	Delete(ctx context.Context, id int) (err error)
}

// PostRepository represent post's repository contract
type PostRepository interface {
	Create(ctx context.Context, post *Post) (id int, err error)
	Update(ctx context.Context, post *Post) (err error)
	GetAll(ctx context.Context) (*[]PostRequestDTO, error)
	GetByID(ctx context.Context, id int) (post *PostRequestDTO, err error)
	GetByCategory(ctx context.Context, category string) (*[]PostRequestDTO, error)
	GetByUserID(ctx context.Context, id int) (*[]PostRequestDTO, error)
	GetLikedByUserID(ctx context.Context, id int) (*[]PostRequestDTO, error)
	Delete(ctx context.Context, id int) (err error)
}

package models

import (
	"context"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	Content   string    `json:"content"`
	PostID    int       `json:"Comment_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentRequestDTO ...
type CommentRequestDTO struct {
	CommentID int       `json:"comment_id"`
	Content   string    `json:"content"`
	PostID    int       `json:"Comment_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentUsecase represents Comment's usecases
type CommentUsecase interface {
	Create(ctx context.Context, comment *Comment) (id int, err error)
	Update(ctx context.Context, comment *Comment) (err error)
	GetByUserID(ctx context.Context, user_id int) (comments []*CommentRequestDTO, err error)
	GetByPostID(ctx context.Context, post_id int) (comments []*CommentRequestDTO, err error)
	Delete(ctx context.Context, id int) (err error)
}

// CommentRepository represent Comment's repository contract
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (id int, err error)
	Update(ctx context.Context, comment *Comment) (err error)
	GetByUserID(ctx context.Context, user_id int) (comments []*CommentRequestDTO, err error)
	GetByPostID(ctx context.Context, post_id int) (comments []*CommentRequestDTO, err error)
	Delete(ctx context.Context, id int) (err error)
}

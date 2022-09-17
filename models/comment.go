package models

import (
	"context"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	Content   string    `json:"comment_content"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentInputRequestDTO to retrieve input data coming from JSON to create comment ...
type CommentInputRequestDTO struct {
	Content string `json:"comment_content"`
	PostID  string `json:"post_id"`
	// UserID  int    `json:"user_id"`
}

// CommentRequestDTO ...
type CommentRequestDTO struct {
	CommentID          int       `json:"comment_id"`
	Content            string    `json:"comment_content"`
	PostID             int       `json:"post_id"`
	UserID             int       `json:"user_id"`
	Username           string    `json:"username"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CommentVoteLike    int       `json:"comment_vote_like"`
	CommentVoteDislike int       `json:"comment_vote_dislike"`
}

// CommentUsecase represents Comment's usecases
type CommentUsecase interface {
	Create(ctx context.Context, comment *Comment) (id int, err error)
	Update(ctx context.Context, comment *Comment) (err error)
	GetByID(ctx context.Context, id int) (comment *CommentRequestDTO, err error)
	GetByUserID(ctx context.Context, user_id int) (comments *[]CommentRequestDTO, err error)
	GetByPostID(ctx context.Context, post_id int) (comments *[]CommentRequestDTO, err error)
	Delete(ctx context.Context, id int) (err error)
}

// CommentRepository represent Comment's repository contract
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (id int, err error)
	Update(ctx context.Context, comment *Comment) (err error)
	GetByID(ctx context.Context, id int) (comment *CommentRequestDTO, err error)
	GetByUserID(ctx context.Context, user_id int) (comments *[]CommentRequestDTO, err error)
	GetByPostID(ctx context.Context, post_id int) (comments *[]CommentRequestDTO, err error)
	Delete(ctx context.Context, id int) (err error)
}

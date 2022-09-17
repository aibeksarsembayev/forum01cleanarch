package models

import (
	"context"
	"time"
)

// CommentVote ...
type CommentVote struct {
	CommentVoteID    int       `json:"comment_vote_id"`
	CommentVoteValue bool      `json:"comment_vote_value"`
	UserID           int       `json:"user_id"`
	CommentID        int       `json:"comment_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CommentVoteInputRequestDTO ...
type CommentVoteInputRequestDTO struct {
	CommentVoteValue string `json:"comment_vote_value"`
	CommentID        string `json:"comment_id"`
}

// CommentVoteCreateRequestDTO ...
type CommentVoteCreateRequestDTO struct {
	CommentVoteValue bool `json:"comment_vote_value"`
	UserID           int  `json:"user_id"`
	CommentID        int  `json:"comment_id"`
}

// CommentResponseDTO ...
type CommentVoteResponseDTO struct {
	CommentID int `json:"comment_id"`
	UserID    int `json:"user_id"`
}

// CommentResponseDTO ...
type CommentVoteCountResponseDTO struct {
	CommentID        int  `json:"comment_id"`
	CommentVoteValue bool `json:"comment_vote_value"`
}

// CommentVoteUsecase ...
type CommentVoteUsecase interface {
	Create(ctx context.Context, commentVote *CommentVote) (id int, err error)
	Update(ctx context.Context, commentVote *CommentVote) (err error)
	Get(ctx context.Context, commentVoteResponse *CommentVoteResponseDTO) (*CommentVote, error)
	GetCount(ctx context.Context, commentVoteResponse *CommentVoteCountResponseDTO) (int, error)
	Delete(ctx context.Context, id int) (err error)
}

// CommentVoteRepository ...
type CommentVoteRepository interface {
	Create(ctx context.Context, commentVote *CommentVote) (id int, err error)
	Update(ctx context.Context, commentVote *CommentVote) (err error)
	Get(ctx context.Context, commentVoteResponse *CommentVoteResponseDTO) (*CommentVote, error)
	GetCount(ctx context.Context, commentVoteResponse *CommentVoteCountResponseDTO) (int, error)
	Delete(ctx context.Context, id int) (err error)
}

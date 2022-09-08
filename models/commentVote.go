package models

import (
	"context"
	"time"
)

// CommentVote ...
type CommentVote struct {
	CommentVoteID int       `json:"comment_vote_id"`
	Value         bool      `json:"value"`
	UserID        int       `json:"user_id"`
	CommentID     int       `json:"comment_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CommentVoteUsecase ...
type CommentVoteUsecase interface {
	Like(ctx context.Context, commentVote *CommentVote) (err error)
	Dislike(ctx context.Context, commentVote *CommentVote) (err error)
}

// CommentVoteRepository ...
type CommentVoteRepository interface {
	Like(ctx context.Context, commentVote *CommentVote) (err error)
	Dislike(ctx context.Context, commentVote *CommentVote) (err error)
}

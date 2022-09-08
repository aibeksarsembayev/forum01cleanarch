package models

import (
	"context"
	"time"
)

// PostVote ...
type PostVote struct {
	PostVoteID    int       `json:"post_vote_id"`
	PostVoteValue bool      `json:"post_vote_value"`
	UserID        int       `json:"user_id"`
	PostID        int       `json:"post_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// PostVoteUsecase respresents postvotevalues usecases
type PostVoteUsecase interface {
	Like(ctx context.Context, postVote *PostVote) (err error)
	Dislike(ctx context.Context, postVote *PostVote) (err error)
	Delete(ctx context.Context, id int) (err error)
}

// PostVoteUsecase respresents postvotevalue's repository contact
type PostVoteRepository interface {
	Like(ctx context.Context, postVote *PostVote) (err error)
	Dislike(ctx context.Context, postVote *PostVote) (err error)
	Delete(ctx context.Context, id int) (err error) // need to check !!!
}

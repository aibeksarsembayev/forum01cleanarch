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

type PostVoteInputRequestDTO struct {
	PostVoteValue string `json:"post_vote_value"`
	PostID        string `json:"post_id"`
}

type PostVoteCreateRequestDTO struct {
	PostVoteValue bool `json:"post_vote_value"`
	UserID        int  `json:"user_id"`
	PostID        int  `json:"post_id"`
}

type PostVoteResponceDTO struct {
	PostVoteValue bool `json:"post_vote_value"`
	UserID        int  `json:"user_id"`
	PostID        int  `json:"post_id"`
}

// PostVoteUsecase respresents postvotevalues usecases
type PostVoteUsecase interface {
	Create(ctx context.Context, postVote *PostVote) (id int, err error)
	Update(ctx context.Context, postVote *PostVote) (err error)
	GetByPostUser(ctx context.Context, postID int, userID int) (*PostVote, error)
	GetCountByPost(ctx context.Context, postID int, value bool) (int, error)
	Delete(ctx context.Context, id int) (err error)
}

// PostVoteUsecase respresents postvotevalue's repository contact
type PostVoteRepository interface {
	Create(ctx context.Context, postVote *PostVote) (id int, err error)
	Update(ctx context.Context, postVote *PostVote) (err error)
	GetByPostUser(ctx context.Context, postID int, userId int) (*PostVote, error)
	GetCountByPost(ctx context.Context, postID int, value bool) (int, error)
	Delete(ctx context.Context, id int) (err error)
}

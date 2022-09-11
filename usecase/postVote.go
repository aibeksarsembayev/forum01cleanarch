package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type postVoteUsecase struct {
	PostVoteRepository models.PostVoteRepository
	// contextTimeout time.Duration
}

// // NewPostVoteUsecase will create new an post Usecase object representation of models.PostUsecase interface will create an object that represents the user. Repository interface
func NewPostVoteUsecase(pv models.PostVoteRepository, timeout time.Duration) models.PostVoteUsecase {
	return &postVoteUsecase{
		PostVoteRepository: pv,
	}
}

// Create post vote ...
func (pv *postVoteUsecase) Create(ctx context.Context, postVote *models.PostVoteCreateRequestDTO) (int, error) {
	id, err := pv.PostVoteRepository.Create(ctx, postVote)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update post vote ...
func (pv *postVoteUsecase) Update(ctx context.Context, postVote *models.PostVote) error {
	err := pv.PostVoteRepository.Update(ctx, postVote)
	if err != nil {
		return err
	}
	return nil
}

// GetByPostUser post vote ...
func (pv *postVoteUsecase) GetByPostUser(ctx context.Context, postID int, userID int) (*models.PostVote, error) {
	postVote, err := pv.PostVoteRepository.GetByPostUser(ctx, postID, userID)
	if err != nil {
		return nil, err
	}
	return postVote, nil
}

// GetCountByPost post vote ...
func (pv *postVoteUsecase) GetCountByPost(ctx context.Context, postID int, value bool) (int, error) {
	voteCount, err := pv.PostVoteRepository.GetCountByPost(ctx, postID, value)
	if err != nil {
		return 0, err
	}
	return voteCount, nil
}

// Delete post vote ...
func (pv *postVoteUsecase) Delete(ctx context.Context, id int) error {
	err := pv.PostVoteRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

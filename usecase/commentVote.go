package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type commentVoteUsecase struct {
	CommentVoteRepository models.CommentVoteRepository
	// contextTimeout time.Duration
}

// // NewComemntVoteUsecase will create new an comment vote Usecase object representation of models.CommentUsecase interface will create an object that represents the comment vote. Repository interface
func NewCommentVoteUsecase(cv models.CommentVoteRepository, timeout time.Duration) models.CommentVoteUsecase {
	return &commentVoteUsecase{
		CommentVoteRepository: cv,
	}
}

// Create comment vote ...
func (cv *commentVoteUsecase) Create(ctx context.Context, commentVote *models.CommentVote) (int, error) {
	id, err := cv.CommentVoteRepository.Create(ctx, commentVote)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update comment vote ...
func (cv *commentVoteUsecase) Update(ctx context.Context, commentVote *models.CommentVote) (err error) {
	err = cv.CommentVoteRepository.Update(ctx, commentVote)
	if err != nil {
		return err
	}
	return nil
}

// Get comment vote by comment id and post id, user id ...
func (cv *commentVoteUsecase) Get(ctx context.Context, commentVoteResponse *models.CommentVoteResponseDTO) (*models.CommentVote, error) {
	commentVote, err := cv.CommentVoteRepository.Get(ctx, commentVoteResponse)
	if err != nil {
		return nil, err
	}
	return commentVote, nil
}

// Get comment vote count by comment id and vote value
func (cv *commentVoteUsecase) GetCount(ctx context.Context, commentVoteResponse *models.CommentVoteCountResponseDTO) (int, error) {
	id, err := cv.CommentVoteRepository.GetCount(ctx, commentVoteResponse)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Delete comment ...
func (cv *commentVoteUsecase) Delete(ctx context.Context, id int) (err error) {
	err = cv.CommentVoteRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

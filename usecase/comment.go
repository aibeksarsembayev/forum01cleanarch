package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type commentUsecase struct {
	CommentRepository models.CommentRepository
}

// NewCommentUsecase will create new an comment Usecase object representation of models.CommentUsecase interface
func NewCommentUsecase(c models.CommentRepository, timeout time.Duration) models.CommentUsecase {
	return &commentUsecase{
		CommentRepository: c,
	}
}

// Create comment ...
func (c *commentUsecase) Create(ctx context.Context, comment *models.Comment) (id int, err error) {
	id, err = c.CommentRepository.Create(ctx, comment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update comment
func (c *commentUsecase) Update(ctx context.Context, comment *models.Comment) (err error) {
	err = c.CommentRepository.Update(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

// GetbyID comment ...
func (c *commentUsecase) GetByID(ctx context.Context, id int) (comment *models.CommentRequestDTO, err error) {
	comment, err = c.CommentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, err
}

// GetByUserID comment ...
func (c *commentUsecase) GetByUserID(ctx context.Context, user_id int) (comments *[]models.CommentRequestDTO, err error) {
	comments, err = c.CommentRepository.GetByUserID(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetByPostID comment ...
func (c *commentUsecase) GetByPostID(ctx context.Context, post_id int) (comments *[]models.CommentRequestDTO, err error) {
	comments, err = c.CommentRepository.GetByPostID(ctx, post_id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Delete comment ...
func (c *commentUsecase) Delete(ctx context.Context, id int) (err error) {
	err = c.CommentRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

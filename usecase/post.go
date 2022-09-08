package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type postUsecase struct {
	PostRepository models.PostRepository
	// contextTimeout time.Duration
}

// NewPostUsecase will create new an post Usecase object representation of models.PostUsecase interface
func NewPostUsecase(p models.PostRepository, timeout time.Duration) models.PostUsecase {
	return &postUsecase{
		PostRepository: p,
	}
}

// Create post ...
func (p *postUsecase) Create(ctx context.Context, post *models.Post) (id int, err error) {
	// ctxt, cancel := context.WithTimeout(ctx, p.contextTimeout)
	// defer cancel()
	id, err = p.PostRepository.Create(ctx, post)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update post ...
func (p *postUsecase) Update(ctx context.Context, post *models.Post) (err error) {
	err = p.PostRepository.Update(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

// GetAll posts ...
func (p *postUsecase) GetAll(ctx context.Context) (*[]models.PostRequestDTO, error) {
	posts, err := p.PostRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// GetbyID the post ...
func (p *postUsecase) GetByID(ctx context.Context, id int) (post *models.PostRequestDTO, err error) {
	post, err = p.PostRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// Delete post ...
func (p *postUsecase) Delete(ctx context.Context, id int) (err error) {
	err = p.PostRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type categoryUsecase struct {
	CategoryRepository models.CategoryRepository
	// contextTimeout time.Duration
}

// NewCategoryUsecase will create new an category Usecase object representation of models.CategoryUsecase interface
func NewCategoryUsecase(c models.CategoryRepository, timeout time.Duration) models.CategoryUsecase {
	return &categoryUsecase{
		CategoryRepository: c,
	}
}

// Create category ...
func (c *categoryUsecase) Create(ctx context.Context, category *models.Category) (id int, err error) {
	return 0, nil
}

// GetAll categories usecase ...
func (c *categoryUsecase) GetAll(ctx context.Context) (*[]models.CategoryRequestDTO, error) {
	categories, err := c.CategoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetbyID category usecase ...
func (c *categoryUsecase) GetByID(ctx context.Context, id int) (category *models.CategoryRequestDTO, err error) {
	return nil, nil
}

// Update category ...
func (c *categoryUsecase) Update(ctx context.Context, category *models.Category) (err error) {
	return nil
}

// Delete category ...
func (c *categoryUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}

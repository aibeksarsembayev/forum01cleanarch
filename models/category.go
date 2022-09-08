package models

import "context"

// Category ...
type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// CategoryRequestDTO ...
type CategoryRequestDTO struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// CategoryResponseDTO ...
type CategoryResponseDTO struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// CategoryUsecase represents category's usecases
type CategoryUsecase interface {
	Create(ctx context.Context, category *Category) (id int, err error)
	GetAll(ctx context.Context) (*[]CategoryRequestDTO, error)
	GetByID(ctx context.Context, id int) (category *CategoryRequestDTO, err error)
	Update(ctx context.Context, category *Category) (err error)
	Delete(ctx context.Context, id int) (err error)
}

// CategoryRepository represents category's repository contract
type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (id int, err error)
	GetAll(ctx context.Context) (*[]CategoryRequestDTO, error)
	GetByID(ctx context.Context, id int) (category *CategoryRequestDTO, err error)
	Update(ctx context.Context, category *Category) (err error)
	Delete(ctx context.Context, id int) (err error)
}

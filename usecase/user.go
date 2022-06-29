package usecase

import (
	"context"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type userUsecase struct {
	UserRepository models.UserRepository
}

// NewUserUsecase will create new user usecase object representation of models.UserUsecase
func NewUserUsecase(u models.UserRepository, timeout time.Duration) models.UserUsecase {
	return &userUsecase{
		UserRepository: u,
	}
}

// Create user ..
func (u *userUsecase) Create(ctx context.Context, user *models.User) (id int, err error) {
	id, err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userUsecase) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	return nil
}

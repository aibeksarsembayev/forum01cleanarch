package models

import (
	"context"
	"time"
)

// User ...
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRequestDTO ...
type UserRequestDTO struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserUsecase represents user's usecase
type UserUsecase interface {
	Create(ctx context.Context, user *User) (int, error)
	Update(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id int) error
}

// UserRepository represents user's repository contract
type UserRepository interface {
	Create(ctx context.Context, user *User) (int, error)
	Update(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id int) error
}

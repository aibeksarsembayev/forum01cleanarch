package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqliteUserRepository struct {
	Conn *sql.DB
}

// NewSqliteUserRepository will create an object that represents the user. Repository interface
func NewSqliteUserRepository(Conn *sql.DB) models.UserRepository {
	return &sqliteUserRepository{
		Conn: Conn,
	}
}

// Create user ..
func (u *sqliteUserRepository) Create(ctx context.Context, user *models.User) (id int, err error) {
	stmt, _ := u.Conn.Prepare("INSERT INTO users (username, password, email, created, updated) VALUES (?, ?, ?, ?, ?)")
	// Time convert
	dateFormat := "2006-01-02T15:04:05Z07:00"
	timeCreated := user.CreatedAt.Format(dateFormat)
	timeUpdated := user.UpdatedAt.Format(dateFormat)

	result, err := stmt.Exec(user.Username, user.Password, user.Email, timeCreated, timeUpdated)
	if err != nil {
		return 0, err
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		return 0, err //nil?
	}

	return int(user_id), nil

}

func (u *sqliteUserRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

// GetByID user ...
func (u *sqliteUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	stmt, _ := u.Conn.Prepare("SELECT * FROM users WHERE user_id = ?")

	row := stmt.QueryRow(id)

	user := &models.User{}

	dateFormat := "2006-01-02T15:04:05Z07:00"
	var timeCreated, timeUpdated string

	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &timeCreated, &timeUpdated)
	user.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	user.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return user, nil
}

// GetByEmail user ...
func (u *sqliteUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	stmt, _ := u.Conn.Prepare("SELECT * FROM users WHERE email = ?")

	row := stmt.QueryRow(email)

	user := &models.User{}

	dateFormat := "2006-01-02T15:04:05Z07:00"
	var timeCreated, timeUpdated string

	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &timeCreated, &timeUpdated)
	user.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	user.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (u *sqliteUserRepository) Delete(ctx context.Context, id int) error {
	return nil
}

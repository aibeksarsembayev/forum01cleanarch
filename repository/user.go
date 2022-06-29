package repository

import (
	"context"
	"database/sql"
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
	result, err := stmt.Exec(user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(user_id), nil

}

func (u *sqliteUserRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

// GetByID user ...
func (u *sqliteUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

// GetByEmail user ...
func (u *sqliteUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	stmt, _ := u.Conn.Prepare("SELECT * FROM users WHERE email = ?")

	row := stmt.QueryRow(email)

	user := &models.User{}

	dateFormat := "2016-10-06 01:50:00 -0700 MST"
	var timeCreated, timeUpdated string

	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &timeCreated, &timeUpdated)
	user.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	user.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *sqliteUserRepository) Delete(ctx context.Context, id int) error {
	return nil
}

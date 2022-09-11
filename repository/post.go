package repository

import (
	"context"
	"database/sql"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqlitePostRepository struct {
	Conn *sql.DB
}

// NewSqliteUserRepository will create an object that represents the user. Repository interface
func NewSqlitePostRepository(Conn *sql.DB) models.PostRepository {
	return &sqlitePostRepository{
		Conn: Conn,
	}
}

// Create post ...
func (p *sqlitePostRepository) Create(ctx context.Context, post *models.Post) (id int, err error) {
	stmt, _ := p.Conn.Prepare("INSERT INTO posts (title, content, user_id, category_name, created, updated) VALUES(?, ?, ?, ?, ?, ?)")
	result, err := stmt.Exec(post.Title, post.Content, post.UserID, post.CategoryName, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return 0, err
	}

	post_id, err := result.LastInsertId()
	if err != nil {
		return 0, err //nil?
	}

	return int(post_id), nil
}

// Update post ...
func (p *sqlitePostRepository) Update(ctx context.Context, post *models.Post) (err error) {
	return nil
}

// GetAll posts ...
func (p *sqlitePostRepository) GetAll(ctx context.Context) (*[]models.PostRequestDTO, error) {
	rows, err := p.Conn.Query("SELECT post_id, title, content, user_id, category_name, created, updated FROM posts ORDER BY post_id DESC")
	if err != nil {
		return nil, err
	}

	var posts []models.PostRequestDTO

	for rows.Next() {
		post := &models.PostRequestDTO{}

		dateFormat := "2016-10-06 01:50:00 -0700 MST"
		var timeCreated, timeUpdated string

		err = rows.Scan(&post.PostID, &post.Title, &post.Content, &post.UserID, &post.CategoryName, &timeCreated, &timeUpdated)
		post.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
		post.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

		if err != nil {
			return nil, err
		}

		// get username by ID
		u := &sqliteUserRepository{Conn: p.Conn}
		user, err := u.GetByID(ctx, post.UserID)
		if err != nil {
			return nil, err
		}
		post.Username = user.Username

		// get post votes by postid and userid
		pv := &sqlitePostVoteRepository{Conn: p.Conn}
		post.VoteLike, err = pv.GetCountByPost(ctx, post.PostID, true)
		if err != nil {
			return nil, err
		}
		post.VoteDislike, err = pv.GetCountByPost(ctx, post.PostID, false)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &posts, nil
}

// GetbyID the post ...
func (p *sqlitePostRepository) GetByID(ctx context.Context, id int) (*models.PostRequestDTO, error) {
	stmt, _ := p.Conn.Prepare("SELECT * FROM posts WHERE post_id = ?")
	row := stmt.QueryRow(id)

	post := &models.PostRequestDTO{}

	dateFormat := "2016-10-06 01:50:00 -0700 MST"
	var timeCreated, timeUpdated string

	err := row.Scan(&post.PostID, &post.Title, &post.Content, &post.UserID, &post.CategoryName, &timeCreated, &timeUpdated)
	post.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	post.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

	if err != nil {
		return nil, err
	}

	// get username
	u := &sqliteUserRepository{Conn: p.Conn}
	user, err := u.GetByID(ctx, post.UserID)
	if err != nil {
		return nil, err
	}
	post.Username = user.Username

	if err = row.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

// Delete post ...
func (p *sqlitePostRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}

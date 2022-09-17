package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqliteCommentRepository struct {
	Conn *sql.DB
}

// NewSqliteCommentRepository will create an object that represents the comment
func NewSqliteCommentRepository(Conn *sql.DB) models.CommentRepository {
	return &sqliteCommentRepository{
		Conn: Conn,
	}
}

// Create comment for post ...
func (c *sqliteCommentRepository) Create(ctx context.Context, comment *models.Comment) (int, error) {
	stmt, _ := c.Conn.Prepare("INSERT INTO comments (comment_body, post_id, user_id, created, updated) VALUES(?, ?, ?, ?, ?)")

	// Time converting to string
	dateFormat := "2006-01-02T15:04:05Z07:00"
	timeCreated := comment.CreatedAt.Format(dateFormat)
	timeUpdated := comment.UpdatedAt.Format(dateFormat)
	result, err := stmt.Exec(comment.Content, comment.PostID, comment.UserID, timeCreated, timeUpdated)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Update comment for post ...
func (c *sqliteCommentRepository) Update(ctx context.Context, comment *models.Comment) (err error) {
	return nil
}

// Get By ID comment ...
func (c *sqliteCommentRepository) GetByID(ctx context.Context, id int) (*models.CommentRequestDTO, error) {
	stmt, _ := c.Conn.Prepare("SELECT comment_id, user_id, post_id, comment_body, created, updated  FROM comments WHERE comment_id = ?")

	row := stmt.QueryRow(id)

	comment := &models.CommentRequestDTO{}

	dateFormat := "2006-01-02T15:04:05Z07:00"
	var timeCreated, timeUpdated string

	err := row.Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.Content, &timeCreated, &timeUpdated)
	comment.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	comment.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	// get comment votes by ...
	cv := &sqliteCommentVoteRepository{Conn: c.Conn}
	comment.CommentVoteLike, err = cv.GetCount(ctx, &models.CommentVoteCountResponseDTO{
		CommentID:        comment.CommentID,
		CommentVoteValue: true,
	})
	if err != nil {
		return nil, err
	}
	comment.CommentVoteDislike, err = cv.GetCount(ctx, &models.CommentVoteCountResponseDTO{
		CommentID:        comment.CommentID,
		CommentVoteValue: false,
	})
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// GetByUserID comment for post ...
func (c *sqliteCommentRepository) GetByUserID(ctx context.Context, user_id int) (comments *[]models.CommentRequestDTO, err error) {
	return nil, nil
}

// GetByPostID comment for post ...
func (c *sqliteCommentRepository) GetByPostID(ctx context.Context, post_id int) (*[]models.CommentRequestDTO, error) {
	stmt, _ := c.Conn.Prepare("SELECT comment_id, user_id, comment_body, created, updated  FROM comments WHERE post_id = ? ORDER BY comment_id DESC")
	rows, err := stmt.Query(post_id)
	if err != nil {
		return nil, err
	}
	var comments []models.CommentRequestDTO

	// iterate each comment
	for rows.Next() {
		comment := &models.CommentRequestDTO{}

		dateFormat := "2006-01-02T15:04:05Z07:00"
		var timeCreated, timeUpdated string

		err = rows.Scan(&comment.CommentID, &comment.UserID, &comment.Content, &timeCreated, &timeUpdated)
		comment.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
		comment.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)

		if err != nil {
			return nil, err
		}

		comment.PostID = post_id

		// get username by ID
		u := &sqliteUserRepository{Conn: c.Conn}
		user, err := u.GetByID(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}
		comment.Username = user.Username

		// get post comment votes
		cv := &sqliteCommentVoteRepository{Conn: c.Conn}
		comment.CommentVoteLike, err = cv.GetCount(ctx, &models.CommentVoteCountResponseDTO{
			CommentID:        comment.CommentID,
			CommentVoteValue: true,
		})
		if err != nil {
			return nil, err
		}
		comment.CommentVoteDislike, err = cv.GetCount(ctx, &models.CommentVoteCountResponseDTO{
			CommentID:        comment.CommentID,
			CommentVoteValue: false,
		})
		if err != nil {
			return nil, err
		}

		comments = append(comments, *comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &comments, nil
}

// Delete comment for post ...
func (c *sqliteCommentRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}

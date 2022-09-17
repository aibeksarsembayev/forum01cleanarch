package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqliteCommentVoteRepository struct {
	Conn *sql.DB
}

// NewSqliteCommentVoteRepository will create an object that represents the comment vote
func NewSqliteCommentVoteRepository(Conn *sql.DB) models.CommentVoteRepository {
	return &sqliteCommentVoteRepository{
		Conn: Conn,
	}
}

// Create comment vote ...
func (cv *sqliteCommentVoteRepository) Create(ctx context.Context, commentVote *models.CommentVote) (int, error) {
	v, err := cv.Get(ctx, &models.CommentVoteResponseDTO{
		CommentID: commentVote.CommentID,
		UserID:    commentVote.UserID,
	})
	// check if empty row or other issues
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // no row, then create vote ...
			stmt, _ := cv.Conn.Prepare("INSERT INTO vote_comment (user_id, comment_id, value, created, updated) VALUES (?, ?, ? ,?, ?)")
			// Time converting to string
			dateFormat := "2006-01-02T15:04:05Z07:00"
			timeCreated := commentVote.CreatedAt.Format(dateFormat)
			timeUpdated := commentVote.UpdatedAt.Format(dateFormat)

			result, err := stmt.Exec(commentVote.UserID, commentVote.CommentID, commentVote.CommentVoteValue, timeCreated, timeUpdated)
			if err != nil {
				return 0, err
			}
			commentVoteID, err := result.LastInsertId()
			if err != nil {
				return 0, err
			}
			return int(commentVoteID), nil
		} else {
			return 0, err
		}
	}
	// if vote already exists, then remove
	if v.CommentVoteValue == commentVote.CommentVoteValue {
		err = cv.Delete(ctx, v.CommentVoteID)
		return 0, err
	}
	// create new comment model vote with updated vote and updated time ...
	vote := &models.CommentVote{
		CommentVoteID:    v.CommentVoteID,
		CommentVoteValue: commentVote.CommentVoteValue,
		UserID:           v.UserID,
		CommentID:        v.CommentID,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        time.Now(),
	}

	// Update vote with new value
	err = cv.Update(ctx, vote)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		} else {
			return 0, err
		}
	}
	return v.CommentVoteID, nil
}

// Update comment vote ...
func (cv *sqliteCommentVoteRepository) Update(ctx context.Context, commentVote *models.CommentVote) (err error) {
	stmt, _ := cv.Conn.Prepare("UPDATE vote_comment SET value = ?, updated = ? WHERE id = ?")
	// Time converting to string
	dateFormat := "2006-01-02T15:04:05Z07:00"
	timeUpdated := commentVote.UpdatedAt.Format(dateFormat)
	_, err = stmt.Exec(commentVote.CommentVoteValue, timeUpdated, commentVote.CommentVoteID)
	if err != nil {
		return err
	}
	return nil
}

// Get comment vote by comment id and user id ...
func (cv *sqliteCommentVoteRepository) Get(ctx context.Context, commentVoteResponse *models.CommentVoteResponseDTO) (*models.CommentVote, error) {
	stmt, _ := cv.Conn.Prepare("SELECT id, user_id, comment_id, value, created, updated FROM vote_comment WHERE comment_id = ? AND user_id = ?")
	row := stmt.QueryRow(commentVoteResponse.CommentID, commentVoteResponse.UserID)
	v := &models.CommentVote{}
	// time converting
	dateFormat := "2006-01-02T15:04:05Z07:00"
	var timeCreated, timeUpdated string

	err := row.Scan(&v.CommentVoteID, &v.UserID, &v.CommentID, &v.CommentVoteValue, &timeCreated, &timeUpdated)
	v.CreatedAt, _ = time.Parse(dateFormat, timeCreated)
	v.UpdatedAt, _ = time.Parse(dateFormat, timeUpdated)
	return v, err
}

// Get comment vote count by comment id and vote value
func (cv *sqliteCommentVoteRepository) GetCount(ctx context.Context, commentVoteResponse *models.CommentVoteCountResponseDTO) (int, error) {
	stmt, _ := cv.Conn.Prepare("SELECT COUNT(value) FROM vote_comment WHERE comment_id = ? AND value = ?")
	row := stmt.QueryRow(commentVoteResponse.CommentID, commentVoteResponse.CommentVoteValue)

	var voteCount int
	err := row.Scan(&voteCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return voteCount, nil
}

// Delete comment vote...
func (cv *sqliteCommentVoteRepository) Delete(ctx context.Context, id int) (err error) {
	stmt, _ := cv.Conn.Prepare("DELETE FROM vote_comment WHERE id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

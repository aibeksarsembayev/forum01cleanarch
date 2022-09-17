package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqlitePostVoteRepository struct {
	Conn *sql.DB
}

// NewSqlitePostvoteRepository will create an object that represents the user. Repository interface
func NewSqlitePostVoteRepository(Conn *sql.DB) models.PostVoteRepository {
	return &sqlitePostVoteRepository{
		Conn: Conn,
	}
}

// Create post vote ...
func (pv *sqlitePostVoteRepository) Create(ctx context.Context, postVote *models.PostVote) (int, error) {
	v, err := pv.GetByPostUser(ctx, postVote.PostID, postVote.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// create row in the table
			// Time converting to string
			dateFormat := "2006-01-02T15:04:05Z07:00"
			CreatedAt := postVote.CreatedAt.Format(dateFormat)
			UpdatedAt := postVote.UpdatedAt.Format(dateFormat)

			stmt, _ := pv.Conn.Prepare("INSERT INTO votes (post_id, user_id, value, created, updated) VALUES (?, ?, ?, ?, ?)")
			result, err := stmt.Exec(postVote.PostID, postVote.UserID, postVote.PostVoteValue, CreatedAt, UpdatedAt)
			if err != nil {
				return 0, err
			}
			vote_id, err := result.LastInsertId()
			if err != nil {
				return 0, err
			}
			return int(vote_id), nil
		} else {
			return 0, err
		}
	}

	// if already liked, to remove
	if v.PostVoteValue == postVote.PostVoteValue {
		err = pv.Delete(ctx, v.PostVoteID)
		return 0, err
	}

	vote := &models.PostVote{
		PostVoteID:    v.PostVoteID,
		PostID:        v.PostID,
		UserID:        v.UserID,
		PostVoteValue: postVote.PostVoteValue,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     time.Now(),
	}

	// update vote if another value
	err = pv.Update(ctx, vote)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		} else {
			return 0, err
		}
	}

	return v.PostVoteID, nil
}

// Update post vote ...
func (pv *sqlitePostVoteRepository) Update(ctx context.Context, postVote *models.PostVote) error {
	stmt, _ := pv.Conn.Prepare("UPDATE votes SET value = ?, updated = ? WHERE id = ?")
	// Time converting to string
	dateFormat := "2006-01-02T15:04:05Z07:00"
	timeUpdated := postVote.UpdatedAt.Format(dateFormat)

	_, err := stmt.Exec(postVote.PostVoteValue, timeUpdated, postVote.PostVoteID)
	if err != nil {
		return err
	}
	return nil
}

// GetByPostUser post vote ...
func (pv *sqlitePostVoteRepository) GetByPostUser(ctx context.Context, postID int, userID int) (*models.PostVote, error) {
	stmt, _ := pv.Conn.Prepare("SELECT id, value FROM votes WHERE post_id = ? AND user_id = ?")
	row := stmt.QueryRow(postID, userID)
	v := &models.PostVote{}
	err := row.Scan(&v.PostVoteID, &v.PostVoteValue)
	return v, err
}

// GetCountByPost post vote ...
func (pv *sqlitePostVoteRepository) GetCountByPost(ctx context.Context, postID int, value bool) (int, error) {
	stmt, _ := pv.Conn.Prepare("SELECT COUNT(value) FROM votes WHERE post_id = ? AND value = ?")
	row := stmt.QueryRow(postID, value)

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

// Delete post vote ...
func (pv *sqlitePostVoteRepository) Delete(ctx context.Context, id int) error {
	stmt, _ := pv.Conn.Prepare("DELETE FROM votes WHERE id = ?")
	_, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

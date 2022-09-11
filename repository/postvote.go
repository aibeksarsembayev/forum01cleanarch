package repository

import (
	"context"
	"database/sql"
	"errors"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqlitePostVoteRepository struct {
	Conn *sql.DB
}

// NewSqliteUPostvoteRepository will create an object that represents the user. Repository interface
func NewSqlitePostVoteRepository(Conn *sql.DB) models.PostVoteRepository {
	return &sqlitePostVoteRepository{
		Conn: Conn,
	}
}

// Create post vote ...
func (pv *sqlitePostVoteRepository) Create(ctx context.Context, postVote *models.PostVoteCreateRequestDTO) (int, error) {
	v, err := pv.GetByPostUser(ctx, postVote.PostID, postVote.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// create row in the table
			stmt, _ := pv.Conn.Prepare("INSERT INTO votes (post_id, user_id, value) VALUES (?, ?, ?)")
			result, err := stmt.Exec(postVote.PostID, postVote.UserID, postVote.PostVoteValue)
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
	// update vote if another value
	err = pv.Update(ctx, &models.PostVote{
		PostVoteID:    v.PostVoteID,
		PostID:        v.PostID,
		UserID:        v.UserID,
		PostVoteValue: v.PostVoteValue,
	})

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
	return nil
}

// GetByPostUser post vote ...
func (pv *sqlitePostVoteRepository) GetByPostUser(ctx context.Context, postID int, userID int) (*models.PostVote, error) {
	return nil, nil
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
	return nil
}

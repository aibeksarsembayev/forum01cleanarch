package httpdelivery

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// commentVote handler for like and dislikes of comments ...
func (ch *Handler) commentVote(w http.ResponseWriter, r *http.Request) {
	// get user information voting for comments ...
	userEmail, isSession := GetSession(r)
	if !isSession {
		JSON(w, http.StatusUnauthorized, ResponseError{Message: "not logged"})
		return
	}
	// get user info ...
	user, err := ch.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	if r.Method != "POST" {
		JSON(w, http.StatusMethodNotAllowed, ResponseError{Message: "Wrong method"})
		return
	} else {
		// JSON decoder ..
		decoder := json.NewDecoder(r.Body)
		var commentVoteInput models.CommentVoteInputRequestDTO
		err := decoder.Decode(&commentVoteInput)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		// comment vote ...
		commentVote := models.CommentVote{
			UserID:    user.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	
		// define like dislike and convert ot boolean
		if commentVoteInput.CommentVoteValue == "like" {
			commentVote.CommentVoteValue = true
		} else {
			commentVote.CommentVoteValue = false //check for 3rd option with error response?
		}
		// comment id converting to int ...
		commentVote.CommentID, err = strconv.Atoi(commentVoteInput.CommentID)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		// create comment vote with checker for liked/disliked comment votes ...
		_, err = ch.CommentVoteUsecase.Create(context.Background(), &commentVote)
		if err != nil {
			ch.ErrorLog.Println(err)
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		// get comment by comment id
		comment, err := ch.CommentUsecase.GetByID(context.Background(), commentVote.CommentID)
		if err != nil {
			ch.ErrorLog.Println(err)
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		JSON(w, http.StatusCreated, comment)
	}
}

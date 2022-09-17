package httpdelivery

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// postVote handler for like/dislikes ...
func (ph *Handler) postVote(w http.ResponseWriter, r *http.Request) {

	userEmail, isSession := GetSession(r)
	if !isSession {

		JSON(w, http.StatusUnauthorized, ResponseError{Message: "not logged"})
		return
	}
	// get user model
	user, err := ph.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	if r.Method != "POST" {
		JSON(w, http.StatusMethodNotAllowed, ResponseError{Message: "Wrong method"}) // custom error?
		return
	} else {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var voteInput models.PostVoteInputRequestDTO
		err := decoder.Decode(&voteInput)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		//
		vote := models.PostVote{
			UserID:    user.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		
		if voteInput.PostVoteValue == "like" {
			vote.PostVoteValue = true
		} else {
			vote.PostVoteValue = false
		}
		//
		vote.PostID, err = strconv.Atoi(voteInput.PostID)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		// create vote with checker for liked/disliked votes ...
		_, err = ph.PostVoteUsecase.Create(context.Background(), &vote)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		post, err := ph.PostUsecase.GetByID(context.Background(), vote.PostID)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		// fmt.Println(post)
		JSON(w, http.StatusCreated, post)
	}

}

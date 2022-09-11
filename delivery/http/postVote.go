package httpdelivery

import (
	"context"
	"encoding/json"
	"net/http"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// postVote handler for like/dislikes ...
func (ph *Handler) postVote(w http.ResponseWriter, r *http.Request) {
	// get post id
	// id, err := strconv.Atoi(path.Base(r.URL.Path))
	// if err != nil || id < 1 {
	// 	ph.renderHTML(w, r, http.StatusNotFound, "404.page.html", ResponseError{Message: err.Error()})
	// 	return
	// }
	// get user info and session ...
	userEmail, isSession := GetSession(r)
	if !isSession {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
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
		vote := models.PostVoteCreateRequestDTO{
			PostVoteValue: voteInput.PostVoteValue,
			PostID:        voteInput.PostID,
			UserID:        user.UserID,
		}
		// create vote with checker for liked/disliked votes ...
		_, err = ph.PostVoteUsecase.Create(context.Background(), &vote)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
		}

		JSON(w, http.StatusCreated, models.Post{PostID: vote.PostID})
	}

}

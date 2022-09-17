package httpdelivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Create comment for post ...
func (ch *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	userEmail, isSession := GetSession(r)
	if !isSession {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}

	// get autorized user, leaving comment ...
	user, err := ch.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	// Check for method of request ...
	if r.Method != "POST" {
		JSON(w, http.StatusMethodNotAllowed, ResponseError{Message: "Wrong method for posting comments. Please do not play with methods and Use POST method"})
		return
	} else {
		// JSON decoder of incoming data ...
		decoder := json.NewDecoder(r.Body)
		var commentInput models.CommentInputRequestDTO
		err := decoder.Decode(&commentInput)

		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			fmt.Println(err)
			return
		}
		// Empty entry checker
		if strings.TrimSpace(commentInput.Content) == "" {
			ch.ErrorLog.Println("issue with entry")
			JSON(w, http.StatusBadRequest, ResponseError{Message: "error: wrong entry"})
			return
		}
		// Transmit from input DTO to DB template model to store
		comment := models.Comment{
			Content: commentInput.Content,
			// PostID:    commentInput.PostID,
			UserID:    user.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		comment.PostID, err = strconv.Atoi(commentInput.PostID)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
		}

		// create comment ...
		_, err = ch.CommentUsecase.Create(context.Background(), &comment)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		// get post by id ...
		// post, err := ch.PostUsecase.GetByID(context.Background(), comment.PostID)
		// if err != nil {
		// 	JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
		// 	return
		// }
		// send comment to frontend ...
		JSON(w, http.StatusCreated, comment)
		// ch.renderHTML(w, r, http.StatusOK, "home.page.html", map[string]interface{}{
		// 	"post": *post,
		// 	// "user":    *user, //does it need to be added here?
		// 	"comment": comment,
		// 	"session": isSession,
		// })
	}

}

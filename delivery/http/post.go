package httpdelivery

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Home page handler for all posts
func (ph *Handler) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ph.renderHTML(w, r, http.StatusNotFound, "404.page.html", map[string]interface{}{})
		return
	}

	p, err := ph.PostUsecase.GetAll(context.Background())

	if err != nil {
		ph.ErrorLog.Println(err)
		ph.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}
	// get logged user info
	email, isSession := GetSession(r)
	// get categories ..
	categories, err := ph.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	// render page ...
	ph.renderHTML(w, r, http.StatusOK, "home.page.html", map[string]interface{}{
		"posts":      *p,
		"user":       models.User{Email: email},
		"session":    isSession,
		"categories": categories,
	})
}

// Create post handler ...
func (ph *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	userEmail, isSession := GetSession(r)

	if !isSession {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	user, err := ph.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	categories, err := ph.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	if r.Method == "GET" {
		// render page ...
		ph.renderHTML(w, r, http.StatusOK, "createpost.page.html", map[string]interface{}{
			"categories": *categories,
		})
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var postInput models.PostCreateRequestDTO
		err := decoder.Decode(&postInput)

		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		// Empty entry checker
		if (strings.TrimSpace(postInput.Title) == "") || (strings.TrimSpace(postInput.Content) == "") {
			ph.ErrorLog.Println("issue with entry")
			JSON(w, http.StatusBadRequest, ResponseError{Message: "error: wrong entry"})
			return
		}
		// Wrong category checker
		categoryChecker := 0
		for _, category := range *categories {
			if category.CategoryName == postInput.CategoryName {
				categoryChecker = 1
			}
		}
		if categoryChecker == 0 {
			ph.ErrorLog.Println("issue with category")
			JSON(w, http.StatusBadRequest, ResponseError{Message: "error: wrong entry"})
			return
		}

		post := models.Post{
			Title:        postInput.Title,
			Content:      postInput.Content,
			CategoryName: postInput.CategoryName,
			UserID:       user.UserID,
			Username:     user.Username,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}



		post.PostID, err = ph.PostUsecase.Create(context.Background(), &post)
		if err != nil {
			JSON(w, getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		JSON(w, http.StatusCreated, post)
	}
}

// Show post handler ...
func (ph *Handler) showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id < 1 {
		ph.renderHTML(w, r, http.StatusNotFound, "404.page.html", map[string]interface{}{})
		return
	}

	post, err := ph.PostUsecase.GetByID(context.Background(), id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			JSON(w, http.StatusNotFound, ResponseError{Message: err.Error()})
		} else {
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
		}
		// ph.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}

	// to implement get comments ...
	comments, err := ph.CommentUsecase.GetByPostID(context.Background(), id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) { // ?
			JSON(w, http.StatusNotFound, ResponseError{Message: err.Error()})
		} else {
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
		}
		return
	}
	// to implement comments number ...
	commentNumber := strconv.Itoa(len(*comments))
	// get session info ...
	email, isSession := GetSession(r)

	ph.renderHTML(w, r, 200, "post.page.html", map[string]interface{}{
		"post":          *post,
		"user":          models.User{Email: email},
		"session":       isSession,
		"comments":      comments,
		"commentNumber": commentNumber,
	})

}

// Show liked post handler ...
func (ph *Handler) showLikedPost(w http.ResponseWriter, r *http.Request) {
	emailInput := path.Base(r.URL.Path)
	userEmail, isSession := GetSession(r)
	if !isSession {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	if emailInput != userEmail {
		JSON(w, http.StatusBadRequest, ResponseError{Message: "please use valid user data"})
	}
	// get user info...
	user, err := ph.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	// get all categories...
	categories, err := ph.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	// get all posts by user liked
	posts, err := ph.PostUsecase.GetLikedByUserID(context.Background(), user.UserID)
	if err != nil {
		ph.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}
	// render page ...
	ph.renderHTML(w, r, http.StatusOK, "profileLiked.page.html", map[string]interface{}{
		"posts":      *posts,
		"user":       user,
		"session":    isSession,
		"categories": categories,
	})
}

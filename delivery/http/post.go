package httpdelivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Home page handler for all posts
func (ph *Handler) home(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "home page is here ...", http.StatusOK)

	if r.URL.Path != "/" {
		ph.renderHTML(w, r, http.StatusNotFound, "404.page.html", map[string]interface{}{})
		return
	}

	p, err := ph.PostUsecase.GetAll(context.Background())

	if err != nil {
		ph.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}

	username, err := GetSession(r)

	if err != nil {
		// JSON(w, 401, err)
		fmt.Println("not in session")
	}

	// render page ...
	ph.renderHTML(w, r, http.StatusOK, "home.page.html", map[string]interface{}{
		"posts": *p,
		"user":  models.User{Username: username},
	})
}

// Create post handler ...
func (ph *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	userEmail, err := GetSession(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	user, err := ph.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, 405, err)
		return
	}

	categories, err := ph.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, 405, err)
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
			JSON(w, 500, err)
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

		// Time converting to string
		dateFormat := "2016-10-06 01:50:00 -0700 MST"
		post.CreatedAt.Format(dateFormat)
		post.UpdatedAt.Format(dateFormat)

		post.PostID, err = ph.PostUsecase.Create(context.Background(), &post)
		if err != nil {
			JSON(w, 500, err)
			return
		}
		JSON(w, 200, post)
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
		// if errors.Is(err, models.ErrNoRecord) {
		// 	JSON(w, 404, err)
		// } else {
		// 	JSON(w, 500, err)
		// }
		ph.renderHTML(w, r, http.StatusNotFound, "500.page.html", map[string]interface{}{})
		return
	}

	// to implement count votes ...

	// to implement get comments ...

	// to implement comments number ...

	// get session info ...
	username, err := GetSession(r)
	if err != nil {
		fmt.Println(err, "showpost")
	}

	ph.renderHTML(w, r, 200, "post.page.html", map[string]interface{}{
		"post": *post,
		"user": models.User{Username: username},
	})

}

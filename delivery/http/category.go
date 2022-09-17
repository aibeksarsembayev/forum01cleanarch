package httpdelivery

import (
	"context"
	"net/http"
	"path"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Show posts by category name ..
func (h *Handler) showPostsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryInput := path.Base(r.URL.Path)
	// get all categories...
	categories, err := h.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	// check if category input exists in DB
	var categoryCheck bool
	for _, c := range *categories {
		if c.CategoryName == categoryInput {
			categoryCheck = true
			break
		}
		categoryCheck = false
	}
	// if not existed, then reply with error
	if !categoryCheck {
		JSON(w, http.StatusBadRequest, ResponseError{Message: "No such category. Please do not play with categories!"})
		return
	}
	posts, err := h.PostUsecase.GetByCategory(context.Background(), categoryInput)
	if err != nil {
		h.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}
	// get logged user info
	email, isSession := GetSession(r)

	// render page ...
	h.renderHTML(w, r, http.StatusOK, "home.page.html", map[string]interface{}{
		"posts":      *posts,
		"user":       models.User{Email: email},
		"session":    isSession,
		"categories": categories,
	})
}

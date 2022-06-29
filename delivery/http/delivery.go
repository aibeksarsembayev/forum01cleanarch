package httpdelivery

import (
	"net/http"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Handler
type Handler struct {
	UserUsecase models.UserUsecase
}

// NewHandler
func NewHandler(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// user authentification handlers ...
	mux.HandleFunc("/signin", handler.signin)
	mux.HandleFunc("/signup", handler.signup)
	mux.HandleFunc("/signout", handler.signout)

	return mux
}

// user endpoints ...

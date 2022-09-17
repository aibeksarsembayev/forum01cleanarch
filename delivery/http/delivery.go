package httpdelivery

import (
	"html/template"
	"log"
	"net/http"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// Handler
type Handler struct {
	UserUsecase        models.UserUsecase
	PostUsecase        models.PostUsecase
	CategoryUsecase    models.CategoryUsecase
	PostVoteUsecase    models.PostVoteUsecase
	CommentUsecase     models.CommentUsecase
	CommentVoteUsecase models.CommentVoteUsecase
	TemplateCache      map[string]*template.Template
	InfoLog            *log.Logger
	ErrorLog           *log.Logger
}

// NewHandler
func NewHandler(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// Home page ...
	mux.HandleFunc("/", handler.home)

	// Post handlers ...
	mux.HandleFunc("/post/create", handler.createPost)
	mux.HandleFunc("/post/", handler.showPost)

	// post vote handlers ...
	mux.HandleFunc("/post/vote/", handler.postVote)

	// post comment handlers ...
	mux.HandleFunc("/post/comment", handler.createComment)

	// post comment vote handlers ...
	mux.HandleFunc("/post/comment/vote/", handler.commentVote)

	// user authentification handlers ...
	mux.HandleFunc("/signin", handler.signin)
	mux.HandleFunc("/signup", handler.signup)
	mux.HandleFunc("/signout", handler.signout)

	// user profile handlers ...
	mux.HandleFunc("/user/profile/", handler.showProfile)
	mux.HandleFunc("/user/liked/", handler.showLikedPost)

	// category handlers ...
	mux.HandleFunc("/category/", handler.showPostsByCategory)

	// static style file ...
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

package httpdelivery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Handler) signin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// http.Error(w, "signin page is OK", 200)

		_, isSession := GetSession(r)
		// sigin page rendering ...
		u.renderHTML(w, r, 200, "signin.page.html", map[string]interface{}{
			"session": isSession,
		})
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var user models.UserLoginRequestDTO
		err := decoder.Decode(&user)

		if err != nil {
			// http.Error(w, "signin: json decoder error", 500)
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
			return
		}

		us, err := u.UserUsecase.GetByEmail(context.Background(), user.Email)
		if err != nil {
			// need tohandle error based on status (500, 401, etc.)
			// http.Error(w, "signin: coudlnt find or unautorized", http.StatusUnauthorized)

			JSON(w, http.StatusUnauthorized, ResponseError{Message: err.Error()})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password)); err != nil {
			// need to handle with 401 error
			// http.Error(w, "signin: password is wrong", http.StatusUnauthorized)

			JSON(w, http.StatusUnauthorized, ResponseError{Message: err.Error()})
			return
		}

		// Create a new random session token
		UUID, _ := uuid.NewV4()
		sessionToken := UUID.String()

		// Set the token in the map, along with the user whom it represents
		Set(us.Email, sessionToken)

		// Finally, we set the client cookie for "session_token" as the session token we just generated
		// we also set an expiry time of 120 seconds, the same as the cache
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    url.QueryEscape(sessionToken),
			Expires:  time.Now().Add(120 * time.Second),
			HttpOnly: true, // Cookies provided only for HTTP(HTTPS) requests only
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (u *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		_, isSession := GetSession(r)
		// signup page rendering ...
		u.renderHTML(w, r, 200, "signup.page.html", map[string]interface{}{
			"session": isSession,
		})
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var userInput models.UserRegisterRequestDTO
		err := decoder.Decode(&userInput)

		if err != nil {
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
			return
		}

		// singnup form checker to be added here ...

		var user models.User
		user.Username = userInput.Username
		user.Email = userInput.Email
		user.Password = userInput.Password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
			return
		}
		user.Password = string(hashedPassword)

		user.UserID, err = u.UserUsecase.Create(context.Background(), &user)
		if err != nil {
			// handle error ...
			// http.Error(w, "signup page: couldnt create user", 500)
			JSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
			return
		}
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		// http.Error(w, "user has been created", http.StatusCreated)
		JSON(w, http.StatusCreated, user) // user or nil?
		return
	} else {
		// method not allowed
		// http.Error(w, "signup page: method not allowed", http.StatusMethodNotAllowed)
		JSON(w, http.StatusMethodNotAllowed, ResponseError{Message: "Wrong method"}) // custom error?
		return
	}
}

func (u *Handler) signout(w http.ResponseWriter, r *http.Request) {
	ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Show user profile ...
func (u *Handler) showProfile(w http.ResponseWriter, r *http.Request) {
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
	user, err := u.UserUsecase.GetByEmail(context.Background(), userEmail)
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	// get all categories...
	categories, err := u.CategoryUsecase.GetAll(context.Background())
	if err != nil {
		JSON(w, http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	// get all posts by user id
	posts, err := u.PostUsecase.GetByUserID(context.Background(), user.UserID)
	if err != nil {
		u.renderHTML(w, r, http.StatusInternalServerError, "500.page.html", map[string]interface{}{})
		return
	}
	// render page ...
	u.renderHTML(w, r, http.StatusOK, "profile.page.html", map[string]interface{}{
		"posts":      *posts,
		"user":       user,
		"session":    isSession,
		"categories": categories,
	})
}

// // Context realization example
// ctx := r.Context()
// fmt.Println("server: hello handler started")
// defer fmt.Println("server: hello handler ended")

// select {
// case <-time.After(10 * time.Second):
// 	fmt.Fprintf(w, "hello\n")
// case <-ctx.Done():

// 	err := ctx.Err()
// 	fmt.Println("server:", err)
// 	internalError := http.StatusInternalServerError
// 	http.Error(w, err.Error(), internalError)
// }

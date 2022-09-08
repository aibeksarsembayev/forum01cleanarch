package httpdelivery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"git.01.alem.school/quazar/forum-authentication/models"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Handler) signin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// http.Error(w, "signin page is OK", 200)
		// sigin page rendering ...
		u.render(w, r, "signin.page.html", nil)
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var user models.UserLoginRequestDTO
		err := decoder.Decode(&user)

		if err != nil {
			// http.Error(w, "signin: json decoder error", 500)
			JSON(w, 500, err)
			return
		}

		us, err := u.UserUsecase.GetByEmail(context.Background(), user.Email)
		if err != nil {
			// need tohandle error based on status (500, 401, etc.)
			// http.Error(w, "signin: coudlnt find or unautorized", http.StatusUnauthorized)

			JSON(w, 401, err)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password)); err != nil {
			// need to handle with 401 error
			// http.Error(w, "signin: password is wrong", http.StatusUnauthorized)

			JSON(w, 401, err)
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
		// JSON(w, 302, nil)
	}
}

func (u *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// http.Error(w, "signup: page is OK", 200)

		u.render(w, r, "signup.page.html", nil)
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var userInput models.UserRegisterRequestDTO
		err := decoder.Decode(&userInput)

		if err != nil {
			// http.Error(w, "signup: json decoder error", 500)
			JSON(w, 500, err)
			return
		}

		// singnup form checker to be added here ...

		var user models.User
		user.Username = userInput.Username
		user.Email = userInput.Username
		user.Password = userInput.Password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(hashedPassword)
		// Time
		dateFormat := "2016-10-06 01:50:00 -0700 MST"
		user.CreatedAt.Format(dateFormat)
		user.UpdatedAt.Format(dateFormat)

		user.UserID, err = u.UserUsecase.Create(context.Background(), &user)
		if err != nil {
			// handle error ...

			// http.Error(w, "signup page: couldnt create user", 500)
			JSON(w, 500, err)
			return
		}

		// http.Redirect(w, r, "/", http.StatusSeeOther)
		// http.Error(w, "user has been created", http.StatusCreated)

		JSON(w, 201, user) // user or nil?

		return
	} else {
		// method not allowed
		// http.Error(w, "signup page: method not allowed", http.StatusMethodNotAllowed)

		JSON(w, 405, nil) // custom error?
		return
	}
}

func (u *Handler) signout(w http.ResponseWriter, r *http.Request) {
	ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

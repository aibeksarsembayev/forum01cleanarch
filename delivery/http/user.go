package httpdelivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"git.01.alem.school/quazar/forum-authentication/models"
	"golang.org/x/crypto/bcrypt"
)

func (u *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "signin page is OK", 200)
		// sigin page rendering ...
	}
	if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, "signin: json decoder error", 500)
			return
		}

		us, err := u.UserUsecase.GetByEmail(context.Background(), user.Email)

		if err != nil {
			// need tohandle error based on status (500, 401, etc.)
			http.Error(w, "signin: coudlnt find or unautorized", 401)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password)); err != nil {
			// need to handle with 401 error
			http.Error(w, "signin: password is wrong", 401)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func (u *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "signup: page is OK", 200)
	} else if r.Method == "POST" {
		// JSON decoder
		decoder := json.NewDecoder(r.Body)
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, "signin: json decoder error", 500)
			return
		}

		// singnup form checker to be added here ...

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
			fmt.Println(err)
			http.Error(w, "signup page: couldnt create user", 500)

			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		// method not allowed
		http.Error(w, "signup page: method not allowed", 405)
		return
	}
}

func (u *Handler) signout(w http.ResponseWriter, r *http.Request) {
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

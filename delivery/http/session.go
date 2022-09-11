package httpdelivery

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var sessions sync.Map

func GetSession(r *http.Request) (string, bool) {
	var userEmail string
	cookie, err := r.Cookie("session_token")
	if err == nil {
		if value, ok := sessions.Load(cookie.Value); ok {
			userEmail = fmt.Sprint(value)
			return userEmail, true
		}
	}
	return userEmail, false
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	if cookieValue, err := getValue(r); err == nil {
		sessions.Delete(cookieValue)
	}
	cookie := &http.Cookie{
		Name:   "your-name",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func Set(userEmail, sessionToken string) {

	// check for existing the same user session
	sessions.Range(func(key, value interface{}) bool {
		log.Println("sessions: ", key, value)
		if userEmail == value {
			sessions.Delete(key)
		}
		return true
	})

	sessions.Store(sessionToken, userEmail)
}

func getValue(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if cookie == nil || err != nil {
		return "", err
	}

	cookieValue := cookie.Value

	return cookieValue, err
}

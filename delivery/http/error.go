package httpdelivery

import (
	"encoding/json"
	"log"
	"net/http"

	"git.01.alem.school/quazar/forum-authentication/models"
)

// ResponseError represent the responce error struct
type ResponseError struct {
	Message string `json:"message"`
}

func errorHandler(status int, err error) {
	log.Println(status, err)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrNoRecord:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// JSON returns given structs as jSON into the responce body ...
func JSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(obj)
}

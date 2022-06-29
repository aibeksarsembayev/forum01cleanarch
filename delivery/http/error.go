package httpdelivery

import (
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
	default:
		return http.StatusInternalServerError
	}
}

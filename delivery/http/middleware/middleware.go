package middleware

import (
	"context"

	"net/http"

	httpdelivery "git.01.alem.school/quazar/forum-authentication/delviery/http"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// Middleware name ...
func (m *GoMiddleware) sessionMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := httpdelivery.GetSession(r)
		if err != nil {
			http.Error(w, "please sign-in", http.StatusUnauthorized)
			httpdelivery.ClearSession(w, r)
		}
		r = r.WithContext(context.WithValue(r.Context(), "session", username))
		handler(w, r)
	}
}

// InitMiddleware initialzie the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

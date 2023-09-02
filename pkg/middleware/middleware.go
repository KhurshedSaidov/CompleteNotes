package middleware

import (
	"awesomeNotes/pkg/service"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	UserService *service.UserService
}

func NewAuthMiddleware(userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService}
}

func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("Username")
		password := r.Header.Get("Password")
		user, err := am.UserService.LoginUser(username, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

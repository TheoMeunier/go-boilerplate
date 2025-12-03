package middleware

import (
	"boilerplate/pkg/jwt"
	"fmt"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, `{"error": "missing token"}`, http.StatusUnauthorized)
		return
	}

	token, err := jwt.VerifyToken(tokenString)
	if err != nil || !token.Valid {
		http.Error(w, `{"error": "invalid token"}`, http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Welcome to the the protected area")
}

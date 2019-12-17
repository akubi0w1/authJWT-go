package handler

import (
	"auth-jwt/server"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// WhoAmI get user name by cookie
func WhoAmI(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("who am i")
		// jwtチェック
		// headerから読み出し
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// tokenの認証
		token, err := VerifyToken(tokenString)
		// token, err := verifyToken(r)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// who am i!
		claims := token.Claims.(jwt.MapClaims)
		server.Success(w, &whoAmIResponse{
			Message: fmt.Sprintf("I am %s", claims["user"]),
		})
	}
}

type whoAmIResponse struct {
	Message string `json:"message"`
}

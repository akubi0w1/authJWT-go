package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"auth-jwt/server"
	"database/sql"

	jwt "github.com/dgrijalva/jwt-go"
)

func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Signup")
		// bodyの読み出し
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var req SignupRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// TODO: データのバリデーション

		// passwordのハッシュ
		hash, err := server.PasswordHash(req.Password)
		if err != nil {
			server.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// dbに登録
		_, err = db.Exec(
			"INSERT INTO users(id, name, password) VALUES (?,?,?)",
			req.ID,
			req.Name,
			hash,
		)
		if err != nil {
			server.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// response
		server.Success(w, &SignupResponse{
			ID:       req.ID,
			Name:     req.Name,
			Password: hash,
		})

	}
}

type SignupRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignupResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login")
		// bodyの読み出し
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var req LoginRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// validation
		var hash string
		row := db.QueryRow("SELECT password FROM users WHERE id=?", req.ID)
		if err = row.Scan(&hash); err != nil {
			server.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = server.PasswordVerify(hash, req.Password)
		if err != nil {
			server.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		log.Println("login success: userid = ", req.ID)

		// tokenの発行
		token, err := server.CreateToken(req.ID)
		if err != nil {
			server.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// response
		server.Success(w, &LoginResponse{
			Token: token,
		})

	}
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Logout(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("logout")
		// headerから読み出し
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// tokenの認証
		token, err := server.VerifyToken(tokenString)
		// token, err := verifyToken(r)
		if err != nil {
			server.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// response
		claims := token.Claims.(jwt.MapClaims)
		server.Success(w, &LogoutResponse{
			Message: fmt.Sprintf("bye %s !", claims["user"]),
		})

	}
}

type LogoutResponse struct {
	Message string `json:"message"`
}

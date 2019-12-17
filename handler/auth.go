package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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
		token, err := createToken(req.ID)
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
		token, err := VerifyToken(tokenString)
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

func createToken(userID string) (string, error) {
	// TODO: tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// TODO: claimsの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// TODO: 署名
	var secretKey = "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return token, err
	}

	// // TODO: 有効期限とかのチェックてしてくれんの？
	// if !token.Valid {
	// 	return token, fmt.Errorf("valid %s", "error")
	// }

	return token, nil

}

// func verifyToken(r *http.Request) (*jwt.Token, error) {
// 	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
// 		b := []byte("secret")
// 		return b, nil
// 	})
// 	return token, err
// }

package server

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID string) (string, error) {
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

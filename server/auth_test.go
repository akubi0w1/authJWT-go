package server

import (
	"testing"
)

func TestJWT(t *testing.T) {
	t.Run("create token", func(t *testing.T) {
		userID := "sample_id"
		token, err := CreateToken(userID)
		if err != nil {
			t.Errorf("fail create token. err: %s, token: %s", err, token)
		}
	})

	t.Run("success verify token", func(t *testing.T) {
		userID := "sample_id"
		tokenString, _ := CreateToken(userID)
		token, err := VerifyToken(tokenString)
		if err != nil {
			t.Errorf("fail verify token. err: %s, token: %v", err, token)
		}
		// TODO: claimの値を確認したい
	})

	// 異なるtokenStringを使って認証した時に弾いて
	t.Run("use wrong token verify...英語...", func(t *testing.T) {
		userID := "sample_id"
		tokenString, _ := CreateToken(userID)
		token, err := VerifyToken(tokenString + "a")
		if err == nil {
			t.Error("fail verify token. token: ", token)
		}
		if token != nil {
			t.Error("tokenに情報が乗ってる. token: ", token)
		}
	})
}

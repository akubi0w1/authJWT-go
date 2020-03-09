package server

import (
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	t.Run("password hash", func(t *testing.T) {
		pw := "hogehoge"
		hash, err := PasswordHash(pw)
		if err == nil {
			t.Error("password hash error. err = ", err)
			t.Errorf("row pass: %s, hash: %s", pw, hash)
		}
	})

	t.Run("too long password hashing", func(t *testing.T) {
		longPass := "hogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehoge"
		wrongPass := "wrongpass"
		hash, err := PasswordHash(longPass)
		if err == nil {
			t.Error("return wrong hash.")
			if err = PasswordVerify(hash, wrongPass); err != nil {
				t.Error("can authorized too long password.")
				t.Errorf("rowPass: {too long...}, pass: %s", wrongPass)
			}
		}
	})

	t.Run("success verify", func(t *testing.T) {
		pw := "hogehoge"
		hash, err := PasswordHash(pw)
		if err = PasswordVerify(hash, pw); err != nil {
			t.Error("passwordVerify error. err = ", err)
		}
	})

	t.Run("wrong pass", func(t *testing.T) {
		pw := "hogehoge"
		hash, err := PasswordHash(pw)
		wrongPass := "wrong pass"
		if err = PasswordVerify(hash, wrongPass); err == nil {
			t.Error("can authorized wrong password.")
			t.Errorf("rowPass: %s, pass: %s", pw, wrongPass)
		}
	})
}

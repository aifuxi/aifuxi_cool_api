package test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := "123456"

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	t.Logf("hash: %s\n", hash)

	t.Logf("compare: %v", bcrypt.CompareHashAndPassword(hash, []byte("123456")))
}

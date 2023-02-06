package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
)

func hashPassword(password string) (hash string) {
	h := sha1.New()
	h.Write([]byte(password))
	hash = hex.EncodeToString(h.Sum(nil))
	return
}

func generateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

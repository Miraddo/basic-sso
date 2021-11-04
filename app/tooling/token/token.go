package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenToken() (string, error) {
	b := make([]byte, 60)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

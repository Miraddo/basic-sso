// Package hash will be used for generate Token and Hashes or,
// other random algorithms.
package hash

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

// GenToken is generate random token with 60 character/length.
func GenToken() (string, error) {
	b := make([]byte, 60)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GenPass will get the password and it return the MD5 hash.
func GenPass(pd string) string {
	hash := md5.Sum([]byte(pd))

	return hex.EncodeToString(hash[:])
}

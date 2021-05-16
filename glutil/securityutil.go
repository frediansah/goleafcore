package glutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func MatchPassword(input, hashed, passKey string) bool {
	mac := hmac.New(sha256.New, []byte(passKey))
	mac.Write([]byte(input))
	expectedMAC := mac.Sum(nil)

	hashedByte, err := hex.DecodeString(hashed)
	if err != nil {
		return false
	}

	return hmac.Equal(hashedByte, expectedMAC)
}

func CreatePassword(password, passKey string) string {
	mac := hmac.New(sha256.New, []byte(passKey))
	mac.Write([]byte(password))
	expectedMAC := mac.Sum(nil)

	sha := hex.EncodeToString(expectedMAC)
	return sha
}

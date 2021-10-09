package endpoints

// contains general functions

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Status Not Found"))
}

func EncryptPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash_password := hex.EncodeToString(hash.Sum(nil))

	return hash_password
}

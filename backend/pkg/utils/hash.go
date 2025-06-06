package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashWithSalt(data, salt string) string {
    hash := sha256.Sum256([]byte(data + salt))
    return fmt.Sprintf("%x", hash)
}

func VerifyHash(data, salt, expectedHash string) bool {
    return HashWithSalt(data, salt) == expectedHash
}

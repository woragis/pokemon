package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
    return string(bytes), err
}

func CheckPassword(hashed, plain string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
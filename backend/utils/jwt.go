package utils

import (
	"fmt"
	"pokemon/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint) (string, error) {
	secret := config.GetJWTSecret()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(72 * time.Hour).Unix(),
    })

    return token.SignedString(secret)
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
    secret := config.GetJWTSecret()

    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(secret), nil
    })
}

package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yikuanzz/ecom/config"
)

func CreateJWT(sercret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(sercret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

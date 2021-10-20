package utils

import (
	"AntarJemput-Be-C/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	secret := config.GetString("JWT_SECRET")
	daysCount := config.GetIntt("JWT_SECRET_KEY_EXPIRE_DAYS_COUNT")

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().AddDate(0, 0, daysCount).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

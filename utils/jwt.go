package utils

import (
	"errors"
	"main.go/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (*string, error) {
	jwtHoursExpireEnv := os.Getenv("JWT_HOURS_EXPIRE")
	if jwtHoursExpireEnv == "" {
		jwtHoursExpireEnv = "24"
	}
	jwtHoursExpire, err := strconv.Atoi(jwtHoursExpireEnv)
	if err != nil {
		return nil, err
	}

	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * time.Duration(jwtHoursExpire)).Unix(),
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret"
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func VerifyToken(token string) (*jwt.MapClaims, error) {
	jwtSecretEnv := os.Getenv("JWT_SECRET")
	if jwtSecretEnv == "" {
		jwtSecretEnv = "secret"
	}

	jwtSecret := []byte(jwtSecretEnv)

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, err
	}

	return &claims, nil
}

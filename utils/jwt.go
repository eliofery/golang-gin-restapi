package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	secretKey = "1It2Is3A4Secret5Key"
)

var (
	ErrUnexpectedSigningMethod = fmt.Errorf("неожиданный метод подписи")
)

func GenerateToken(email string, userId int) (string, error) {
	op := "utils.GenerateToken"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return jwtToken, nil
}

func VerifyToken(token string) (int, error) {
	op := "utils.VerifyToken"

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: %w", op, ErrUnexpectedSigningMethod)
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	email, ok := claims["email"].(string)
	if !ok {
		email = string(email)
	}

	userId, ok := claims["userId"].(float64)

	return int(userId), nil
}

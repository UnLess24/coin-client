package jwttoken

import (
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

func Parse(bearerToken string, JWTSecretKey []byte) error {
	jwtToken := bearerToken[len("Bearer "):]

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("token verification", "errorMessage", "Unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return JWTSecretKey, nil
	})
	if err != nil {
		slog.Error("token verification", "errorMessage", err, "token", jwtToken)
		return fmt.Errorf("error while parsing token %w", err)
	}

	if !token.Valid {
		slog.Error("token verification", "token", jwtToken)
		return fmt.Errorf("token is invalid")
	}

	return nil
}

package utils

import (
	"log/slog"
	"os"

	"dadandev.com/dcbt/internal/dto"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strt, err := token.SignedString([]byte(os.Getenv("JWT_AUTH_SECRET")))
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	return strt, err
}
func ValidateJwt(jwtToken string) (dto.UserData, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_AUTH_SECRET")), nil
	})
	if err != nil {
		slog.Error(err.Error())
		return dto.UserData{}, err
	}
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return dto.UserData{
				Id:    claims["id"].(string),
				Email: claims["email"].(string),
				Name:  claims["name"].(string),
			}, err
		}
	}
	return dto.UserData{}, err
}

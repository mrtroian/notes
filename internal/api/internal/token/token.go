package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrtroian/notes/internal/config"
	"github.com/mrtroian/notes/internal/user"
)

func Generate(user *user.User) (string, error) {
	claims := jwt.MapClaims{
		"user":   user.Serialize(),
		"exp":    time.Now().AddDate(0, 0, 7).Unix(),
		"issuer": "notes-app-auth",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	key := config.GetSecret()
	signing, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return signing, nil
}

func Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return config.GetSecret(), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}

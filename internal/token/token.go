package token

import (
	"fmt"
	"time"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/mrtroian/notes/internal/rts"
	"github.com/mrtroian/notes/internal/common"
	"github.com/mrtroian/notes/internal/database/models"
)

func Generate(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user.Serialize(),
	    "exp": time.Now().AddDate(0, 0, 7).Unix(),
	    "issuer":  "notes-app-auth",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	key := rts.GetSecret()
	signing, err := token.SignedString([]byte(key))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return signing, nil
}

func Validate(tokenString string) (common.JSON, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return rts.GetSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}


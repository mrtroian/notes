package hash

import (
    "golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

    if err != nil {
        return "", err
    }

    return string(bytes), nil
}

func Validate(password string, hash string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

    if err != nil {
        return err
    }

    return nil
}

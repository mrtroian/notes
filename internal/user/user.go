package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
}

func NewUser() *User {
	return new(User)
}

// Stub for pooling
func (_ *User) Reset() {}

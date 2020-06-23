package models

import (
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/common"
)

type User struct {
    gorm.Model
    Username     string
    Email        string
    PasswordHash string
}

func (u *User) Serialize() common.JSON {
    return common.JSON{
        "id":       u.ID,
        "username": u.Username,
        "email":    u.Email,
    }
}

func (u *User) Read(m common.JSON) {
    // @TODO: check values
    u.ID = uint(m["id"].(float64))
    u.Username = m["username"].(string)
    u.Email = m["email"].(string)
}

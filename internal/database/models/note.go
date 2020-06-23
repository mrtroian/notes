package models

import (
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/common"
)

type Note struct {
    gorm.Model
    Title  string
    Text   string
    User   User `gorm:"foreignkey:UserID"`
    UserID uint
}

func (n Note) Serialize() common.JSON {
    return common.JSON{
        "id":         n.ID,
        "title":      n.Title,
        "text":       n.Text,
        "user":       n.User.Serialize(),
        "created_at": n.CreatedAt,
    }
}

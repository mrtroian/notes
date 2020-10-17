package note

import (
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/user"
)

type Note struct {
    gorm.Model
    Title  string
    Text   string
    User   user.User `gorm:"foreignkey:UserID"`
    UserID uint
}

func (n Note) Serialize() map[string]interface{} {
    return map[string]interface{}{
        "id":         n.ID,
        "title":      n.Title,
        "text":       n.Text,
        "user":       n.User.Serialize(),
        "created_at": n.CreatedAt,
    }
}

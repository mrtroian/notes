package note

import (
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/user"
)

// @TODO: merge into user as pkg is dependent
type Note struct {
    gorm.Model
    Title  string
    Text   string
    User   *user.User `gorm:"foreignkey:UserID"`
    UserID uint
}

// func NewNote() *Note {
//     return &Note{}
// }

// func (*Note) Reset() {}

package note

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

var (
	ErrForbidden = errors.New("note: entry is not allowed for user")
)

func (n *Note) Create() error {
	database.NewRecord(n)
	database.Create(n)
	return nil
}

func LoadByUserID(uid uint, limit, offset int) ([]Note, error) {
	var notes []Note

	err := database.Preload("User").
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&notes).
		Error

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func Load(uid uint, id int) (*Note, error) {
	n := &Note{}
	err := database.Set("gorm:auto_preload", true).Where("id = ?", id).First(n).Error

	if err != nil {
		return nil, err
	}

	return n, nil
}

func Delete(uid uint, id int) error {
	n := &Note{}

	if err := database.Where("id = ?", id).First(n).Error; err != nil {
		return err
	}

	if uid != n.UserID {
		return ErrForbidden
	}

	return nil
}

func Init(db *gorm.DB) {
	database = db
}

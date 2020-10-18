package user

import "github.com/jinzhu/gorm"

var database *gorm.DB

func FindByUsername(u string) (*User, error) {
	user := NewUser()
	err := database.Where("username = ?", u).First(user).Error
	user.Reset()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func ExistsWithUsername(u string) bool {
	user := NewUser()
	err := database.Where("username = ?", u).First(user).Error
	user.Reset()

	if err != nil {
		// treat err as true
		// @TODO: validation
		return true
	}

	return false
}

func Save(user *User) error {
	database.NewRecord(user)
	database.Create(&user)
	return nil
}

func Init(db *gorm.DB) {
	database = db
}

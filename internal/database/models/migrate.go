package models

import (
    "log"

    "github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{}, &Note{})
    db.Model(&Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
    log.Println("Auto Migration has beed processed")
}

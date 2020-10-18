package database

import (
    "log"
    "os/user"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/mrtroian/notes/internal/config"
    "github.com/mrtroian/notes/internal/note"
)

func autoMigrate(db *gorm.DB) {
    db.AutoMigrate(&user.User{}, &note.Note{})
    db.Model(&note.Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
    log.Println("Auto Migration has beed processed")
}

// Initialize the database
func Init() (*gorm.DB, error) {
    dbPath := config.GetDBPath()
    db, err := gorm.Open("sqlite3", dbPath)

    if err != nil {
        return nil, err
    }

    db.LogMode(true)
    autoMigrate(db)

    log.Println("Connected to database")

    return db, nil
}

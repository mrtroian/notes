package database

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/mrtroian/notes/internal/database/models"
    "github.com/mrtroian/notes/internal/rts"
)

// Initialize the database
func Initialize() (*gorm.DB, error) {
    dbPath := rts.GetDBPath()
    db, err := gorm.Open("sqlite3", dbPath)

    if err != nil {
        return nil, err
    }

    db.LogMode(true)
    models.Migrate(db)

    log.Println("Connected to database")

    return db, nil
}

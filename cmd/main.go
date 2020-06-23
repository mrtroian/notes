package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/api"
    web "github.com/mrtroian/notes/internal/application"
    "github.com/mrtroian/notes/internal/database"
    "github.com/mrtroian/notes/internal/middleware"
    "github.com/mrtroian/notes/internal/rts"
)

func main() {
    if err := rts.IsValid(); err != nil {
        log.Fatal(err)
    }

    // initialize database
    db, err := database.Initialize()

    if err != nil {
        log.Fatal(err)
    }

    app := gin.Default()
    app.Use(database.Add(db))
    app.Use(middleware.JWTMiddleware())
    web.ApplyRoute(app)
    api.ApplyRoutes(app)
    app.Run(":" + rts.GetPort())
}

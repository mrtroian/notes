package main

import (
    "context"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/api"
    web "github.com/mrtroian/notes/internal/application"
    "github.com/mrtroian/notes/internal/database"
    "github.com/mrtroian/notes/internal/middleware"
    "github.com/mrtroian/notes/internal/rts"
    "github.com/mrtroian/notes/internal/signals"
)

func main() {
    if err := rts.IsValid(); err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    router := gin.Default()
    db, err := database.Initialize()

    if err != nil {
        log.Fatal(err)
    }

    signals.Handle(cancel)
    router.Use(database.Add(db))
    router.Use(middleware.JWTMiddleware())
    api.ApplyRoutes(router)
    web.ApplyRoute(router)
    web.UseRouter(router)
    web.Start()
    defer web.Stop()

    <-ctx.Done()
}

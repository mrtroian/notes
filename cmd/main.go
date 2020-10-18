package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/api"
    "github.com/mrtroian/notes/internal/config"
    "github.com/mrtroian/notes/internal/database"
    "github.com/mrtroian/notes/internal/note"
    "github.com/mrtroian/notes/internal/user"
)

var sigChannel = make(chan os.Signal)

func ShutdownHandler(cancelFunc context.CancelFunc) {
    go func() {
        // kill -9 doesn't sound like a graceful shutdown
        // add syscall.SIGKILL otherwise
        signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

        for sig := range sigChannel {
            switch sig {
            case os.Interrupt:
                cancelFunc()
                return
            }
        }
    }()
}

func main() {
    if err := config.IsValid(); err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    db, err := database.Init()

    if err != nil {
        log.Fatal(err)
    }

    user.Init(db)
    note.Init(db)

    ShutdownHandler(cancel)
    api.Start(gin.Default())
    defer api.Stop()

    <-ctx.Done()
}

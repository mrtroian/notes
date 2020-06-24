package application

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/rts"
)

var srv *http.Server

func UseRouter(r *gin.Engine) {
    srv.Handler = r
}

func Start() {
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("server: %s\n", err)
        }
    }()
}

func Stop() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Println("server: forced shutdown:", err)
    } else {
        log.Println("server: shutted down correctly")
    }
}

func init() {
    srv = new(http.Server)
    srv.Addr = fmt.Sprintf("%s:%s", rts.GetHost(), rts.GetPort())
}

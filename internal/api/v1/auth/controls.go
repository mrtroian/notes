package auth

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/common"
    "github.com/mrtroian/notes/internal/database/models"
    "github.com/mrtroian/notes/internal/hash"
    "github.com/mrtroian/notes/internal/request"
    "github.com/mrtroian/notes/internal/token"
)

// duration of one week to be set in cookies
const weekTime int = 60 * 60 * 24 * 7

// sign up a new user
func signup(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    body := new(request.Register)
    exists := new(models.User)

    if err := c.BindJSON(body); err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    if err := db.Where("username = ?", body.Username).First(exists).Error; err == nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusConflict)
        return
    }

    hs, err := hash.Generate(body.Password)

    if err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    user := models.User{
        Username:     body.Username,
        Email:        body.Email,
        PasswordHash: hs,
    }

    db.NewRecord(user)
    db.Create(&user)

    t, err := token.Generate(user)

    if err != nil {
        log.Println(err)
    }

    c.SetCookie("token", t, weekTime, "/", "", false, true)

    c.JSON(http.StatusOK, common.JSON{
        "user":  user.Serialize(),
        "token": t,
    })
}

// sign in a user
func signin(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    body := new(request.Login)
    user := new(models.User)

    if err := c.BindJSON(body); err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    if err := db.Where("username = ?", body.Username).First(user).Error; err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    if err := hash.Validate(body.Password, user.PasswordHash); err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    t, err := token.Generate(*user)

    if err != nil {
        log.Println(err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.SetCookie("token", t, weekTime, "/", "", false, true)
    c.JSON(http.StatusOK, common.JSON{
        "user":  user.Serialize(),
        "token": t,
    })
}

// renew token when token life is less than 3 days
func check(c *gin.Context) {
    userRaw, ok := c.Get("user")

    if !ok {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    user := userRaw.(models.User)
    expiration := int64(c.MustGet("token_expire").(float64))

    if time.Now().AddDate(0, 0, -3).Unix() > expiration {
        t, err := token.Generate(user)

        if err != nil {
            log.Println(err)
        }
        c.SetCookie("token", t, weekTime, "/", "", false, true)
        c.JSON(http.StatusOK, common.JSON{
            "token": t,
            "user":  user.Serialize(),
        })
        return
    }

    c.JSON(http.StatusOK, common.JSON{
        "token": nil,
        "user":  user.Serialize(),
    })
}

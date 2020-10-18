package notes

import (
    "errors"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/user"
)

func readUser(c *gin.Context) (*user.User, error) {
    userInt, ok := c.Get("user")

    if !ok {
        return nil, errors.New("user not in context")
    }

    user, ok := userInt.(*user.User)

    if !ok {
        return nil, errors.New("interface{} to (*user.User) casting error")
    }

    return user, nil
}

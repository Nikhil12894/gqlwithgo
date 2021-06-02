package auth

import (
	"fmt"
	"net/http"

	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/jwt"
	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		// Allow unauthenticated users in
		if header == "" {
			c.JSON(http.StatusForbidden, "token is empety")
			return
		}

		//validate jwt token
		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid token")
		}

		// create user and check if user exists in db
		id, err := db.UserByEmail(username)
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid token")
		} else {
			fmt.Print(id.Email)
		}
		c.Next()
	}
}

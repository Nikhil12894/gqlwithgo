package auth

import (
	"fmt"
	"net/http"

	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/jwt"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		header := c.Request.Header.Get("token")
		fmt.Println(header)
		// Allow unauthenticated users in
		if header == "" {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf(" ******************************************************************** Token is empety"))
		}

		//validate jwt token
		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Invalid token"))

		}

		// create user and check if user exists in db
		id, err := db.UserByEmail(username)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid token"))
		} else {
			fmt.Println(id.Email)
			c.Next()
		}

	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}

	}
}

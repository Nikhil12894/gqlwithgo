package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/graph"
	"github.com/Nikhil12894/gqlwithgo/graph/generated"
	"github.com/Nikhil12894/gqlwithgo/graph/model"
	"github.com/Nikhil12894/gqlwithgo/hash"
	"github.com/Nikhil12894/gqlwithgo/jwt"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db.DB,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterUser(c *gin.Context) {
	var userInput *model.User
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse key heartstrength from http post request error",
		})
		return
	}

	// Create user/
	err := db.DB.Create(&userInput).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	token, err := jwt.GenerateToken(userInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, token)
}

func Login(c *gin.Context) {
	var login *model.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse key heartstrength from http post request error",
		})
		return
	}
	correct := hash.AuthenticateLogin(login, db.UserPasswordByName(login.Username))
	if !correct {
		c.JSON(http.StatusBadRequest, &hash.WrongUsernameOrPasswordError{})
	}
	token, err := jwt.GenerateToken(login.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, token)
}

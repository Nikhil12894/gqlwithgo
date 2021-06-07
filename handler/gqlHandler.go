package handler

import (
	"fmt"
	"net/http"
	"strings"

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
	user, _ := db.UserByEmail(userInput.Email)
	if user != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("user alredy exist"))
	}

	userInput.Password = hash.HashPassword(userInput.Password)
	if strings.HasPrefix(userInput.FirstName, "admin") || strings.HasPrefix(userInput.FirstName, "ADMIN") {
		userInput.UserType = "ADMIN"
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
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("parse key heartstrength from http post request error"))
	}
	correct := hash.AuthenticateLogin(login, db.UserPasswordByName(login.Username))
	if !correct {
		c.AbortWithError(http.StatusBadRequest, &hash.WrongUsernameOrPasswordError{})

	}
	token, err := jwt.GenerateToken(login.Username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

	}
	c.JSON(http.StatusOK, token)
}

func AppLogin(c *gin.Context) {
	var login *model.Login
	if err := c.BindJSON(&login); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("parse key heartstrength from http post request error"))
	}
	user, err := db.UserByEmail(login.Username)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &hash.WrongUsernameOrPasswordError{})
	}

	correct := hash.AuthenticateLogin(login, user.Password)
	if !correct {
		c.AbortWithError(http.StatusBadRequest, &hash.WrongUsernameOrPasswordError{})

	}
	token, err := jwt.GenerateToken(login.Username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

	}
	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"mobile":    user.Mobile,
		"image":     user.Image,
		"usertype":  user.UserType,
		"id":        user.ID,
	})
}

func Refrace(c *gin.Context) {
	// var m *map{}}{}
	// if err := c.BindJSON(&m); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "parse key heartstrength from http post request error",
	// 	})
	// 	return
	// }
	// correct := hash.AuthenticateLogin(login, db.UserPasswordByName(login.Username))
	// if !correct {
	// 	c.JSON(http.StatusBadRequest, &hash.WrongUsernameOrPasswordError{})
	// }
	// token, err := jwt.GenerateToken(login.Username)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// }
	// c.JSON(http.StatusOK, token)
}

func Testquery(c *gin.Context) {
	reqbody := map[string]interface{}{}
	if err := c.BindJSON(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse key heartstrength from http post request error",
		})
		return
	}
	query := reqbody["query"].(string)
	if strings.HasPrefix(query, "select") {
		result := []map[string]interface{}{}
		err := db.DB.Raw(query, 3).Scan(&result).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, result)
	} else {
		result := map[string]interface{}{}
		err := db.DB.Exec(query, &result).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, result)
	}
}

func BookingVachiIdAndPrice(c *gin.Context) {
	reqbody := map[string]interface{}{}
	if err := c.BindJSON(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse key heartstrength from http post request error",
		})
		return
	}
	m := make(map[string]interface{})
	allBookedVachil := fmt.Sprintf("%v", reqbody["allBookedVachil"])
	allBookedPrice := fmt.Sprintf("%v", reqbody["allBookedPrice"])
	if len(allBookedVachil) > 2 {
		allBookedVachil = strings.Replace(allBookedVachil, "[", "(", -1)
		allBookedVachil = strings.Replace(allBookedVachil, "]", ")", -1)
		allBookedVachil = strings.Replace(allBookedVachil, " ", ",", -1)
		result := []map[string]interface{}{}
		err := db.DB.Raw(fmt.Sprintf("select id,images FROM vachils WHERE id in %v", allBookedVachil), 3).Scan(&result).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		m["allBookedVachil"] = result
	}

	if len(allBookedPrice) > 2 {
		allBookedPrice = strings.Replace(allBookedPrice, "[", "(", -1)
		allBookedPrice = strings.Replace(allBookedPrice, "]", ")", -1)
		allBookedPrice = strings.Replace(allBookedPrice, " ", ",", -1)
		result := []map[string]interface{}{}
		err := db.DB.Raw(fmt.Sprintf("SELECT * FROM total_prices WHERE id in %v", allBookedPrice), 3).Scan(&result).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		m["allBookedPrice"] = result
	}
	c.JSON(http.StatusOK, m)
}

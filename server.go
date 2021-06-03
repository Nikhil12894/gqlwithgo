package main

import (
	"log"
	"os"

	"github.com/Nikhil12894/gqlwithgo/auth"
	"github.com/Nikhil12894/gqlwithgo/config"
	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/handler"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	runApp()
}

func runApp() {
	configuration, err := config.FromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	db.ConnectDataBase(configuration.DBConfig.Vender, configuration.DBConfig.DbName, configuration.IsProd)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Setting up Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(auth.CORSMiddleware())
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.Login)
	r.POST("/refrace", handler.Refrace)
	r.POST("/testquery", handler.Testquery)
	v1 := r.Group("v1")
	v1.Use(auth.Middleware())
	v1.POST("/query", handler.GraphqlHandler())
	r.GET("/", handler.PlaygroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", configuration.Server.Port)
	log.Fatal(r.Run(":" + configuration.Server.Port))
}

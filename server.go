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
	configuration, err := config.FromFile("config.json")
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
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.Login)
	r.POST("/query", handler.GraphqlHandler(), auth.Middleware())
	r.GET("/", handler.PlaygroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}

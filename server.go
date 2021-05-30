package main

import (
	"log"
	"os"

	"github.com/Nikhil12894/gqlwithgo/handler"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	runApp()
}

func runApp() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Setting up Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/query", handler.GraphqlHandler())
	r.GET("/", handler.PlaygroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}

// with mix
/* func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} */

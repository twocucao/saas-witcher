package main

import (
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	//r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}

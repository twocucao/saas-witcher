package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/playground"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// Defining the Graphql handler
//func graphqlHandler() gin.HandlerFunc {
//	// NewExecutableSchema and Config are in the generated.go file
//	// Resolver is in the resolver.go file
//	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
//
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
//}

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

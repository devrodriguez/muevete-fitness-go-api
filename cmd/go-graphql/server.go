package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph/generated"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL playground", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	server.POST("/query", func(c *gin.Context) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

		srv.ServeHTTP(c.Writer, c.Request)
	})

	server.Run(defaultPort)
}

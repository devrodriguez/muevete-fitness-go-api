package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/graphql/graph"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/graphql/graph/generated"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/rest"
	"github.com/devrodriguez/muevete-fitness-go-api/middlewares"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(middlewares.EnableCORS())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.POST("/auth/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, rest.APIResponse{
			Data: gin.H{
				"token": "456ty.fybuijnk.87568hdsfs",
				"c":     "1235",
			},
		})
	})
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

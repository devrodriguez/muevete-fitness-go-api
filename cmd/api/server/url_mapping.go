package server

import (
	"net/http"

	"github.com/devrodriguez/muevete-fitness-go-api/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapUrls(server *gin.Engine, dbcli *mongo.Client) {
	pubRouter := server.Group("/public")
	pubRouter.Use(middlewares.EnableCORS())
	{
		pubRouter.GET("/handshake", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello")
		})
	}

	authRouter := server.Group("/")
	authRouter.Use(middlewares.EnableCORS())
	{
		authRouter.GET("/routines", func(c *gin.Context) {
			c.JSON(http.StatusOK, "some routines")
		})
	}
}

package server

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/rest"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/session"
	"net/http"

	"github.com/devrodriguez/muevete-fitness-go-api/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapUrls(server *gin.Engine, dbcli *mongo.Client) {

	// Sessions
	sesRepo := dbmongo.NewDbSessionCrud(dbcli)
	sesUc := sessions.NewCrudSession(sesRepo)
	sesHand := rest.NewSessionHand(sesUc)

	// Routines
	rtRepo := dbmongo.NewDbRoutineCrud(dbcli)
	rtUc := routines.NewCrudRoutine(rtRepo)
	rtHand := rest.NewRoutineHand(rtUc)

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
		authRouter.GET("/sessions", sesHand.GetAllSessions)
		authRouter.POST("/sessions", sesHand.CreateSession)

		authRouter.GET("/routines", rtHand.GetAllRoutines)
		authRouter.POST("/routines", rtHand.CreateRoutine)
	}
}

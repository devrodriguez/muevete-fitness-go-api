package server

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/categories"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/customers"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/rest"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/sessions"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/weeklies"
	"net/http"

	"github.com/devrodriguez/muevete-fitness-go-api/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapUrls(server *gin.Engine, dbCli *mongo.Client) {

	// Sessions
	sesRepo := dbmongo.NewDbSessionCrud(dbCli)
	sesUc := sessions.NewCrudSession(sesRepo)
	sesHand := rest.NewSessionHand(sesUc)

	// Routines
	rtRepo := dbmongo.NewDbRoutineCrud(dbCli)
	rtUc := routines.NewCrudRoutine(rtRepo)
	rtHand := rest.NewRoutineHand(rtUc)

	// Categories
	catRepo := dbmongo.NewDbCategoryCrud(dbCli)
	catUc := categories.NewCategoryCrud(catRepo)
	catHand := rest.NewCategoryHand(catUc)

	// Customers
	cusRepo := dbmongo.NewDbCustomerCrud(dbCli)
	cusUc := customers.NewCustomerCrud(cusRepo)
	cusHand := rest.NewCustomerHand(cusUc)

	// Routine Schedule
	rsRepo := dbmongo.NewDbRoutineSchedule(dbCli)
	rsUc := routines.NewRoutineSchedule(rsRepo)
	rsHand := rest.NewRoutineScheduleHand(rsUc)

	// Weekly
	wkRepo := dbmongo.NewDbWeeklyCrud(dbCli)
	wkUc := weeklies.NewCustomerCrud(wkRepo)
	wkHand := rest.NewWeeklyHand(wkUc)

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
		authRouter.POST("/routines/schedule", rsHand.CreateRoutineSchedule)

		authRouter.GET("/categories", catHand.GetAllCategories)
		authRouter.POST("/categories", catHand.CreateCategory)

		authRouter.GET("/customers", cusHand.GetAllCustomers)
		authRouter.POST("/customers", cusHand.CreateCustomer)

		authRouter.GET("/weeklies", wkHand.GetAllWeeklies)
		authRouter.POST("/weeklies", wkHand.CreateRoutine)
	}
}

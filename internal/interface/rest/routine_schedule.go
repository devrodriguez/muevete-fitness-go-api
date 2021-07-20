package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoutineScheduleHand struct {
	uc routines.IRoutineSchedule
}

func NewRoutineScheduleHand(uc routines.IRoutineSchedule) RoutineScheduleHand {
	return RoutineScheduleHand{
		uc,
	}
}

func (rsh *RoutineScheduleHand) GetRoutineSchedule(c *gin.Context) {
	rss, err := rsh.uc.GetSchedule(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    rss,
	})
}

func (rsh *RoutineScheduleHand) CreateRoutineSchedule(c *gin.Context) {
	var rs domain.RoutineScheduleMod

	if err := c.BindJSON(&rs); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := rsh.uc.CreateSchedule(c, rs); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

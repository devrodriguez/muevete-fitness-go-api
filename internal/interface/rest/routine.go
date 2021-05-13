package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoutineHand struct {
	uc routines.ICrudRoutine
}

func NewRoutineHand(uc routines.ICrudRoutine) RoutineHand {
	return RoutineHand{
		uc,
	}
}

func (rh *RoutineHand) CreateRoutine(c *gin.Context) {
	var r domain.Routine

	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := rh.uc.Create(c, r); err !=  nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (rh *RoutineHand) GetAllRoutines(c *gin.Context) {

	rs, err := rh.uc.GetAll(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    rs,
	})
}
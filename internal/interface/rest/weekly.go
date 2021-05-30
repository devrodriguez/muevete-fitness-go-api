package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/weeklies"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeeklyHand struct {
	uc weeklies.IWeeklyCrud
}

func NewWeeklyHand(uc weeklies.IWeeklyCrud) WeeklyHand {
	return WeeklyHand{
		uc,
	}
}

func (wh *WeeklyHand) CreateRoutine(c *gin.Context) {
	var w domain.WeeklyMod

	if err := c.BindJSON(&w); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := wh.uc.CreateWeekly(c, w); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (wh *WeeklyHand) GetAllWeeklies(c *gin.Context) {

	ds, err := wh.uc.GetAllWeeklies(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    ds,
	})
}

package rest

import (
	"net/http"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/sessions"
	"github.com/gin-gonic/gin"
)

type SessionScheduleHand struct {
	uc sessions.ISessionSchedule
}

func NewSessionScheduleHand(ss sessions.ISessionSchedule) SessionScheduleHand {
	return SessionScheduleHand{
		ss,
	}
}

func (ssh *SessionScheduleHand) GetSessionsSchedule(c *gin.Context) {

	sch, err := ssh.uc.GetSessionsSchedule(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    sch,
	})
}

func (ssh *SessionScheduleHand) CreateSessionSchedule(c *gin.Context) {
	var newSsch domain.SessionScheduleMod

	if err := c.BindJSON(&newSsch); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := ssh.uc.CreateSessionsSchedule(c, newSsch); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Reason:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

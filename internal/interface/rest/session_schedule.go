package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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

	sch, err := ssh.uc.GetScheduleSessions(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    sch,
	})
}

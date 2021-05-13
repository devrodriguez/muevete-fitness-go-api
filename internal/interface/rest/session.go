package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/session"
	"github.com/gin-gonic/gin"
	"net/http"
)


type SessionHand struct {
	uc sessions.ICrudSession
}

func NewSessionHand(uc sessions.ICrudSession) SessionHand {
	return SessionHand{
		uc,
	}
}

func (sh *SessionHand) CreateSession(c *gin.Context) {
	var session domain.Session

	if err := c.BindJSON(&session); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := sh.uc.Create(c, session); err !=  nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (sh *SessionHand) GetAllSessions(c *gin.Context) {

	ds, err := sh.uc.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data: ds,
	})
}

package sessions

import (
"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
"github.com/gin-gonic/gin"
)

type ISessionSchedule interface {
	GetScheduleSessions(c *gin.Context) ([]domain.ScheduleSession, error)
}

type ImpSessionSchedule struct {
	dbImp dbmongo.IDbSessionSchedule
}

func NewSessionSchedule(dbImp dbmongo.IDbSessionSchedule) ISessionSchedule {
	return &ImpSessionSchedule{
		dbImp,
	}
}

func (cs *ImpSessionSchedule) GetScheduleSessions(c *gin.Context) ([]domain.ScheduleSession, error) {
	schs, err := cs.dbImp.GetSessionSchedule(c)

	if err != nil {
		return nil, err
	}

	return schs, err
}


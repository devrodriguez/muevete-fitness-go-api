package sessions

import (
"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
"github.com/gin-gonic/gin"
)

type ISessionSchedule interface {
	GetSessionsSchedule(*gin.Context) ([]domain.SessionSchedule, error)
	CreateSessionsSchedule(*gin.Context, domain.SessionScheduleMod) error
}

type ImpSessionSchedule struct {
	dbImp dbmongo.IDbSessionSchedule
}

func NewSessionSchedule(dbImp dbmongo.IDbSessionSchedule) ISessionSchedule {
	return &ImpSessionSchedule{
		dbImp,
	}
}

func (cs *ImpSessionSchedule) GetSessionsSchedule(c *gin.Context) ([]domain.SessionSchedule, error) {
	schs, err := cs.dbImp.GetSessionSchedule(c)

	if err != nil {
		return nil, err
	}

	return schs, err
}

func (cs *ImpSessionSchedule) CreateSessionsSchedule(c *gin.Context, ss domain.SessionScheduleMod) error {
	if err := cs.dbImp.SaveSessionSchedule(c, ss); err != nil {
		return err
	}

	return nil
}


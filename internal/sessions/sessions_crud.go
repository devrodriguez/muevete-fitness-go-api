package sessions

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICrudSession interface {
	GetAllSessions(c *gin.Context) ([]domain.Session, error)
	CreateSession(c *gin.Context, ds domain.Session) error
}

type ImpCrudSession struct {
	dbImp dbmongo.IDbSessionCrud
}

func NewCrudSession(dbImp dbmongo.IDbSessionCrud) ICrudSession {
	return &ImpCrudSession{
		dbImp,
	}
}

func (cs *ImpCrudSession) GetAllSessions(c *gin.Context) ([]domain.Session, error) {
	ses, err := cs.dbImp.GetAllSessions(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cs *ImpCrudSession) CreateSession(c *gin.Context, ses domain.Session) error {
	err := cs.dbImp.InsertSession(c, ses)

	if err != nil {
		return err
	}

	return nil
}

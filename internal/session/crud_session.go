package sessions

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICrudSession interface {
	GetAll(c *gin.Context) ([]domain.Session, error)
	Create(c *gin.Context, ds domain.Session) error
}

type ImpCrudSession struct {
	dbImp dbmongo.IDbSessionCrud
}

func NewCrudSession(dbImp dbmongo.IDbSessionCrud) ICrudSession {
	return &ImpCrudSession{
		dbImp,
	}
}

func (cs *ImpCrudSession) GetAll(c *gin.Context) ([]domain.Session, error) {
	ses, err := cs.dbImp.GetAll(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cs *ImpCrudSession) Create(c *gin.Context, ses domain.Session) error {
	err := cs.dbImp.Create(c, ses)

	if err != nil {
		return err
	}

	return nil
}

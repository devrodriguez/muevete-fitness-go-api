package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbSessionCrud interface {
	GetAllSessions(*gin.Context) ([]domain.Session, error)
	CreateSession(*gin.Context, domain.Session) error
}

type ImpDbSessionCrud struct {
	*mongo.Client
}

func NewDbSessionCrud(cli *mongo.Client) IDbSessionCrud {
	return &ImpDbSessionCrud{
		cli,
	}
}

func (sc *ImpDbSessionCrud) GetAllSessions(c *gin.Context) ([]domain.Session, error) {
	var sess []domain.Session

	findOpt := options.Find()
	docRef := sc.Client.Database("fitness").Collection("sessions")
	cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var ses domain.Session

		if err := cursor.Decode(&ses); err != nil {
			panic(err)
		}

		sess = append(sess, ses)
	}

	return sess, nil
}

func (sc *ImpDbSessionCrud) CreateSession(c *gin.Context, ses domain.Session) error {
	docRef := sc.Client.Database("fitness").Collection("sessions")

	_, err := docRef.InsertOne(c, ses)

	if err != nil {
		return err
	}

	return nil
}

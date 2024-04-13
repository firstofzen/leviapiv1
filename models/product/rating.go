package product

import (
	"github.com/gocql/gocql"
	"levi-apis/models"
	"time"
)

type Rating struct {
	Id          gocql.UUID          `json:"id"`
	RatingUsers []models.RatingUser `json:"rating_users" db:"rating_users"`
}

func (r Rating) New() *Rating {
	return &Rating{
		Id:          gocql.UUIDFromTime(time.Now()),
		RatingUsers: []models.RatingUser{},
	}
}

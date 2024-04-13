package models

import (
	"github.com/gocql/gocql"
	"time"
)

type RatingUser struct {
	EmailUser string  `json:"email_user" db:"email_user"`
	Feedback  string  `json:"feedback" db:"feedback"`
	Score     float32 `json:"score" db:"score"`
}
type ProductIdentify struct {
	Id        gocql.UUID `json:"id" db:"id"`
	Category2 string     `json:"category_2" db:"category_2"`
}
type OrderStatus struct {
	LastUpdate time.Time `json:"last_update" db:"last_update"`
	Status     string    `json:"status" db:"status"`
}

package response

import (
	"github.com/gocql/gocql"
	"time"
)

type RespProduct struct {
	Category1   string     `json:"category_1"`
	Category2   string     `json:"category_2"`
	Id          gocql.UUID `json:"id"`
	Email       string     `json:"email"`
	AvatarUrl   string     `json:"avatar_url"`
	PicturesUrl []string   `json:"pictures_url"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	TotalRating float64    `json:"total_rating"`
	Sale        float64    `json:"sale"`
	TtlSale     int        `json:"ttl_sale"`
	Sold        int        `json:"sold"`
	Inventory   int        `json:"inventory"`
	PublishAt   time.Time  `json:"publish_at"`
}
type RespProductRating struct {
}

package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"levi-apis/models/customerror"
	"time"
)

var valErr customerror.ValidateError

type Product struct {
	Category1   string     `json:"category_1" db:"category_1"`
	Category2   string     `json:"category_2" db:"category_2"`
	Id          gocql.UUID `json:"id" db:"id"`
	Email       string     `json:"email" db:"email"`
	AvatarUrl   string     `json:"avatar_url" db:"avatar_url"`
	PicturesUrl []string   `json:"pictures_url" db:"pictures_url"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Price       float64    `json:"price" db:"price"`
	TotalRating float64    `json:"total_rating" db:"total_rating"`
	Sale        float64    `json:"sale" db:"sale"`
	Sold        int        `json:"sold" db:"sold"`
	Inventory   int        `json:"inventory" db:"inventory"`
	PublishAt   time.Time  `json:"publish_at" db:"publish_at"`
}

func (p Product) New(c1 string, c2 string, email string, avatarUrl string, pictureUrls []string, title string, description string, price float64, sale float64) *Product {

	return &Product{
		Category1:   c1,
		Category2:   c2,
		Id:          gocql.UUIDFromTime(time.Now()),
		Email:       email,
		AvatarUrl:   avatarUrl,
		PicturesUrl: pictureUrls,
		Title:       title,
		Description: description,
		Price:       price,
		TotalRating: 0,
		Sale:        sale,
		Sold:        0,
		Inventory:   0,
		PublishAt:   time.Now(),
	}
}

func (p Product) validate() error {
	return validator.New().Struct(p)
}

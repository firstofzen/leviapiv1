package trade

import (
	"levi-apis/models"
	"levi-apis/models/customerror"
)

var validateErr customerror.ValidateError

type Booth struct {
	Email                 string                   `json:"email" db:"email"`
	Brand                 string                   `json:"brand" db:"brand"`
	IsAuthentic           bool                     `json:"is_authentic" db:"is_authentic"`
	Description           string                   `json:"description" db:"description"`
	MainCategory          []string                 `json:"main_category" db:"main_category"`
	ProductIdentifies     []models.ProductIdentify `json:"product_identifies" db:"product_identifies"`
	ProductSoldIdentifies []models.ProductIdentify `json:"product_sold_identifies" db:"product_sold_identifies"`
}

func (b Booth) New(email string, des string, brand string) *Booth {
	return &Booth{
		Email:                 email,
		Brand:                 brand,
		IsAuthentic:           false,
		Description:           des,
		ProductIdentifies:     []models.ProductIdentify{},
		ProductSoldIdentifies: []models.ProductIdentify{},
	}
}

package product

import "github.com/gocql/gocql"

type Ranking struct {
	ProductCategory2 string     `json:"product_category_2,omitempty" db:"product_category_2"`
	ProductId        gocql.UUID `json:"product_id,omitempty" db:"product_id"`
	Sold             int        `json:"sold,omitempty" db:"sold"`
	TotalRating      float64    `json:"total_rating,omitempty" db:"total_rating"`
}

func (r Ranking) New(productCategory2 string, productId gocql.UUID) *Ranking {
	return &Ranking{ProductId: productId, ProductCategory2: productCategory2, Sold: 0, TotalRating: 0}
}

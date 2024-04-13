package trade

import (
	"github.com/go-playground/validator/v10"
)

type Cart struct {
	Email    string   `json:"email" db:"order_ids"`
	OrderIds []string `json:"order_ids" db:"order_ids"`
}

func (C Cart) New(email string) *Cart {
	return &Cart{Email: email, OrderIds: []string{}}
}

func (c *Cart) validate() error {
	val := validator.New()
	return val.Struct(c)
}

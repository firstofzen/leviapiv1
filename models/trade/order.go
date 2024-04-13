package trade

import (
	"github.com/gocql/gocql"
	"levi-apis/models"
	"time"
)

type Order struct {
	Email           string                 `json:"email" db:"email"`
	ProductIdentify models.ProductIdentify `json:"product_identify" db:"product_identify"`
	Id              gocql.UUID             `json:"id" db:"id"`
	CreateAt        time.Time              `json:"create_at" db:"create_at"`
	TotalAmount     float64                `json:"total_amount" db:"total_amount"`
	Status          models.OrderStatus     `json:"status" db:"status"`
}

func (o Order) New(email string, pc2 string, pid gocql.UUID, totalAmount float64) *Order {
	return &Order{
		Email: email,
		ProductIdentify: models.ProductIdentify{
			Id:        pid,
			Category2: pc2,
		},
		Id:          gocql.UUIDFromTime(time.Now()),
		CreateAt:    time.Now(),
		TotalAmount: totalAmount,
		Status: models.OrderStatus{
			LastUpdate: time.Now(),
			Status:     "pending",
		},
	}
}

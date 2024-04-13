package user

import (
	"github.com/gocql/gocql"
	"time"
)

type Account struct {
	Email     string     `json:"email" db:"email"`
	IdCart    gocql.UUID `json:"id_cart" db:"id_cart"`
	IdBooth   gocql.UUID `json:"id_booth" db:"id_booth"`
	IdDetail  gocql.UUID `json:"id_detail" db:"id_detail"`
	Name      string     `json:"name" db:"name"`
	AvatarUrl string     `json:"avatar_url" db:"avatar_url"`
}

type Detail struct {
	Id        gocql.UUID `json:"id" db:"id"`
	Phone     string     `json:"phone" db:"phone"`
	Address   string     `json:"address" db:"address"`
	IsViolate bool       `json:"is_violate" db:"is_violate"`
}

func (a Account) New(email string, name string, avatarUrl string) *Account {
	return &Account{
		Email: email, Name: name, AvatarUrl: avatarUrl, IdDetail: uuidFromTime(), IdBooth: uuidFromTime(), IdCart: uuidFromTime(),
	}

}
func (d Detail) New(phone string, address string) *Detail {
	return &Detail{Id: uuidFromTime(), IsViolate: false, Phone: phone, Address: address}
}
func uuidFromTime() gocql.UUID {
	return gocql.UUIDFromTime(time.Now())
}

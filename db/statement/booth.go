package statement

import (
	"fmt"
)

type BoothStmt struct {
}

func (b BoothStmt) Add(email string, brand string, des string) string {
	return fmt.Sprintf("insert into trade.booth(email, brand, is_authentic, description) values ('%s', '%s', false, '%s')", email, brand, des)
}
func (b BoothStmt) Get(email string) string {
	return fmt.Sprintf("select * from trade.booth where email='%s'", email)
}
func (b BoothStmt) UpdateProduct(c2 string, id string, email string, isInsert bool) string {
	if isInsert {
		return fmt.Sprintf("update trade.booth set product_identifies = product_identifies + {category_2: '%s', id: %s} where email='%s'", c2, id, email)
	}
	return fmt.Sprintf("update trade.booth set product_identifies = product_identifies - {category_2: '%s', id: %s} where email='%s'", c2, id, email)
}
func (b BoothStmt) UpdateSoldProduct(c2 string, id string, email string) string {
	return fmt.Sprintf("update trade.booth set product_sold_identifies = product_sold_identifies + {category_2: '%s', id: %s} where email='%s'", c2, id, email)
}
func (b BoothStmt) UpdateAField(email string, name string, value string) string {
	return fmt.Sprintf("update trade.booth set %s = %s where email = '%s'", name, value, email)
}
func (b BoothStmt) Delete(email string) string {
	return fmt.Sprintf("delete from trade.booth where email = '%s'", email)
}

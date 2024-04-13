package statement

import (
	"fmt"
	"strconv"
	"time"
)

type OrderStmt struct {
}

func (o OrderStmt) Add(email string, id string, pc2 string, pid string, totalAmount float64) string {
	return fmt.Sprintf("insert into trade.orders(email, id, product_identify, create_at, total_amount, status) values ('%s', %s, {category_2: '%s', id: %s}, "+strconv.FormatInt(time.Now().UnixNano(), 10)+", %f, 'pending')", email, id, pc2, pid, totalAmount)
}
func (o OrderStmt) GetByEmail(email string) string {
	return fmt.Sprintf("select * from trade.orders where email = '%s'", email)
}
func (o OrderStmt) UpdateStatus(status string, email string, id string) string {
	return fmt.Sprintf("update trade.orders set status='%s' where email='%s' and id = %s", status, email, id)
}
func (o OrderStmt) Delete(email string, id string) string {
	return fmt.Sprintf("delete from trade.orders where email='%s' and id = %s", email, id)
}

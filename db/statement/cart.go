package statement

import "fmt"

type CartStmt struct {
}
func (c CartStmt) Add(email string) string {
	return fmt.Sprintf("insert into trade.cart(email) values ('%s')", email)
}
func (c CartStmt) Get(email string) string {
	return fmt.Sprintf("select * from trade.cart where email='%s'", email)
}
func (c CartStmt) UpdateOrderIds(isInsert bool, email string, orderId string) string {
	if isInsert {
		return fmt.Sprintf("update trade.cart set order_ids = order_ids + %s where email='%s'", orderId, email)
	}
	return fmt.Sprintf("update trade.cart set order_ids = order_ids - %s where email='%s'", orderId, email)
}
func (c CartStmt) Delete(email string) string {
	return fmt.Sprintf("delete from trade.cart where email = '%s'", email)
}

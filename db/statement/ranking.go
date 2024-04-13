package statement

import "fmt"

type RankingStmt struct{}

func AddProduct(c2 string, id string, sold int, totalRating float64) string {
	return fmt.Sprintf("insert into product.ranking (product_category_2, product_id, sold, total_rating) values ('%s', %s, %d, %f)", c2, id, sold, totalRating)
}
func GetTopByC(size int, c2 string) string {
	return fmt.Sprintf("select * from product.ranking where product_category_2 = '%s' limit %d", c2, size)
}
func DeleteProduct(c2 string, id string) string {
	return fmt.Sprintf("delete from product.ranking where product_category_2='%s' and product_id=%s", c2, id)
}

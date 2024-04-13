package statement

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ProductStmt struct {
}

func (p ProductStmt) Add(category2 string, category1 string, id string, email string, avatarUrl string, picturesUrl []string, title string, price float64, sale float64, inven int) string {
	return fmt.Sprintf("insert into product.product(category_1, category_2, id, email, avatar_url, pictures_url, title, price, total_rating, sale, sold, inventory, publish_at) values ('%s', '%s', %s, '%s', '%s', "+arrStrToStr(picturesUrl)+", '%s', %f, %f, %f, %d, %d, %s)", category1, category2, id, email, avatarUrl, title, price, 0, sale, 0, 0, strconv.FormatInt(time.Now().UnixNano(), 10))
}
func (p ProductStmt) Get(c2 string, id string) string {
	return fmt.Sprintf("select * from product.product where category_2='%s' and id = %s", c2, id)
}
func (p ProductStmt) GetProductByCategory2(c2 string, size int) string {
	return fmt.Sprintf("select * from product.product where category_2='%s' limit %d", c2, size)
}
func (p ProductStmt) GetCountProductByCategory2(c2 string) string {
	return fmt.Sprintf("select count(*) from product.product where category_2='%s'", c2)
}
func (p ProductStmt) UpdateSale(s float64, ttl int, c2 string, id string) string {
	return fmt.Sprintf("update product.product using ttl %d set sale = %f where category_2='%s' and id=%s", ttl, s, c2, id)
}
func (p ProductStmt) UpdateSold(isIncrease bool, c2 string, id string) string {
	if isIncrease {
		return fmt.Sprintf("update product.product set sold = sold + 1 where category_2 = '%s' and id = %s", c2, id)
	}
	return fmt.Sprintf("update product.product set sold = sold - 1 where category_2 = '%s' and id = %s", c2, id)
}
func (p ProductStmt) UpdateInventory(i int, c2 string, id string) string {
	return fmt.Sprintf("update product.product set inventory = %d where category_2 = '%s' and id = %s", i, c2, id)
}
func (p ProductStmt) UpdateTotalRating(c2 string, id string, t float64) string {
	return fmt.Sprintf("update product.product set total_rating=%f where category_2='%s' and id = %s", t, c2, id)
}
func (p ProductStmt) Delete(c2 string, id string) string {
	return fmt.Sprintf("delete from product.product where category_2='%s' and id=%s", c2, id)
}
func arrStrToStr(arr []string) string {
	var sb strings.Builder

	sb.WriteString("{")
	for _, e := range arr {
		sb.WriteString("'")
		sb.WriteString(e)
		sb.WriteString("',")
	}
	sb.WriteString("}")
	return sb.String()
}

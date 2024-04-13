package statement

import (
	"fmt"
	"levi-apis/models"
)

type RatingStmt struct {
}

func (r RatingStmt) Add(productCategory2 string, productId string) string {
	return fmt.Sprintf("insert into product.rating(product_category_2, product_id) values ('%s', %s)", productCategory2, productId)
}
func (r RatingStmt) Get(productCategory2 string, productId string) string {
	return fmt.Sprintf("select * from product.rating where product_category_2='%s' and product_id=%s", productCategory2, productId)
}
func (r RatingStmt) Delete(productCategory2 string, productId string) string {
	return fmt.Sprintf("delete from product.rating where product_category_2='%s' and product_id=%s", productCategory2, productId)
}
func (r RatingStmt) AddRatingUser(ratingUsr models.RatingUser, productCategory2 string, productId string) string {
	return fmt.Sprintf("update product.rating set rating_users = {email_user:'%s',feedback: '%s', score: %f} where product_category_2= '%s' and product_id=%s", ratingUsr.EmailUser, ratingUsr.Feedback, ratingUsr.Score, productCategory2, productId)
}

package router

import (
	"github.com/gin-gonic/gin"
	handler2 "levi-apis/apis/handler"
)

type ProductRouter struct {
	E *gin.Engine
}

func (r *ProductRouter) Route() {
	var handler handler2.ProductHandler
	r.E.POST("/product/add", func(c *gin.Context) {
		handler.AddProduct(c)
	})
	r.E.GET("/product/get/id", func(c *gin.Context) {
		handler.GetProductByID(c)
	})
	r.E.GET("/product/get/category2", func(c *gin.Context) {
		handler.GetProductByCategory2(c)
	})
	r.E.PATCH("/product/update/sale", func(c *gin.Context) {
		handler.UpdateSaleProduct(c)
	})
	r.E.GET("/product/delete", func(c *gin.Context) {
		handler.DeleteProduct(c)
	})

	r.E.GET("/product/rating/get", func(c *gin.Context) {
		handler.GetProductRating(c)
	})
	r.E.POST("/product/rating/add_rating_user", func(c *gin.Context) {
		handler.AddRatingUser(c)
	})
	r.E.DELETE("/product/rating/remove_rating_user", func(c *gin.Context) {
		handler.RemoveRatingUser(c)
	})

	r.E.GET("/product/ranking/get", func(c *gin.Context) {
		handler.GetProductRanking(c)
	})
}

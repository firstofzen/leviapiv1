package router

import (
	"github.com/gin-gonic/gin"
	handler2 "levi-apis/apis/handler"
)

type TradeRouter struct {
	E *gin.Engine
}

func (r *TradeRouter) Route() {
	var handler handler2.TradeHandler
	r.E.GET("/booth/get", func(c *gin.Context) {
		handler.BoothGet(c)
	})
	r.E.PATCH("/booth/update/description", func(c *gin.Context) {
		handler.BoothUpdateDescription(c)
	})
	r.E.PATCH("/booth/update/brand", func(c *gin.Context) {
		handler.BoothUpdateBrand(c)
	})
	r.E.GET("/cart/get", func(c *gin.Context) {
		handler.CartGet(c)
	})
	r.E.POST("/order/add", func(c *gin.Context) {
		handler.OrderAdd(c)
	})
	r.E.GET("/order/get/email", func(c *gin.Context) {
		handler.OrderGetByEmail(c)
	})
}

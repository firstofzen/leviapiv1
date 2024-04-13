package handler

import (
	"github.com/gin-gonic/gin"
	"levi-apis/apis/request"
	"levi-apis/apis/response"
	"levi-apis/service"
	"levi-apis/utils"
	"levi-apis/utils/validate"
)

type ProductHandler struct{}

func (h ProductHandler) AddProduct(c *gin.Context) {
	var pReq request.ReqAddProduct
	var reqVal validate.RequestValidate
	var jwt utils.JwtToken
	failres := response.FailResponse{C: c}
	if err := c.ShouldBindJSON(&pReq); err != nil {
		failres.CannotBindBodyReq(err.Error())
	} else {
		if err := reqVal.Default(pReq); err != nil {
			failres.CannotValidateBodyReq(err.Error())
		} else {
			if isExpire, err := jwt.TokenIsExpire(pReq.AccessToken); err != nil {
				failres.InvalidToken()
			} else {
				if isExpire {
					failres.TokenIsExpire()
				} else {
					if claims, err := jwt.ParseToken(pReq.AccessToken); err != nil {
						failres.CannotParseToken(err.Error())
					} else {
						var pService service.ProductService
						if err := pService.AddProduct(pReq.Category2, pReq.Category1, pReq.Description, pReq.Sale, pReq.Price, pReq.Title, pReq.AvatarUrl, pReq.PicturesUrl, claims["email"].(string)); err != nil {
							failres.ErrorPersist(err.Error())
						} else {
							c.String(200, "oke")
						}
					}
				}
			}
		}
	}
}

func (h ProductHandler) GetProductByID(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	c2 := c.Query("category_2")
	id := c.Query("id")
	if id != "" && c2 != "" {
		var pService service.ProductService
		if rs, err := pService.GetProductById(c2, id); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.JSON(200, rs)
		}
	} else {
		failRes.FieldGivenEmpty("category_2", "id")
	}
}

func (h ProductHandler) GetProductByCategory2(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	c2 := c.Query("category_2")
	if c2 != "" {
		var pService service.ProductService
		if rs, err := pService.GetProductByCategory2(c2); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.JSON(200, rs)
		}
	} else {
		failRes.FieldGivenEmpty("category_2")
	}
}

func (h ProductHandler) UpdateSaleProduct(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	var bodyReq request.ReqUpdateSaleProduct
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		var pService service.ProductService
		if err := pService.UpdateSaleProduct(bodyReq); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.String(200, "oke")
		}
	}
}

func (h ProductHandler) DeleteProduct(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	c2 := c.Query("category_2")
	id := c.Query("id")
	if c2 != "" && id != "" {
		var pService service.ProductService
		if err := pService.DeleteProduct(c2, id); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.String(200, "oke")
		}
	} else {
		failRes.FieldGivenEmpty("category_2", "id")
	}
}

func (h ProductHandler) GetProductRating(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	c2 := c.Query("category_2")
	id := c.Query("id")
	if c2 != "" && id != "" {
		var pService service.ProductService
		if rs, err := pService.GetRating(c2, id); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.JSON(200, rs)
		}
	} else {
		failRes.FieldGivenEmpty("category_2", "id")
	}
}

func (h ProductHandler) AddRatingUser(c *gin.Context) {
	var bodyReq request.ReqAddRatingUser
	failRes := response.FailResponse{C: c}
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		var pService service.ProductService
		if err := pService.AddRatingUser(bodyReq); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.String(200, "oke")
		}
	}
}

func (h ProductHandler) RemoveRatingUser(c *gin.Context) {
	var bodyReq request.ReqAddRatingUser
	failRes := response.FailResponse{C: c}
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		var pService service.ProductService
		if err := pService.RemoveRatingUser(bodyReq); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.String(200, "oke")
		}
	}
}

func (h ProductHandler) GetProductRanking(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	c2 := c.Query("category_2")
	if c2 != "" {
		var pService service.ProductService
		if rs, err := pService.GetProductRanking(c2); err != nil {
			failRes.ErrorPersist(err.Error())
		} else {
			c.JSON(200, rs)
		}
	} else {
		failRes.FieldGivenEmpty("category_2")
	}
}

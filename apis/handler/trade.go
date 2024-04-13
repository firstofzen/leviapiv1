package handler

import (
	"github.com/gin-gonic/gin"
	"levi-apis/apis/request"
	"levi-apis/apis/response"
	"levi-apis/service"
	"levi-apis/utils"
	"net/http"
)

type TradeHandler struct{}

func (h TradeHandler) BoothGet(c *gin.Context) {
	var jwt utils.JwtToken
	var failRes response.FailResponse
	token := c.Query("access_token")
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.Login(err.Error())
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var tService service.TradeService
				if rs, err := tService.BoothGet(claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.JSON(200, rs)
				}
			}
		}
	}

}

func (h TradeHandler) BoothUpdateDescription(c *gin.Context) {
	var jwt utils.JwtToken
	failRes := response.FailResponse{C: c}
	var bodyReq request.ReqBoothUpdateDescription
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		if isExpire, err := jwt.TokenIsExpire(bodyReq.AccessToken); err != nil {
			failRes.Login(err.Error())
		} else {
			if isExpire {
				failRes.TokenIsExpire()
			} else {
				if claims, err := jwt.ParseToken(bodyReq.AccessToken); err != nil {
					failRes.CannotParseToken(err.Error())
				} else {
					var tService service.TradeService
					if err := tService.BoothUpdateDescription(claims["email"].(string), bodyReq.Description); err != nil {
						failRes.ErrorPersist(err.Error())
					} else {
						c.JSON(200, "oke")
					}
				}
			}
		}
	}
}

func (h TradeHandler) OrderAdd(c *gin.Context) {
	failRes := response.FailResponse{C: c}
	var bodyReq request.ReqOrderAdd
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		var jwt utils.JwtToken
		if isExpire, err := jwt.TokenIsExpire(bodyReq.AccessToken); err != nil {
			c.String(http.StatusBadRequest, "Invalid token")
		} else {
			if isExpire {
				failRes.TokenIsExpire()
			} else {
				if claims, err := jwt.ParseToken(bodyReq.AccessToken); err != nil {
					failRes.CannotParseToken(err.Error())
				} else {
					var tService service.TradeService
					if err := tService.OrderAdd(claims["email"].(string), bodyReq.TotalAmount, bodyReq.ProductCategory2, bodyReq.ProductId); err != nil {
						failRes.ErrorPersist(err.Error())
					} else {
						c.JSON(200, "oke")
					}
				}
			}
		}
	}
}

func (h TradeHandler) OrderGetByEmail(c *gin.Context) {
	var jwt utils.JwtToken
	failRes := response.FailResponse{C: c}
	tok := c.Query("access_token")
	if isExpire, err := jwt.TokenIsExpire(tok); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(tok); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var tService service.TradeService
				if rs, err := tService.OrderGetByEmail(claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.JSON(200, rs)
				}
			}
		}
	}

}

func (h TradeHandler) CartGet(c *gin.Context) {
	var jwt utils.JwtToken
	failRes := response.FailResponse{C: c}
	tok := c.Query("access_token")
	if isExpire, err := jwt.TokenIsExpire(tok); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(tok); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var tService service.TradeService
				if rs, err := tService.CartGet(claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.JSON(200, rs)
				}
			}
		}
	}
}

func (h TradeHandler) BoothUpdateBrand(c *gin.Context) {
	var jwt utils.JwtToken
	failRes := response.FailResponse{C: c}
	var bodyReq request.ReqBoothUpdateBrand
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		failRes.CannotBindBodyReq(err.Error())
	} else {
		if isExpire, err := jwt.TokenIsExpire(bodyReq.AccessToken); err != nil {
			failRes.Login(err.Error())
		} else {
			if isExpire {
				failRes.TokenIsExpire()
			} else {
				if claims, err := jwt.ParseToken(bodyReq.AccessToken); err != nil {
					failRes.CannotParseToken(err.Error())
				} else {
					var tService service.TradeService
					if err := tService.BoothUpdateBrand(claims["email"].(string), bodyReq.Brand); err != nil {
						failRes.ErrorPersist(err.Error())
					} else {
						c.JSON(200, "oke")
					}
				}
			}
		}
	}
}

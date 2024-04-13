package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"io"
	"levi-apis/apis/response"
	config2 "levi-apis/config"
	"levi-apis/service"
	"levi-apis/utils"
	"os"
)

type UserHandler struct{}

func (h UserHandler) Login(c *gin.Context) {
	config := config2.Oauth2{}.Default()
	url := config.AuthCodeURL("state-oke")
	c.Redirect(302, url)
}

func (h UserHandler) LoginGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	config := config2.Oauth2{}.Default()
	client := req.R()
	failRes := response.FailResponse{C: c}
	if tok, err := config.Exchange(context.Background(), code); err != nil {
		failRes.Login(err.Error())
	} else {
		if resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken); err != nil {
			failRes.Login(err.Error())
		} else {
			var m map[string]interface{}
			if body, err := io.ReadAll(resp.Body); err != nil {
				failRes.Login(err.Error())
			} else {
				if err := json.Unmarshal(body, &m); err != nil {
					failRes.Login(err.Error())
				} else {
					var uService service.UserService
					if err := uService.AddAccountIfExists(m["email"].(string), m["name"].(string), m["picture"].(string)); err != nil {
						failRes.Login(err.Error())
					} else {
						var jwt utils.JwtToken
						if at, err := jwt.CreateAccessToken(m["email"].(string)); err != nil {
							failRes.Login(err.Error())
						} else {
							c.Redirect(302, os.Getenv("CLIENT_URL")+"?access_token="+at)
						}
					}
				}
			}
		}
	}
}

func (h UserHandler) LoginAccessToken(c *gin.Context) {
	token := c.Query("access_token")
	failRes := response.FailResponse{C: c}
	var jwt utils.JwtToken
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var uService service.UserService
				if acc, err := uService.GetAccountByEmail(claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.JSON(200, acc)
				}
			}
		}
	}
}

func (h UserHandler) UpdateUserAccountName(c *gin.Context) {
	var jwt utils.JwtToken
	token := c.Query("access_token")
	name := c.Query("name")
	failRes := response.FailResponse{C: c}
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if claims, err := jwt.ParseToken(token); err != nil {
			failRes.CannotParseToken(err.Error())
		} else {
			var uService service.UserService
			if isExpire {
				failRes.TokenIsExpire()
			} else {
				if err := uService.UpdateUserAccountName(name, claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.String(200, "oke")
				}
			}
		}
	}
}

func (h UserHandler) UpdateUserAccountAvatarURL(c *gin.Context) {
	var jwt utils.JwtToken
	token := c.Query("access_token")
	avatar := c.Query("avatar_url")
	failRes := response.FailResponse{C: c}
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if claims, err := jwt.ParseToken(token); err != nil {
			failRes.CannotParseToken(err.Error())
		} else {
			var uService service.UserService
			if isExpire {
				failRes.TokenIsExpire()
			} else {
				if err := uService.UpdateUserAccountAvatarURL(avatar, claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.String(200, "oke")
				}
			}
		}
	}
}

func (h UserHandler) DeleteUserAccount(c *gin.Context) {
	var jwt utils.JwtToken
	token := c.Query("access_token")
	failRes := response.FailResponse{C: c}
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var uService service.UserService
				if err := uService.DeleteUserAccount(claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.String(200, "oke")
				}
			}
		}
	}
}

func (h UserHandler) UpdateUserDetailPhone(c *gin.Context) {
	var jwt utils.JwtToken
	token := c.Query("access_token")
	phone := c.Query("phone")
	failRes := response.FailResponse{C: c}
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var uService service.UserService
				if err := uService.UpdateUserDetailPhone(phone, claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.String(200, "oke")
				}
			}
		}
	}
}

func (h UserHandler) UpdateUserDetailAddress(c *gin.Context) {
	var jwt utils.JwtToken
	token := c.Query("access_token")
	address := c.Query("address")
	failRes := response.FailResponse{C: c}
	if isExpire, err := jwt.TokenIsExpire(token); err != nil {
		failRes.InvalidToken()
	} else {
		if isExpire {
			failRes.TokenIsExpire()
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				failRes.CannotParseToken(err.Error())
			} else {
				var uService service.UserService
				if err := uService.UpdateUserDetailAddress(address, claims["email"].(string)); err != nil {
					failRes.ErrorPersist(err.Error())
				} else {
					c.String(200, "oke")
				}
			}
		}
	}
}

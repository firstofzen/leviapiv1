package router

import (
	"github.com/gin-gonic/gin"
	handler2 "levi-apis/apis/handler"
)

type UserRouter struct {
	E *gin.Engine
}

func (u *UserRouter) Route() {
	var handler handler2.UserHandler
	u.E.GET("/login", func(c *gin.Context) {
		handler.Login(c)
	})
	u.E.GET("/login/callback/google", func(c *gin.Context) {
		handler.LoginGoogleCallback(c)
	})
	u.E.GET("/login/access_token", func(c *gin.Context) {
		handler.LoginAccessToken(c)
	})
	u.E.PATCH("/user/account/update/name", func(c *gin.Context) {
		handler.UpdateUserAccountName(c)
	})
	u.E.PATCH("/user/account/update/avatar_url", func(c *gin.Context) {
		handler.UpdateUserAccountAvatarURL(c)
	})
	u.E.DELETE("/user/account/delete", func(c *gin.Context) {
		handler.DeleteUserAccount(c)
	})
	u.E.PATCH("/user/detail/update/phone", func(c *gin.Context) {
		handler.UpdateUserDetailPhone(c)
	})
	u.E.PATCH("/user/detail/update/address", func(c *gin.Context) {
		handler.UpdateUserDetailAddress(c)
	})
}

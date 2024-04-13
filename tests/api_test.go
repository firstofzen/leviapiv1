package tests

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
	"testing"
)

func TestApi(t *testing.T) {
	r := gin.Default()
	conf := &oauth2.Config{ClientID: "1094721977853-r6ve3sn0qm6lr5kpoq6t2n2i2llnonvs.apps.googleusercontent.com", ClientSecret: "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", RedirectURL: "http://localhost:1111/login/success", Endpoint: google.Endpoint, Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"}}
	r.GET("/login", func(c *gin.Context) {
		url := conf.AuthCodeURL("state")
		c.Redirect(http.StatusSeeOther, url)
	})
	r.GET("/login/success", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		if state != "state" {
			c.JSON(http.StatusInternalServerError, "state dont match")
		}

		if tok, err := conf.Exchange(context.Background(), code); err != nil {
			c.String(http.StatusInternalServerError, "can not exchange code")
		} else {
			if resp, err := req.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken); err != nil {
				c.String(http.StatusInternalServerError, "cant get resource")
			} else {
				if data, err := io.ReadAll(resp.Body); err != nil {
					c.String(http.StatusInternalServerError, "cant read body")
				} else {
					var m map[string]interface{}
					if err := json.Unmarshal(data, &m); err != nil {
						c.String(http.StatusInternalServerError, "cant unmarshal")
					}
					c.JSON(http.StatusOK, m)
				}
			}

		}
	})
	r.Run(":1111")
}

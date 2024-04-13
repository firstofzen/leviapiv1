package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type Oauth2 struct {
}

func (o Oauth2) Default() *oauth2.Config {
	clientId := os.Getenv("GG_CLIENT_ID")
	clientSecret := os.Getenv("GG_CLIENT_SECRET")
	redirectUrl := os.Getenv("GG_REDIRECT_URL")
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectUrl,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
	}
}

package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Config struct {
	GoogleLoginConfig   oauth2.Config
	FacebookLoginConfig oauth2.Config
}

var AppConfig Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const OauthFacebookUrlAPI = "https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token="

func LoadConfig(id string, secret string, redirect_url string) {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirect_url,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	// Oauth configuration for Facebook
	AppConfig.FacebookLoginConfig = oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirect_url,
		Endpoint:     facebook.Endpoint,
		Scopes: []string{
			"email",
			"public_profile",
		},
	}
}

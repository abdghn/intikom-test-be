package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig() oauth2.Config {

	config := oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     "",
		ClientSecret: "",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return config
}

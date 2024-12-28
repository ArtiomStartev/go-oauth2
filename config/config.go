package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"os"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
	GitHubLoginConfig oauth2.Config
}

var AppConfig Config

func GoogleConfig() (oauth2.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file: ", err)
		return oauth2.Config{}, err
	}

	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/google_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig, nil
}

func GitHubConfig() (oauth2.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file: ", err)
		return oauth2.Config{}, err
	}

	AppConfig.GitHubLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/github_callback",
		Scopes:       []string{"user", "repo"},
		Endpoint:     github.Endpoint,
	}

	return AppConfig.GitHubLoginConfig, nil
}

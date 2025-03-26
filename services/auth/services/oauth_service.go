package services

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

var (
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     oauth2.Endpoint{AuthURL: "https://accounts.google.com/o/oauth2/auth", TokenURL: "https://oauth2.googleapis.com/token"},
	}

	GitHubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
)

func GetOAuthURL(provider, state string) string {
	switch provider {
	case "google":
		return GoogleOAuthConfig.AuthCodeURL(state)
	case "github":
		return GitHubOAuthConfig.AuthCodeURL(state)
	default:
		return ""
	}
}

func ExchangeOAuthCode(provider, code string) (*oauth2.Token, error) {
	switch provider {
	case "google":
		return GoogleOAuthConfig.Exchange(context.Background(), code)
	case "github":
		return GitHubOAuthConfig.Exchange(context.Background(), code)
	default:
		return nil, nil
	}
}

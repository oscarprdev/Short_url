package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuthProvider() {
	// key := os.Getenv("AUTH_KEY")

	// Google
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv(("GOOGLE_CLIENT_SECRET"))
	googleCallbackUrl := "http://localhost:8080/auth/callback?provider=google"

	// Github
	githubClientId := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv(("GITHUB_CLIENT_SECRET"))
	githubCallbackUrl := "http://localhost:8080/auth/callback?provider=github"

	store := sessions.NewCookieStore([]byte("my secure key"))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, googleCallbackUrl),
		github.New(githubClientId, githubClientSecret, githubCallbackUrl),
	)
}

package authenticator

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config

	store *session.Store
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) verifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func NewAuthenticator() (*Authenticator, error) {
	domain := viper.GetString("auth0.domain")
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+domain+"/",
	)
	if err != nil {
		return nil, err
	}

	clientID := viper.GetString("auth0.client_id")
	if clientID == "" {
		return nil, errors.New("client_id is required")
	}

	clientSecret := viper.GetString("auth0.client_secret")
	if clientSecret == "" {
		return nil, errors.New("client_secret is required")
	}

	redirectURI := viper.GetString("auth0.redirect_uri")
	if redirectURI == "" {
		return nil, errors.New("redirect_uri is required")
	}

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURI,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	store := session.New()

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		store:    store,
	}, nil
}

// Define middelware
func (a *Authenticator) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Verify the JWT
		token, err := a.verifyIDToken(c.Context(), c.Locals("token").(*oauth2.Token))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
		}

		// Set the claims in the context
		c.Locals("claims", token.Claims)

		return c.Next()
	}
}

// Handler for redirecting to the login page.
func (a *Authenticator) Login(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		state, err := generateRandomState()
		if err != nil {
			return err
		}

		session, err := a.store.Get(c)
		if err != nil {
			return err
		}

		// Save the state inside the session.
		session.Set("state", state)
		if err := session.Save(); err != nil {
			return err
		}

		location := a.AuthCodeURL(state)

		return c.Redirect(location, fiber.StatusTemporaryRedirect)
	}
}

func (a *Authenticator) Callback(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := a.store.Get(c)
		if err != nil {
			return err
		}

		if c.Query("state") != session.Get("state") {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid state parameter.")
		}

		// Exchange an authorization code for a token.
		token, err := a.Exchange(c.Context(), c.Query("code"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Failed to convert an authorization code into a token.")
		}

		idToken, err := a.verifyIDToken(c.Context(), token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to verify ID Token.")
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err := session.Save(); err != nil {
			return err
		}

		return c.Redirect("/api/user", fiber.StatusTemporaryRedirect)
	}
}

func (a *Authenticator) Logout(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		domain := viper.GetString("auth0.domain")

		logoutUrl := "https://" + domain + "/v2/logout"

		scheme := "http"
		if c.Protocol() == "https" {
			scheme = "https"
		}

		returnTo := scheme + "://" + c.Hostname()
		client_id := viper.GetString("auth0.client_id")
		parameters := "returnTo=" + returnTo + "&client_id=" + client_id

		return c.Redirect(logoutUrl + "?" + parameters)
	}
}

func (a *Authenticator) GetUser(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := a.store.Get(c)
		if err != nil {
			return err
		}

		profile := session.Get("profile")

		return c.JSON(profile)
	}
}

// // IsAuthenticated is a middleware that checks if
// // the user has already been authenticated previously.
//
//	func IsAuthenticated(ctx *gin.Context) {
//		if sessions.Default(ctx).Get("profile") == nil {
//			ctx.Redirect(http.StatusSeeOther, "/")
//		} else {
//			ctx.Next()
//		}
//	}
func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

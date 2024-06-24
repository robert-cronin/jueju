// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authenticator

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

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

func NewAuthenticator(store *session.Store) (*Authenticator, error) {
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

	callbackURI := viper.GetString("auth0.callback_url")
	if callbackURI == "" {
		return nil, errors.New("callback_url is required")
	}

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURI,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		store:    store,
	}, nil
}

// AuthRequired is a middleware that checks if the user is authenticated
func (a *Authenticator) AuthRequired(c *fiber.Ctx) error {
	session, err := a.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	if session.Get("profile") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}

// Login initiates the authentication process
func (a *Authenticator) Login(c *fiber.Ctx) error {
	state, err := generateRandomState()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate state"})
	}

	session, err := a.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	session.Set("state", state)
	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session"})
	}

	return c.Redirect(a.AuthCodeURL(state), fiber.StatusTemporaryRedirect)
}

func (a *Authenticator) deleteSession(c *fiber.Ctx) error {
	session, err := a.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	if err := session.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to destroy session"})
	}

	return nil
}

// Callback handles the callback from Auth0
func (a *Authenticator) Callback(c *fiber.Ctx) error {
	session, err := a.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	fmt.Println("state:", session.Get("state"), "query:", c.Query("state"))

	if c.Query("state") != session.Get("state") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state parameter"})
	}

	token, err := a.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	idToken, err := a.verifyIDToken(c.Context(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify ID token"})
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session"})
	}

	// Get the redirect url
	redirectURL := viper.GetString("auth0.redirect_url")
	if redirectURL == "" {
		redirectURL = "/"
	}

	return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
}

// Logout logs the user out
func (a *Authenticator) Logout(c *fiber.Ctx) error {
	domain := viper.GetString("auth0.domain")
	clientID := viper.GetString("auth0.client_id")
	returnTo := viper.GetString("auth0.redirect_url")

	logoutURL := url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   "/v2/logout",
		RawQuery: url.Values{
			"client_id": {clientID},
			"returnTo":  {returnTo},
		}.Encode(),
	}

	if err := a.deleteSession(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete session"})
	}

	return c.Redirect(logoutURL.String(), fiber.StatusTemporaryRedirect)
}

// GetUser returns the user profile
func (a *Authenticator) GetUser(c *fiber.Ctx) error {
	session, err := a.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	profile := session.Get("profile")
	if profile == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
	}

	return c.JSON(profile)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

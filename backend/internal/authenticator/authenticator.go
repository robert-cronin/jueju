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
	"net/url"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/robert-cronin/jueju/backend/internal/database"
	"github.com/robert-cronin/jueju/backend/internal/models"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config

	Store *session.Store
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
		Store:    store,
	}, nil
}

// AuthRequired is a middleware that checks if the user is authenticated
func (a *Authenticator) AuthRequired(c *fiber.Ctx) error {
	session, err := a.Store.Get(c)
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

	session, err := a.Store.Get(c)
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
	session, err := a.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	if err := session.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to destroy session"})
	}

	return nil
}

func (a *Authenticator) GetOrCreateUser(profile map[string]interface{}) (*models.User, error) {
	var user models.User

	auth0ID, ok := profile["sub"].(string)
	if !ok {
		return nil, errors.New("invalid Auth0 ID")
	}

	// Try to find the user by Auth0 ID
	result := database.DB.Where("auth0_id = ?", auth0ID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// User doesn't exist, create a new one
			user = models.User{
				Auth0ID:       auth0ID,
				Email:         profile["email"].(string),
				EmailVerified: profile["email_verified"].(bool),
				Name:          profile["name"].(string),
				Nickname:      profile["nickname"].(string),
				Picture:       profile["picture"].(string),
				LastLogin:     time.Now(),
			}
			if err := database.DB.Create(&user).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, result.Error
		}
	} else {
		// User exists, update last login
		user.LastLogin = time.Now()
		if err := database.DB.Save(&user).Error; err != nil {
			return nil, err
		}
	}

	return &user, nil
}

// Callback handles the callback from Auth0
func (a *Authenticator) Callback(c *fiber.Ctx) error {
	session, err := a.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	if session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session is nil"})
	}

	stateFromSession := session.Get("state")
	if stateFromSession == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "State not found in session"})
	}

	if c.Query("state") != stateFromSession.(string) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state parameter"})
	}

	token, err := a.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to exchange token: " + err.Error()})
	}

	idToken, err := a.verifyIDToken(c.Context(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify ID token: " + err.Error()})
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info: " + err.Error()})
	}

	user, err := a.GetOrCreateUser(profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create or update user: " + err.Error()})
	}

	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User is nil after GetOrCreateUser"})
	}

	session.Set("user_id", user.ID.String())
	session.Set("auth0_id", user.Auth0ID)
	session.Set("access_token", token.AccessToken)

	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session: " + err.Error()})
	}

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

// GetUser returns the user data
func (a *Authenticator) GetUser(c *fiber.Ctx) error {
	session, err := a.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
	}

	userID := session.Get("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	return c.JSON(user)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// AuthMiddleware is a middleware that checks if the user is authenticated and sets the user ID in the context
func (a *Authenticator) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := a.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get session"})
		}

		userID := session.Get("user_id")
		if userID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
		}

		// Parse the user ID from string to uuid.UUID
		parsedUserID, err := uuid.Parse(userID.(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user ID format"})
		}

		// Set the user ID in the context
		c.Locals("userID", parsedUserID)

		return c.Next()
	}
}

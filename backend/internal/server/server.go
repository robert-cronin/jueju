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

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/robert-cronin/jueju/backend/internal/api"
	"github.com/robert-cronin/jueju/backend/internal/authenticator"
	"github.com/robert-cronin/jueju/backend/internal/redis"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	auth  *authenticator.Authenticator
	store *session.Store
}

// Callback implements api.ServerInterface.
func (s *Server) Callback(c *fiber.Ctx) error {
	return s.auth.Callback(c)
}

// Logout implements api.ServerInterface.
func (s *Server) Logout(c *fiber.Ctx) error {
	return s.auth.Logout(c)
}

// GetUser implements ServerInterface.
func (s *Server) GetUser(c *fiber.Ctx) error {
	return s.auth.GetUser(c)
}

// Login implements ServerInterface.
func (s *Server) Login(c *fiber.Ctx) error {
	return s.auth.Login(c)
}

// AuthMiddleware checks if the user is authenticated
func (s *Server) AuthMiddleware() fiber.Handler {
	return s.auth.AuthRequired
}

// NewServer creates a new server
func NewServer() (*Server, error) {
	// Create a new session store
	fiberStorage := redis.NewFiberClient("fiber_session_")
	store := session.New(
		session.Config{
			Storage:        fiberStorage,
			Expiration:     24 * 60 * 60 * 1000, // 24 hours
			CookieSameSite: "Lax",
			CookieSecure:   false,
		},
	)

	auth, err := authenticator.NewAuthenticator(store)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		auth:  auth,
		store: store,
	}

	return srv, nil
}

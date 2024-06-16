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
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	auth *authenticator.Authenticator
	store *session.Store
}

// Callback implements api.ServerInterface.
func (s *Server) Callback(c *fiber.Ctx) error {
	handler := s.auth.Callback(c)
	return handler(c)
}

// Logout implements api.ServerInterface.
func (s *Server) Logout(c *fiber.Ctx) error {
	handler := s.auth.Logout(c)
	return handler(c)
}

// GetUser implements ServerInterface.
func (s Server) GetUser(c *fiber.Ctx) error {
	// TODO: implement GetUser
	// for now we'll just return a mock response
	return c.JSON(map[string]string{
		"message": "Hello, World!",
	})
}

// Login implements ServerInterface.
func (s Server) Login(c *fiber.Ctx) error {
	handler := s.auth.Login(c)
	return handler(c)
}

func NewServer() *Server {
	auth, err := authenticator.NewAuthenticator()
	if err != nil {
		panic(err)
	}

	store := session.New()

	srv := &Server{
		auth: auth,
		store: store,
	}

	return srv
}

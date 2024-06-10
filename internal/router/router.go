package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/robert-cronin/jueju-backend/internal/app/callback"
	"github.com/robert-cronin/jueju-backend/internal/app/login"
	"github.com/robert-cronin/jueju-backend/internal/app/logout"
	"github.com/robert-cronin/jueju-backend/internal/app/user"

	"github.com/robert-cronin/jueju-backend/internal/authenticator"
	"github.com/robert-cronin/jueju-backend/internal/middleware"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}

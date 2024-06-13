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

	api := router.Group("/api")
	{
		// To store custom types in our cookies,
		// we must first register them using gob.Register
		gob.Register(map[string]interface{}{})

		store := cookie.NewStore([]byte("secret"))
		api.Use(sessions.Sessions("auth-session", store))

		api.GET("/login", login.Handler(auth))
		api.GET("/callback", callback.Handler(auth))
		api.GET("/user", middleware.IsAuthenticated, user.Handler)
		api.GET("/logout", logout.Handler)
	}

	return router
}

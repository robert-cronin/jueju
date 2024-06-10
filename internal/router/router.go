package router

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"time"

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

	// Custom Logger Middleware
	router.Use(func(c *gin.Context) {
		// Request details
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		headers := c.Request.Header
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Read body for logging (and reset reader for Gin)
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		latency := time.Since(start)
		status := c.Writer.Status()
		responseHeaders := c.Writer.Header()

		// Log format
		log.Printf("REQUEST: method=%s path=%s query=%s headers=%v body=%s clientIP=%s userAgent=%s",
			method, path, query, headers, string(bodyBytes), clientIP, userAgent)
		log.Printf("RESPONSE: status=%d headers=%v latency=%s",
			status, responseHeaders, latency)
	})

	return router
}

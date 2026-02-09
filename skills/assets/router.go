// internal/shared/server/http.go
package server

import "github.com/gin-gonic/gin"

func NewRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()

	// Global middleware
	r.Use(gin.Recovery())
	// Add logging, tracing, CORS via middleware — see go-observability skill

	return r
}

// In cmd/server/main.go:
func setupRoutes(r *gin.Engine, bookingHandler *handler.BookingHandler) {
	// Health check — no version prefix
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1
	v1 := r.Group("/api/v1")
	{
		bookingHandler.RegisterRoutes(v1)
	}
}

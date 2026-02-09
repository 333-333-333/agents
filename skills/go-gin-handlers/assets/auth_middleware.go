// internal/shared/middleware/auth.go
package middleware

import (
	"api/booking/internal/shared/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthRequired(tokenValidator TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			server.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "Missing or invalid authorization header")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := tokenValidator.Validate(c.Request.Context(), token)
		if err != nil {
			server.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

// TokenValidator is a port — implementation is in infrastructure
type TokenValidator interface {
	Validate(ctx context.Context, token string) (*Claims, error)
}

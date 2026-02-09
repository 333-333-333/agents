// cmd/server/main.go
import (
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "api/booking/docs" // Generated docs
)

func setupRoutes(r *gin.Engine) {
    // Swagger UI — only in non-production
    if cfg.Env != "production" {
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }

    // API routes...
}

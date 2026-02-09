// internal/shared/server/errors.go
package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Domain error sentinel types — defined in each domain's error.go
type NotFoundError interface{ NotFound() }
type ConflictError interface{ Conflict() }
type ForbiddenError interface{ Forbidden() }
type ValidationError interface{ Validation() }

func HandleDomainError(c *gin.Context, err error) {
	var (
		notFound   NotFoundError
		conflict   ConflictError
		forbidden  ForbiddenError
		validation ValidationError
	)

	switch {
	case errors.As(err, &notFound):
		Fail(c, http.StatusNotFound, "NOT_FOUND", err.Error())
	case errors.As(err, &conflict):
		Fail(c, http.StatusConflict, "CONFLICT", err.Error())
	case errors.As(err, &forbidden):
		Fail(c, http.StatusForbidden, "FORBIDDEN", err.Error())
	case errors.As(err, &validation):
		Fail(c, http.StatusUnprocessableEntity, "DOMAIN_VALIDATION", err.Error())
	default:
		Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred")
	}
}

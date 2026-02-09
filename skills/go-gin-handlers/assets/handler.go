// internal/booking/infrastructure/handler/http.go
package handler

import (
	"api/booking/internal/booking/application"
	"api/booking/internal/shared/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingHandler struct {
	service *application.BookingService
}

func NewBookingHandler(service *application.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

// RegisterRoutes mounts all routes for this domain.
func (h *BookingHandler) RegisterRoutes(rg *gin.RouterGroup) {
	bookings := rg.Group("/bookings")
	{
		bookings.POST("", h.Create)
		bookings.GET("/:id", h.GetByID)
		bookings.GET("", h.ListByOwner)
		bookings.PATCH("/:id/cancel", h.Cancel)
	}
}

type CreateBookingRequest struct {
	CaregiverID string   `json:"caregiver_id" binding:"required,uuid"`
	ServiceType string   `json:"service_type" binding:"required,oneof=walk hosting visit specialized"`
	StartAt     string   `json:"start_at" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndAt       string   `json:"end_at" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	PetIDs      []string `json:"pet_ids" binding:"required,min=1,dive,uuid"`
}

func (h *BookingHandler) Create(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		server.FailWithDetails(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid request body", err.Error())
		return
	}

	// Extract authenticated user from context (set by auth middleware)
	ownerID := c.GetString("user_id")

	output, err := h.service.CreateBooking(c.Request.Context(), application.CreateBookingInput{
		OwnerID:     ownerID,
		CaregiverID: req.CaregiverID,
		ServiceType: req.ServiceType,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
		PetIDs:      req.PetIDs,
	})
	if err != nil {
		handleDomainError(c, err)
		return
	}

	server.OK(c, http.StatusCreated, output)
}

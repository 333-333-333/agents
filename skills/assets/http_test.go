// internal/booking/infrastructure/handler/http_test.go
package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/booking/internal/booking/infrastructure/handler"
	"github.com/gin-gonic/gin"
)

func TestBookingHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	svc := &MockBookingService{} // mock application service
	h := handler.NewBookingHandler(svc)

	router := gin.New()
	v1 := router.Group("/api/v1")
	h.RegisterRoutes(v1)

	body, _ := json.Marshal(map[string]any{
		"caregiver_id": "caregiver-1",
		"service_type": "walk",
		"start_at":     "2026-02-02T10:00:00Z",
		"end_at":       "2026-02-02T11:00:00Z",
		"pet_ids":      []string{"pet-1"},
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/bookings", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rec.Code)
	}
}

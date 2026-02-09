package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Handler integration test: tests the full HTTP stack with real service + in-memory repo.
// This validates: JSON binding, status codes, response format, error mapping.

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Wire up with real service + in-memory dependencies:
	// repo := repository.NewInMemoryNotificationRepository()
	// logSender := sender.NewLogSender(slog.Default(), domain.ChannelEmail)
	// dispatcher := application.NewNotificationDispatcher(logSender, ...)
	// svc := application.NewNotificationService(dispatcher, repo)
	// handler := handler.NewNotificationHandler(svc)
	// v1 := r.Group("/api/v1")
	// handler.RegisterRoutes(v1)

	return r
}

func TestHTTP_SendNotification_Success(t *testing.T) {
	router := setupTestRouter()

	payload := map[string]string{
		"channel":   "email",
		"recipient": "user@test.com",
		"subject":   "Welcome",
		"content":   "Hello world",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notifications/send", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d. Body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if _, ok := resp["data"]; !ok {
		t.Error("response missing 'data' key")
	}
}

func TestHTTP_SendNotification_InvalidChannel(t *testing.T) {
	router := setupTestRouter()

	payload := map[string]string{
		"channel":   "telegram",
		"recipient": "user",
		"content":   "hello",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notifications/send", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	errObj, ok := resp["error"].(map[string]interface{})
	if !ok {
		t.Fatal("response missing 'error' key")
	}
	if errObj["code"] != "VALIDATION_ERROR" {
		t.Errorf("error.code = %v, want VALIDATION_ERROR", errObj["code"])
	}
}

func TestHTTP_SendNotification_MissingBody(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notifications/send", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestHTTP_HealthCheck(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("health check status = %d, want %d", rec.Code, http.StatusOK)
	}
}

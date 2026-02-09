package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Producer contract test: verifies that the API response SHAPE matches the contract.
// This test does NOT validate business logic — only the JSON structure.

func TestContract_SendNotification_SuccessResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Set up router with real handler + in-memory dependencies
	// router := setupTestRouter()

	// Make request
	body := `{"channel":"email","recipient":"user@test.com","subject":"Test","content":"Hello"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/notifications/send", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// router.ServeHTTP(rec, req)

	// Contract: status must be 200
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	// Contract: response body must have "data" key with required fields
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatal("response missing 'data' key")
	}

	// Verify required fields exist in contract
	requiredFields := []string{"id", "channel", "recipient", "status", "created_at"}
	for _, field := range requiredFields {
		if _, exists := data[field]; !exists {
			t.Errorf("response.data missing required field: %s", field)
		}
	}

	// Verify field types
	if _, ok := data["id"].(string); !ok {
		t.Error("response.data.id must be a string")
	}
	if _, ok := data["channel"].(string); !ok {
		t.Error("response.data.channel must be a string")
	}
	if status, ok := data["status"].(string); !ok || (status != "sent" && status != "pending") {
		t.Errorf("response.data.status must be 'sent' or 'pending', got %v", data["status"])
	}
}

func TestContract_SendNotification_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// router := setupTestRouter()

	// Missing required field
	body := `{"channel":"email","recipient":"user@test.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/notifications/send", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// router.ServeHTTP(rec, req)

	// Contract: validation errors return 400
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}

	// Contract: error response envelope
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}

	errObj, ok := response["error"].(map[string]interface{})
	if !ok {
		t.Fatal("response missing 'error' key")
	}

	if _, ok := errObj["code"].(string); !ok {
		t.Error("error.code must be a string")
	}
	if _, ok := errObj["message"].(string); !ok {
		t.Error("error.message must be a string")
	}
}

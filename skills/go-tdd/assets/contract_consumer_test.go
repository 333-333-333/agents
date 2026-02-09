package client_test

import (
	"encoding/json"
	"testing"
	"time"
)

// Consumer contract test: verifies that the CLIENT can correctly parse
// the response from the notification service.
// This test runs in the CONSUMER's codebase (e.g., booking service calling notifications).

// NotificationResponse is the consumer's model of what the notification service returns.
type NotificationResponse struct {
	Data *NotificationData `json:"data"`
}

type NotificationData struct {
	ID        string     `json:"id"`
	Channel   string     `json:"channel"`
	Recipient string     `json:"recipient"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	SentAt    *time.Time `json:"sent_at,omitempty"`
}

// ErrorResponse is the consumer's model of error responses.
type ErrorResponse struct {
	Error *ErrorData `json:"error"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func TestConsumerContract_ParseSuccessResponse(t *testing.T) {
	// Simulate what the notification service returns (copy from actual response)
	rawJSON := `{
		"data": {
			"id": "550e8400-e29b-41d4-a716-446655440000",
			"channel": "email",
			"recipient": "user@example.com",
			"status": "sent",
			"created_at": "2026-02-04T23:55:55.202381-03:00",
			"sent_at": "2026-02-04T23:55:55.202415-03:00"
		}
	}`

	var resp NotificationResponse
	if err := json.Unmarshal([]byte(rawJSON), &resp); err != nil {
		t.Fatalf("failed to parse notification response: %v", err)
	}

	if resp.Data == nil {
		t.Fatal("expected data to be present")
	}
	if resp.Data.ID == "" {
		t.Error("expected non-empty ID")
	}
	if resp.Data.Channel != "email" {
		t.Errorf("channel = %q, want %q", resp.Data.Channel, "email")
	}
	if resp.Data.Status != "sent" {
		t.Errorf("status = %q, want %q", resp.Data.Status, "sent")
	}
	if resp.Data.SentAt == nil {
		t.Error("expected sent_at to be present for sent notification")
	}
}

func TestConsumerContract_ParseErrorResponse(t *testing.T) {
	rawJSON := `{
		"error": {
			"code": "VALIDATION_ERROR",
			"message": "invalid channel: must be push, email, sms, or whatsapp"
		}
	}`

	var resp ErrorResponse
	if err := json.Unmarshal([]byte(rawJSON), &resp); err != nil {
		t.Fatalf("failed to parse error response: %v", err)
	}

	if resp.Error == nil {
		t.Fatal("expected error to be present")
	}
	if resp.Error.Code == "" {
		t.Error("expected non-empty error code")
	}
	if resp.Error.Message == "" {
		t.Error("expected non-empty error message")
	}
}

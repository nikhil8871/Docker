package routes

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheckEndpoints(t *testing.T) {
	app := fiber.New()
	RegisterRoute(app, nil)

	testCases := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "healthz check",
			path:       "/healthz",
			wantStatus: fiber.StatusOK,
			wantBody:   "healthy",
		},
		{
			name:       "health check",
			path:       "/health",
			wantStatus: fiber.StatusOK,
			wantBody:   "healthy",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.path, nil)
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("Failed to execute request for %s: %v", tc.path, err)
			}

			if resp.StatusCode != tc.wantStatus {
				t.Errorf("Expected status %d, got %d", tc.wantStatus, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			var result map[string]string
			if err := json.Unmarshal(body, &result); err != nil {
				t.Fatalf("Response is not valid JSON: %v", err)
			}

			if result["status"] != tc.wantBody {
				t.Errorf("Expected status '%s', got '%s'", tc.wantBody, result["status"])
			}
		})
	}
}

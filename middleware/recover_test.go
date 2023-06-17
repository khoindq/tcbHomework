package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
)

func TestRecover(t *testing.T) {
	// Create a new Gin router and use the Recover middleware.
	router := gin.New()
	router.Use(Recover())

	// Register a route that will panic.
	router.GET("/panic", func(c *gin.Context) {
		panic(errors.New("something went wrong"))
	})

	// Create a mock HTTP request for the "/panic" route.
	req, err := http.NewRequest("GET", "/panic", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response.
	rec := httptest.NewRecorder()

	// Serve the request using the router.
	router.ServeHTTP(rec, req)

	// Check the response status code.
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	// Check the response body.
	expectedBody := "{\"status_code\":500,\"message\":\"something went wrong in the server\",\"log\":\"something went wrong\",\"error_key\":\"ErrInternal\"}"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, rec.Body.String())
	}
}

// Mock common.AppError type and ErrInternal function for testing purposes.

func TestRecoverWithCustomError(t *testing.T) {
	// Create a new Gin router and use the Recover middleware.
	router := gin.New()
	router.Use(Recover())

	// Register a route that will panic with a custom error.
	router.GET("/panic-custom", func(c *gin.Context) {
		panic(&common.AppError{
			StatusCode: http.StatusBadRequest,
		})
	})

	// Create a mock HTTP request for the "/panic-custom" route.
	req, err := http.NewRequest("GET", "/panic-custom", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response.
	rec := httptest.NewRecorder()

	// Serve the request using the router.
	router.ServeHTTP(rec, req)

	// Check the response status code.
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}

	// Check the response body.
	expectedBody := "{\"status_code\":400,\"message\":\"\",\"log\":\"\",\"error_key\":\"\"}"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, rec.Body.String())
	}
}

package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestEmailSignIn_InvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/v1/users/signin", strings.NewReader("not json"))
	c.Request.Header.Set("Content-Type", "application/json")

	EmailSignIn(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Invalid input") {
		t.Errorf("expected 'Invalid input' in body, got %s", w.Body.String())
	}
}

func TestEmailSignIn_MissingFields(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := `{"email":"test@test.com"}`
	c.Request = httptest.NewRequest("POST", "/api/v1/users/signin", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	EmailSignIn(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestEmailSignIn_MissingEntClient(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := `{"email":"test@test.com","password":"password123"}`
	c.Request = httptest.NewRequest("POST", "/api/v1/users/signin", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	EmailSignIn(c)

	// Should get 500 because entClient is not in context
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "internal server error") {
		t.Errorf("expected 'internal server error' in body, got %s", w.Body.String())
	}
}

func TestEmailSignIn_ResponseDoesNotLeakPII(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := `{"email":"sensitive@email.com","password":"password123"}`
	c.Request = httptest.NewRequest("POST", "/api/v1/users/signin", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	EmailSignIn(c)

	// Regardless of error, the response should not contain the email
	if strings.Contains(w.Body.String(), "sensitive@email.com") {
		t.Error("response body should not contain user email (PII)")
	}
}

package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGenerateOAuthPassword_Success(t *testing.T) {
	pw, err := generateOAuthPassword()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 32 random bytes -> 64 hex chars
	if len(pw) != 64 {
		t.Errorf("expected 64-char hex string, got %d chars", len(pw))
	}
}

func TestGenerateOAuthPassword_Uniqueness(t *testing.T) {
	pw1, _ := generateOAuthPassword()
	pw2, _ := generateOAuthPassword()
	if pw1 == pw2 {
		t.Error("two consecutive passwords should not be identical")
	}
}

func TestGetFrontendURL_FromContext(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Set("frontendURL", "https://myapp.com")

	url := getFrontendURL(c)
	if url != "https://myapp.com" {
		t.Errorf("expected https://myapp.com, got %s", url)
	}
}

func TestGetFrontendURL_Default(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	url := getFrontendURL(c)
	if url != "http://localhost:3000" {
		t.Errorf("expected http://localhost:3000, got %s", url)
	}
}

func TestGetFrontendURL_WrongType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Set("frontendURL", 12345) // wrong type

	url := getFrontendURL(c)
	if url != "http://localhost:3000" {
		t.Errorf("expected fallback to default, got %s", url)
	}
}

func TestAuthCallback_MissingDB(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/auth/google/callback", nil)
	c.Params = gin.Params{{Key: "provider", Value: "google"}}

	// AuthCallback needs gothic to complete auth first, which will fail
	// without a proper OAuth session. This tests that the handler doesn't
	// panic even when called with minimal context.
	AuthCallback(c)

	// Should get an error response, not a panic
	if w.Code != http.StatusUnauthorized && w.Code != http.StatusInternalServerError {
		t.Errorf("expected 401 or 500, got %d", w.Code)
	}
}

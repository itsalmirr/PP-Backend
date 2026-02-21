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

func TestAuthCallback_AuthFailure(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/auth/google/callback", nil)
	c.Params = gin.Params{{Key: "provider", Value: "google"}}

	// Without a valid OAuth session, gothic.CompleteUserAuth fails and
	// AuthCallback returns 401 before reaching the db lookup.
	AuthCallback(c)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

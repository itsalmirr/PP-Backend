package api

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

func TestGetEntClient_Missing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	_, ok := getEntClient(c)
	if ok {
		t.Error("expected ok=false when entClient is missing")
	}
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "internal server error") {
		t.Errorf("expected 'internal server error' in body, got %s", w.Body.String())
	}
}

func TestGetImageService_Missing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	_, ok := getImageService(c)
	if ok {
		t.Error("expected ok=false when imageService is missing")
	}
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/v1/users/", strings.NewReader("not json"))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateUser(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Invalid input") {
		t.Errorf("expected 'Invalid input' in body, got %s", w.Body.String())
	}
}

func TestCreateUser_MissingEntClient(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := `{"email":"test@test.com","username":"testuser","full_name":"Test User","password":"password123"}`
	c.Request = httptest.NewRequest("POST", "/api/v1/users/", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateUser(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "internal server error") {
		t.Errorf("expected 'internal server error' in body, got %s", w.Body.String())
	}
}

func TestCreateRealtor_InvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/v1/realtors/", strings.NewReader("{invalid}"))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateRealtor(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestGetListings_InvalidQuery(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/v1/properties/buy?page_size=abc", nil)

	GetListings(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestDeleteListing_MissingID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("DELETE", "/api/v1/properties/", nil)

	DeleteListing(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Missing ID") {
		t.Errorf("expected 'Missing ID' in body, got %s", w.Body.String())
	}
}

func TestUpdateListing_InvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PATCH", "/api/v1/properties/update", strings.NewReader("bad"))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateListing(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestParseListingForm_MissingFields(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/", strings.NewReader(""))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := parseListingForm(c)
	if err == nil {
		t.Error("expected error for missing fields")
	}
	if !strings.Contains(err.Error(), "missing required fields") {
		t.Errorf("expected 'missing required fields', got %s", err.Error())
	}
}

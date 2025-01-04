package router_test

import (
	"bytes"
	"gin-samples/internal/mock"
	"gin-samples/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHelloRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &mock.MockHelloController{}

	r := gin.Default()

	router.AddHelloRoutes(r, mockController)

	req, _ := http.NewRequest(http.MethodGet, "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message":"Mocked Hello, World!"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestAddHelloRoutes_CreateGreeting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &mock.MockHelloController{}

	r := gin.Default()

	router.AddHelloRoutes(r, mockController)

	// Mocked Request Body
	body := []byte(`{"message":"Mocked POST Greeting!"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/hello", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expectedResponse := `{"message":"Mocked POST Greeting!"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestAddHelloRoutes_GetAllGreetings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &mock.MockHelloController{}

	r := gin.Default()

	router.AddHelloRoutes(r, mockController)

	req, _ := http.NewRequest(http.MethodGet, "/api/hello/all", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `[
		{"message":"Mocked Hello, World!"},
		{"message":"Mocked Hi!"}
	]`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

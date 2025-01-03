package router_test

import (
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

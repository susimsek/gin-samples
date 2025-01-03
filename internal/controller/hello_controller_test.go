package controller

import (
	"gin-samples/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &testutils.MockHelloService{}
	controller := NewHelloController(mockService)

	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Mock Hello"}`, w.Body.String())
}

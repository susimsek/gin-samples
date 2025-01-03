package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-samples/internal/controller"
	"gin-samples/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	helloService := service.NewHelloService()
	helloController := controller.NewHelloController(helloService)
	r.GET("/api/hello", helloController.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200")

	expectedResponse := `{"message":"Hello, World!"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Expected response body to match")
}

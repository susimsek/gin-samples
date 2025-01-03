package router_test

import (
	"gin-samples/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-samples/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddHelloRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &testutils.MockHelloController{}

	r := gin.Default()

	router.AddHelloRoutes(r, mockController)

	req, _ := http.NewRequest(http.MethodGet, "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message":"Mocked Hello, World!"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

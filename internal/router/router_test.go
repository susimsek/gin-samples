package router_test

import (
	"gin-samples/internal/router"
	"gin-samples/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &testutils.MockHelloController{}

	r := router.SetupRouter(mockController)

	req, _ := http.NewRequest(http.MethodGet, "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message":"Mocked Hello, World!"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-samples/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code 200 expected")
}

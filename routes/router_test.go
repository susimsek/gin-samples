package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter_HelloEndpoint(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "HTTP status code should be 200")

	expectedResponse := `{"message":"Hello, World!"}`

	assert.JSONEq(t, expectedResponse, resp.Body.String(), "Response body should match the expected JSON")
}

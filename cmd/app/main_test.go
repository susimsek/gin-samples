package main

import (
	"gin-samples/testutils"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	mockContainer := testutils.NewMockContainer()
	go func() {
		run(":8081", mockContainer)
	}()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8081/mocked")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expectedResponse := `{"message": "This is a mocked container"}`
	body := getResponseBody(resp)
	assert.JSONEq(t, expectedResponse, body)
}

func getResponseBody(resp *http.Response) string {
	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(bodyBytes)
}

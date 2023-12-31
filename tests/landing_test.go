package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josephkipkemoi/politicapp/server"
	"github.com/stretchr/testify/assert"
)

// Test can view landing route message
func TestLanding(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should return success status code: 200")
}

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josephkipkemoi/politicapp/server"
	"github.com/stretchr/testify/assert"
)

func TestCanRegisterUser(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	input := struct {
		FirstName       string
		LastName        string
		PhoneNumber     int
		Password        string
		ConfirmPassword string
	}{
		FirstName:       "Joseph",
		LastName:        "Ngetich",
		PhoneNumber:     254700545727,
		Password:        "1234",
		ConfirmPassword: "1234",
	}

	// Json encode input struct
	d, err := json.Marshal(&input)
	if err != nil {
		fmt.Println(err)
	}
	bodyReader := bytes.NewReader(d)
	req, _ := http.NewRequest("POST", "/api/v1/register", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code, "Should return status code 201")
}

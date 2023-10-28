package websvc

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var p = &Parameters{Port: 8080}

func TestHealthCheckHandlerOK(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	w, _ := New(p)
	pg := gomonkey.ApplyGlobalVar(&Web, w)

	req, err := http.NewRequest("GET", URIHealth, nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `{"web-service-status":"UP","ldap-service-status":"DOWN"}`
	assert.Equal(t, rr.Body.String(), expected)
	pg.Reset()
}

func TestHealthCheckHandlerWebNil(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	Web = nil

	req, err := http.NewRequest("GET", URIHealth, nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the response body is what we expect.
	expected := `{"web-service-status":"DOWN","ldap-service-status":"DOWN"}`
	assert.Equal(t, expected, rr.Body.String())
}

package websvc

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewFail(t *testing.T) {
	p0 := gomonkey.ApplyFuncReturn(mux.NewRouter, nil)
	w, _ := New(&Parameters{Port: 8080})
	assert.Nil(t, w)
	p0.Reset()
}

func TestNewSuccess(t *testing.T) {
	w, _ := New(&Parameters{Port: 8080})
	assert.NotNil(t, w)
}

func TestWebsvc_Start(t *testing.T) {
	w, _ := New(&Parameters{Port: 8080})
	p0 := gomonkey.ApplyMethodReturn(w.srv, "ListenAndServe", http.ErrServerClosed)
	w.Start()
	p0.Reset()

	p0 = gomonkey.ApplyMethodReturn(w.srv, "ListenAndServe", errors.New("different error"))
	w.Start()
	p0.Reset()
}

func TestWebsvc_StopFail(t *testing.T) {
	w, _ := New(&Parameters{Port: 8080})
	p0 := gomonkey.ApplyMethodReturn(w.srv, "Shutdown", errors.New("shutdown error"))
	assert.Error(t, w.Stop())
	p0.Reset()
}

func TestWebsvc_StopSuccess(t *testing.T) {
	w, _ := New(&Parameters{Port: 8080})
	p0 := gomonkey.ApplyMethodReturn(w.srv, "Shutdown", nil)
	assert.Nil(t, w.Stop())
	p0.Reset()
}

func TestWebsvc_AddRoute(t *testing.T) {
	var w *Websvc
	assert.Error(t, w.AddRoute("", nil))

	w, _ = New(&Parameters{Port: 8080})
	assert.Error(t, w.AddRoute("", nil))
	assert.Error(t, w.AddRoute("", LDAPQueryHandler))
	assert.NoError(t, w.AddRoute("test", LDAPQueryHandler))
}

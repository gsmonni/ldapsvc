package websvc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func setJsonResponse(data interface{}, w http.ResponseWriter) {
	var j []byte
	var err error

	j, err = json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		j = []byte(fmt.Sprintf("error while converting response data (%v)", err.Error()))
	}
	_, _ = fmt.Fprintf(w, string(j))
}

func LDAPQueryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := QueryResponse{
		QueryAttributeType:  vars[QueryAttributeType],
		QueryAttributeValue: vars[QueryAttributeValue],
	}
	setJsonResponse(res, w)
}

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	s := ServerStatus{ReturnStatusCode: http.StatusOK}
	if Web == nil {
		s.LDAPServiceStatus = LDAPStatusDown
		s.WebSvcStatus = ServiceStatusDown
		s.ReturnStatusCode = http.StatusInternalServerError
		w.WriteHeader(s.ReturnStatusCode)
	} else {
		s.LDAPServiceStatus = LDAPStatusDown
		s.WebSvcStatus = ServiceStatusUp
	}
	setJsonResponse(s, w)
}

func StopRequestHandler(w http.ResponseWriter, _ *http.Request) {
	s := ServerStatus{ReturnStatusCode: http.StatusOK}
	if Web != nil {
		setJsonResponse(s, w)
		time.Sleep(200 * time.Millisecond)
		if err := Web.Stop(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			_, _ = fmt.Fprintf(w, "server stopped")
			w.WriteHeader(http.StatusOK)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

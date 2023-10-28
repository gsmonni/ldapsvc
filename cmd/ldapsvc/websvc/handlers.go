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

func LdaPQueryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := QueryResponse{
		QueryAttributeType:  vars[QueryAttributeType],
		QueryAttributeValue: vars[QueryAttributeValue],
	}
	setJsonResponse(res, w)
}

func Health(w http.ResponseWriter, _ *http.Request) {
	s := ServerStatus{ReturnStatusCode: http.StatusOK}
	if Web == nil {
		s.LDAPServiceStatus = LDAPStatusDown
		s.WebSvcStatus = ServiceStatusDown
		s.ReturnStatusCode = http.StatusInternalServerError
	} else {
		s.LDAPServiceStatus = LDAPStatusDown
		s.WebSvcStatus = ServiceStatusUp
	}
	setJsonResponse(s, w)
	w.WriteHeader(s.ReturnStatusCode)
}

func Stop(w http.ResponseWriter, _ *http.Request) {
	s := ServerStatus{ReturnStatusCode: http.StatusOK}
	if Web != nil {
		setJsonResponse(s, w)
		time.Sleep(1 * time.Second)
		Web.Stop()
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

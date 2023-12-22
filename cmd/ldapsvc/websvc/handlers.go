package websvc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/ldapbackend"
	"log"
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
	q := fmt.Sprintf("%s=%s", vars[QueryAttributeType], vars[QueryAttributeValue])
	v := ldapbackend.QueryResult{}
	if !v.IsValidFieldName(vars[QueryAttributeType]) {
		s := ReturnMessage{Message: fmt.Sprintf("%s is not a valid property-name", vars[QueryAttributeType])}
		setJsonResponse(s, w)
		w.WriteHeader(s.Code)
	}
	if Provider == nil {
		s := ReturnMessage{"LDAP Provider not initialized", http.StatusBadRequest}
		setJsonResponse(s, w)
		w.WriteHeader(s.Code)
	}
	if r, err := Provider.Query(q); err != nil {
		s := ReturnMessage{fmt.Sprintf("ldap query error (%v)", err.Error()), http.StatusBadRequest}
		setJsonResponse(s, w)
		w.WriteHeader(s.Code)
	} else {
		log.Printf("%v", r)
		if len(*r) > 0 {
			setJsonResponse(*r, w)
		} else {
			s := ReturnMessage{fmt.Sprintf("ldap query returned empty result"), http.StatusOK}
			setJsonResponse(s, w)
		}
	}
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
		s.LDAPMock = Web.p.LDAP.Mock
		if s.LDAPMock {
			s.MockDataFile = Web.p.LDAP.MockDataFile
		}
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

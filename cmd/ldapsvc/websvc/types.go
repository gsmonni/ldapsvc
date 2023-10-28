package websvc

import (
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type (
	IWebSvc interface {
		Start()
		Stop()
	}
	Websvc struct {
		r   *mux.Router
		wg  *sync.WaitGroup
		srv *http.Server
	}
	QueryResponse struct {
		QueryAttributeType  string   `json:",omitempty"`
		QueryAttributeValue string   `json:",omitempty"`
		CommonName          string   `json:",omitempty"`
		Groups              []string `json:",omitempty"`
		Roles               []string `json:",omitempty"`
		Country             []string `json:",omitempty"`
		Uid                 string   `json:",omitempty"`
		SN                  string   `json:",omitempty"`
	}
	ServerStatus struct {
		WebSvcStatus      string `json:"web-service-status"`
		LDAPServiceStatus string `json:"ldap-service-status"`
		ReturnStatusCode  int    `json:"-"`
	}
)

package websvc

import (
	"github.com/ardanlabs/conf/v3"
	"github.com/gorilla/mux"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/ldapbackend"
	"net/http"
	"sync"
)

type (
	ReturnMessage struct {
		Message string
		Code    int
	}

	Parameters struct {
		Certificate    common.TCertificate
		LocalAddress   string `json:"local-address"`
		Port           int    `json:"port"`
		SaveLastConfig bool   `json:"save-last-config"`
		LDAP           ldapbackend.LDAPParameters
		conf.Version   `json:"-" conf:"-"`
	}

	IWebSvc interface {
		Start()
		Stop() error
		Walk()
	}
	Websvc struct {
		p   *Parameters
		r   *mux.Router
		wg  *sync.WaitGroup
		srv *http.Server
		IWebSvc
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
		LDAPMock          bool   `json:"ldap-mock"`
		MockDataFile      string `json:"mock-file"`
		ReturnStatusCode  int    `json:"-"`
	}
)

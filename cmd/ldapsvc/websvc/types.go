package websvc

import (
	"github.com/ardanlabs/conf/v3"
	"github.com/gorilla/mux"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"net/http"
	"sync"
)

type (
	Parameters struct {
		Certificate  common.TCertificate
		LocalAddress string
		Port         int          `conf:"default:8080"`
		Version      conf.Version `conf:"-"`
	}

	IWebSvc interface {
		Start()
		Stop()
	}
	Websvc struct {
		p   *Parameters
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

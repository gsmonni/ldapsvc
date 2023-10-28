package common

import (
	"crypto/tls"
)

type (
	TPort uint16

	TCertificate struct {
		UseTLS  bool `json:"use-tls" conf:"default:false"`
		UseMTLS bool `json:"use-mtls" conf:"default:false"`

		CAFile   string `json:"ca-file" conf:"default:ca.cert"`
		CertFile string `json:"cert-file" conf:"default:/Users/gianstefanomonni/git/ladapsvc/data/certs/cert.crt"`
		KeyFile  string `json:"key-file" conf:"default:/Users/gianstefanomonni/git/ladapsvc/data/certs/private.key"`
		CAPath   string `json:"ca-path" conf:"default:."`

		Cert tls.Certificate `json:"-" conf:"-"`
	}
)

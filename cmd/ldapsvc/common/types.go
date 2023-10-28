package common

import (
	"crypto/tls"
)

type (
	TPort uint16

	TCertificate struct {
		UseTLS  bool `conf:"default:false"`
		UseMTLS bool `conf:"default:false"`
		Strict  bool `conf:"default:false"`

		CAFile   string `conf:"default:ca.cert"`
		CertFile string `conf:"default:/Users/gianstefanomonni/git/ladapsvc/data/certs/cert.crt"`
		KeyFile  string `conf:"default:/Users/gianstefanomonni/git/ladapsvc/data/certs/private.key"`

		Cert tls.Certificate `conf:"-"`

		CAPath string `conf:"default:."`
	}
)

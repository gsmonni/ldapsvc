package common

import (
	"crypto/tls"
)

type (
	TPort uint16

	TCertificate struct {
		UseTLS  bool `json:"use-tls"`
		UseMTLS bool `json:"use-mtls"`

		CAFile   string `json:"ca-file"`
		CertFile string `json:"cert-file"`
		KeyFile  string `json:"key-file"`
		CAPath   string `json:"ca-path"`

		Cert tls.Certificate `json:"-" conf:"-"`
	}
)

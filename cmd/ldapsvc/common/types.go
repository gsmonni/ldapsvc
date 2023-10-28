package common

import (
	"crypto/tls"
)

type (
	TPort uint16

	TCertificate struct {
		UseTLS  bool
		UseMTLS bool
		Strict  bool

		CAFile   string
		CertFile string
		KeyFile  string

		Cert tls.Certificate

		CAPath string
	}
)

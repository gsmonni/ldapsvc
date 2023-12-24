package common

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCertificate_Validate(t *testing.T) {
	var c *TCertificate
	assert.Error(t, c.Validate())

	c = &TCertificate{
		UseTLS:   false,
		UseMTLS:  false,
		CAFile:   "",
		CertFile: "",
		KeyFile:  "",
		CAPath:   "",
		Cert:     tls.Certificate{},
	}
	assert.NoError(t, c.Validate())

	c.UseMTLS = true
	assert.Error(t, c.Validate())
}

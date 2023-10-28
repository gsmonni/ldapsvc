package common

import (
	"crypto/tls"
	"fmt"
	"path/filepath"
)

func (c *TCertificate) Validate() error {
	if c.UseMTLS && !c.UseTLS {
		c.UseTLS = true
	}
	if c.UseTLS {
		var err error
		if !FileExists(c.KeyFile) {
			return fmt.Errorf("invalid key-file %s", c.KeyFile)
		}
		if !FileExists(c.CertFile) {
			return fmt.Errorf("invalid key-file %s", c.CertFile)
		}
		if c.Cert, err = tls.LoadX509KeyPair(c.CertFile, c.KeyFile); err != nil {
			return err
		}
		if c.UseMTLS {
			if !FileExists(c.CAFile) {
				return fmt.Errorf("invalid ca-file %s", c.CAFile)
			}
			if c.CAPath == "" {
				c.CAPath = filepath.Dir(c.KeyFile)
			}
			if !IsDir(c.CAPath) {
				return fmt.Errorf("invalid CAPath %s", c.CAPath)
			}
		}
	}
	return nil
}

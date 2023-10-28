package main

import (
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
	"log"
)

func main() {
	p := websvc.Parameters{
		Port: 8443,
		Certificate: common.TCertificate{
			UseTLS:   true,
			CertFile: "/Users/gianstefanomonni/git/ladapsvc/data/certs/cert.crt",
			KeyFile:  "/Users/gianstefanomonni/git/ladapsvc/data/certs/private.key",
		},
	}
	var err error
	if websvc.Web, err = websvc.New(&p); err != nil {
		log.Fatalf("cannot build web-service(%v)", err.Error())
	}
	websvc.Web.Start()
}

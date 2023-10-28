package main

import (
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/cli"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
	"log"
)

func main() {
	var cfg websvc.Parameters

	if err := cli.Parse(&cfg); err != nil {
		log.Fatalf("main : error while parsing Config : %v", err)
	}

	var err error
	if websvc.Web, err = websvc.New(&cfg); err != nil {
		log.Fatalf("cannot build web-service(%v)", err.Error())
	}
	websvc.Web.Start()
}

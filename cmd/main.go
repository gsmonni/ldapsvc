package main

import (
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
)

func main() {
	websvc.Web = websvc.New("localhost", "8080")
	websvc.Web.Start()
}

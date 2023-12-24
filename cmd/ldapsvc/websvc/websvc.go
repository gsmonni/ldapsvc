package websvc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/ldapbackend"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
// load CA certificate file and add it to list of client CAs

	caCertFile, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
	    log.Fatalf("error reading CA certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertFile)

	// serve on port 9090 of local host
	server := http.Server{
	    Addr:    ":9090",
	    Handler: handler,
	    TLSConfig: &tls.Config{
	        ClientAuth: tls.RequireAndVerifyClientCert,
	        ClientCAs:  certPool,
	        MinVersion: tls.VersionTLS12,
	    },
	}
*/
func buildServer(p *Parameters, h http.Handler) (*http.Server, error) {
	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("cannot start server. invalid parameters (%v)", err.Error())
	}
	var s *http.Server

	if p.Certificate.UseTLS {
		TLSConf := &tls.Config{
			Certificates: []tls.Certificate{p.Certificate.Cert},
		}
		if p.Certificate.UseMTLS {
			caCertFile, err := os.ReadFile(p.Certificate.CAFile)
			if err != nil {
				return nil, fmt.Errorf("error reading CA certificate: %v", err)
			}
			certPool := x509.NewCertPool()
			certPool.AppendCertsFromPEM(caCertFile)
			TLSConf.ClientAuth = tls.RequireAndVerifyClientCert
			TLSConf.ClientCAs = certPool
			TLSConf.MinVersion = tls.VersionTLS12
		}
		s = &http.Server{
			Addr:      fmt.Sprintf("%s:%d", p.LocalAddress, p.Port),
			Handler:   h,
			TLSConfig: TLSConf,
		}
	} else {
		s = &http.Server{
			Addr:         fmt.Sprintf("%s:%d", p.LocalAddress, p.Port),
			Handler:      h,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
	}
	return s, nil
}

func New(p *Parameters) (*Websvc, error) {
	var (
		w   = new(Websvc)
		err error
	)
	if err = p.Validate(); err != nil {
		return nil, err
	}

	w.p = p
	w.r = mux.NewRouter()
	if w.r == nil {
		return nil, fmt.Errorf("cannot build mux-router")
	}
	w.r.Use(commonMiddleware)
	w.r.StrictSlash(true)

	sw := w.r.PathPrefix("/swaggerui").Subrouter()
	sw.Use(SwaggerMiddleware)
	swaggerpath := filepath.Join(common.Datapath, "data/swagger/")
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir(swaggerpath)))
	sw.PathPrefix("/").Handler(sh)
	sw.StrictSlash(true)
	apir := w.r.PathPrefix("/api/v1").Subrouter()
	apir.Use(apiMiddleware)
	_ = AddRoute(apir, URILDAPQuery, LDAPQueryHandler)
	_ = AddRoute(apir, URIHealth, HealthCheckHandler)
	_ = AddRoute(apir, URIStop, StopRequestHandler)

	if w.srv, err = buildServer(p, w.r); err != nil {
		return nil, err
	}
	// add waitgroup
	w.wg = &sync.WaitGroup{}
	w.wg.Add(1)

	if Provider, err = ldapbackend.New(&p.LDAP); err != nil {
		return nil, fmt.Errorf("cannot build LDAP provider (%v)", err.Error())
	}

	return w, nil
}

func (w *Websvc) Start() {
	if w == nil || w.srv == nil {
		if w == nil {
			log.Fatal("webservice set to nil. cannot start")
		}
		log.Fatal("server set to nil. cannot start")
	}
	go func() {
		defer w.wg.Done() // let main know we are done cleaning up
		// always returns error. ErrServerClosed on graceful close
		if w.p.Certificate.UseTLS {
			if err := w.srv.ListenAndServeTLS(w.p.Certificate.CertFile, w.p.Certificate.KeyFile); !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("cannot start https server (%v)", err)
			}
		} else {
			if err := w.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				// unexpected error. port in use?
				log.Fatalf("cannot start http server (%v)", err)
			}
		}

	}()
	// wait for goroutine started in startHttpServer() to stop
	w.wg.Wait()
}

func (w *Websvc) Stop() error {
	if err := w.srv.Shutdown(context.TODO()); err != nil {
		return err // failure/timeout shutting down the server gracefully
	}
	log.Printf("main: done. exiting")
	return nil
}

func (w *Websvc) Walk() {
	_ = w.r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
}

func AddRoute(r *mux.Router, uri string, h http.HandlerFunc) error {
	if r == nil {
		return fmt.Errorf("invalid web service")
	}
	if h == nil {
		return fmt.Errorf("invalid handle func")
	}
	if uri == "" {
		return fmt.Errorf("invalid uri")
	}
	r.HandleFunc(uri, h)
	return nil
}

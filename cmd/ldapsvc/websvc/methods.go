package websvc

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

func buildServer(p *Parameters, h http.Handler) (*http.Server, error) {
	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("cannot start server. invalid parameters (%v)", err.Error())
	}
	var s *http.Server

	if p.Certificate.UseTLS {
		s = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", p.LocalAddress, p.Port),
			Handler: h,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{p.Certificate.Cert},
			},
		}
	} else {
		s = &http.Server{
			Addr:         fmt.Sprintf("%s:%d", p.LocalAddress, p.Port),
			Handler:      h,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
	}
	if s == nil {
		return nil, fmt.Errorf("cannot create server")
	}
	return s, nil
}

// Simple wrapper to Allow CORS.
func withCORS(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fn.ServeHTTP(w, r)
	}
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

	sw := w.r.PathPrefix("/swaggerui").Subrouter()
	sw.Use(SwaggerMiddleware)
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./data/swagger/")))
	sw.PathPrefix("/").Handler(sh)

	apir := w.r.PathPrefix("/api/v1").Subrouter()
	apir.Use(apiMiddleware)
	_ = AddRoute(apir, URILDAPQuery, LDAPQueryHandler)
	_ = AddRoute(apir, URIHealth, HealthCheckHandler)
	_ = AddRoute(apir, URIStop, StopRequestHandler)

	w.Walk()

	if w.srv, err = buildServer(p, w.r); err != nil {
		return nil, err
	}
	// add waitgroup
	w.wg = &sync.WaitGroup{}
	w.wg.Add(1)

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
			if err := w.srv.ListenAndServeTLS(w.p.Certificate.CertFile, w.p.Certificate.KeyFile); err != http.ErrServerClosed {
				log.Fatalf("cannot start https server (%v)", err)
			}
		} else {
			if err := w.srv.ListenAndServe(); err != http.ErrServerClosed {
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

func (w *Websvc) Walk() {
	_ = w.r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})

}

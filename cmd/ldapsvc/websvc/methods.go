package websvc

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

func New(server string, port string) *Websvc {
	var w = new(Websvc)
	w.r = mux.NewRouter()
	if w.r == nil {
		return nil
	}
	w.r.Use(commonMiddleware)

	w.r.HandleFunc(URILDAPQuery, LdaPQueryHandler)
	w.r.HandleFunc(URIHealth, Health)
	w.r.HandleFunc(URIStop, Stop)

	// add waitgroup
	w.wg = &sync.WaitGroup{}
	w.wg.Add(1)

	w.srv = &http.Server{
		Handler: w.r,
		Addr:    fmt.Sprintf("%s:%s", server, port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return w
}

func (w *Websvc) Start() {
	go func() {
		defer w.wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := w.srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	// wait for goroutine started in startHttpServer() to stop
	w.wg.Wait()

}

func (w *Websvc) Stop() {
	if err := w.srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	log.Printf("main: done. exiting")
}

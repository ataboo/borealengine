package network

import (
	"context"
	"net/http"
	"github.com/ataboo/borealengine/config"
	"time"
	"log"
)

var tlsCertPath string
var tlsKeyPath string

func init() {
	tlsCertPath = "server.crt"
	tlsKeyPath = "server.key"

	// TODO flaggy if needed
}

type NetworkService struct {

}

func (n *NetworkService) Start(ctx context.Context) {
	srv := &http.Server{Addr: config.Config.Ws.HostAddress}

	srv.Handler = http.HandlerFunc(n.handler)

	go func() {
		srv.ListenAndServeTLS(tlsCertPath, tlsKeyPath)
	}()

	log.Printf("\nNetwork Service active on %s", config.Config.Ws.HostAddress)

	for {
		select {
		case <-ctx.Done():
			//cleanup?
			srv.Shutdown(nil)
			return
		default:
			//
		}
	}
}

func (n *NetworkService) handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		w.Write([]byte("Not Found"))
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Tadaa!"))
}

func (n *NetworkService) Update(delta time.Duration) {
	//
}

package network

import (
	"context"
	"net/http"
	. "github.com/ataboo/borealengine/config"
	"time"
	"log"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/ataboo/gowssrv/models"
	"github.com/satori/go.uuid"
	"github.com/op/go-logging"
	"github.com/ataboo/gowssrv/accountstorage"
)

var tlsCertPath string
var tlsKeyPath string

var Logger = logging.MustGetLogger(Config.General.LogName)

func init() {
	tlsCertPath = "server.crt"
	tlsKeyPath = "server.key"

	// TODO flaggy if needed
}

type NetworkService struct {
	upgrader websocket.Upgrader
}

func (n *NetworkService) Start(ctx context.Context) {
	var checkOriginFunc func(r *http.Request) bool
	if !Config.Ws.EnforceOrigin {
		// In development, using certain ws clients might show as a false positive for a cross-site attack.
		checkOriginFunc = func(r *http.Request) bool {
			return true
		}
	}

	n.upgrader = websocket.Upgrader{
		CheckOrigin: checkOriginFunc,
		ReadBufferSize: Config.Ws.ReadBuffer,
		WriteBufferSize: Config.Ws.WriteBuffer,
	}

	srv := &http.Server{Addr: Config.Ws.HostAddress}

	srv.Handler = tokenMiddleware(n.handler)

	go func() {
		srv.ListenAndServeTLS(tlsCertPath, tlsKeyPath)
	}()

	log.Printf("\nNetwork Service active on %s", Config.Ws.HostAddress)

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
		jsonErrorResponse(w, 404, "Not Found")
		return
	}

	conn , err := n.upgrader.Upgrade(w, r, w.Header())

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Failed to upgrade connection"))
	}

	w.WriteHeader(200)
	w.Write([]byte("Tadaa!"))
}

func tokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := userFromToken(r.Header.Get("Authorization"))
		if err != nil {
			jsonErrorResponse(w, 401, "not authorized")
			return
		}

		// Add user to request for future handlers.
		r = r.WithContext(context.WithValue(r.Context(), "user", user))

		next(w, r)
	})
}

func userFromToken(token string) (models.User, error) {
	parsed, err := uuid.FromString(token)
	if err != nil || parsed.Version() != uuid.V4 {
		Logger.Debug("Token not valid UUID4")
	}

	user, err := accountstorage.Users.BySession(token)
	if err != nil {
		Logger.Debug("No user found with session id")
		return user, err
	}

	return user, nil
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{})  {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func jsonErrorResponse(w http.ResponseWriter, code int, payload interface{}) {
	jsonResponse(w, code, map[string]interface{}{"error": payload})
}

func (n *NetworkService) Update(delta time.Duration) {
	//
}

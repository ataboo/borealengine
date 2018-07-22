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
)

var tlsCertPath string
var tlsKeyPath string

func init() {
	tlsCertPath = "server.crt"
	tlsKeyPath = "server.key"

	// TODO flaggy if needed
}

type UserConn struct {
	User models.UserPublic
	Conn *websocket.Conn
}

type ConnectChan chan UserConn

type Service struct {
	upgrader websocket.Upgrader
	OnConnect ConnectChan
}

func NewService() *Service {
	var checkOriginFunc func(r *http.Request) bool
	if !Config.Ws.EnforceOrigin {
		// In development, using certain ws clients might show as a false positive for a cross-site attack.
		checkOriginFunc = func(r *http.Request) bool {
			return true
		}
	}

	return &Service{
		websocket.Upgrader{
			CheckOrigin: checkOriginFunc,
			ReadBufferSize: Config.Ws.ReadBuffer,
			WriteBufferSize: Config.Ws.WriteBuffer,
		},
		make(ConnectChan),
	}
}

func (n *Service) Start(ctx context.Context) {
	srv := &http.Server{Addr: Config.Ws.HostAddress}

	srv.Handler = n.tokenMiddleware(n.upgradeWs)

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

func (n *Service) upgradeWs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		jsonErrorResponse(w, 404, "Not Found")
		return
	}

	conn , err := n.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Failed to upgrade connection"))
	}

	select {
		case n.OnConnect <- UserConn{r.Context().Value("user").(models.UserPublic), conn}:
			return
		case <-time.After(time.Second * 1):
			Logger.Errorf("timed out passing user connect to engine")
			conn.Close()
			jsonErrorResponse(w, 500, "error connecting")
			return
	}
}

func (n *Service) tokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := n.userFromToken(r.Header.Get("Authorization"))
		if err != nil {
			jsonErrorResponse(w, 401, "not authorized")
			return
		}

		// Add user to request for future handlers.
		r = r.WithContext(context.WithValue(r.Context(), "user", user.ToPublic()))

		next(w, r)
	})
}

func (n *Service) userFromToken(token string) (models.User, error) {
	parsed, err := uuid.FromString(token)
	if err != nil || parsed.Version() != uuid.V4 {
		Logger.Debug("Token not valid UUID4")
	}

	user, err := accounts.GetByToken(token)
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

func (n *Service) Update(delta time.Duration) {
	//
}

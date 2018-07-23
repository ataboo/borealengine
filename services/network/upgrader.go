package network

import (
	"net/http"
	"github.com/gorilla/websocket"
	. "github.com/ataboo/borealengine/config"
)

func getUpgraderForEnv() Upgrader {
	switch Config.Environment() {
	case EnvProd:
		return prodUpgrader()
	case EnvTest:
		return testUpgrader()
	default:
		return devUpgrader()
	}
}

type Upgrader interface {
	Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error)
}

type upgradeHandler func(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error)

func prodUpgrader() Upgrader {
	return &websocket.Upgrader{
		// Default behavior matches domain to origin
		CheckOrigin: nil,
		ReadBufferSize: Config.Ws.ReadBuffer,
		WriteBufferSize: Config.Ws.WriteBuffer,
	}
}

func devUpgrader() Upgrader {
	return &websocket.Upgrader{
		// Origin may mismatch in dev with certain clients.  False positive for cross-site attack.
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize: Config.Ws.ReadBuffer,
		WriteBufferSize: Config.Ws.WriteBuffer,
	}
}

func testUpgrader() Upgrader {
	upgrader := FakeUpgrader{
		// Replace this method for testing connections.
		func(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
			return nil, nil
		},
	}

	return &upgrader
}

type FakeUpgrader struct {
	UpgradeHandler upgradeHandler
}

func (f *FakeUpgrader) Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
	return f.UpgradeHandler(w, r, h)
}

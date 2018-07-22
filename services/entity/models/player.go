package models

import (
	"github.com/gorilla/websocket"
	"github.com/ataboo/gowssrv/models"
	"time"
	"sync"
)

type Player struct {
	Lock sync.Mutex
	Creature *Creature
	Comms Communication
	User models.UserPublic
}

type Communication struct {
	Conn      *websocket.Conn
	LastPong  time.Time
	EventsOut UpdateQueue
	EventsIn  UpdateQueue
}

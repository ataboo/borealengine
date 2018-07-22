package entity

import (
	"context"
	"time"
	"github.com/ataboo/borealengine/services/entity/models"
	"gopkg.in/mgo.v2/bson"
	srvmodels "github.com/ataboo/gowssrv/models"
	"github.com/gorilla/websocket"
	"github.com/ataboo/borealengine/services/entity/playercontrol"
	"github.com/ataboo/borealengine/services/entity/npccontrol"
)

type PlayerMap map[bson.ObjectId]playercontrol.Control
type NPCMap map[bson.ObjectId]npccontrol.Control

type Service struct {
	Players       PlayerMap
	NPCs          NPCMap
}

func (s *Service) Start(ctx context.Context) {
	// Load snapshot

	// Call Start on entities

	//
}

func (s *Service) Update(delta time.Duration) {
	for _, p := range s.Players {
		p.HandleIncomingEvents(delta)
	}

	for _, n := range s.NPCs {
		n.UpdateBehavior(delta)
	}

	for _, p := range s.Players {
		p.ResolvePhysics(delta)
	}

	for _, n := range s.NPCs {
		n.ResolvePhysics(delta)
	}
}

func (s *Service) PlayerConnected(user srvmodels.UserPublic, conn *websocket.Conn) error {
	ctl, ok := s.Players[user.ID]
	if !ok {
		player := models.Player{
			Creature: nil,
			Comms: models.Communication{
				Conn: conn,
				LastPong: time.Now(),
				EventsOut: models.UpdateQueue{},
				EventsIn: models.UpdateQueue{},
			},
			User: user,
		}

		ctl := playercontrol.Control{&player}
		s.Players[user.ID] = ctl

		// init?
	}

	ctl.Connect(conn)

	return nil
}

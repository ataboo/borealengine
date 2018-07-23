package playercontrol

import (
	"time"
	"github.com/ataboo/borealengine/services/entity/models"
	"github.com/ataboo/borealengine/config"
	"github.com/gorilla/websocket"
	"github.com/ataboo/borealengine/services/entity"
)

type Control struct {
	player *models.Player
}

var Logger = config.Logger

func (c *Control) HandleIncomingEvents(delta time.Duration)  {
	comms := c.player.Comms

	for {
		if !c.player.Comms.EventsIn.HasNext() {
			return
		}

		event, _ := comms.EventsIn.Pop()

		switch event.GetType() {
		case models.UServer:
			c.applyServerUpdate(event.(models.ServerUpdate))
		default:
			Logger.Errorf("event type not supported as incoming: %+v", event)
		}
	}
}

func (c *Control) applyServerUpdate(update models.ServerUpdate) {
	//TODO: Validate update

	c.player.Creature.Transform = update.Transform

	//TODO: attacking

	//TODO: chat
	// TODO: transition
}

func (c *Control) UpdateNeighbours(delta time.Duration) {

}

func (c *Control) FindNeighbours(delta time.Duration) []entity.NPCControl {
	//There's probably a better way to do this.


}

func (c *Control) ResolvePhysics(delta time.Duration) {
	//TODO: physic...lol
}

func (c *Control) Connect(conn *websocket.Conn) {
	c.player.Lock.Lock()
	defer c.player.Lock.Unlock()

	comms := c.player.Comms
	if comms.Conn != nil {
		comms.Conn.Close()
	}

	comms.Conn = conn

	StartReader(c.player.User.Username, &comms)
	StartWriter(c.player.User.Username, &comms)

	// TODO: send init?
}




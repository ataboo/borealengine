package npccontrol

import (
	"github.com/ataboo/borealengine/services/entity/models"
	"time"
)

type Control struct {
	npc models.NPC
	LastUpdate time.Time
}

func (c *Control) UpdateBehavior(delta time.Duration) {

}

func (c *Control) ResolvePhysics(delta time.Duration) {

}

package entity

import (
	"time"
	"github.com/ataboo/borealengine/services/entity/models"
)

type NPCControl struct {
	ctx interface{}
	eventQueue models.UpdateQueue
}

func (c *NPCControl) Start(ctx interface{}) {
	c.ctx = ctx
}

func (c *NPCControl) ResolveEvents(creature models.Creature, delta time.Duration)  {

}

func (c *NPCControl) PhysicsUpdate(creature models.Creature, delta time.Duration) {

}


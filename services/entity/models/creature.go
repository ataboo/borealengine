package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/ataboo/borealengine/mathies"
	"time"
)

type Cooldowns struct {
	Attack time.Time
	Transition time.Time
	Chat time.Time
}


type Creature struct {
	ID bson.ObjectId
	Name string

	Transform Transform
	Vitals    Vitals

	Velocity mathies.Vector2

	// Attached To

	Species Species

	Attacking bool
	Cooldowns Cooldowns

}

func (c *Creature) GetId() bson.ObjectId {
	return c.ID
}

func (c *Creature) GetTransform() Transform {
	return c.Transform
}

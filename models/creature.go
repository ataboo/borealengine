package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Controller interface {
	Update(delta time.Duration, ctx interface{})
}

type Creature struct {
	ID bson.ObjectId
	Name string

	Transform Transform
	Vitals    Vitals

	// Attached To
	// Contacts

	Species Species

	Control Controller
}

func (c *Creature) GetId() bson.ObjectId {
	return c.ID
}

func (c *Creature) GetTransform() Transform {
	return c.Transform
}
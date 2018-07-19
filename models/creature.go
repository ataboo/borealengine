package models

import "gopkg.in/mgo.v2/bson"

type Creature struct {
	ID bson.ObjectId
	Name string

	Transform Transform
	Vitals    Vitals

	// Attached To
	// Contacts

	Species Species
}

func (c *Creature) GetId() bson.ObjectId {
	return c.ID
}

func (c *Creature) GetTransform() Transform {
	return c.Transform
}
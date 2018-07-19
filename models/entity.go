package models

import "gopkg.in/mgo.v2/bson"

type Entity interface {
	GetId() bson.ObjectId
	GetTransform() Transform
}
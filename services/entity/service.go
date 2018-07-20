package entity

import (
	"context"
	"time"
	"github.com/ataboo/borealengine/models"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type CreatureMap map[bson.ObjectId]models.Creature


type Service struct {
	Creatures CreatureMap
	lock sync.Mutex
}

func (s *Service) Start(ctx context.Context) {

}

func (s *Service) Update(delta time.Duration) {

}

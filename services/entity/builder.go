package entity

import (
	"github.com/ataboo/borealengine/models"
	"fmt"
	"strings"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

var CreatureBuilder = cBuilder{}

type cBuilder struct {
	Template models.Creature
	errors   []string
}

func (c *cBuilder) Start(speciesName string) *cBuilder {
	c.errors = []string{}
	c.Template = models.Creature{}
	spec, ok := SpeciesConfigs[speciesName]
	if ok {
		c.Template.Species = spec
		c.setDefaultVitals()
	} else {
		c.errors = append(c.errors, fmt.Sprintf("species %s is invalid", speciesName))
	}

    return c
}

func (c *cBuilder) setDefaultVitals() {
	c.Template.Vitals = models.Vitals{
		Health: 100,
		Food: 50,
		Water: 50,
		Energy: 50,
		Stamina: 100,
		Age: 0,
	}
}

func (c *cBuilder) MakeMany(count int) ([]models.Creature, error) {
	if len(c.errors) > 0 {
		joined := strings.Join(c.errors, ", ")
		return make([]models.Creature, 1), fmt.Errorf(joined)
	}

	rand.Seed(time.Now().Unix())

	creatures := make([]models.Creature, 0)
	for i := 0; i < count; i++ {
		creature := models.Creature(c.Template)
		creature.ID = bson.NewObjectId()
		creature.Name = Names[rand.Intn(len(Names))]
		creatures = append(creatures, creature)
	}

	return creatures, nil
}

func (c *cBuilder) MakeOne() (models.Creature, error) {
	creatures, err := c.MakeMany(1)
	if err != nil {
		return models.Creature{}, err
	}

	return creatures[0], nil
}
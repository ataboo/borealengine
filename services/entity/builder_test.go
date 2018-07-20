package entity

import (
	"testing"
	"github.com/ataboo/borealengine/models"
)

func TestBuildsSingleBear(t *testing.T) {
	builder := CreatureBuilder.Start("bear")

	bear, bearror := builder.MakeOne()

	if bearror != nil {
		t.Errorf("failed to build a bear %s", bearror)
	}

	validateCreature(bear, t)
}

func TestBuildThreeBears(t *testing.T) {
	builder := CreatureBuilder.Start("bear")
	bears, bearror := builder.MakeMany(3)

	if bearror != nil {
		t.Errorf("failed to build 3 bears with %s", bearror)
	}

	if len(bears) != 3 {
		t.Errorf("wrong number of bears built %d; expected 3", len(bears))
	}

	for _, bear := range bears {
		validateCreature(bear, t)
	}
}

func validateCreature(c models.Creature, t *testing.T) {
	if !c.ID.Valid() {
		t.Errorf("creature id invalid %+v", c)
	}

	if c.Name == "" {
		t.Errorf("creature built without a name %+v", c)
	}

	if c.Vitals.Health == 0 {
		t.Errorf("creature built with health at 0; vitals not initialized? %+v", c)
	}

	if c.Species.Base.Size.R == 0 {
		t.Errorf("creature built without size; species not loaded? %+v", c)
	}
}

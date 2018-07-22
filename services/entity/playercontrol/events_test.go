package playercontrol

import (
	"testing"
	"github.com/ataboo/borealengine/services/entity/models"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func TestRawEventParsing(t *testing.T) {
	target := bson.NewObjectId()
	jString := `{"type":"`+models.UDamage+`","payload": {"target":"`+target.Hex()+`","amount": 42.5}}`

	event, ok := parseEvent(jString)

	if !ok {
		t.Errorf("failed to parse event")
	}

	if event.Type != models.UDamage {
		t.Errorf("expected damage update")
	}

	dmg := models.DamageUpdate{}
	err := json.Unmarshal(event.Payload, &dmg)

	if err != nil {
		t.Errorf("failed to parse damage event with: "+err.Error())
	}

	if dmg.Target != target {
		t.Errorf("target id doesn't match: %s, %s", target, dmg.Target)
	}

	if dmg.Amount != 42.5 {
		t.Errorf("damage amount doesn't match: %s, %s", 42.5, dmg.Amount)
	}
}

func TestUpdateMarshaling(t *testing.T) {
	update := models.DamageUpdate{
		Target: bson.NewObjectId(),
		Amount: 42.5,
	}
	expected := `{"type":"dmg","payload":{"target":"`+update.Target.Hex()+`","amount":42.5}}`

	marshaled, err := marshalUpdate(&update)
	if err != nil {
		t.Errorf("failed to marshal the update")
	}

	if string(marshaled) != expected {
		t.Errorf("marshalled doesn't match %s, %s", expected, string(marshaled))
	}
}

package playercontrol

import (
	"github.com/ataboo/borealengine/services/entity/models"
	"encoding/json"
)

func parseEvent(eventString string) (*models.RawEvent, bool) {
	event := models.RawEvent{}

	err := json.Unmarshal([]byte(eventString), &event)

	return &event, err == nil
}

func parseServerUpdate(updatePayload []byte) (update models.ServerUpdate, ok bool) {
	err := json.Unmarshal(updatePayload, &update)

	return update, err == nil
}

func marshalUpdate(u models.Update) ([]byte, error) {
	rawEvent, err := toRawUpdate(u)
	if err != nil {
		return nil, err
	}

	return json.Marshal(rawEvent)
}

func toRawUpdate(u models.Update) (models.RawEvent, error) {
	marshaledUpdate, err := json.Marshal(u)
	if err != nil {
		return models.RawEvent{}, err
	}

	return models.RawEvent{
		Type: u.GetType(),
		Payload: marshaledUpdate,
	}, nil
}

func marshalUpdates(u []models.Update) ([]byte, error) {
	rawEvents := make([]models.RawEvent, 0)

	for _, event := range u {
		rawEvent, err := toRawUpdate(event)
		if err != nil {
			rawEvents = append(rawEvents, rawEvent)
		}
	}

	return json.Marshal(rawEvents)
}
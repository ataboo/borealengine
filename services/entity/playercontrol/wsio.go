package playercontrol

import (
	"github.com/ataboo/borealengine/services/entity/models"
	"github.com/gorilla/websocket"
	"time"
	"github.com/ataboo/borealengine/config"
)

func StartReader(userName string, comms *models.Communication) {
	go func() {
		defer comms.Conn.Close()

		for {
			msgType, msg, err := comms.Conn.ReadMessage()
			if err != nil {
				Logger.Debugf("Player %s stopped reading because of err: %s", userName, err)
				return
			}

			if msgType != websocket.TextMessage {
				Logger.Debugf("invalid WS message type: %d", msgType)
				continue
			}

			raw, ok := parseEvent(string(msg))
			if !ok {
				Logger.Debugf("malformed event: %s", msg)
				continue
			}

			if raw.Type == models.UServer {
				update, ok := parseServerUpdate(raw.Payload)
				if !ok {
					Logger.Debugf("malformed update: %s", raw.Payload)
					continue
				}

				comms.LastPong = time.Now()
				comms.EventsIn.Push(&update)
			}
		}
	}()
}

func StartWriter(userName string, comms *models.Communication) {
	go func() {
		defer comms.Conn.Close()
		events := comms.EventsOut

		for {
			select {
			case <- time.Tick(time.Microsecond * time.Duration(config.Config.Ws.SendDeltaMicro)):
				if !events.HasNext() {
					continue
				}

				popped := events.PopAll()
				marshaled, err := marshalUpdates(popped)
				if err != nil {
					err := comms.Conn.WriteMessage(websocket.TextMessage, marshaled)
					if err != nil {
						Logger.Debugf("Player %s stopped writing because of err: %s", userName, err)
						return
					}
				} else {
					Logger.Debugf("failed to marshal events: %s", err)
				}
			default:
				//
			}
		}
	}()
}
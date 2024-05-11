package models

import (
	"encoding/json"
	"time"

	"github.com/mastik5h/LLD/eventbus/utils"
)

type EventBusTier string

const (
	Basic   EventBusTier = "basic"
	Premium EventBusTier = "premium"
)

type EventBusMetadata struct {
	Name        string
	Description string
	Tier        EventBusTier
}
type Event struct {
	Id           string
	Name         string
	CreationTime time.Time
	Attributes   map[string]interface{}
}

func NewEvent(id, name string, creationTime time.Time, attributes map[string]interface{}) *Event {
	return &Event{
		Id:           id,
		Name:         name,
		CreationTime: creationTime,
		Attributes:   attributes,
	}
}

func ProcessEventsJson(eventsJson []string) ([]*Event, error) {
	events := make([]*Event, 0)
	for _, e := range eventsJson {
		var event Event
		err := json.Unmarshal([]byte(e), &event)
		if err != nil {
			return events, err
		}

		if event.Id == "" {
			event.Id = utils.GetUniqueIdInString(event.Name)
		}
		if event.CreationTime == (time.Time{}) {
			event.CreationTime = utils.GetCurrentTime()
		}

		events = append(events, &event)
	}
	return events, nil
}

type PublishEventResponse struct {
}

type SubscribeEventResponse struct {
	EventIds []string
}

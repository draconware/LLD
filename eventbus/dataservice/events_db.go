package dataservice

import "github.com/mastik5h/LLD/eventbus/models"

type EventsDB struct {
	events map[string]*models.Event
}

func NewEventsDB() *EventsDB {
	return &EventsDB{
		events: make(map[string]*models.Event, 0),
	}
}

func (edb *EventsDB) InsertEvents(insertEvents []*models.Event) error {
	for _, event := range insertEvents {
		edb.events[event.Id] = event
	}
	return nil
}

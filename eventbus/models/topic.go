package models

import (
	"errors"
	"fmt"
	"time"
)

type Topic struct {
	Id                    string
	Metadata              *TopicMetadata
	EventIdList           []string
	NextEventId           map[string]string
	EventTimestampList    []time.Time
	EventTimestampToIdMap map[time.Time]string
}

type TopicMetadata struct {
	Name        string
	Description string
	Capacity    int
}

type TopicStatus struct {
	Id string
}

func NewTopic(id, name, description string, capacity int) *Topic {
	return &Topic{
		Id: id,
		Metadata: &TopicMetadata{
			Name:        name,
			Description: description,
			Capacity:    capacity,
		},
		EventIdList:           make([]string, 0),
		NextEventId:           make(map[string]string, 0),
		EventTimestampList:    make([]time.Time, 0),
		EventTimestampToIdMap: make(map[time.Time]string, 0),
	}
}

func (t *Topic) AddEvents(events []*Event) error {
	for _, event := range events {
		if err := t.AddEventId(event.Id, event.CreationTime); err != nil {
			return err
		}
	}
	return nil
}

func (t *Topic) AddEventId(eventId string, timestamp time.Time) error {
	t.EventIdList = append(t.EventIdList, eventId)
	t.EventTimestampList = append(t.EventTimestampList, timestamp)
	t.EventTimestampToIdMap[timestamp] = eventId
	totalEvents := len(t.EventIdList)
	if len(t.EventIdList) > 1 {
		t.NextEventId[t.EventIdList[totalEvents-2]] = eventId
	}
	t.NextEventId[eventId] = "null"
	return nil
}

func (t *Topic) GetLastEvents(noOfEvents int) ([]string, error) {
	result := make([]string, 0)
	totalEvents := len(t.EventIdList)
	if totalEvents < noOfEvents {
		return result, fmt.Errorf("Topic: %s doesn't contain more than %d events", t.Metadata.Name, totalEvents)
	}
	for i := 0; i < noOfEvents; i++ {
		result = append(result, t.EventIdList[totalEvents-1-i])
	}
	return result, nil
}

func (t *Topic) GetLastEventsFromId(eventId string, noOfEvents int) ([]string, error) {
	result := make([]string, 0)
	result = append(result, eventId)
	tmpEventId := eventId
	for len(result) != noOfEvents {
		if nextEventId, ok := t.NextEventId[tmpEventId]; !ok || nextEventId == "null" {
			break
		} else {
			result = append(result, nextEventId)
			tmpEventId = nextEventId
		}
	}
	if len(result) != noOfEvents {
		return result, fmt.Errorf("Topic: %s doesn't contain more than %d events after eventId: %s", t.Metadata.Name, len(result), eventId)
	}
	return result, nil
}

func (t *Topic) GetFirstEventIdFromTimestamp(timestamp time.Time) (string, error) {
	nearestTimestamp, err := nearestPossibleTimestamp(timestamp, t.EventTimestampList)
	if err != nil {
		return "", err
	}
	if eventId, ok := t.EventTimestampToIdMap[nearestTimestamp]; ok {
		return eventId, nil
	} else {
		return "", errors.New("eventId doesn't exist at or after given timestamp")
	}
}

func nearestPossibleTimestamp(timestamp time.Time, possibleTimestamps []time.Time) (time.Time, error) {
	for _, pts := range possibleTimestamps {
		if pts.Compare(timestamp) != -1 {
			return pts, nil
		}
	}
	return time.Time{}, errors.New("eventId doesn't exist at or after given timestamp")
}

package v1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mastik5h/LLD/eventbus/models"
	"github.com/mastik5h/LLD/eventbus/utils"
)

func TestNewEventBus(t *testing.T) {
	utils.IntializeGenerator(0)
	eventbusName := "e1"
	eb := NewEventBus(eventbusName, "", "basic")
	if eb == nil {
		t.Fatalf("TestNewEventBus failed: cannot get new eventbus with eventbusname: %s", eventbusName)
	}
}

func TestCreateTopic(t *testing.T) {
	utils.IntializeGenerator(0)
	eventbusName := "e1"
	eb := NewEventBus(eventbusName, "", "basic")
	_, err := eb.CreateTopic("t1", "", 1000)
	if err != nil {
		t.Fatalf("TestCreateTopic failed: cannot create new topic: %s in eventbus: %s", "t1", eventbusName)
	}
}

func TestPublishEvents(t *testing.T) {
	utils.IntializeGenerator(0)
	eventbusName := "e1"
	eb := NewEventBus(eventbusName, "", "basic")
	status, err := eb.CreateTopic("t1", "", 1000)
	if err != nil {
		t.Fatalf("TestCreateTopic failed: cannot create new topic: %s in eventbus: %s", "t1", eventbusName)
	}
	eventsList := make([]string, 0)
	eventsList = append(eventsList, getMockEvent("e1", nil))
	eventsList = append(eventsList, getMockEvent("e2", nil))
	eventsList = append(eventsList, getMockEvent("e3", nil))

	_, err = eb.PublishEvents(status.Id, eventsList)
	if err != nil {
		t.Fatalf("TestPublishEvents failed: cannot publish events: %v to topic: %s in eventbus: %s", eventsList, "t1", eventbusName)
	}
}

func TestPushEvents(t *testing.T) {
	utils.IntializeGenerator(0)
	eventbusName := "e1"
	eb := NewEventBus(eventbusName, "", "basic")
	status, err := eb.CreateTopic("t1", "", 1000)
	if err != nil {
		t.Fatalf("TestPushEvents failed: cannot create new topic: %s in eventbus: %s", "t1", eventbusName)
	}
	eventsList := make([]string, 0)
	eventsList = append(eventsList, getMockEvent("e1", nil))
	eventsList = append(eventsList, getMockEvent("e2", nil))
	eventsList = append(eventsList, getMockEvent("e3", nil))

	_, err = eb.PublishEvents(status.Id, eventsList)
	if err != nil {
		t.Fatalf("TestPushEvents failed: cannot publish events: %v to topic: %s in eventbus: %s", eventsList, "t1", eventbusName)
	}

	response, err := eb.PushEvents(status.Id, "sub1", 2)
	if err != nil {
		t.Fatalf("TestPushEvents failed: cannot push events to subscriber: %s", "sub1")
	}
	if len(response.EventIds) != 2 {
		t.Fatalf("TestPushEvents failed: requested events count: %d, response events count: %d", 2, len(response.EventIds))
	}
}

func TestPullEventsById(t *testing.T) {
	utils.IntializeGenerator(0)
	eventbusName := "e1"
	eb := NewEventBus(eventbusName, "", "basic")
	status, err := eb.CreateTopic("t1", "", 1000)
	if err != nil {
		t.Fatalf("TestPullEventsById failed: cannot create new topic: %s in eventbus: %s", "t1", eventbusName)
	}
	event1Time := time.Now()
	event2Time := event1Time.Add(60 * time.Second)
	event3Time := event1Time.Add(120 * time.Second)
	eventsList := make([]string, 0)
	eventsList = append(eventsList, getMockEvent("e1", &event1Time))
	eventsList = append(eventsList, getMockEvent("e2", &event2Time))
	eventsList = append(eventsList, getMockEvent("e3", &event3Time))

	_, err = eb.PublishEvents(status.Id, eventsList)
	if err != nil {
		t.Fatalf("TestPullEventsById failed: cannot publish events: %v to topic: %s in eventbus: %s", eventsList, "t1", eventbusName)
	}

	response, err := eb.PollEventsWithId(status.Id, "sub2", "e1", 3)
	if err != nil {
		t.Fatalf("TestPullEventsById failed: cannot poll events after event id: %s with count: %d from topic: %s", "e1", 3, "t1")
	}
	if len(response.EventIds) != len(eventsList) {
		t.Fatalf("TestPullEventsById failed: expected events: %v, actual events: %v", eventsList, response.EventIds)
	}

	_, err = eb.PollEventsWithId(status.Id, "sub3", "e4", 5)
	if err == nil {
		t.Fatalf("TestPullEventsById failed: cannot poll events after event id: %s with count: %d from topic: %s.Error: %s", "e4", 5, "t1", err.Error())
	}

	_, err = eb.PollEventsWithId(status.Id, "sub4", "e2", 2)
	if err != nil {
		t.Fatalf("TestPullEventsById failed: cannot poll events after event id: %s with count: %d from topic: %s", "e2", 2, "t1")
	}
}

func getMockEvent(name string, eventTime *time.Time) string {
	mockEvent := &models.Event{
		Id:   name,
		Name: name,
		Attributes: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}
	if eventTime != nil {
		mockEvent.CreationTime = *eventTime
	}
	mockEventBytes, _ := json.Marshal(&mockEvent)
	return string(mockEventBytes)
}

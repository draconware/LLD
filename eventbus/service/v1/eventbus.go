package v1

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mastik5h/LLD/eventbus/dataservice"
	"github.com/mastik5h/LLD/eventbus/models"
	"github.com/mastik5h/LLD/eventbus/utils"
)

type EventBus struct {
	Id          string
	Metadata    *models.EventBusMetadata
	Topics_Db   *dataservice.TopicsDB
	Events_Db   *dataservice.EventsDB
	ExectorPool *utils.ExecutorPool
}

func NewEventBus(name, description, tier string) *EventBus {
	return &EventBus{
		Id: utils.GetUniqueId(),
		Metadata: &models.EventBusMetadata{
			Name:        name,
			Description: description,
			Tier:        models.EventBusTier(tier),
		},
		Topics_Db:   dataservice.NewTopicsDB(),
		Events_Db:   dataservice.NewEventsDB(),
		ExectorPool: utils.NewExecutorPool(10),
	}
}

// CreateOrUpdate topic
func (eb *EventBus) CreateTopic(name, description string, capacity int) (*models.TopicStatus, error) {
	topic := models.NewTopic(utils.GetUniqueId(), name, description, capacity)
	if topic == nil {
		return nil, fmt.Errorf("could not create a topic with submitted configuration. Name: %s, Description: %s, Capacity: %d", name, description, capacity)
	}
	err := eb.Topics_Db.AddTopic(topic)
	if err != nil {
		return nil, err
	}
	return &models.TopicStatus{
		Id: topic.Id,
	}, nil
}

// Delete topic
func (eb *EventBus) DeleteTopic(topicId string) error {
	topic, err := eb.Topics_Db.GetTopic(topicId)
	if err != nil {
		return err
	}
	err = eb.Topics_Db.DeleteTopic(topic)
	if err != nil {
		return err
	}
	return nil
}

// Publish Events to Eventbus
func initiatePublishEvents(eb *EventBus, topicId string, eventsJson []string) error {
	topic, err := eb.Topics_Db.GetTopic(topicId)
	if err != nil {
		return err
	}

	events, err := models.ProcessEventsJson(eventsJson)
	if err != nil {
		return err
	}

	err = eb.Events_Db.InsertEvents(events)
	if err != nil {
		return err
	}

	err = topic.AddEvents(events)
	if err != nil {
		return err
	}
	return nil
}
func (eb *EventBus) PublishEvents(topicId string, eventsJson []string) (*models.PublishEventResponse, error) {
	taskResponse := make(chan map[string]interface{}, 1)
	publishTask := NewPublishEventTask(eb, topicId, eventsJson, taskResponse)
	_, err := eb.ExectorPool.Submit(utils.HandlePublishEventWorker, publishTask)
	if err != nil {
		return nil, err
	}
	response := publishTask.Wait()
	if v, ok := response["error"]; ok {
		return nil, errors.New(v.(string))
	}
	return &models.PublishEventResponse{}, nil
}

// Subscribe for Push Events From Eventbus
func initiatePushEventTask(eb *EventBus, topicId, subscriberId string, noOfEvents int) ([]string, error) {
	eventIds := make([]string, noOfEvents)
	topic, err := eb.Topics_Db.GetTopic(topicId)
	if err != nil {
		return eventIds, err
	}
	eventIds, err = topic.GetLastEvents(noOfEvents)
	if err != nil {
		return eventIds, err
	}
	return eventIds, nil
}
func (eb *EventBus) PushEvents(topic, subscriberId string, noOfEvents int) (*models.SubscribeEventResponse, error) {
	pushEventTaskResponse := make(chan map[string]interface{}, 1)
	pushEventTask := NewPushEventTask(eb, topic, subscriberId, pushEventTaskResponse, noOfEvents)
	_, err := eb.ExectorPool.Submit(utils.HandlePushEventWorker, pushEventTask)
	if err != nil {
		return nil, err
	}
	response := pushEventTask.Wait()
	eventsResult := make([]string, noOfEvents)
	if errStr, ok := response["error"].(string); ok {
		return nil, errors.New(errStr)
	} else {
		for k, v := range response {
			idx, _ := strconv.Atoi(k)
			eventsResult[idx] = v.(string)
		}
	}
	if err != nil {
		return nil, err
	}
	return &models.SubscribeEventResponse{
		EventIds: eventsResult,
	}, nil
}

// Subscribe for Poll Events From Eventbus By EventId/Timestamp
func initiatePollEventsById(eb *EventBus, topicId, eventId string, noOfEvents int) ([]string, error) {
	eventIds := make([]string, noOfEvents)
	topic, err := eb.Topics_Db.GetTopic(topicId)
	if err != nil {
		return eventIds, err
	}
	eventIds, err = topic.GetLastEventsFromId(eventId, noOfEvents)
	if err != nil {
		return eventIds, err
	}
	return eventIds, nil
}
func (eb *EventBus) PollEventsWithTimestamp(topicId, subscriberId string, timestamp time.Time, noOfEvents int) (*models.SubscribeEventResponse, error) {
	topic, err := eb.Topics_Db.GetTopic(topicId)
	if err != nil {
		return nil, err
	}
	sourceEventId, err := topic.GetFirstEventIdFromTimestamp(timestamp)
	if err != nil {
		return nil, err
	}
	return eb.PollEventsWithId(topicId, subscriberId, sourceEventId, noOfEvents)
}
func (eb *EventBus) PollEventsWithId(topicId, subscriberId, eventId string, noOfEvents int) (*models.SubscribeEventResponse, error) {
	pollEventWithTSTaskResp := make(chan map[string]interface{}, 1)
	pollEventWithTSTask := NewPollEventFromIdTask(eb, topicId, subscriberId, eventId, noOfEvents, pollEventWithTSTaskResp)
	_, err := eb.ExectorPool.Submit(utils.HandlePullEventByTimestampWorker, pollEventWithTSTask)
	if err != nil {
		return nil, err
	}
	response := pollEventWithTSTask.Wait()
	eventsResult := make([]string, noOfEvents)
	if errStr, ok := response["error"]; ok {
		err = errors.New(errStr.(string))
	} else {
		for k, v := range response {
			idx, _ := strconv.Atoi(k)
			eventsResult[idx] = v.(string)
		}
	}
	if err != nil {
		return nil, err
	}
	return &models.SubscribeEventResponse{
		EventIds: eventsResult,
	}, nil
}

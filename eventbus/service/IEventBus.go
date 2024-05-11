package service

import "github.com/mastik5h/LLD/eventbus/models"

type IEventBus interface {
	PublishEvents(topicId string, events []string) (*models.PublishEventResponse, error)
	PushEvents(topicId, subscriberId string) (*models.SubscribeEventResponse, error)
	PollEventsWithTimestamp(topicId, subscriberId string, timestamp int64) (*models.SubscribeEventResponse, error)
	PollEventsWithId(topicId, subscriberId, eventId string) (*models.SubscribeEventResponse, error)
	CreateTopic(name, description string, capacity int) (*models.TopicStatus, error)
	DeleteTopic(topicId string) error
}

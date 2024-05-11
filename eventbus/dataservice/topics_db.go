package dataservice

import (
	"errors"
	"fmt"

	"github.com/mastik5h/LLD/eventbus/models"
)

type TopicsDB struct {
	topics map[string]*models.Topic
}

func NewTopicsDB() *TopicsDB {
	return &TopicsDB{
		topics: make(map[string]*models.Topic, 0),
	}
}

func (tdb *TopicsDB) GetTopic(topicId string) (*models.Topic, error) {
	topic, found := tdb.topics[topicId]
	if !found {
		return nil, fmt.Errorf("topic not found against given topic_id: %s", topicId)
	}
	return topic, nil
}

func (tdb *TopicsDB) AddTopic(topic *models.Topic) error {
	if _, ok := tdb.topics[topic.Id]; ok {
		return errors.New("topic already exists with given id")
	}
	tdb.topics[topic.Id] = topic
	return nil
}

func (tdb *TopicsDB) DeleteTopic(topic *models.Topic) error {
	if topic == nil {
		return errors.New("no topic found to delete")
	}
	delete(tdb.topics, topic.Id)
	return nil
}

package v1

import (
	"time"

	"github.com/mastik5h/LLD/eventbus/utils"
)

type PublishEventTask struct {
	Eb         *EventBus
	TopicId    string
	EventsJson []string
	Response   chan map[string]interface{}
}

func NewPublishEventTask(eb *EventBus, topicId string, eventsJson []string, response chan map[string]interface{}) utils.ITask {
	return &PublishEventTask{
		Eb:         eb,
		TopicId:    topicId,
		EventsJson: eventsJson,
		Response:   response,
	}
}

func (publish_et *PublishEventTask) Execute() {
	err := initiatePublishEvents(publish_et.Eb, publish_et.TopicId, publish_et.EventsJson)
	if err != nil {
		publish_et.Response <- map[string]interface{}{
			"error": err.Error(),
		}
	}
	publish_et.Response <- map[string]interface{}{
		"code": "success",
	}
}

func (publish_et *PublishEventTask) GetId() string {
	return publish_et.TopicId
}

func (publish_et *PublishEventTask) Wait() (result map[string]interface{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

LOOP:
	for {
		select {
		case <-ticker.C:
			if resp := <-publish_et.Response; resp != nil {
				result = resp
				break LOOP
			}
		}
	}
	return
}

package v1

import (
	"strconv"
	"time"

	"github.com/mastik5h/LLD/eventbus/utils"
)

type PushEventTask struct {
	TopicId      string
	SubscriberId string
	Eb           *EventBus
	Response     chan map[string]interface{}
	NoOfEvents   int
}

func NewPushEventTask(eb *EventBus, topicId, subscriberId string, response chan map[string]interface{}, noOfEvents int) utils.ITask {
	return &PushEventTask{
		TopicId:      topicId,
		SubscriberId: subscriberId,
		Eb:           eb,
		Response:     response,
		NoOfEvents:   noOfEvents,
	}
}

func (push_et *PushEventTask) Execute() {
	response := make(map[string]interface{}, 0)
	event_ids, err := initiatePushEventTask(push_et.Eb, push_et.TopicId, push_et.SubscriberId, push_et.NoOfEvents)
	if err != nil {
		response["error"] = err.Error()
	} else {
		for i, eid := range event_ids {
			response[strconv.Itoa(i)] = eid
		}
	}
	push_et.Response <- response
}

func (push_et *PushEventTask) GetId() string {
	return push_et.TopicId + "-" + push_et.SubscriberId
}

func (push_et *PushEventTask) Wait() (result map[string]interface{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
LOOP:
	for {
		select {
		case <-ticker.C:
			if resp := <-push_et.Response; resp != nil {
				result = resp
			}
			break LOOP
		}
	}
	return
}

/*
e1, e2, e3, e4, e5

*/

package v1

import (
	"strconv"
	"time"

	"github.com/mastik5h/LLD/eventbus/utils"
)

type PollEventFromIdTask struct {
	TopicId      string
	SubscriberId string
	LastEventId  string
	Eb           *EventBus
	Response     chan map[string]interface{}
	NoOfEvents   int
}

func NewPollEventFromIdTask(eb *EventBus, topicId, subscriberId, eventId string, noOfEvents int, response chan map[string]interface{}) utils.ITask {
	return &PollEventFromIdTask{
		TopicId:      topicId,
		SubscriberId: subscriberId,
		LastEventId:  eventId,
		Eb:           eb,
		Response:     response,
		NoOfEvents:   noOfEvents,
	}
}

func (poll_et *PollEventFromIdTask) Execute() {
	response := make(map[string]interface{}, 0)
	event_ids, err := initiatePollEventsById(poll_et.Eb, poll_et.TopicId, poll_et.LastEventId, poll_et.NoOfEvents)
	if err != nil {
		response["error"] = err.Error()
	} else {
		for i, eid := range event_ids {
			response[strconv.Itoa(i)] = eid
		}
	}
	poll_et.Response <- response
}

func (push_et *PollEventFromIdTask) GetId() string {
	return push_et.TopicId + "-" + push_et.SubscriberId
}

func (push_et *PollEventFromIdTask) Wait() (result map[string]interface{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
LOOP:
	for {
		select {
		case <-ticker.C:
			if resp := <-push_et.Response; resp != nil {
				result = resp
				break LOOP
			}
		}
	}
	return
}

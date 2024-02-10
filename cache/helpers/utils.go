package helpers

import "time"

func GetTimeFromCurrentTime(duration *int, defaultDuration time.Duration) time.Time {
	durationTime := defaultDuration
	if duration != nil {
		durationTime = time.Duration(int64(*duration)) * time.Second
	}
	return time.Now().Add(durationTime)
}

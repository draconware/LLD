package utils

import (
	"hash/fnv"
	"time"
)

func GetCurrentTime() time.Time {
	return time.Now()
}

func calculateHash(input string) int {
	h := fnv.New32a()
	h.Write([]byte(input))
	return int(h.Sum32())
}

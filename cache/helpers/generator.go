package helpers

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
)

type StringGenerator struct {
	id int
}

var stringGenerator *StringGenerator

func InitializeStringGenerator() {
	if stringGenerator == nil {
		stringGenerator = &StringGenerator{
			id: 0,
		}
	}
}

func (sg *StringGenerator) SetId(id int) {
	sg.id = id
}

func GetUniqueIdString() string {
	res := stringGenerator.id
	stringGenerator.SetId(res + 1)
	return strconv.FormatInt(int64(res), 10)
}

func GenerateConsistentHashKey(str string) uint64 {
	hash := sha256.Sum256([]byte(str))
	hashValue := binary.BigEndian.Uint64(hash[:8])
	return hashValue
}

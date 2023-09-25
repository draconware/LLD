package generator

import "strconv"

type StringGenerator struct {
	id int
}

var stringGenerator *StringGenerator

func InitializeStringGenerator() {
	if stringGenerator == nil {
		stringGenerator = new(StringGenerator)
		stringGenerator.SetId(1)
	}
}

func GetUniqueIdString() string {
	var uniqueId string
	if stringGenerator != nil {
		uniqueId = strconv.Itoa(stringGenerator.id)
		stringGenerator.SetId(stringGenerator.id + 1)
	}
	return uniqueId
}

func (st *StringGenerator) SetId(id int) {
	st.id = id
}

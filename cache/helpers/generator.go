package helpers

import "strconv"

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

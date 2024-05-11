package utils

import "strconv"

type generator struct {
	id int
}

var idGenerator *generator

func IntializeGenerator(start int) {
	if idGenerator == nil {
		idGenerator = &generator{
			id: start,
		}
	}
}
func (g *generator) SetId(id int) {
	g.id = id
}

func GetUniqueIdInString(name string) string {
	uniqueId := strconv.Itoa(idGenerator.id)
	idGenerator.SetId(idGenerator.id + 1)
	return uniqueId + "-" + name
}
func GetUniqueId() string {
	uniqueId := idGenerator.id
	return strconv.Itoa(uniqueId)
}

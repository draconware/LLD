package models

type Entity struct {
	Id          string
	Name        string
	Description string
	EntityType  EntityTypeEnum
}

type EntityTypeEnum string

const (
	PUBLISHER  EntityTypeEnum = "Publisher"
	SUBSCRIBER EntityTypeEnum = "Subscriber"
	UNKNOWN    EntityTypeEnum = ""
)

func NewEntity(id, name, description, entityType string) *Entity {
	return &Entity{
		Id:          id,
		Name:        name,
		Description: description,
		EntityType:  EntityTypeEnum(entityType),
	}
}

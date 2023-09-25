package data_service

import (
	"errors"

	"github.com/mastik5h/LLD/splitwise/models"
)

var groupDbMap = make(map[string]*models.Group, 0)

func CreateGroup(group *models.Group) error {
	if group == nil {
		return errors.New("no group information found")
	}
	groupDbMap[group.Id] = group
	return nil
}

func GetGroup(groupId string) (*models.Group, error) {
	if group, ok := groupDbMap[groupId]; ok {
		return group, nil
	}
	return nil, errors.New("no group found")
}

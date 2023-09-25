package data_service

import (
	"errors"

	"github.com/mastik5h/LLD/splitwise/models"
)

var userIdDbMap = make(map[string]*models.User, 0)

func CreateUserEntry(user *models.User) error {
	if user == nil {
		return errors.New("no group information found")
	}
	userIdDbMap[user.Id] = user
	return nil
}

func GetUserById(userId string) (*models.User, error) {
	if user, ok := userIdDbMap[userId]; ok {
		return user, nil
	}
	return nil, errors.New("no group found")
}

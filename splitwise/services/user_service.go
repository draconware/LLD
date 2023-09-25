package services

import (
	data_service "github.com/mastik5h/LLD/splitwise/dataServices"
	generator "github.com/mastik5h/LLD/splitwise/helpers"
	"github.com/mastik5h/LLD/splitwise/models"
)

func CreateUser(metaData *models.UserMetaData) models.CreateUserResponse {
	resp := models.CreateUserResponse{}
	if err := validateUserMetadata(metaData); err != nil {
		resp.ErrorDetails = "invalid user details: " + err.Error()
		return resp
	}

	userId := generator.GetUniqueIdString()

	user := &models.User{
		Id:       userId,
		MetaData: metaData,
	}
	if isUserExists(user) {
		resp.ErrorDetails = "user already exists."
		return resp
	}
	if err := data_service.CreateUserEntry(user); err != nil {
		resp.ErrorDetails = "internal error: " + err.Error()
		return resp
	}
	resp.UserId = user.Id
	return resp
}

func ValidateUsersId(users []string) bool {
	for _, user := range users {
		userObj, err := data_service.GetUserById(user)
		if err != nil || !isUserExists(userObj) {
			return false
		}
	}
	return true
}

func isUserExists(user *models.User) bool {
	if user == nil {
		return false
	}
	userState, err := data_service.GetUserById(user.Id)
	if err != nil || userState == nil {
		return false
	}
	return true
}

func validateUserMetadata(metadata *models.UserMetaData) error {
	// not implemented yet
	return nil
}

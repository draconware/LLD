package models

type User struct {
	Id       string
	MetaData *UserMetaData
}

type UserMetaData struct {
	Name     string
	Email    string
	PhoneNum string
}

type CreateUserResponse struct {
	UserId       string
	ErrorDetails string
}

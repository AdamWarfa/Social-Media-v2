package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
)

func GetUsers() ([]models.User, error) {
	return repositories.GetUsers()
}

func GetUser(id string) (models.User, error) {
	return repositories.GetUser(id)
}

func SaveUser(body models.User) (models.User, error) {
	user := models.User{Id: body.Id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}
	return repositories.SaveUser(user)
}

func UpdateUser(id string, body models.User) (models.User, error) {
	user := models.User{Id: id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}
	return repositories.UpdateUser(id, user)
}

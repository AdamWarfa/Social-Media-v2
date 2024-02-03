package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	SaveUser(user models.User) (models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type ProdUserService struct {
	repo repositories.UserRepository
}

func NewProdUserService(repo repositories.UserRepository) *ProdUserService {
	return &ProdUserService{
		repo: repo,
	}
}

func (us *ProdUserService) GetUsers() ([]models.User, error) {
	return us.repo.GetUsers()
}

func (us *ProdUserService) GetUser(id string) (models.User, error) {
	return us.repo.GetUser(id)
}

func (us *ProdUserService) SaveUser(body models.User) (models.User, error) {
	user := models.User{Id: body.Id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}
	return us.repo.SaveUser(user)
}

func (us *ProdUserService) UpdateUser(id string, body models.User) (models.User, error) {
	user := models.User{Id: id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}
	return us.repo.UpdateUser(id, user)
}

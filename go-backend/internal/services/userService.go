package services

import (
	"errors"
	"somev2/internal/models"
	"somev2/internal/repositories"
	"somev2/internal/security"
	"somev2/internal/utilities"

	"go.uber.org/zap"
)

type UserServiceI interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	RegisterUser(username, email, password, avatar string) (*models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
	DeleteUser(username string) (string, error)
	VerifyLogin(email, password string) (string, string, error)
}

type UserService struct {
	repo   repositories.UserRepositoryI
	logger *zap.Logger
}

func NewUserService(repo repositories.UserRepositoryI) *UserService {
	return &UserService{
		repo:   repo,
		logger: utilities.NewLogger(),
	}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	return us.repo.GetUsers()
}

func (us *UserService) GetUser(id string) (models.User, error) {
	return us.repo.GetUser(id)
}

func (us *UserService) RegisterUser(username, email, password, avatar string) (*models.User, error) {

	// Check if username already exists
	existingUser, _ := us.repo.FindByUsername(username)

	if existingUser != nil {
		us.logger.Info("Username already taken", zap.String("username", username))
		return nil, errors.New("the username is already taken")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		us.logger.Error("Failed to hash password", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	// Create the user
	return us.repo.CreateUser(username, email, hashedPassword, avatar)
}

func (us *UserService) UpdateUser(id string, body models.User) (models.User, error) {
	user := models.User{Id: id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}
	return us.repo.UpdateUser(id, user)
}

func (us *UserService) DeleteUser(username string) (string, error) {
	err := us.repo.DeleteUser(username)
	if err != nil {
		us.logger.Error("Failed to delete user", zap.Error(err))
		return "", errors.New("internal server error")
	}

	return username + " has been deleted", nil
}

func (us *UserService) VerifyLogin(email, password string) (string, string, error) {
	user, err := us.repo.FindByEmail(email)
	if err != nil {
		us.logger.Error("Failed to find user by username", zap.Error(err))
		return "", "", errors.New("internal server error")
	}

	if err := security.ComparePasswords(user.Password, password); err != nil {
		us.logger.Error("Failed to compare passwords", zap.Error(err))
		return "", "", errors.New("internal server error")
	}

	token, err := security.GenerateJWT(user.Id, user.Username)
	if err != nil {
		us.logger.Error("Failed to generate JWT token", zap.Error(err))
		return "", "", errors.New("could not generate token")
	}

	return token, user.Username, nil

}

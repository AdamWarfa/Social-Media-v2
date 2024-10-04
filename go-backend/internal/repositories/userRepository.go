package repositories

import (
	"somev2/internal/models"
	"somev2/internal/utilities"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryI interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(username, email, password, avatar string) (*models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
	DeleteUser(username string) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db     *gorm.DB
	logger *zap.Logger
	UserRepositoryI
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: utilities.NewLogger(),
	}
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	result := ur.db.Find(&users)

	if result.Error != nil {
		ur.logger.Error("Failed to fetch users (repo)", zap.Error(result.Error))
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) GetUser(id string) (models.User, error) {
	var user models.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		ur.logger.Error("Failed to fetch (repo)", zap.Error(err))
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(username, email, password, avatar string) (*models.User, error) {
	user := models.User{
		Id:        uuid.New().String(),
		Username:  username,
		Email:     email,
		Password:  password,
		Avatar:    avatar,
		Followers: 0,
		Posts:     []models.Post{},
	}

	// Create the user in the database
	err := ur.db.Create(&user).Error
	if err != nil {
		ur.logger.Error("Error while creating user", zap.Error(err))
		return nil, err
	}

	// The user.ID is now populated after Create()
	ur.logger.Info("User created successfully", zap.String("username", user.Username))
	return &user, nil
}

func (ur *UserRepository) UpdateUser(id string, user models.User) (models.User, error) {
	result := ur.db.Save(&user)
	if result.Error != nil {
		ur.logger.Error("Failed to update user (repo)", zap.Error(result.Error))
		return models.User{}, result.Error
	}

	ur.logger.Info("User updated successfully (repo)")
	return user, nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	var user models.User

	result := ur.db.Where("username = ?", username).First(&user).Delete(&user)

	if result.Error != nil {
		ur.logger.Error("Error while deleting user by username", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		ur.logger.Error("Error while fetching user by username", zap.Error(result.Error))
		return nil, result.Error
	}
	return &user, nil

}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		ur.logger.Error("Error while fetching user by email", zap.Error(result.Error))
		return nil, result.Error
	}
	return &user, nil
}

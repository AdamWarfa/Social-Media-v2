package repositories

import (
	"somev2/internal/models"
	"somev2/internal/utilities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	SaveUser(user models.User) (models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type ProdUserRepository struct {
	db     *gorm.DB
	logger *zap.Logger
	UserRepository
}

func NewProdUserRepository(db *gorm.DB) *ProdUserRepository {
	return &ProdUserRepository{
		db:     db,
		logger: utilities.NewLogger(),
	}
}

func (ur *ProdUserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	result := ur.db.Find(&users)

	if result.Error != nil {
		ur.logger.Error("Failed to fetch users (repo)", zap.Error(result.Error))
		return nil, result.Error
	}

	return users, nil
}

func (ur *ProdUserRepository) GetUser(id string) (models.User, error) {
	var user models.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		ur.logger.Error("Failed to fetch (repo)", zap.Error(err))
		return models.User{}, err
	}

	return user, nil
}

func (ur *ProdUserRepository) SaveUser(user models.User) (models.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		ur.logger.Error("Failed to save user (repo)", zap.Error(result.Error))
		return models.User{}, result.Error
	}

	ur.logger.Info("User saved successfully (repo)")
	return user, nil
}

func (ur *ProdUserRepository) UpdateUser(id string, user models.User) (models.User, error) {
	result := ur.db.Save(&user)
	if result.Error != nil {
		ur.logger.Error("Failed to update user (repo)", zap.Error(result.Error))
		return models.User{}, result.Error
	}

	ur.logger.Info("User updated successfully (repo)")
	return user, nil
}

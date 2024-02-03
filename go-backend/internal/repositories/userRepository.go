package repositories

import (
	"somev2/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	SaveUser(user models.User) (models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type ProdUserRepository struct {
	db *gorm.DB
	UserRepository
}

func NewProdUserRepository(db *gorm.DB) *ProdUserRepository {
	return &ProdUserRepository{
		db: db,
	}
}

func (ur *ProdUserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	result := ur.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (ur *ProdUserRepository) GetUser(id string) (models.User, error) {

	var user models.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *ProdUserRepository) SaveUser(user models.User) (models.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (ur *ProdUserRepository) UpdateUser(id string, user models.User) (models.User, error) {
	result := ur.db.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

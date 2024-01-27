package repositories

import (
	"somev2/internal/initializers"
	"somev2/internal/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	result := initializers.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUser(id string) (models.User, error) {

	var user models.User

	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SaveUser(user models.User) (models.User, error) {
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func UpdateUser(id string, user models.User) (models.User, error) {
	result := initializers.DB.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

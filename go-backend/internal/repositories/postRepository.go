package repositories

import (
	"fmt"
	"somev2/internal/initializers"
	"somev2/internal/models"
)

func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func GetPost(id string) (models.Post, error) {
	var post models.Post

	if err := initializers.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func GetPostsByAuthor(id string) ([]models.Post, error) {
	var posts []models.Post

	result := initializers.DB.Where("author = ?", id).Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func CreatePost(post models.Post) (models.Post, error) {

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}
	fmt.Printf("Post %v saved in DB", &post.Id)

	return post, nil
}

func LikePost(id string, post *models.Post) (models.Post, error) {

	result := initializers.DB.Save(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return *post, nil
}

func DeletePost(id string) error {

	result := initializers.DB.Where("id = ?", id).Delete(&models.Post{})

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Post deleted from DB")

	return nil
}

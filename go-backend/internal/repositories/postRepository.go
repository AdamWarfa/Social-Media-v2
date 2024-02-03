package repositories

import (
	"fmt"
	"somev2/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	GetPostsByAuthor(id string) ([]models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	LikePost(id string, post *models.Post) (models.Post, error)
	DeletePost(id string) error
}

type ProdPostRepository struct {
	db *gorm.DB
	PostRepository
}

func NewProdPostRepository(db *gorm.DB) *ProdPostRepository {
	return &ProdPostRepository{
		db: db,
	}
}

func (pr *ProdPostRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	result := pr.db.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (pr *ProdPostRepository) GetPost(id string) (models.Post, error) {
	var post models.Post

	if err := pr.db.Where("id = ?", id).First(&post).Error; err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (pr *ProdPostRepository) GetPostsByAuthor(id string) ([]models.Post, error) {
	var posts []models.Post

	result := pr.db.Where("author = ?", id).Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (pr *ProdPostRepository) CreatePost(post models.Post) (models.Post, error) {

	result := pr.db.Create(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}
	fmt.Printf("Post %v saved in DB", &post.Id)

	return post, nil
}

func (pr *ProdPostRepository) LikePost(id string, post *models.Post) (models.Post, error) {

	result := pr.db.Save(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return *post, nil
}

func (pr *ProdPostRepository) DeletePost(id string) error {

	result := pr.db.Where("id = ?", id).Delete(&models.Post{})

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Post deleted from DB")

	return nil
}

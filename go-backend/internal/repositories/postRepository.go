package repositories

import (
	"somev2/internal/models"
	"somev2/internal/utilities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostRepositoryI interface {
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	GetPostsByAuthor(id string) ([]models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	LikePost(id string, post *models.Post) (models.Post, error)
	DeletePost(id string) error
}

type PostRepository struct {
	db     *gorm.DB
	logger *zap.Logger
	PostRepositoryI
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db:     db,
		logger: utilities.NewLogger(),
	}
}

func (pr *PostRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	result := pr.db.Find(&posts)

	if result.Error != nil {
		pr.logger.Error("Failed to fetch posts", zap.Error(result.Error))
		return nil, result.Error
	}

	return posts, nil
}

func (pr *PostRepository) GetPost(id string) (models.Post, error) {
	var post models.Post

	if err := pr.db.Where("id = ?", id).First(&post).Error; err != nil {
		pr.logger.Error("Failed to fetch post (repo)", zap.Error(err))
		return models.Post{}, err
	}
	return post, nil
}

func (pr *PostRepository) GetPostsByAuthor(id string) ([]models.Post, error) {
	var posts []models.Post

	result := pr.db.Where("author = ?", id).Find(&posts)

	if result.Error != nil {
		pr.logger.Error("Failed to fetch posts by author (repo)", zap.Error(result.Error))
		return nil, result.Error
	}

	return posts, nil
}

func (pr *PostRepository) CreatePost(post models.Post) (models.Post, error) {
	result := pr.db.Create(&post)

	if result.Error != nil {
		pr.logger.Error("Failed to save post in database (repo)", zap.Error(result.Error))
		return models.Post{}, result.Error
	}

	pr.logger.Info("Post saved in DB (repo)", zap.String("id", post.Id))
	return post, nil
}

func (pr *PostRepository) LikePost(id string, post *models.Post) (models.Post, error) {
	result := pr.db.Save(&post)

	if result.Error != nil {
		pr.logger.Error("Failed to like post in database (repo)", zap.Error(result.Error))
		return models.Post{}, result.Error
	}

	pr.logger.Info("Post liked in database (repo)", zap.String("id", post.Id))
	return *post, nil
}

func (pr *PostRepository) DeletePost(id string) error {
	result := pr.db.Where("id = ?", id).Delete(&models.Post{})

	if result.Error != nil {
		pr.logger.Error("Failed to delete post from database (repo)", zap.Error(result.Error))
		return result.Error
	}

	pr.logger.Info("Post deleted from database (repo)", zap.String("id", id))
	return nil
}

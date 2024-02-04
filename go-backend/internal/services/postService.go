package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
	"somev2/internal/utilities"

	"go.uber.org/zap"
)

// PostService is a contract for the PostService
type PostService interface {
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	GetPostsByAuthor(id string) ([]models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	LikePost(id string, post *models.Post) (models.Post, error)
	DeletePost(id string) error
}

// ProdPostService is a struct for the PostService
type ProdPostService struct {
	repo   repositories.PostRepository
	logger *zap.Logger
}

// NewProdPostService is a constructor for the ProdPostService
func NewProdPostService(repo repositories.PostRepository) *ProdPostService {
	return &ProdPostService{
		repo:   repo,
		logger: utilities.NewLogger(),
	}
}

// GetPosts is a method to get all posts
func (ps *ProdPostService) GetPosts() ([]models.Post, error) {
	return ps.repo.GetPosts()
}

// GetPost is a method to get a post by id
func (ps *ProdPostService) GetPost(id string) (models.Post, error) {
	return ps.repo.GetPost(id)
}

func (ps *ProdPostService) GetPostsByAuthor(id string) ([]models.Post, error) {
	return ps.repo.GetPostsByAuthor(id)
}

func (ps *ProdPostService) CreatePost(post models.Post) (models.Post, error) {
	return ps.repo.CreatePost(post)
}

func (ps *ProdPostService) LikePost(id string, post *models.Post) (models.Post, error) {

	post.Likes = post.Likes + 1

	likedPost, err := ps.repo.LikePost(id, post)
	if err != nil {
		ps.logger.Error("Failed to like post (service)", zap.Error(err))
		return models.Post{}, err
	}

	return likedPost, nil
}

func (ps *ProdPostService) DeletePost(id string) error {
	return ps.repo.DeletePost(id)
}

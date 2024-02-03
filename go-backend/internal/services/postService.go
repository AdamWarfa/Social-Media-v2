package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
)

type PostService interface {
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	GetPostsByAuthor(id string) ([]models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	LikePost(id string, post *models.Post) (models.Post, error)
	DeletePost(id string) error
}

type ProdPostService struct {
	repo repositories.PostRepository
}

func NewProdPostService(repo repositories.PostRepository) *ProdPostService {
	return &ProdPostService{
		repo: repo,
	}
}

func (ps *ProdPostService) GetPosts() ([]models.Post, error) {
	return ps.repo.GetPosts()
}

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

	return ps.repo.LikePost(id, post)
}

func (ps *ProdPostService) DeletePost(id string) error {
	return ps.repo.DeletePost(id)
}

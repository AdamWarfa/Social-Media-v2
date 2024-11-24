package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
	"somev2/internal/utilities"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// PostService is a contract for the PostService
type PostServiceI interface {
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	GetPostsByAuthor(id string) ([]models.Post, error)
	CreatePost(PostRequest models.PostRequest) (models.Post, error)
	DeletePost(id string) error
}

// ProdPostService is a struct for the PostService
type PostService struct {
	repo   repositories.PostRepositoryI
	logger *zap.Logger
}

// NewProdPostService is a constructor for the PostService
func NewPostService(repo repositories.PostRepositoryI) *PostService {
	return &PostService{
		repo:   repo,
		logger: utilities.NewLogger(),
	}
}

// GetPosts is a method to get all posts
func (ps *PostService) GetPosts() ([]models.Post, error) {
	return ps.repo.GetPosts()
}

// GetPost is a method to get a post by id
func (ps *PostService) GetPost(id string) (models.Post, error) {
	return ps.repo.GetPost(id)
}

func (ps *PostService) GetPostsByAuthor(id string) ([]models.Post, error) {
	return ps.repo.GetPostsByAuthor(id)
}

func (ps *PostService) CreatePost(PostRequest models.PostRequest) (models.Post, error) {

	newUUID := uuid.New()
	date := time.Now().UTC()
	// Format to ISO 8601 (RFC 3339)
	formattedDate := date.Format("2006-01-02T15:04:05.000Z")

	post := models.Post{Id: newUUID.String(), Text: PostRequest.Text, ImgSrc: PostRequest.ImgSrc, AuthorId: PostRequest.AuthorId, Likes: []models.Like{}, PostDate: formattedDate}

	return ps.repo.CreatePost(post)
}

func (ps *PostService) DeletePost(id string) error {
	return ps.repo.DeletePost(id)
}

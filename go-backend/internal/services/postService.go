package services

import (
	"somev2/internal/models"
	"somev2/internal/repositories"
)

func GetPosts() ([]models.Post, error) {
	return repositories.GetPosts()
}

func GetPost(id string) (models.Post, error) {
	return repositories.GetPost(id)
}

func GetPostsByAuthor(id string) ([]models.Post, error) {
	return repositories.GetPostsByAuthor(id)
}

func CreatePost(post models.Post) (models.Post, error) {
	return repositories.CreatePost(post)
}

package services

import "somev2/internal/repositories"

type LikeServiceI interface {
	LikePost(userID, postID string) error
	UnlikePost(userID, postID string) error
	GetLikeCount(postID string) (int64, error)
	HasUserLikedPost(userID, postID string) (bool, error)
}

type LikeService struct {
	repo repositories.LikeRepositoryI
}

func NewLikeService(repo repositories.LikeRepositoryI) *LikeService {
	return &LikeService{repo: repo}
}

func (ls *LikeService) LikePost(userID, postID string) error {
	// Add like
	return ls.repo.AddLike(userID, postID)
}

func (ls *LikeService) UnlikePost(userID, postID string) error {
	// Remove like
	return ls.repo.RemoveLike(userID, postID)
}

func (ls *LikeService) GetLikeCount(postID string) (int64, error) {
	// Get like count
	return ls.repo.CountLikes(postID)
}

func (ls *LikeService) HasUserLikedPost(userID, postID string) (bool, error) {
	// Check if user has liked the post
	return ls.repo.HasLiked(userID, postID)
}

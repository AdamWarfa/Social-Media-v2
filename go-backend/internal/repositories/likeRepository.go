package repositories

import (
	"errors"
	"somev2/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeRepositoryI interface {
	AddLike(userId, postId string) error
	RemoveLike(userID, postId string) error
	HasLiked(userID, postId string) (bool, error)
	CountLikes(postId string) (int64, error)
}

type LikeRepository struct {
	db *gorm.DB
	LikeRepositoryI
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

func (lr *LikeRepository) AddLike(userId, postId string) error {
	// Check if user already liked the post
	var like models.Like
	err := lr.db.Where("user_id = ? AND post_id = ?", userId, postId).First(&like).Error
	if err == nil {
		return errors.New("user already liked this post")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Add like
	like = models.Like{
		Id:     uuid.NewString(),
		UserId: userId,
		PostId: postId,
	}
	return lr.db.Create(&like).Error
}

func (lr *LikeRepository) RemoveLike(userId, postId string) error {
	return lr.db.Where("user_id = ? AND post_id = ?", userId, postId).Delete(&models.Like{}).Error
}

func (lr *LikeRepository) HasLiked(userId, postId string) (bool, error) {
	var count int64
	err := lr.db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", userId, postId).Count(&count).Error
	return count > 0, err
}

func (lr *LikeRepository) CountLikes(postId string) (int64, error) {
	var count int64
	err := lr.db.Model(&models.Like{}).Where("post_id = ?", postId).Count(&count).Error
	return count, err
}

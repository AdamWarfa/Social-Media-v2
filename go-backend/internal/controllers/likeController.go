package controllers

import (
	"fmt"
	"somev2/internal/services"

	"github.com/gofiber/fiber/v2"
)

type LikeControllerI interface {
	LikePost(ctx *fiber.Ctx) error
	UnlikePost(ctx *fiber.Ctx) error
	GetLikeCount(ctx *fiber.Ctx) error
	HasUserLiked(ctx *fiber.Ctx) error
}

type LikeController struct {
	service services.LikeServiceI
}

func NewLikeController(service services.LikeServiceI) *LikeController {
	return &LikeController{service: service}
}

func (lc *LikeController) LikePost(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(string)
	postID := ctx.Params("postId")

	fmt.Println("uid", userID, "pid", postID)

	if err := lc.service.LikePost(userID, postID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Post liked successfully",
	})
}

func (lc *LikeController) UnlikePost(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(string)
	postID := ctx.Params("postId")

	if err := lc.service.UnlikePost(userID, postID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Post unliked successfully",
	})
}

func (lc *LikeController) GetLikeCount(ctx *fiber.Ctx) error {
	postID := ctx.Params("postId")

	count, err := lc.service.GetLikeCount(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"postId": postID,
		"likes":  count,
	})
}

func (lc *LikeController) HasUserLiked(ctx *fiber.Ctx) error {
	// Hent userId fra kontekst, med validering
	userID, ok := ctx.Locals("userId").(string)
	if !ok || userID == "" {
		fmt.Println("uid", userID)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID is missing or invalid",
		})
	}

	// Hent postId fra URL-parametre
	postID := ctx.Params("postId")
	if postID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	// Tjek om brugeren har liket posten
	hasLiked, err := lc.service.HasUserLikedPost(userID, postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Returner resultatet
	return ctx.JSON(fiber.Map{
		"postId":   postID,
		"hasLiked": hasLiked,
	})
}

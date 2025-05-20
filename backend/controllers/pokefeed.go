package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Post a new PokePost (image URL + caption)
func PostPokePost(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uuid.UUID)

    var body struct {
        Caption  string `json:"caption"`
        ImageURL string `json:"image_url"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if body.ImageURL == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Image URL is required"})
    }

    post := models.PokePost{
        ID:        uuid.New(),
        UserID:    userID,
        Caption:   body.Caption,
        ImageURL:  body.ImageURL,
        CreatedAt: time.Now(),
    }

    if err := database.DB.Create(&post).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
    }

    return c.JSON(post)
}

// Get paginated feed of posts
func GetPokeFeed(c *fiber.Ctx) error {
    page, _ := strconv.Atoi(c.Query("page", "1"))
    pageSize := 20
    offset := (page - 1) * pageSize

    var posts []models.PokePost
    if err := database.DB.Preload("User").
        Order("created_at desc").
        Offset(offset).
        Limit(pageSize).
        Find(&posts).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch feed"})
    }

    return c.JSON(posts)
}

// Get posts by specific user
func GetUserPokePosts(c *fiber.Ctx) error {
    userID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    var posts []models.PokePost
    if err := database.DB.Where("user_id = ?", userID).
        Order("created_at desc").
        Find(&posts).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user posts"})
    }

    return c.JSON(posts)
}

// Like a PokePost
func LikePokePost(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uuid.UUID)
    postID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }

    like := models.PokePostLike{
        ID:         uuid.New(),
        UserID:     userID,
        PokePostID: postID,
    }

    if err := database.DB.Create(&like).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to like post"})
    }

    return c.JSON(fiber.Map{"message": "Post liked"})
}

func CommentOnPokePost(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uuid.UUID)
    postID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }

    var body struct {
        Content string `json:"content"`
    }
    if err := c.BodyParser(&body); err != nil || body.Content == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid content"})
    }

    comment := models.PokePostComment{
        ID:         uuid.New(),
        UserID:     userID,
        PokePostID: postID,
        Content:    body.Content,
        CreatedAt:  time.Now(),
    }

    if err := database.DB.Create(&comment).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to comment"})
    }

    return c.JSON(comment)
}

func UnlikePokePost(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uuid.UUID)
    postID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }

    if err := database.DB.Where("user_id = ? AND poke_post_id = ?", userID, postID).
        Delete(&models.PokePostLike{}).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unlike post"})
    }

    return c.JSON(fiber.Map{"message": "Post unliked"})
}

func FollowTrainer(c *fiber.Ctx) error {
    followerID := c.Locals("user_id").(uuid.UUID)
    followedID, err := uuid.Parse(c.Params("id"))
    if err != nil || followerID == followedID {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid follow"})
    }

    follow := models.UserFollow{
        ID:         uuid.New(),
        FollowerID: followerID,
        FollowedID: followedID,
    }

    if err := database.DB.Create(&follow).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to follow"})
    }

    return c.JSON(fiber.Map{"message": "Followed"})
}

func UnfollowTrainer(c *fiber.Ctx) error {
    followerID := c.Locals("user_id").(uuid.UUID)
    followedID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid unfollow"})
    }

    if err := database.DB.Where("follower_id = ? AND followed_id = ?", followerID, followedID).
        Delete(&models.UserFollow{}).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unfollow"})
    }

    return c.JSON(fiber.Map{"message": "Unfollowed"})
}

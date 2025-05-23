package controllers

import (
	"time"

	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateForumTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	var input struct {
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		CategoryID uuid.UUID `json:"category_id"`
		Pinned     bool      `json:"pinned"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	topic := models.ForumTopic{
		Title:      input.Title,
		AuthorID:   userID,
		CategoryID: input.CategoryID,
		Pinned:     input.Pinned,
		CreatedAt:  time.Now(),
	}
	if err := database.DB.Create(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topic)
}

func GetForumTopics(c *fiber.Ctx) error {
	var topics []models.ForumTopic
	if err := database.DB.Preload("Author").Preload("Category").Find(&topics).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topics)
}

func GetForumTopicById(c *fiber.Ctx) error {
	id := c.Params("id")
	topicID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.Preload("Author").Preload("Category").First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(topic)
}

func EditForumTopicById(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	if topic.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	var input struct {
		Title      *string    `json:"title"`
		Content    *string `json:"content"`
		CategoryID *uuid.UUID `json:"category_id"`
		Pinned     *bool      `json:"pinned"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if input.Title != nil {
		topic.Title = *input.Title
	}
	if input.Content != nil {
		topic.Content = *input.Content
	}
	if input.CategoryID != nil {
		topic.CategoryID = *input.CategoryID
	}
	if input.Pinned != nil {
		topic.Pinned = *input.Pinned
	}
	topic.UpdatedAt = time.Now()
	if err := database.DB.Save(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topic)
}

func DeleteForumTopicById(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	if topic.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	if err := database.DB.Delete(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func CreateForumCategory(c *fiber.Ctx) error {
	var input models.ForumCategory
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(input)
}

func GetForumCategories(c *fiber.Ctx) error {
	var categories []models.ForumCategory
	if err := database.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

func GetForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var category models.ForumCategory
	if err := database.DB.First(&category, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(category)
}

func EditForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var input models.ForumCategory
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	var category models.ForumCategory
	if err := database.DB.First(&category, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	category.Name = input.Name
	category.Color = input.Color
	category.Description = input.Description
	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(category)
}

func DeleteForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	if err := database.DB.Delete(&models.ForumCategory{}, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Topic interactions
func LikeForumTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic UUID"})
	}

	like := models.ForumTopicLike{
		UserID:  userID,
		TopicID: topicID,
	}

	if err := database.DB.Create(&like).Error; err != nil {
		// Prevent duplicate likes
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "already liked"})
	}

	// Optionally update like counter
	database.DB.Model(&models.ForumTopic{}).Where("id = ?", topicID).UpdateColumn("likes_count", gorm.Expr("likes_count + 1"))

	return c.SendStatus(fiber.StatusCreated)
}

func GetForumTopicComments(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic UUID"})
	}

	// Parse pagination query params with defaults
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	var comments []models.ForumTopicComment
	var total int64

	// Count total comments
	if err := database.DB.
		Model(&models.ForumTopicComment{}).
		Where("topic_id = ?", topicID).
		Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Fetch paginated comments with user info
	if err := database.DB.
		Preload("User").
		Where("topic_id = ?", topicID).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"comments": comments,
		"pagination": fiber.Map{
			"total":  total,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func CommentOnForumTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic UUID"})
	}

	var input struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&input); err != nil || input.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	comment := models.ForumTopicComment{
		UserID:  userID,
		TopicID: topicID,
		Content: input.Content,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Optionally update reply count
	database.DB.Model(&models.ForumTopic{}).Where("id = ?", topicID).UpdateColumn("replies_count", gorm.Expr("replies_count + 1"))

	return c.JSON(comment)
}

func ViewForumTopic(c *fiber.Ctx) error {
	var userID *uuid.UUID
	if uid, ok := c.Locals("user_id").(uuid.UUID); ok {
		userID = &uid
	}

	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic UUID"})
	}

	ip := c.IP()

	view := models.ForumTopicView{
		TopicID:   topicID,
		UserID:    userID,
		IPAddress: ip,
	}

	// Prevent duplicate view from same user/IP (optional)
	_ = database.DB.Create(&view)

	// Update views count (can be approximate)
	database.DB.Model(&models.ForumTopic{}).Where("id = ?", topicID).UpdateColumn("views_count", gorm.Expr("views_count + 1"))

	return c.SendStatus(fiber.StatusNoContent)
}

func UnlikeForumTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic UUID"})
	}

	if err := database.DB.Where("user_id = ? AND topic_id = ?", userID, topicID).Delete(&models.ForumTopicLike{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Model(&models.ForumTopic{}).Where("id = ?", topicID).UpdateColumn("likes_count", gorm.Expr("likes_count - 1"))

	return c.SendStatus(fiber.StatusNoContent)
}

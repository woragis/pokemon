package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func parseUUIDParam(c *fiber.Ctx, param string) (uuid.UUID, error) {
	return uuid.Parse(c.Params(param))
}

// ✅ Working
// ✅ Tested (not yet)
func PostShout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	var body struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	shout := models.Shout{
		ID:        uuid.New(),
		UserID:    userID,
		Content:   body.Content,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&shout).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shout)
}

// ✅ Working
// ✅ Tested (not yet)
func GetShoutFeed(c *fiber.Ctx) error {
	var shouts []models.Shout

	if err := database.DB.
		Preload("User").
		Preload("Comments").
		Preload("Likes").
		Preload("ReshoutOf").
		Order("created_at desc").
		Limit(50).
		Find(&shouts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shouts)
}

// ✅ Working
// ✅ Tested (not yet)
func LikeShout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID) // ✅ fixed key
	shoutID, err := parseUUIDParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest // ✅ handle user error gracefully
	}

	// Optional: prevent duplicate likes
	var existing models.ShoutLike
	if err := database.DB.
		Where("user_id = ? AND shout_id = ?", userID, shoutID).
		First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Already liked",
		})
	}

	like := models.ShoutLike{
		UserID:  userID,
		ShoutID: shoutID,
	}

	if err := database.DB.Create(&like).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "liked"})
}

// ✅ Working
// ✅ Tested (not yet)
// Add a comment to a shout
func CommentOnShout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	shoutID, err := parseUUIDParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var body struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	comment := models.ShoutComment{
		UserID:  userID,
		ShoutID: shoutID,
		Content: body.Content,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(comment)
}

// ✅ Working
// ✅ Tested (not yet)
// User's timeline
func GetUserTimeline(c *fiber.Ctx) error {
	userID, err := parseUUIDParam(c, "id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	var shouts []models.Shout
	if err := database.DB.
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Comments").
		Preload("Likes").
		Order("created_at desc").
		Limit(30).
		Find(&shouts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shouts)
}

func GetPaginatedFeed(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize := 20
	offset := (page - 1) * pageSize

	var shouts []models.Shout
	if err := database.DB.
		Preload("User").
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&shouts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shouts)
}

func Reshout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID) // correct key and type

	shoutID, err := parseUUIDParam(c, "id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid shout ID")
	}

	var body struct {
		Quote string `json:"quote"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	shout := models.Shout{
		UserID:      userID,
		ReshoutOfID: &shoutID,
	}

	if body.Quote != "" {
		shout.QuoteContent = &body.Quote
	}

	if err := database.DB.Create(&shout).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shout)
}

func GetInfiniteFeed(c *fiber.Ctx) error {
	cursor := c.Query("cursor")
	pageSize := 20

	var shouts []models.Shout
	query := database.DB.Preload("User").Order("created_at desc").Limit(pageSize)

	if cursor != "" {
		parsedTime, err := time.Parse(time.RFC3339, cursor)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid cursor format")
		}
		query = query.Where("created_at < ?", parsedTime)
	}

	if err := query.Find(&shouts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shouts)
}

func DeleteReshout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	shoutID, err := parseUUIDParam(c, "id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid shout ID")
	}

	var shout models.Shout
	if err := database.DB.First(&shout, "id = ?", shoutID).Error; err != nil {
		return fiber.ErrNotFound
	}

	if shout.UserID != userID {
		return fiber.NewError(fiber.StatusForbidden, "Not your shout")
	}

	if shout.ReshoutOfID == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Not a reshout")
	}

	if err := database.DB.Delete(&shout).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "Reshout deleted"})
}

// // AI-generated reply (just a mocked example)
// func AIReplyToShout(c *fiber.Ctx) error {
// 	userID := c.Locals("userID").(uint)
// 	shoutID := parseUint(c.Params("id"))

// 	var original models.Shout
// 	if err := database.DB.First(&original, shoutID).Error; err != nil {
// 		return fiber.ErrNotFound
// 	}

// 	aiReply := "That's a cool shout! What Pokémon team are you building?" // This should be replaced with a real call to an AI

// 	reply := models.Shout{
// 		UserID:         userID,
// 		Content:        aiReply,
// 		AIReplyContent: &aiReply,
// 	}

// 	if err := database.DB.Create(&reply).Error; err != nil {
// 		return fiber.ErrInternalServerError
// 	}

// 	return c.JSON(reply)
// }

package forum

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s topicService
	categoryService topicCategoryService
	commentService topicCommentService
	commentLikeService commentLikeService
	viewService topicViewService
	likeService topicLikeService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newTopicRepository(db)
	service := newTopicService(repo, redis)

	categoryRepo := newTopicCategoryRepository(db)
	categoryService := newTopicCategoryService(categoryRepo, redis)

	commentRepo := newTopicCommentRepository(db)
	commentService := newTopicCommentService(commentRepo, redis)

	commentLikeRepo := newCommentLikeRepository(db)
	commentLikeService := newCommentLikeService(commentLikeRepo, redis)

	viewRepo := newTopicViewRepository(db)
	viewService := newTopicViewService(viewRepo, redis)

	likeRepo := newTopicLikeRepository(db)
	likeService := newTopicLikeService(likeRepo, redis)

	return &handler{
		s: service,
		categoryService: categoryService,
		commentService: commentService,
		commentLikeService: commentLikeService,
		viewService: viewService,
		likeService: likeService,
	}
}

// POST /topics
func (h *handler) createTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid user ID")
	}

	var topic Topic
	if err := c.BodyParser(&topic); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	topic.UserID = userID

	if err := topic.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.create(&topic); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(topic)
}

// GET /topics/:id
func (h *handler) getTopicByID(c *fiber.Ctx) error {
	id := c.Params("id")
	topic, err := h.s.getByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "topic not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(topic)
}

// PUT /topics/:id
func (h *handler) updateTopic(c *fiber.Ctx) error {
	id := c.Params("id")

	var input Topic
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	input.ID, _ = uuid.Parse(id)

	if err := input.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.update(&input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(input)
}

// DELETE /topics/:id
func (h *handler) deleteTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.s.delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// GET /topics
func (h *handler) listTopic(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	topics, err := h.s.list(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(topics)
}

// GET /topics/user/user_id
func (h *handler) listTopicByUser(c *fiber.Ctx) error {
	userIDParam := c.Params("user_id")

	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id format")
	}

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	topics, err := h.s.listByUser(userID, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(topics)
}

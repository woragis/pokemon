package user

import (
	"pokemon/pkg/utils"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
    service Service
}

func NewHandler(db *gorm.DB, redis *redis.Client) *Handler {
    repo := NewRepository(db, redis)
    service := NewService(repo)
    
    return &Handler{
        service: service,
    }
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
    var req CreateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
    }
    
    if err := utils.ValidateStruct(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
    user, err := h.service.CreateUser(c.Context(), &req)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }
    
    return utils.SuccessResponse(c, fiber.StatusCreated, "User created successfully", user)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
    }
    
    user, err := h.service.GetUser(c.Context(), uint(id))
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
    }
    
    return utils.SuccessResponse(c, fiber.StatusOK, "User retrieved successfully", user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
    }
    
    var req UpdateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
    }
    
    user, err := h.service.UpdateUser(c.Context(), uint(id), &req)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }
    
    return utils.SuccessResponse(c, fiber.StatusOK, "User updated successfully", user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
    }
    
    err = h.service.DeleteUser(c.Context(), uint(id))
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }
    
    return utils.SuccessResponse(c, fiber.StatusOK, "User deleted successfully", nil)
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    offset, _ := strconv.Atoi(c.Query("offset", "0"))
    
    users, err := h.service.ListUsers(c.Context(), limit, offset)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }
    
    return utils.SuccessResponse(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func (h *Handler) Login(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
    }
    
    if err := utils.ValidateStruct(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
    token, err := h.service.Login(c.Context(), &req)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
    }
    
    return utils.SuccessResponse(c, fiber.StatusOK, "Login successful", map[string]string{
        "token": token,
    })
}

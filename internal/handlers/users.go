package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler interface {
	GetAll(c fiber.Ctx) error
	GetByID(c fiber.Ctx) error
	GetByName(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{service: userService}
}

func (h *userHandler) GetAll(c fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return err
	}

	c.JSON(users)

	return nil
}

func (h *userHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.service.GetByID(id)
	if err != nil {
		return err
	}

	c.JSON(user)

	return nil
}

func (h *userHandler) GetByName(c fiber.Ctx) error {
	name := c.Params("name")

	user, err := h.service.GetByName(name)
	if err != nil {
		return err
	}

	c.JSON(user)

	return nil
}

func (h *userHandler) Create(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return err
	}

	user, err := h.service.Create(*user)
	if err != nil {
		return err
	}

	c.JSON(user)

	return nil
}

func (h *userHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	rowsAffected, err := h.service.Delete(id)
	if err != nil {
		return err
	}

	c.JSON(fiber.Map{"status": "ok", "rowsAffected": rowsAffected})

	return nil
}

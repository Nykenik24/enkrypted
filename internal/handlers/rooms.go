package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/services"
	"github.com/gofiber/fiber/v3"
)

type RoomHandler interface {
	GetAll(c fiber.Ctx) error
	GetByID(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type roomHandler struct {
	service services.RoomService
}

func NewRoomHandler(roomService services.RoomService) RoomHandler {
	return &roomHandler{service: roomService}
}

func (h *roomHandler) GetAll(c fiber.Ctx) error {
	rooms, err := h.service.GetAll()
	if err != nil {
		return err
	}

	c.JSON(rooms)

	return nil
}

func (h *roomHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	room, err := h.service.GetByID(id)
	if err != nil {
		return err
	}

	c.JSON(room)

	return nil
}

func (h *roomHandler) Create(c fiber.Ctx) error {
	room := new(models.Room)
	if err := c.Bind().Body(room); err != nil {
		return err
	}

	room, err := h.service.Create(*room)
	if err != nil {
		return err
	}

	c.JSON(room)

	return nil
}

func (h *roomHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	rowsAffected, err := h.service.Delete(id)
	if err != nil {
		return err
	}

	c.JSON(fiber.Map{"status": "ok", "rowsAffected": rowsAffected})

	return nil
}

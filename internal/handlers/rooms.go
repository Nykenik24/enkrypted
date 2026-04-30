package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/services"
	"github.com/gofiber/fiber/v3"
)

func GetAllRooms(c fiber.Ctx) error {
	rooms, err := services.GetAllRooms()
	if err != nil {
		c.Status(500)
		return err
	}

	c.JSON(rooms)

	return nil
}

func CreateRoom(c fiber.Ctx) error {
	room := new(models.Room)

	if err := c.Bind().Body(room); err != nil {
		return err
	}

	err := services.CreateRoom(*room)
	if err != nil {
		return err
	}

	c.SendString("ok")

	return nil
}

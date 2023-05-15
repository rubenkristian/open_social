package handlers

import (
	"rubenkristian/open-social/app/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Router fiber.Router
	Models *models.Models
}

func BuildHandler(router fiber.Router, models *models.Models) *Handler {
	return &Handler{
		Router: router,
		Models: models,
	}
}

func (h *Handler) Run() {

}

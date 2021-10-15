package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RootHandler(c *fiber.Ctx) error
	AboutHandler(c *fiber.Ctx) error
}

type DefaultHandler struct {
}

func (h *DefaultHandler) RootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *DefaultHandler) AboutHandler(c *fiber.Ctx) error {
	return c.SendString("About Page")
}

package adapters

import (
	"github.com/gofiber/fiber/v2"
)

type Fiber interface {
	Get(path string, handlers ...fiber.Handler) fiber.Router
	Listen(addr string) error
}

type FiberCtx interface {
	SendString(body string) error
}

func NewFiber() Fiber {
	return fiber.New()
}

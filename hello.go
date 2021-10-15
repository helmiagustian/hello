package hello

import (
	"log"

	"github.com/helmiagustian/hello/pkg/adapters"
	"github.com/helmiagustian/hello/pkg/handlers"
)

type IHelo interface {
	SetupRouters(h handlers.Handler)
	ListenAndServe()
}

type Hello struct {
	Fiber adapters.Fiber
}

func NewHello() IHelo {
	return &Hello{
		Fiber: adapters.NewFiber(),
	}
}

func (hello *Hello) SetupRouters(h handlers.Handler) {
	hello.Fiber.Get("/", h.RootHandler)
	hello.Fiber.Get("/about", h.AboutHandler)
}

func (hello *Hello) ListenAndServe() {
	log.Fatal(hello.Fiber.Listen(":3000"))
}

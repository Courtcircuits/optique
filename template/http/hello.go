package http

import "github.com/gofiber/fiber/v2"

type HelloController interface {
	Hello() fiber.Handler
	Register(app *fiber.App)
}

type helloController struct{}

func NewHelloController() *helloController {
	return &helloController{}
}

func (h *helloController) Hello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}

func (h *helloController) Register(app *fiber.App) {
	app.Get("/", h.Hello())
}

package http

import "github.com/gofiber/fiber/v2"

type Validator interface {
	Validate(c *fiber.Ctx) error
}

package requestbody

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

func BodyHandler(c *fiber.Ctx) error {
	body := c.Body()
	log.Info(body)
	return c.SendString(string(body))
}

func BodyParseHandler(c *fiber.Ctx) error {

	user := &User{}
	if err := c.BodyParser(user); err != nil {
		return err
	}
	log.Info(user)
	return c.JSON(user)
}

func ValidateHandler(c *fiber.Ctx) error {
	user := &User{}
	if err := c.BodyParser(user); err != nil {
		log.Error(err)
		return err
	}

	if err := ValidCheck(c, user); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(user)
}

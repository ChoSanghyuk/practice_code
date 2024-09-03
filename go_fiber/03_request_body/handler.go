package requestbody

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

func BodyHandler(c *fiber.Ctx) error {
	body := c.Body()
	log.Info(string(body))
	return c.SendString(string(body))
}

func BodyParserHandler(c *fiber.Ctx) error {

	user := &User{}
	if err := c.BodyParser(user); err != nil {
		log.Error(user)
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

	if err := validCheck(user); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(user)
}

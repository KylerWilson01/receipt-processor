package controllers

import "github.com/gofiber/fiber/v2"

type ReceiptHandler struct{}

func (*ReceiptHandler) GetReceipts(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}

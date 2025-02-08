// Package routes defines the endpoints for the app
package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/KylerWilson01/receipt-processor.git/controllers"
)

// ReceiptRoute holdes the endpoints for receipts
func ReceiptRoute(a fiber.Router) {
	h := controllers.ReceiptHandler{}
	r := a.Group("/receipt")

	r.Get("/", h.GetReceipts)
}

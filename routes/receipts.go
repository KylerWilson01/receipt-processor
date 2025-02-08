// Package routes defines the endpoints for the app
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/KylerWilson01/receipt-processor.git/controllers"
	"github.com/KylerWilson01/receipt-processor.git/models"
)

// ReceiptRoute holdes the endpoints for receipts
func ReceiptRoute(a fiber.Router) {
	// Creates the ReceiptHandler with an empty map
	h := controllers.ReceiptHandler{
		Receipts: make(map[uuid.UUID]models.Receipt),
	}
	// Groups the routes
	r := a.Group("/receipts")
	// Creates the endpoints
	r.Get("/:id/points", h.GetReceiptPoints)
	r.Post("/process", h.ProcessReceipt)
}

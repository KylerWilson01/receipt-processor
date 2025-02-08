// Package controllers is used to store all the controllers for this program
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/KylerWilson01/receipt-processor.git/models"
	"github.com/KylerWilson01/receipt-processor.git/utils"
)

const (
	// ReceiptNotFound is the error message if a receipt cannot be found with the given id
	ReceiptNotFound = "No receipt found for that ID."
	// InvalidReceipt is the error message if a given receipt is unable to be parsed
	InvalidReceipt = "The receipt is invalid."
)

// ReceiptHandler is the struct that will store the receipts in memory and has all the endpoint functions attached to it
type ReceiptHandler struct {
	Receipts map[uuid.UUID]models.Receipt
}

// GetReceiptPoints takes an id and uses it to find the points for a receipt
func (h *ReceiptHandler) GetReceiptPoints(c *fiber.Ctx) error {
	// Grabs the id from the params
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		// If there is no given id we error out saying we cannot find the receipt
		return c.Status(400).JSON(ReceiptNotFound)
	}

	// Checks to make sure the receipt is stored in memory
	val, ok := h.Receipts[id]
	if !ok {
		// If not we error out saying we cannot find the receipt
		return c.Status(400).JSON(ReceiptNotFound)
	}

	// Return the points for the given id
	return c.JSON(fiber.Map{"points": val.Points})
}

// ProcessReceipt takes a receipt, parses it out, and adds up all the points
func (h *ReceiptHandler) ProcessReceipt(c *fiber.Ctx) error {
	// Create the new id
	guid := uuid.New()

	// Parse the body
	receipt := new(models.Receipt)
	if err := c.BodyParser(receipt); err != nil {
		return c.Status(400).JSON(err)
	}

	// Set to the new guid
	receipt.ID = guid

	// Set the points for the receipt for later use
	totalPoints := 0
	pUtil := utils.PointUtil{}

	totalPoints += pUtil.CheckRetailerName(receipt.Retailer)
	totalPoints += pUtil.CheckDate(receipt.PurchaseDate)
	totalPoints += pUtil.CheckTime(receipt.PurchaseTime)
	totalPoints += pUtil.CheckRoundDollar(receipt.Total)
	totalPoints += pUtil.CheckMultiple(receipt.Total)
	totalPoints += pUtil.CountLengthOfItems(receipt.Items)

	for _, i := range receipt.Items {
		totalPoints += pUtil.CheckDescriptionLength(i)
	}

	receipt.Points = totalPoints

	// Add the receipt to the map
	h.Receipts[guid] = *receipt

	// return the id
	return c.JSON(fiber.Map{"id": guid})
}

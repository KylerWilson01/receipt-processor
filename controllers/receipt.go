package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/KylerWilson01/receipt-processor.git/models"
	"github.com/KylerWilson01/receipt-processor.git/utils"
)

const (
	ReceiptNotFound = "No receipt found for that ID."
	InvalidReceipt  = "The receipt is invalid."
)

type ReceiptHandler struct {
	Receipts map[uuid.UUID]models.Receipt
}

func (h *ReceiptHandler) GetReceiptPoints(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(ReceiptNotFound)
	}

	val, ok := h.Receipts[id]

	if !ok {
		return c.Status(400).JSON(ReceiptNotFound)
	}

	return c.JSON(fiber.Map{"points": val.Points})
}

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

	// Parse the receipt and set the points
	totalPoints := 0
	pUtil := utils.PointUtil{}

	totalPoints += pUtil.CheckRetailerName(receipt.Retailer)
	fmt.Println("Retailer Name: ", totalPoints)
	totalPoints += pUtil.CheckDate(receipt.PurchaseDate)
	fmt.Println("Purchase Date: ", totalPoints)
	totalPoints += pUtil.CheckTime(receipt.PurchaseTime)
	fmt.Println("Purchase Time: ", totalPoints)
	totalPoints += pUtil.CheckRoundDollar(receipt.Total)
	fmt.Println("Round Dollar: ", totalPoints)
	totalPoints += pUtil.CheckMultiple(receipt.Total)
	fmt.Println("Multiple of a Quarter: ", totalPoints)
	totalPoints += pUtil.CountLengthOfItems(receipt.Items)
	fmt.Println("Length of Items: ", totalPoints)

	for _, i := range receipt.Items {
		totalPoints += pUtil.CheckDescriptionLength(i)
	}
	fmt.Println(totalPoints)

	receipt.Points = totalPoints

	// Add the receipt to the map
	h.Receipts[guid] = *receipt

	// return the id
	return c.JSON(fiber.Map{"id": guid})
}

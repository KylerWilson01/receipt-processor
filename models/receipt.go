package models

import (
	"github.com/google/uuid"
)

// Receipt is what gets parsed from the user
type Receipt struct {
	ID           uuid.UUID `json:"id"`
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        string    `json:"total"`
	Points       int       `json:"-"`
}
